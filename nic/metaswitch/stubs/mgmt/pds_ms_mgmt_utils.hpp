// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// Purpose: Common helper APIs header file for metaswitch stub programming 

#ifndef __PDS_MS_MGMT_UTILS_HPP__
#define __PDS_MS_MGMT_UTILS_HPP__
#include <nbase.h>
extern "C" {
#include <a0spec.h>
#include <o0mac.h>
#include <a0cust.h>
#include <a0glob.h>
#include <a0mib.h>
#include <ambips.h>
#include <a0stubs.h>
#include <a0cpif.h>
#include "smsiincl.h"
#include "smsincl.h"
}
#include "li_mgmt_if.h"
#include "lim_mgmt_if.h"
#include "psm_prod.h"
#include "ftm_mgmt_if.h"
#include "nrm_mgmt_if.h"
#include "nrm_prod.h"
#include "psm_mgmt_if.h"
#include "include/sdk/ip.hpp"
#include "nic/metaswitch/stubs/mgmt/pds_ms_config.hpp"
#include "nic/metaswitch/stubs/mgmt/pds_ms_config.hpp"
#include "nic/metaswitch/stubs/mgmt/pds_ms_ctm.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_util.hpp"
#include "gen/proto/types.pb.h"
#include "gen/proto/internal.pb.h"
#include "gen/proto/internal_bgp.pb.h"
#include "gen/proto/internal_evpn.pb.h"
#include "gen/proto/internal_cp_route.pb.h"
#include "gen/proto/epoch.pb.h"
#include "nic/metaswitch/stubs/mgmt/gen/mgmt/pds_ms_internal_bgp_utils_gen.hpp"
#include "nic/metaswitch/stubs/mgmt/gen/mgmt/pds_ms_internal_evpn_utils_gen.hpp"
#include "gen/proto/cp_test.pb.h"
#include "nic/metaswitch/stubs/mgmt/gen/mgmt/pds_ms_internal_utils_gen.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_tbl_idx.hpp"
#include "nic/apollo/api/include/pds.hpp"

#define PDS_MS_CTM_GRPC_CORRELATOR 0x101
#define PDS_MS_CTM_STUB_INIT_CORRELATOR 0x42
using namespace std;

const char* pds_ms_api_ret_str (types::ApiStatus api_err);
sdk_ret_t pds_ms_api_to_sdk_ret(types::ApiStatus api_err);
types::ApiStatus pds_ms_sdk_ret_to_api_status(sdk_ret_t sdk_ret);

void ip_addr_to_spec(const ip_addr_t *ip_addr,
                     types::IPAddress *ip_addr_spec);

NBB_VOID pds_ms_convert_amb_ip_addr_to_ip_addr (NBB_BYTE      *amb_ip_addr,
                                                NBB_LONG      type,
                                                NBB_ULONG     len,
                                                ip_addr_t     *pds_ms_ip_addr);
NBB_VOID  pds_ms_convert_ip_addr_to_amb_ip_addr (const ip_addr_t& pds_ms_ip_addr,
                                                 NBB_LONG      *type,
                                                 NBB_ULONG     *len,
                                                 NBB_BYTE      *ambip_addr,
                                                 bool          is_zero_ip_valid);

bool ip_addr_spec_to_ip_addr (const types::IPAddress& in_ipaddr,
                              ip_addr_t *out_ipaddr);

NBB_LONG pds_ms_nbb_get_long(NBB_BYTE *byteVal);

static inline NBB_LONG pds_ms_get_amb_bool (NBB_LONG amb_bool) {
    // convert from amb_bool to bool
    return (amb_bool == AMB_TRUE) ? 1 : 0;
}
static inline NBB_VOID pds_ms_set_amb_bool(NBB_LONG *amb_bool,
                                           NBB_LONG bool_val) {
    // convert bool to amb_bool
    *amb_bool = bool_val ? AMB_TRUE : AMB_FALSE;
}

