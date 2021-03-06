//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// LI VXLAN Tunnel HAL integration
//---------------------------------------------------------------

#ifndef __PDS_MS_HALS_ROUTE_HPP__
#define __PDS_MS_HALS_ROUTE_HPP__

#include "nic/metaswitch/stubs/common/pds_ms_cookie.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_util.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_defs.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_state.hpp"
#include "nic/apollo/api/include/pds_nexthop.hpp"
#include "nic/sdk/include/sdk/eth.hpp"
#include <nbase.h>
extern "C"
{
#include <a0build.h>
#include <a0glob.h>
#include <a0spec.h>
#include <o0mac.h>
#include <ropi.h>
}

namespace pds_ms {

using pds_ms::ms_ifindex_t;
using pds_ms::mac_addr_wr_t;
using pds_ms::cookie_t;
using pds_ms::pds_batch_ctxt_guard_t;

class hals_route_t {
public:
    NBB_BYTE handle_add_upd_ips(ATG_ROPI_UPDATE_ROUTE* add_upd_route_ips);
    void handle_delete(ATG_ROPI_ROUTE_ID route_id);

private:
    struct ips_info_t {
        uint32_t vrf_id;
        ip_prefix_t pfx;
        uint32_t ecmp_id; // Overlay ECMP ID DP correlator
        uint32_t pathset_id; // Overlay Pathset ID
    };

private:
    std::unique_ptr<cookie_t> cookie_uptr_;
    ips_info_t  ips_info_;
    bool op_delete_ = false;
    bool op_create_ = false;
    pds_route_t route_ = {0};
    pds_route_t prev_route_ = {0};
    pds_obj_key_t rttbl_key_ = {0};

private:
    pds_batch_ctxt_guard_t make_batch_pds_spec_(state_t* state,
                                                const pds_obj_key_t&
                                                rttable_key);
    void populate_route_id(ATG_ROPI_ROUTE_ID* route_id);
    bool parse_ips_info_(ATG_ROPI_UPDATE_ROUTE* route_add_upd);
    bool make_pds_rttable_spec_(state_t* state,
                                pds_route_table_spec_t &rttbl,
                                const pds_obj_key_t& rttable_key);
    pds_obj_key_t make_pds_rttable_key_(state_t*);
    void overlay_route_del_(void);
    sdk_ret_t underlay_route_add_upd_(void);
    sdk_ret_t underlay_route_del_(void);
};

} // End namespace
#endif
