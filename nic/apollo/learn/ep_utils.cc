//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// utilities to handle EP MAC and IP entries
///
//----------------------------------------------------------------------------

#include "nic/sdk/include/sdk/eth.hpp"
#include "nic/sdk/include/sdk/l2.hpp"
#include "nic/sdk/lib/event_thread/event_thread.hpp"
#include "nic/apollo/api/include/pds_batch.hpp"
#include "nic/apollo/api/internal/pds_mapping.hpp"
#include "nic/apollo/api/subnet.hpp"
#include "nic/apollo/api/utils.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/learn/ep_utils.hpp"
#include "nic/apollo/learn/learn_impl_base.hpp"
#include "nic/apollo/learn/learn_state.hpp"
#include "nic/apollo/learn/utils.hpp"

namespace event = sdk::event_thread;

namespace learn {

static sdk_ret_t
delete_ip_mapping (ep_ip_entry *ip_entry, pds_batch_ctxt_t bctxt)
{
    pds_mapping_key_t mapping_key;
    const ep_ip_key_t *ep_ip_key = ip_entry->key();

    mapping_key.type = PDS_MAPPING_TYPE_L3;
    mapping_key.vpc = ep_ip_key->vpc;
    mapping_key.ip_addr = ep_ip_key->ip_addr ;
    return api::pds_local_mapping_delete(&mapping_key, bctxt);
}

static sdk_ret_t
delete_ip_entry (ep_ip_entry *ip_entry, ep_mac_entry *mac_entry)
{
    sdk_ret_t ret;

    event::timer_stop(ip_entry->timer());
    ip_entry->set_state(EP_STATE_DELETED);
    mac_entry->del_ip(ip_entry);

    ret = ip_entry->del_from_db();
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to delete EP %s from db, error code %u",
                      ip_entry->key2str().c_str(), ret);
        return ret;
    }
    return ip_entry->delay_delete();
}

sdk_ret_t
delete_ip_from_ep (ep_ip_entry *ip_entry, pds_batch_ctxt_t bctxt)
{
    sdk_ret_t ret;

    ret = delete_ip_mapping(ip_entry, bctxt);
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to delete IP mapping for EP %s, error code %u",
                      ip_entry->key2str().c_str(), ret);
        return ret;
    }

    PDS_TRACE_INFO("Deleting IP mapping %s", ip_entry->key2str().c_str());
    return delete_ip_entry(ip_entry, ip_entry->mac_entry());
}

static bool
delete_ip_mapping_cb (void *obj, void *ctxt)
{
    ep_ip_entry *ip_entry = (ep_ip_entry *)obj;
    pds_batch_ctxt_t bctxt = PDS_BATCH_CTXT_INVALID;

    if (ctxt) {
        bctxt = *((pds_batch_ctxt_t *)ctxt);
    }
    return (delete_ip_mapping(ip_entry, bctxt) == SDK_RET_OK);
}

static bool
delete_ip_entry_cb (void *obj, void *ctxt)
{
    ep_ip_entry *ip_entry = (ep_ip_entry *)obj;
    ep_mac_entry *mac_entry = (ep_mac_entry *)ctxt;

    return (delete_ip_entry(ip_entry, mac_entry) == SDK_RET_OK);
}

static sdk_ret_t
delete_vnic (ep_mac_entry *mac_entry, pds_batch_ctxt_t bctxt)
{
    pds_obj_key_t vnic_key;

    vnic_key = api::uuid_from_objid(mac_entry->vnic_obj_id());
    return pds_vnic_delete(&vnic_key, bctxt);
}

static sdk_ret_t
delete_mac_entry (ep_mac_entry *mac_entry)
{
    sdk_ret_t ret;

    timer_stop(mac_entry->timer());
    mac_entry->set_state(EP_STATE_DELETED);
    learn_db()->vnic_obj_id_free(mac_entry->vnic_obj_id());

    ret = mac_entry->del_from_db();
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to delete EP %s from db, error code %u",
                      mac_entry->key2str().c_str(), ret);
        return ret;
    }
    return mac_entry->delay_delete();
}

// note: caller needs to check if it is expected that there be no IPs associated
// with the EP before deleting it
// note: if caller provides batch, we commit it here TODO: split delete api into
// hardware state delete and software state delete, this would help us batch all
// deletes into a single batch when deleting all endpoints under a subnet
sdk_ret_t
delete_ep (ep_mac_entry *mac_entry, pds_batch_ctxt_t bctxt)
{
    sdk_ret_t ret;

    // delete all IP mappings and then MAC mapping
    if (bctxt == PDS_BATCH_CTXT_INVALID) {
        // start a batch so that all IP mappings and vnic can be deleted
        // together, if successful, then only we will delete sw state
        // this way, we do not end up with inconsistent hw and sw states
        pds_batch_params_t batch_params {learn_db()->epoch_next(), false,
                                         nullptr, nullptr};
        bctxt = pds_batch_start(&batch_params);
        if (unlikely(bctxt == PDS_BATCH_CTXT_INVALID)) {
            PDS_TRACE_ERR("Failed to create api batch");
            return SDK_RET_ERR;
        }
    }

    mac_entry->walk_ip_list(delete_ip_mapping_cb, (void *)&bctxt);
    ret = delete_vnic(mac_entry, bctxt);
    if (ret != SDK_RET_OK) {
        pds_batch_destroy(bctxt);
        PDS_TRACE_ERR("Failed to delete EP %s, error code %u",
                      mac_entry->key2str().c_str(), ret);
        return ret;
    }

    ret = pds_batch_commit(bctxt);
    LEARN_COUNTER_INCR(api_calls);
    if (unlikely(ret != SDK_RET_OK)) {
        PDS_TRACE_ERR("Failed to commit API batch, error code %u", ret);
        LEARN_COUNTER_INCR(api_failure);
        return SDK_RET_ERR;
    }
    PDS_TRACE_INFO("Deleted EP %s", mac_entry->key2str().c_str());

    // delete sw state for all IP entries
    mac_entry->walk_ip_list(delete_ip_entry_cb, mac_entry);
    return delete_mac_entry(mac_entry);
}

