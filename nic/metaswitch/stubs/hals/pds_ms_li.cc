//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// PDSA Implementation of Metaswitch LI stub integration 
//---------------------------------------------------------------
 
#include "nic/metaswitch/stubs/hals/pds_ms_li.hpp"
#include "nic/metaswitch/stubs/common/pdsa_cookie.hpp"
#include "nic/metaswitch/stubs/common/pdsa_linux_util.hpp"
#include "nic/metaswitch/stubs/common/pdsa_state.hpp"
#include "nic/apollo/api/include/pds_tep.hpp"
#include "nic/apollo/api/include/pds_nexthop.hpp"
#include "nic/sdk/include/sdk/ip.hpp"
#include "nic/sdk/lib/logger/logger.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_li_vxlan_tnl.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_li_intf.hpp"

namespace pds_ms {

using pdsa_stub::Error;
using pdsa_stub::convert_ipaddr_ms_to_pdsa;
using pdsa_stub::vrfname_2_vrfid;

li_integ_subcomp_t* li_is () 
{
    static li_integ_subcomp_t g_li_is; 
    return &g_li_is;
}

//----------------------------------
// Physical port (Front panel port)
//---------------------------------
NBB_BYTE li_integ_subcomp_t::port_add_update(ATG_LIPI_PORT_ADD_UPDATE* port_add_upd_ips) {
    try {
        li_intf_t intf;
        intf.handle_add_upd_ips (port_add_upd_ips);
    } catch (Error& e) {
        SDK_TRACE_ERR ("Interface Add Update processing failed %s", e.what());
        port_add_upd_ips->return_code = ATG_UNSUCCESSFUL;
    }
    // Always return ATG_OK - fill the actual return code in the IPS
    return ATG_OK;
}

NBB_BYTE li_integ_subcomp_t::port_delete(NBB_ULONG port_ifindex) {
     try {
        li_intf_t intf;
        intf.handle_delete (port_ifindex);
    } catch (Error& e) {
        SDK_TRACE_ERR ("Interface Delete processing failed %s", e.what());
    }
    // Deletes are assummed to be synchronous and always successful in MS
    return ATG_OK;
}

//----------------------------------
// VRF
//---------------------------------
NBB_BYTE li_integ_subcomp_t::vrf_add_update(ATG_LIPI_VRF_ADD_UPDATE* vrf_add_upd_ips) {
    try {
        auto vrf_id = vrfname_2_vrfid(vrf_add_upd_ips->vrf_name, vrf_add_upd_ips->vrf_name_len);
        vrf_id = vrf_id;
    } catch (Error& e) {
        SDK_TRACE_ERR("VRF Add Update failed %s", e.what());
    }
    // Always return ATG_OK - fill the actual return code in the IPS
    return ATG_OK;
}

NBB_BYTE li_integ_subcomp_t::vrf_delete(const NBB_BYTE* vrf_name, NBB_ULONG vrf_len) {
    // Deletes are assummed to be synchronous and always successful in MS
    return ATG_OK;
}

//----------------------------------
// VXLAN tunnel (TEP)
//---------------------------------
NBB_BYTE li_integ_subcomp_t::vxlan_add_update(ATG_LIPI_VXLAN_ADD_UPDATE* vxlan_tnl_add_upd_ips) {
    try {
        li_vxlan_tnl vxtnl;
        vxtnl.handle_add_upd_ips (vxlan_tnl_add_upd_ips);
    } catch (Error& e) {
        SDK_TRACE_ERR ("Vxlan Tunnel Add Update processing failed %s", e.what());
        vxlan_tnl_add_upd_ips->return_code = ATG_UNSUCCESSFUL;
    }
    // Always return ATG_OK - fill the actual return code in the IPS
    return ATG_OK;
}
     
NBB_BYTE li_integ_subcomp_t::vxlan_delete(NBB_ULONG vxlan_tnl_ifindex) {
    try {
        li_vxlan_tnl vxtnl;
        vxtnl.handle_delete (vxlan_tnl_ifindex);
    } catch (Error& e) {
        SDK_TRACE_ERR ("Vxlan Tunnel Delete processing failed %s", e.what());
    }
    // Deletes are assummed to be synchronous and always successful in MS
    return ATG_OK;
}

//----------------------------------
// VXLAN port (TEP, VNI)
//---------------------------------
NBB_BYTE li_integ_subcomp_t::vxlan_port_add_update(ATG_LIPI_VXLAN_PORT_ADD_UPD* vxlan_port_add_upd_ips) {
    return ATG_OK;
}
NBB_BYTE li_integ_subcomp_t::vxlan_port_delete(NBB_ULONG vxlan_port_ifindex) {
    return ATG_OK;
}

//----------------------------------
// IRB (SVI for overlay BD)
//---------------------------------
NBB_BYTE li_integ_subcomp_t::irb_add_update(ATG_LIPI_IRB_ADD_UPDATE* irb_add_upd_ips) {
    return ATG_OK;
}
NBB_BYTE li_integ_subcomp_t::irb_delete(NBB_ULONG irb_ifindex) {
    return ATG_OK;
}

//--------------------------------------------------
// Software interface (Loopback and Dummy LIFs)
//-------------------------------------------------
NBB_BYTE li_integ_subcomp_t::softwif_add_update(ATG_LIPI_SOFTWIF_ADD_UPDATE* swif_add_upd_ips) {
    SDK_TRACE_INFO("Loopback interface create IfIndex 0x%lx Ifname %s SwType %d", 
                   swif_add_upd_ips->id.if_index, swif_add_upd_ips->id.if_name, 
                   swif_add_upd_ips->softwif_type);
    return ATG_OK;
}

NBB_BYTE li_integ_subcomp_t::softwif_delete(NBB_ULONG if_index, 
                                            const NBB_CHAR (&if_name) [ATG_LIPI_NAME_MAX_LEN], 
                                            NBB_ULONG softwif_type) {
    SDK_TRACE_INFO("Loopback interface delete IfIndex 0x%lx Ifname %s SwType %d", 
                   if_index, if_name, softwif_type);
    return ATG_OK;
}

//--------------------------------------------------
// Software interface IP (Loopback)
//-------------------------------------------------
NBB_BYTE li_integ_subcomp_t::softwif_addr_set(const NBB_CHAR *if_name,
                                              ATG_LIPI_L3_IP_ADDR *ip_addr,
                                              NBB_BYTE *vrf_name) {
    try {
        ip_addr_t ip; 
        convert_ipaddr_ms_to_pdsa (ip_addr->inet_addr, &ip);
        SDK_TRACE_INFO("Loopback interface IP address set request %s %s", 
                       if_name, ipaddr2str(&ip));
        if (ip.af == IP_AF_IPV6) {
            SDK_TRACE_INFO("Ignore IPv6 address");
            return ATG_OK;
        } 
        pdsa_stub::config_linux_loopback_ip(ip, ip_addr->prefix_len);
    } catch (Error& e) {
        SDK_TRACE_ERR ("Loopback interface IP address add failed %s", e.what());
        return ATG_UNSUCCESSFUL;
    }
    return ATG_OK;
}

NBB_BYTE li_integ_subcomp_t::softwif_addr_del(const NBB_CHAR *if_name,
                                              ATG_LIPI_L3_IP_ADDR *ip_addr,
                                              NBB_BYTE *vrf_name) {
    try {
        ip_addr_t ip; 
        convert_ipaddr_ms_to_pdsa (ip_addr->inet_addr, &ip);
        SDK_TRACE_INFO("Loopback interface IP address delete request %s %s", 
                       if_name, ipaddr2str(&ip));
        if (ip.af == IP_AF_IPV6) {
            SDK_TRACE_INFO("Ignore IPv6 address");
            return ATG_OK;
        } 
        pdsa_stub::config_linux_loopback_ip(ip, ip_addr->prefix_len, true);
    } catch (Error& e) {
        SDK_TRACE_ERR ("Loopback interface IP address delete failed %s", e.what());
        return ATG_UNSUCCESSFUL;
    }
    return ATG_OK;
}

} // End namespace