static inline NBB_VOID pds_ms_set_evpn_cfg(NBB_LONG *autoCfg,
                                           pds_ms::EvpnCfg val) {
    if (val == pds_ms::EVPN_CFG_MANUAL) {
        *autoCfg = AMB_EVPN_CONFIGURED;
    } else if (val == pds_ms::EVPN_CFG_AUTO) {
        *autoCfg = AMB_EVPN_AUTO;
    } else {
        *autoCfg = AMB_EVPN_NONE;
    }
}
static inline pds_ms::EvpnCfg pds_ms_get_evpn_cfg(NBB_LONG autoCfg) {
    if (autoCfg == AMB_EVPN_CONFIGURED) { return pds_ms::EVPN_CFG_MANUAL; }
    if (autoCfg == AMB_EVPN_AUTO) { return pds_ms::EVPN_CFG_AUTO;}
    return pds_ms::EVPN_CFG_NONE;
}
static inline NBB_VOID pds_ms_set_evpn_rttype(NBB_LONG *rttype,
                                              pds_ms::EvpnRtType val) {
    if (val == pds_ms::EVPN_RT_NONE) {
        // map rt_none (0) to MS rt_none (4)
        *rttype = AMB_EVPN_RT_NONE;
    } else {
        // remaining values of the rt types are one to one match
        *rttype = (NBB_LONG)val;
    }
}
static inline pds_ms::EvpnRtType pds_ms_get_evpn_rttype(NBB_LONG rttype) {
    if (rttype == AMB_EVPN_RT_NONE) {return pds_ms::EVPN_RT_NONE;}
    return (pds_ms::EvpnRtType)rttype;
}

NBB_VOID
pds_ms_get_uuid(pds_obj_key_t *out_uuid, const string& in_str);

NBB_VOID
pds_ms_set_string_in_byte_array_with_len(NBB_BYTE *field,
                                       NBB_ULONG *len,
                                       string in_str);

NBB_VOID pds_ms_validate_byte_array (string in_str, string in_msg,
                                     string in_field, int min_len, int max_len);
NBB_VOID  pds_ms_print_byte_array (string in_str, string in_msg,
                                   string in_field);

NBB_VOID
pds_ms_set_string_in_byte_array_with_len_oid(NBB_ULONG *oid,
                                           string in_str,
                                           NBB_LONG setKeyOidIdx,
                                           NBB_LONG setKeyOidLenIdx);

string
pds_ms_get_string_in_byte_array_with_len(NBB_BYTE *in_str,
                                       NBB_ULONG len);

NBB_VOID
pds_ms_get_string_in_byte_array_with_len_oid(NBB_ULONG *oid,
                                           string in_str,
                                           NBB_LONG getKeyOidIdx,
                                           NBB_LONG getKeyOidLenIdx);

NBB_VOID
pds_ms_set_string_in_byte_array(NBB_BYTE *field,
                              string in_str);

NBB_VOID
pds_ms_set_string_in_byte_array_oid(NBB_ULONG *oid,
                                  string in_str,
                                  NBB_LONG setKeyOidIdx);

string
pds_ms_get_string_in_byte_array(NBB_BYTE *val,
                              NBB_ULONG len);

NBB_VOID
pds_ms_get_string_in_byte_array_oid(NBB_ULONG *oid,
                                  string in_str,
                                  NBB_LONG getKeyOidIdx);
