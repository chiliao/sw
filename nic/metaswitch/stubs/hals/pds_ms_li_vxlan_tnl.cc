//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// LI VXLAN Tunnel HAL integration
//---------------------------------------------------------------

#include <thread>
#include "nic/metaswitch/stubs/hals/pds_ms_li_vxlan_tnl.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_li_vxlan_port.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_util.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_hal_init.hpp"
#include "nic/metaswitch/stubs/mgmt/pds_ms_mgmt_state.hpp"
#include "nic/sdk/lib/logger/logger.hpp"
#include <li_fte.hpp>
#include <li_lipi_slave_join.hpp>
#include <li_vxlan.hpp>

extern NBB_ULONG li_proc_id;

//-------------------------------------------------------------------
//       MS obj                  PDS HAL Spec
// a) Pathset containing
//    L3 interfaces  -> Underlay ECMP NHs (referenced by TEP entries)
// b) VXLAN Tunnel   -> TEP entry -> Overlay ECMP entry (ref by Type2 MAC,IP)
// c) L2 VXLAN Port  -> Unused (since Type 2 VNI comes from Egress BD)
// d) L3 VXLAN Port  -> TEP,VNI entry (referenced by Type5 Overlay ECMP)
// e) Pathset containing
//    L3 VXLAN Ports -> Overlay ECMP entry (ref by Type5 Prefix routes)
//--------------------------------------------------------------------