void
send_arp_probe (ep_ip_entry *ip_entry)
{
    void *mbuf;
    char *tx_hdr;
    eth_hdr_t *eth_hdr;
    arp_hdr_t *arp_hdr;
    arp_data_ipv4_t *arp_data;
    impl::p4_tx_info_t tx_info = { 0 };
    pds_obj_key_t vnic_key;
    vnic_entry *vnic;
    pds_obj_key_t subnet_key;
    subnet_entry *subnet;

    vnic_key = api::uuid_from_objid(ip_entry->vnic_obj_id());
    vnic = vnic_db()->find(&vnic_key);
    subnet_key = vnic->subnet();
    subnet = subnet_db()->find(&subnet_key);

    mbuf = learn_lif_alloc_mbuf();
    if (unlikely(mbuf == nullptr)) {
        PDS_TRACE_ERR("Failed to allocate pkt buffer for ARP probe, EP %s, "
                      "%s, %s", ip_entry->key2str().c_str(),
                      subnet->key2str().c_str(), vnic->key2str().c_str());
        return;
    }
    tx_hdr = learn_lif_mbuf_append_data(mbuf, ARP_PKT_ETH_FRAME_LEN +
                                        impl::arm_to_p4_hdr_sz());

    // fill Ethernet header, P4 adds encap header if required
    eth_hdr = (eth_hdr_t *)(tx_hdr + impl::arm_to_p4_hdr_sz());
    MAC_ADDR_COPY(eth_hdr->dmac, vnic->mac());
    MAC_ADDR_COPY(eth_hdr->smac, subnet->vr_mac());
    eth_hdr->eth_type = htons(ETH_TYPE_ARP);

    // fill ARP header
    arp_hdr = (arp_hdr_t *)(eth_hdr + 1);
    arp_hdr->htype = htons(ARP_HRD_TYPE_ETHER);
    arp_hdr->ptype = htons(ETH_TYPE_IPV4);
    arp_hdr->hlen = ETH_ADDR_LEN;
    arp_hdr->plen = IP4_ADDR8_LEN;
    arp_hdr->op = htons(ARP_OP_REQUEST);

    // fill ARP data
    arp_data = (arp_data_ipv4_t *)(arp_hdr + 1);
    MAC_ADDR_COPY(arp_data->smac, subnet->vr_mac());
    arp_data->sip = 0;
    memset(arp_data->tmac, 0, ETH_ADDR_LEN);
    arp_data->tip = htonl(ip_entry->key()->ip_addr.addr.v4_addr);

    // padding
    memset((arp_data + 1), 0, ARP_PKT_ETH_FRAME_LEN - ARP_PKT_LEN);

    // add ARM to P4 tx header and send the pkt
    tx_info.nh_type = impl::LEARN_NH_TYPE_VNIC;
    tx_info.vnic_key = vnic_key;
    impl::arm_to_p4_tx_hdr_fill(tx_hdr, &tx_info);
    learn_lif_send_pkt(mbuf);
}

static void
fill_learn_event (event_t *event, event_id_t learn_event, vnic_entry *vnic,
                  ep_ip_entry *ip_entry)
{
    core::learn_event_info_t *info = &event->learn;

    event->event_id = learn_event;
    info->subnet = vnic->subnet();
    info->ifindex = api::objid_from_uuid(vnic->host_if());
    MAC_ADDR_COPY(info->mac_addr, vnic->mac());
    if (ip_entry) {
        info->vpc = ip_entry->key()->vpc;
        info->ip_addr = ip_entry->key()->ip_addr;
    } else {
        info->vpc = { 0 };
        info->ip_addr = { 0 };
    }
}

void
fill_mac_event (event_t *event, event_id_t learn_event, ep_mac_entry *mac_entry)
{
    pds_obj_key_t vnic_key;
    vnic_entry *vnic;

    vnic_key = api::uuid_from_objid(mac_entry->vnic_obj_id());
    vnic = vnic_db()->find(&vnic_key);
    fill_learn_event(event, learn_event, vnic, nullptr);
}

void
fill_ip_event (event_t *event, event_id_t learn_event, ep_ip_entry *ip_entry)
{
    pds_obj_key_t vnic_key;
    vnic_entry *vnic;

    vnic_key = api::uuid_from_objid(ip_entry->vnic_obj_id());
    vnic = vnic_db()->find(&vnic_key);
    fill_learn_event(event, learn_event, vnic, ip_entry);
}

}    // namespace learn