namespace pds_ms {
types::ApiStatus bgp_clear_route_action_func (const pds::BGPClearRouteRequest *req,
                                              pds::BGPClearRouteResponse  *resp);
NBB_VOID bgp_rm_ent_get_fill_func (BGPSpec &req,
                                   NBB_ULONG*   oid);

NBB_VOID bgp_rm_ent_set_fill_func (BGPSpec        &req, 
                                   AMB_GEN_IPS    *mib_msg, 
                                   AMB_BGP_RM_ENT *v_amb_bgp_rm_ent, 
                                   NBB_LONG       row_status);
NBB_VOID bgp_rm_ent_pre_set (BGPSpec &req, NBB_LONG row_status,
                             NBB_ULONG correlator, NBB_VOID* kh, bool op_update);

NBB_VOID bgp_peer_pre_set(BGPPeerSpec &req, NBB_LONG row_status,
                          NBB_ULONG correlator, NBB_VOID* kh, bool op_update=false);

NBB_VOID
bgp_peer_afi_safi_pre_set(BGPPeerAfSpec &req, NBB_LONG row_status,
                          NBB_ULONG correlator, NBB_VOID* kh, bool op_update=false);

NBB_VOID bgp_rm_ent_pre_get(BGPSpec &req, BGPGetResponse* resp, NBB_VOID* kh);
NBB_VOID bgp_peer_pre_get(BGPPeerSpec &req, BGPPeerGetResponse* resp, NBB_VOID* kh);
NBB_VOID bgp_peer_afi_safi_pre_get(BGPPeerAfSpec &req,
                                   BGPPeerAfGetResponse* resp,
                                   NBB_VOID* kh);
bool bgp_rm_ent_pre_fill_get (amb_bgp_rm_ent *data);
bool bgp_peer_afi_safi_pre_fill_get (amb_bgp_peer_afi_safi *data);
NBB_VOID bgp_peer_status_get_fill_func (BGPPeerSpec& req,
                                        NBB_ULONG*          oid);

NBB_VOID bgp_peer_get_fill_func (BGPPeerSpec&   req,
                                 NBB_ULONG*           oid);

NBB_VOID bgp_peer_set_fill_func (BGPPeerSpec& req,
                                 AMB_GEN_IPS  *mib_msg,
                                 AMB_BGP_PEER *v_amb_bgp_peer,
                                NBB_LONG      row_status);

NBB_VOID
bgp_peer_af_status_get_fill_func (BGPPeerAfSpec &req,
                           NBB_ULONG*    oid);

NBB_VOID
bgp_peer_af_get_fill_func (BGPPeerAfSpec &req,
                           NBB_ULONG*    oid);

NBB_VOID bgp_peer_af_set_fill_func (BGPPeerAfSpec&        req,
                                    AMB_GEN_IPS           *mib_msg,
                                    AMB_BGP_PEER_AFI_SAFI *v_amb_bgp_peer_af,
                                    NBB_LONG               row_status);
NBB_VOID evpn_evi_pre_set (EvpnEviSpec  &req,
                           NBB_LONG     row_status,
                           NBB_ULONG    test_correlator,
                           NBB_VOID*    kh,
                           bool         op_update=false);
NBB_VOID evpn_evi_pre_get (EvpnEviSpec &req, EvpnEviGetResponse* resp,
                           NBB_VOID* kh);
NBB_VOID evpn_evi_rt_pre_set (EvpnEviRtSpec  &req,
                              NBB_LONG       row_status,
                              NBB_ULONG      test_correlator,
                              NBB_VOID*      kh,
                              bool           op_update=false);
NBB_VOID evpn_evi_rt_pre_get (EvpnEviRtSpec &req, EvpnEviRtGetResponse* resp,
                              NBB_VOID* kh);
NBB_VOID evpn_evi_post_getall (EvpnEviGetResponse* resp);
NBB_VOID evpn_evi_post_get (EvpnEviSpec &req, EvpnEviGetResponse* resp,
                            NBB_VOID* kh);
NBB_VOID evpn_evi_rt_post_getall (EvpnEviRtGetResponse* resp);
NBB_VOID evpn_evi_rt_post_get (EvpnEviRtSpec &req, EvpnEviRtGetResponse* resp,
                               NBB_VOID* kh);
NBB_VOID evpn_ip_vrf_pre_set (EvpnIpVrfSpec  &req,
                              NBB_LONG       row_status,
                              NBB_ULONG      test_correlator,
                              NBB_VOID*      kh,
                              bool           op_update=false);
NBB_VOID evpn_ip_vrf_post_getall (EvpnIpVrfGetResponse *resp);
NBB_VOID evpn_ip_vrf_post_get (EvpnIpVrfSpec &req, EvpnIpVrfGetResponse *resp,
                               NBB_VOID* kh);
NBB_VOID evpn_ip_vrf_pre_get (EvpnIpVrfSpec &req, EvpnIpVrfGetResponse *resp,
                              NBB_VOID* kh);
NBB_VOID evpn_ip_vrf_rt_pre_set (EvpnIpVrfRtSpec  &req,
                                 NBB_LONG       row_status,
                                 NBB_ULONG      test_correlator,
                                 NBB_VOID*      kh,
                                 bool           op_update=false);
NBB_VOID evpn_ip_vrf_rt_pre_get (EvpnIpVrfRtSpec &req, EvpnIpVrfRtGetResponse *resp,
                                 NBB_VOID *kh);
NBB_VOID evpn_ip_vrf_rt_post_getall (EvpnIpVrfRtGetResponse *resp);
NBB_VOID evpn_ip_vrf_rt_post_get (EvpnIpVrfRtSpec &req, EvpnIpVrfRtGetResponse *resp,
                                  NBB_VOID *kh);
NBB_VOID evpn_evi_get_fill_func (EvpnEviSpec&    req,
                                 NBB_ULONG*       oid);
NBB_VOID evpn_evi_status_get_fill_func (EvpnEviSpec& req,
                                        NBB_ULONG*     oid);
NBB_VOID evpn_evi_set_fill_func (EvpnEviSpec&    req,
                                 AMB_GEN_IPS     *mib_msg,
                                 AMB_EVPN_EVI    *data,
                                 NBB_LONG        row_status);
NBB_VOID  evpn_ip_vrf_get_fill_func (EvpnIpVrfSpec&   req,
                                     NBB_ULONG*       oid);
NBB_VOID  evpn_ip_vrf_status_get_fill_func (EvpnIpVrfSpec&   req,
                                            NBB_ULONG*       oid);
NBB_VOID  evpn_ip_vrf_set_fill_func (EvpnIpVrfSpec&   req,
                                     AMB_GEN_IPS      *mib_msg,
                                     AMB_EVPN_IP_VRF  *data,
                                     NBB_LONG         row_status);
NBB_VOID evpn_evi_rt_set_fill_func (EvpnEviRtSpec&   req,
                                    AMB_GEN_IPS      *mib_msg,
                                    AMB_EVPN_EVI_RT  *data,
                                    NBB_LONG         row_status);
NBB_VOID evpn_evi_rt_get_fill_func (EvpnEviRtSpec&   req,
                                    NBB_ULONG*         oid);
NBB_VOID evpn_ip_vrf_rt_get_fill_func (EvpnIpVrfRtSpec&      req,
                                   NBB_ULONG*         oid);
NBB_VOID evpn_ip_vrf_rt_set_fill_func (EvpnIpVrfRtSpec&      req,
                                       AMB_GEN_IPS           *mib_msg,
                                       AMB_EVPN_IP_VRF_RT    *data,
                                       NBB_LONG         row_status);
NBB_VOID evpn_mac_ip_get_fill_func (EvpnMacIpStatus& req, NBB_ULONG *oid);
NBB_VOID evpn_ip_vrf_fill_name_oid (EvpnIpVrfSpec& req, NBB_ULONG *oid);
NBB_VOID evpn_ip_vrf_fill_name_field (EvpnIpVrfSpec& req, AMB_GEN_IPS *mib_msg);
NBB_VOID evpn_ip_vrf_rt_fill_name_field (EvpnIpVrfRtSpec& req,
                                         AMB_GEN_IPS *mib_msg);
NBB_VOID evpn_ip_vrf_rt_fill_name_oid (EvpnIpVrfRtSpec& req, NBB_ULONG *oid);
NBB_VOID evpn_ip_vrf_get_name_field (EvpnIpVrfSpec* req, AMB_EVPN_IP_VRF *data);
NBB_VOID evpn_ip_vrf_rt_get_name_field (EvpnIpVrfRtSpec* req,
                                        AMB_EVPN_IP_VRF_RT *data);
NBB_VOID rtm_strt_set_fill_func (CPStaticRouteSpec&      req,
                                 AMB_GEN_IPS             *mib_msg,
                                 AMB_CIPR_RTM_STATIC_RT  *data,
                                 NBB_LONG                row_status);
NBB_VOID rtm_strt_get_fill_func (CPStaticRouteSpec& req, NBB_ULONG *oid);
NBB_VOID lim_intf_addr_fill_func (LimInterfaceAddrSpec&  req,
                                  AMB_GEN_IPS           *mib_msg,
                                  AMB_LIM_L3_IF_ADDR    *data,
                                  NBB_LONG              row_status);
NBB_VOID lim_sw_intf_fill_func (LimInterfaceSpec&    req,
                                AMB_GEN_IPS         *mib_msg,
                                AMB_LIM_SOFTWARE_IF *data,
                                NBB_LONG            row_status);
types::ApiStatus fill_epoch_get_response(const EpochGetRequest *req,
                                         EpochGetResponse *resp);
types::ApiStatus l2f_test_local_mac_ip_add (const CPL2fTestCreateSpec   *req,
                                            CPL2fTestResponse *resp);
types::ApiStatus l2f_test_local_mac_ip_del (const CPL2fTestDeleteSpec   *req,
                                            CPL2fTestResponse *resp);
types::ApiStatus ip_track_add (const CPIPTrackTestCreateSpec   *req,
                               CPIPTrackTestResponse *resp);
types::ApiStatus ip_track_del (const CPIPTrackTestDeleteSpec   *req,
                               CPIPTrackTestResponse *resp);
NBB_VOID lim_l3_if_addr_pre_set(LimInterfaceAddrSpec &req,
                                NBB_LONG row_status,
                                NBB_ULONG correlator,
                                NBB_VOID* kh,
                                bool      op_update=false);
NBB_VOID cp_route_pre_set(CPStaticRouteSpec &req, NBB_LONG row_status,
                          NBB_ULONG correlator, NBB_VOID* kh, bool op_update=false);
NBB_VOID update_bgp_route_map_table (NBB_ULONG correlator);
NBB_VOID pds_ms_rtm_redis_connected (pds_ms::pds_ms_config_t *conf);
NBB_VOID pds_ms_li_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_l2f_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_smi_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_sck_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_hals_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_nar_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_ft_stub_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_lim_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_ftm_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_nrm_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_psm_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_rtm_ft_stub_join(pds_ms_config_t *conf, int entity_index);
NBB_VOID pds_ms_rtm_create (pds_ms_config_t *conf, int entity_index,
                            bool is_default);
NBB_VOID pds_ms_rtm_ftm_join (pds_ms_config_t *conf, int entity_index);
NBB_VOID pds_ms_bgp_create (pds_ms_config_t *conf);
NBB_VOID pds_ms_evpn_create (pds_ms_config_t *conf);
void pds_ms_evpn_rtm_join(pds_ms_config_t *conf, int rtm_entity_index);
void populate_lim_addr_spec (ip_prefix_t           *ip_prefix,
                             LimInterfaceAddrSpec& req,
                             uint32_t              if_type,
                             uint32_t              if_id);
} // end namespace pds_ms

namespace pds_ms_test {
NBB_VOID pds_ms_test_row_update_l2f_mac_ip_cfg (ip_addr_t ip_addr,
                                              NBB_ULONG host_ifindex);
} // end namespace pds_ms_test
#endif /*__PDS_MS_MGMT_UTILS_HPP__*/