namespace pds_ms {

void li_vxlan_tnl::fetch_store_info_(pds_ms::state_t* state) {
    store_info_.tun_if_obj = state->if_store().get(ips_info_.if_index);
    if (op_delete_) {
        if (unlikely(store_info_.tun_if_obj == nullptr)) {
            throw Error("VXLAN Tunnel delete for unknown IfIndex " + std::to_string(ips_info_.if_index));
        }
        auto& tun_prop = store_info_.tun_if_obj->vxlan_tunnel_properties();
        store_info_.tep_obj = state->tep_store().get(tun_prop.tep_ip);
        SDK_ASSERT (store_info_.tep_obj != nullptr);
    } else {
        store_info_.tep_obj = state->tep_store().get(ips_info_.tep_ip);
    }
}

pds_tep_spec_t li_vxlan_tnl::make_pds_tep_spec_(void) {
    pds_tep_spec_t spec = {0};
    auto& tep_prop = store_info_.tep_obj->properties();
    spec.key = make_pds_tep_key_();
    spec.remote_ip = tep_prop.tep_ip;
    spec.ip_addr = tep_prop.src_ip;
    spec.nh_type = PDS_NH_TYPE_UNDERLAY_ECMP;
    spec.nh_group = msidx2pdsobjkey(tep_prop.hal_uecmp_idx, true);
    spec.type = PDS_TEP_TYPE_WORKLOAD;
    spec.nat = false;
    return spec;
}

pds_nexthop_group_spec_t li_vxlan_tnl::make_pds_nhgroup_spec_(void) {
    pds_nexthop_group_spec_t spec = {0};
    auto& tep_prop = store_info_.tep_obj->properties();
    spec.key = make_pds_nhgroup_key_();
    spec.type = PDS_NHGROUP_TYPE_OVERLAY_ECMP;
    spec.num_nexthops = 1;
    spec.nexthops[0].type = PDS_NH_TYPE_OVERLAY;
    // Use the TEP MS IfIndex as the TEP Index
    spec.nexthops[0].tep = msidx2pdsobjkey(tep_prop.hal_tep_idx);
    return spec;
}

void li_vxlan_tnl::parse_ips_info_(ATG_LIPI_VXLAN_ADD_UPDATE* vxlan_tnl_add_upd_ips) {
    ips_info_.if_index = vxlan_tnl_add_upd_ips->id.if_index;
    ATG_INET_ADDRESS& ms_dest_ip = vxlan_tnl_add_upd_ips->vxlan_settings.dest_ip;
    ms_to_pds_ipaddr(ms_dest_ip, &ips_info_.tep_ip);
    ATG_INET_ADDRESS& ms_src_ip = vxlan_tnl_add_upd_ips->vxlan_settings.source_ip;
    ms_to_pds_ipaddr(ms_src_ip, &ips_info_.src_ip);
    NBB_CORR_GET_VALUE(ips_info_.hal_uecmp_idx,
                       vxlan_tnl_add_upd_ips->vxlan_settings.dp_pathset_correlator);
    ips_info_.tep_ip_str = ipaddr2str(&ips_info_.tep_ip);
    NBB_CORR_GET_VALUE(ips_info_.uecmp_ps, vxlan_tnl_add_upd_ips->vxlan_settings.pathset_id);
}

void li_vxlan_tnl::cache_obj_in_cookie_for_create_op_(void) {
    if (likely (store_info_.tun_if_obj == nullptr)) {
        // Create new If Object but do not save it in the Global State yet
        // This automatically allocates a HAL Overlay ECMP table index
        std::unique_ptr<if_obj_t> new_if_obj
            (new if_obj_t(if_obj_t::vxlan_tunnel_properties_t
                              {ips_info_.if_index,
                              ips_info_.tep_ip}));
        // Update the local store info context so that the make_pds_spec
        // refers to the latest fields
        store_info_.tun_if_obj = new_if_obj.get();
        // Cache the new object in the cookie to revisit asynchronously
        // when the PDS API response is received
        cookie_uptr_->objs.push_back(std::move (new_if_obj));
    } else {
        auto& tnl_if_prop = store_info_.tun_if_obj->vxlan_tunnel_properties();
        // Dest IP cannot change for existing tunnel
        SDK_ASSERT(IPADDR_EQ (&tnl_if_prop.tep_ip, &ips_info_.tep_ip));
    }

    // Create new Tep Object but do not save it in the Global State yet
    // Use the MS Tunnel IfIndex as the HAL index for TEP table
    // ECMP Table index is allocated in constructor for every new TEP object
    std::unique_ptr<tep_obj_t> new_tep_obj
        (new tep_obj_t(ips_info_.tep_ip, ips_info_.src_ip,
                       store_info_.ll_direct_pathset, ips_info_.if_index));
    // Update the local store info context so that the make_pds_spec
    // refers to the latest fields
    store_info_.tep_obj = new_tep_obj.get();
    // Cache the new object in the cookie to revisit asynchronously
    // when the PDS API response is received
    cookie_uptr_->objs.push_back(std::move(new_tep_obj));
}

bool
li_vxlan_tnl::cache_obj_in_cookie_for_update_op_(void) {
    // Updating existing tunnel - check all properties to see what has changed
    auto& tnl_if_prop = store_info_.tun_if_obj->vxlan_tunnel_properties();
    // Dest IP cannot change for existing tunnel
    SDK_ASSERT(IPADDR_EQ (&tnl_if_prop.tep_ip, &ips_info_.tep_ip));

    if (unlikely (store_info_.ll_direct_pathset ==
                  store_info_.tep_obj->properties().hal_uecmp_idx)) {
        // No change in TEP
        return false;
    }

    // Create a new object with the updated fields
    // but do not save it in the Global State yet
    // This does not allocate a new Overlay ECMP index
    // since the old obj is copied to the new obj first
    std::unique_ptr<tep_obj_t> new_tep_obj
        (new tep_obj_t(*(store_info_.tep_obj)));
    new_tep_obj->properties().hal_uecmp_idx = store_info_.ll_direct_pathset;
    // Update the local store info context so that the make_pds_spec
    // refers to the latest fields
    store_info_.tep_obj = new_tep_obj.get();
    // Cache the new object in the cookie to revisit asynchronously
    // when the PDS API response is received
    cookie_uptr_->objs.push_back(std::move(new_tep_obj));
    return true;
}

pds_batch_ctxt_guard_t li_vxlan_tnl::make_batch_pds_spec_(state_t::context_t&
                                                          state_ctxt) {
    pds_batch_ctxt_guard_t bctxt_guard_;
    sdk_ret_t ret = SDK_RET_OK;

    SDK_ASSERT(cookie_uptr_); // Cookie should not be empty
    pds_batch_params_t bp {PDS_BATCH_PARAMS_EPOCH, PDS_BATCH_PARAMS_ASYNC,
                           pds_ms::hal_callback,
                           cookie_uptr_.get()};
    auto bctxt = pds_batch_start(&bp);

    if (unlikely (!bctxt)) {
        throw Error("PDS Batch Start failed for TEP "
                    + ips_info_.tep_ip_str);
    }
    bctxt_guard_.set (bctxt);

    if (op_delete_) { // Delete
        // First delete Overlay ECMP entry before TEP entry to ensure
        // Overlay ECMP does not point to deleted TEP in hardware.
        auto nhgroup_key = make_pds_nhgroup_key_();
        if (!PDS_MOCK_MODE()) {
            ret = pds_nexthop_group_delete(&nhgroup_key, bctxt);
        }
        if (unlikely (ret != SDK_RET_OK)) {
            throw Error(std::string("PDS ECMP Delete failed for TEP ")
                        .append(ipaddr2str(&store_info_.tep_obj->properties().tep_ip))
                        .append(" err=").append(std::to_string(ret)));
        }

        auto tep_key = make_pds_tep_key_();
        if (!PDS_MOCK_MODE()) {
            ret = pds_tep_delete(&tep_key, bctxt);
        }
        if (unlikely (ret != SDK_RET_OK)) {
            throw Error(std::string("PDS TEP Delete failed for TEP ")
                        .append(ipaddr2str(&store_info_.tep_obj->properties().tep_ip))
                        .append(" err=").append(std::to_string(ret)));
        }

    } else { // Add or update
        auto tep_spec = make_pds_tep_spec_();
        if (op_create_) {
            if (!PDS_MOCK_MODE()) {
                ret = pds_tep_create(&tep_spec, bctxt);
            }
        } else {
            if (!PDS_MOCK_MODE()) {
                ret = pds_tep_update(&tep_spec, bctxt);
            }

            // For TEP uecmp update need to update all L3 VXLAN Ports
            auto l_tep_obj = store_info_.tep_obj;
            store_info_.tep_obj->walk_l3_vxlan_ports(
                [&state_ctxt, l_tep_obj, bctxt] (ms_ifindex_t vxp_ifindex) -> bool {
                    auto vxp_if_obj = state_ctxt.state()->if_store().get(vxp_ifindex);
                    SDK_ASSERT(vxp_if_obj != nullptr);
                    auto& vxp_prop = vxp_if_obj->vxlan_port_properties();
                    PDS_TRACE_DEBUG("TEP %s Updating Underlay ECMP Index for MS"
                                    " L3 VXLAN Port %d",
                                    ipaddr2str(&vxp_prop.tep_ip), vxp_ifindex);
                    li_vxlan_port vxp;
                    vxp.add_pds_tep_spec(bctxt, vxp_if_obj, l_tep_obj,
                                         false /* Op Update */);
                    return false; // continue loop
                });
        }
        if (unlikely (ret != SDK_RET_OK)) {
            throw Error(std::string("PDS TEP Create or Update failed for TEP ")
                        .append(ips_info_.tep_ip_str)
                        .append(" err=").append(std::to_string(ret)));
        }

        if (op_create_) {
            // Create Overlay ECMP entry for a new TEP
            // No change in Overlay ECMP for TEP update
            auto nhgroup_spec = make_pds_nhgroup_spec_();
            if (!PDS_MOCK_MODE()) {
                ret = pds_nexthop_group_create(&nhgroup_spec, bctxt);
            }
            if (unlikely (ret != SDK_RET_OK)) {
                throw Error(std::string("PDS ECMP Create failed for TEP ")
                            .append(ips_info_.tep_ip_str)
                            .append(" err=").append(std::to_string(ret)));
            }
        }
    }
    return bctxt_guard_;
}

static ms_ps_id_t map_indirect_ps_2_tep_ip_(state_t* state, ms_ps_id_t indirect_pathset,
                                         const ip_addr_t& tep_ip) {
    auto indirect_ps_obj = state->indirect_ps_store().get(indirect_pathset);
    if (indirect_ps_obj == nullptr) {
        PDS_TRACE_DEBUG("Map new indirect underlay pathset %d to TEP %s",
                        indirect_pathset, ipaddr2str(&tep_ip));
        std::unique_ptr<indirect_ps_obj_t> indirect_ps_obj_uptr 
            (new indirect_ps_obj_t(tep_ip));
        state->indirect_ps_store().add_upd(indirect_pathset,
                                           std::move(indirect_ps_obj_uptr));
        return PDS_MS_ECMP_INVALID_INDEX;
    }

    if (!ip_addr_is_zero(&(indirect_ps_obj->tep_ip))) {
        // Assert there is only 1 TEP referring to each indirect Pathset
        if (!ip_addr_is_equal (&(indirect_ps_obj->tep_ip), &tep_ip)) {
            PDS_TRACE_ERR("Attempt to stitch TEP %s to MS indirect pathset %d"
                          " that is already stitched to TEP %s",
                          ipaddr2str(&tep_ip), indirect_pathset,
                          ipaddr2str(&(indirect_ps_obj->tep_ip)));
            SDK_ASSERT(0);
        }
        return indirect_ps_obj->ll_direct_pathset;
    }
    PDS_TRACE_DEBUG("Stitch TEP %s to indirect pathset %d direct pathset %d",
                    ipaddr2str(&tep_ip), indirect_pathset,
                    indirect_ps_obj->ll_direct_pathset);
    indirect_ps_obj->tep_ip = tep_ip;
    return indirect_ps_obj->ll_direct_pathset;
}

static void unmap_indirect_ps_2_tep_ip_(state_t* state,
                                        ms_ps_id_t indirect_pathset) {
    PDS_TRACE_DEBUG("Unstitch from indirect pathset %d",
                    indirect_pathset);
    auto indirect_ps_obj = state->indirect_ps_store().get(indirect_pathset);
    if (indirect_ps_obj == nullptr) {
        return;
    }
    ip_addr_t zero_ip = {0};
    indirect_ps_obj->tep_ip = zero_ip;
}

NBB_BYTE li_vxlan_tnl::handle_add_upd_() {

    NBB_BYTE rc = ATG_OK;
    pds_batch_ctxt_guard_t pds_bctxt_guard;

    { // Enter thread-safe context to access/modify global state
    auto state_ctxt = pds_ms::state_t::thread_context();

    fetch_store_info_(state_ctxt.state());

    // Associate Indirect ECMP Pathset -> TEP
    // This will not be deleted even if the Tunnel create fails
    // Entry will be erased when the Pathset is deleted from Metaswitch
    store_info_.ll_direct_pathset =
        map_indirect_ps_2_tep_ip_(state_ctxt.state(),
                                  ips_info_.uecmp_ps, ips_info_.tep_ip);

    if (store_info_.ll_direct_pathset != ips_info_.hal_uecmp_idx) {
        PDS_TRACE_DEBUG("Detected parallel update to TEP %s (ecmp: %d) and "
                        "indirect pathset %d (ecmp: %d). prefer indirect pathset",
                        ips_info_.tep_ip_str.c_str(), ips_info_.hal_uecmp_idx,
                        ips_info_.uecmp_ps, store_info_.ll_direct_pathset);
    if (store_info_.ll_direct_pathset == PDS_MS_ECMP_INVALID_INDEX) {
        PDS_TRACE_DEBUG("Ignore VXLAN Tunnel %s update pointing to"
                        " blackholed pathset %d",
                        ips_info_.tep_ip_str.c_str(),ips_info_.uecmp_ps);
        return rc;
    }
    }

    if (store_info_.tep_obj != nullptr) {
        // Update Tunnel
        auto old_uecmp_ps = store_info_.tep_obj->properties().uecmp_ps;
        if (old_uecmp_ps != ips_info_.uecmp_ps) {
            PDS_TRACE_DEBUG("TEP %s change in indirect pathset, "
                            "unmap from old indirect ps %d",
                            ips_info_.tep_ip_str.c_str(), old_uecmp_ps);
            unmap_indirect_ps_2_tep_ip_(state_ctxt.state(), old_uecmp_ps);
        }
        PDS_TRACE_INFO ("TEP %s Update IPS Underlay ECMP Pathset %d DP Corr %d",
                        ips_info_.tep_ip_str.c_str(), ips_info_.uecmp_ps,
                        store_info_.ll_direct_pathset);
        if (unlikely(!cache_obj_in_cookie_for_update_op_())) {
            // No change
            return rc;
        }
    } else {
        // Create Tunnel
        PDS_TRACE_INFO ("TEP %s Create IPS Underlay ECMP Pathset %d DP Corr %d",
                        ips_info_.tep_ip_str.c_str(), ips_info_.uecmp_ps, store_info_.ll_direct_pathset);
        op_create_ = true;
        cache_obj_in_cookie_for_create_op_();
    }
    store_info_.tep_obj->properties().uecmp_ps = ips_info_.uecmp_ps;

    pds_bctxt_guard = make_batch_pds_spec_(state_ctxt);

    // If we have batched multiple IPS earlier flush it now
    // Cannot add a Tunnel create to an existing batch
    state_ctxt.state()->flush_outstanding_pds_batch();

    } // End of state thread_context
      // Do Not access/modify global state after this

    // All processing complete, only batch commit remains -
    // safe to release the cookie_uptr_ unique_ptr
    rc = ATG_ASYNC_COMPLETION;
    auto cookie = cookie_uptr_.release();
    auto ret = pds_batch_commit(pds_bctxt_guard.release());
    if (unlikely (ret != SDK_RET_OK)) {
        delete cookie;
        throw Error(std::string("Batch commit failed for Add-Update TEP ")
                    .append(ips_info_.tep_ip_str)
                    .append(" err=").append(std::to_string(ret)));
    }
    PDS_TRACE_DEBUG ("TEP %s: Add/Upd PDS Batch commit successful",
                     ips_info_.tep_ip_str.c_str());
    if (PDS_MOCK_MODE()) {
        // Call the HAL callback in PDS mock mode
        std::thread cb(pds_ms::hal_callback, SDK_RET_OK, cookie);
        cb.detach();
    }
    return rc;
}

// Underlying Indirect Pathset for the tunnel has been updated with a new
// DP correlator (ECMP 1 to 2 case)
NBB_BYTE li_vxlan_tnl::handle_uecmp_update(state_t::context_t&& state_ctxt,
                                           tep_obj_t* tep_obj,
                                           ms_ps_id_t pathset_id,
                                           ms_ps_id_t ll_dp_corr_id,
                                           cookie_uptr_t&& cookie_uptr) {
    // Mock as if VXLAN Tunnel update IPS is received
    ips_info_.src_ip = tep_obj->properties().src_ip;
    ips_info_.tep_ip = tep_obj->properties().tep_ip;
    // MS LIM IfIndex for the VXLAN Tunnel is used as the HAL TEP Index
    ips_info_.if_index = tep_obj->properties().hal_tep_idx;
    store_info_.ll_direct_pathset = ll_dp_corr_id;
    ips_info_.hal_uecmp_idx = ll_dp_corr_id;
    NBB_CORR_GET_VALUE(ips_info_.uecmp_ps, pathset_id);
    state_ctxt.release();

    cookie_uptr_ = std::move(cookie_uptr);
    return handle_add_upd_();
}

NBB_BYTE li_vxlan_tnl::handle_add_upd_ips(ATG_LIPI_VXLAN_ADD_UPDATE* vxlan_tnl_add_upd_ips) {

    parse_ips_info_(vxlan_tnl_add_upd_ips);
    // Alloc new cookie and cache IPS
    cookie_uptr_.reset (new cookie_t);

    cookie_uptr_->send_ips_reply =
        [vxlan_tnl_add_upd_ips] (bool pds_status, bool ips_mock) -> void {
            // ----------------------------------------------------------------
            // This block is executed asynchronously when PDS response is rcvd
            // ----------------------------------------------------------------
            if (unlikely(ips_mock)) return; // UT

            NBB_CREATE_THREAD_CONTEXT
            NBS_ENTER_SHARED_CONTEXT(li_proc_id);
            NBS_GET_SHARED_DATA();

            auto key = li::VxLan::get_key(*vxlan_tnl_add_upd_ips);
            auto& vxlan_store = li::Fte::get().get_lipi_join()->get_vxlan_store();
            auto it = vxlan_store.find(key);
            if (it == vxlan_store.end()) {
                // MS Stub Stateless mode
                auto send_response = li::VxLan::set_ips_rc(&vxlan_tnl_add_upd_ips->ips_hdr,
                                                          (pds_status) ? ATG_OK : ATG_UNSUCCESSFUL);
                SDK_ASSERT(send_response);
                PDS_TRACE_DEBUG("++++++ VXLAN Tunnel 0x%x: Send Async IPS reply %s stateless mode ++++++",
                                key, (pds_status) ? "Success": "Failure");
                li::Fte::get().get_lipi_join()->send_ips_reply(&vxlan_tnl_add_upd_ips->ips_hdr);
            } else {
                // MS Stub Stateful mode
                if (pds_status) {
                    PDS_TRACE_DEBUG("VXLAN Tunnel 0x%x: Send Async IPS Reply success stateful mode",
                                     key);
                    (*it)->update_complete(ATG_OK);
                } else {
                    PDS_TRACE_DEBUG("VXLAN Tunnel 0x%x: Send Async IPS Reply failure stateful mode",
                                     key);
                    (*it)->update_failed(ATG_UNSUCCESSFUL);
                }
            }
            NBS_RELEASE_SHARED_DATA();
            NBS_EXIT_SHARED_CONTEXT();
            NBB_DESTROY_THREAD_CONTEXT
        };

    return handle_add_upd_();
}

void li_vxlan_tnl::handle_delete(NBB_ULONG tnl_ifindex) {
    pds_batch_ctxt_guard_t  pds_bctxt_guard;
    op_delete_ = true;

    // MS stub Integration APIs do not support Async callback for deletes.
    // However since we should not block the MS NBase main thread
    // the HAL processing is always asynchronous even for deletes.
    // Assuming that Deletes never fail the Store is also updated
    // in a synchronous fashion for deletes so that it is in sync
    // if there is a subsequent create from MS.

    ips_info_.if_index = tnl_ifindex;
    // Empty cookie to force async PDS.
    cookie_uptr_.reset (new cookie_t);
    ip_addr_t tep_ip;

    { // Enter thread-safe context to access/modify global state
        auto state_ctxt = pds_ms::state_t::thread_context();

        fetch_store_info_(state_ctxt.state());
        if (unlikely (store_info_.tep_obj == nullptr &&
                      store_info_.tun_if_obj == nullptr)) {
            // No change
            return;
        }
        tep_ip = store_info_.tep_obj->properties().tep_ip;
        PDS_TRACE_INFO ("TEP %s: Delete IPS", ipaddr2str(&tep_ip));

        pds_bctxt_guard = make_batch_pds_spec_ (state_ctxt);

        // If we have batched multiple IPS earlier flush it now
        // VXLAN Tunnel deletion cannot be appended to an existing batch
        state_ctxt.state()->flush_outstanding_pds_batch();

        unmap_indirect_ps_2_tep_ip_(state_ctxt.state(),
            store_info_.tep_obj->properties().uecmp_ps);

    } // End of state thread_context
      // Do Not access/modify global state after this

    cookie_uptr_->send_ips_reply =
        [tep_ip] (bool pds_status, bool ips_mock) -> void {
            // ----------------------------------------------------------------
            // This block is executed asynchronously when PDS response is rcvd
            // ----------------------------------------------------------------
            PDS_TRACE_DEBUG("+++++++ TEP %s Delete: Rcvd Async PDS response %s +++++++++",
                            ipaddr2str(&tep_ip), (pds_status)?"Success": "Failure");

        };

    // All processing complete, only batch commit remains -
    // safe to release the cookie_uptr_ unique_ptr
    auto cookie = cookie_uptr_.release();
    auto ret = pds_batch_commit(pds_bctxt_guard.release());
    if (unlikely (ret != SDK_RET_OK)) {
        delete cookie;
        throw Error(std::string("Batch commit failed for delete TEP ")
                    .append(ipaddr2str(&tep_ip))
                    .append(" err=").append(std::to_string(ret)));
    }
    PDS_TRACE_DEBUG ("TEP %s: Delete PDS Batch commit successful",
                     ipaddr2str(&tep_ip));

    { // Enter thread-safe context to access/modify global state
        auto state_ctxt = pds_ms::state_t::thread_context();
        // Deletes are synchronous - Delete the store entry immediately
        state_ctxt.state()->tep_store().erase(tep_ip);
        state_ctxt.state()->if_store().erase(ips_info_.if_index);
    }
}

} // End namespace
