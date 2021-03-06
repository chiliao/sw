// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// Purpose: Common helper APIs for metaswitch stub programming 

#include "nic/metaswitch/stubs/mgmt/pds_ms_mgmt_utils.hpp"
#include "include/sdk/base.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_error.hpp"
#include "nic/metaswitch/stubs/mgmt/pds_ms_mgmt_state.hpp"
#include "gen/proto/epoch.pb.h"

#define SHARED_DATA_TYPE CSS_LOCAL
using namespace types;

void
ip_addr_to_spec (const ip_addr_t *ip_addr,
                 types::IPAddress *ip_addr_spec)
{
    if (ip_addr->af == IP_AF_IPV4) {
        ip_addr_spec->set_af(types::IP_AF_INET);
        ip_addr_spec->set_v4addr(ip_addr->addr.v4_addr);
    } else if (ip_addr->af == IP_AF_IPV6) {
        ip_addr_spec->set_af(types::IP_AF_INET6);
        ip_addr_spec->set_v6addr(ip_addr->addr.v6_addr.addr8, IP6_ADDR8_LEN);
    }

    return;
}

bool
ip_addr_spec_to_ip_addr (const types::IPAddress& in_ipaddr,
                         ip_addr_t *out_ipaddr)
{
    memset(out_ipaddr, 0, sizeof(ip_addr_t));
    if (in_ipaddr.af() == types::IP_AF_INET) {
        out_ipaddr->af = IP_AF_IPV4;
        out_ipaddr->addr.v4_addr = in_ipaddr.v4addr();
    } else if (in_ipaddr.af() == types::IP_AF_INET6) {
        out_ipaddr->af = IP_AF_IPV6;
        memcpy(out_ipaddr->addr.v6_addr.addr8,
               in_ipaddr.v6addr().c_str(),
               IP6_ADDR8_LEN);
    } else {
        return false;
    }

    return true;
}

NBB_VOID
pds_ms_convert_ip_addr_to_amb_ip_addr (const ip_addr_t&   pds_ms_ip_addr, 
                                       NBB_LONG      *type, 
                                       NBB_ULONG     *len, 
                                       NBB_BYTE      *amb_ip_addr,
                                       bool          is_zero_ip_valid)
{
    switch (pds_ms_ip_addr.af)
    {
        case IP_AF_IPV4:
            *type = AMB_INETWK_ADDR_TYPE_IPV4;
            *len = AMB_MAX_IPV4_ADDR_LEN;
            break;

        case IP_AF_IPV6:
            *type = AMB_INETWK_ADDR_TYPE_IPV6;
            *len = AMB_MAX_IPV6_ADDR_LEN;
            break;

        default:
            *type = *len = 0;
            return;
    }

    if (!is_zero_ip_valid && ip_addr_is_zero(&pds_ms_ip_addr)) {
        *len = 0;
        return;
    }

    if (pds_ms_ip_addr.af == IP_AF_IPV4) {
        NBB_PUT_LONG (amb_ip_addr, pds_ms_ip_addr.addr.v4_addr);
    } else {
        NBB_MEMCPY (amb_ip_addr, &pds_ms_ip_addr.addr, *len);
    }
    return;
}

// Converts from byte-order to network-order IP Address
NBB_VOID
pds_ms_convert_amb_ip_addr_to_ip_addr (NBB_BYTE      *amb_ip_addr,
                                     NBB_LONG      type,
                                     NBB_ULONG     len,
                                     ip_addr_t     *proto_ip_addr)
{
    switch (type)
    {
        case AMB_INETWK_ADDR_TYPE_IPV4:
            proto_ip_addr->af = IP_AF_IPV4;
            break;

        case AMB_INETWK_ADDR_TYPE_IPV6:
            proto_ip_addr->af = IP_AF_IPV6;
            break;

        case AMB_INETWK_ADDR_TYPE_OTHER:
            return;

        default:
            assert(0);
    }
    if (proto_ip_addr->af == IP_AF_IPV4) {
        NBB_GET_LONG (proto_ip_addr->addr.v4_addr, amb_ip_addr);
    } else {
        NBB_MEMCPY (&proto_ip_addr->addr, amb_ip_addr, len);
    }
    return;
}

NBB_LONG
pds_ms_nbb_get_long(NBB_BYTE *byteVal)
{
    NBB_LONG val;
    NBB_GET_LONG(val, byteVal);
    return val;
}

NBB_VOID
pds_ms_set_string_in_byte_array_with_len(NBB_BYTE *field, NBB_ULONG *len, string in_str)
{
    *len = in_str.length();
    NBB_MEMCPY(field, in_str.c_str(), *len);
}

NBB_VOID
pds_ms_set_string_in_byte_array_with_len_oid(NBB_ULONG *oid, string in_str, NBB_LONG setKeyOidIdx, NBB_LONG setKeyOidLenIdx)
{
    oid[setKeyOidLenIdx] = (NBB_ULONG)in_str.length();

    auto str = (const unsigned char*) in_str.c_str();
    for (NBB_ULONG i=0; i<in_str.length(); i++) {
        oid[setKeyOidIdx + i] = (NBB_ULONG)str[i];
    }
}

NBB_VOID
pds_ms_validate_byte_array (string in_str, string in_msg, string in_field,
                            int min_len, int max_len)
{
    if (in_str.length() == 0) {
        // null string. no need to validate. if this is required field, let
        // metaswitch detect it
        return;
    }

    if (in_str.length() < min_len || in_str.length() > max_len) {
        throw pds_ms::Error (in_msg.append(".").append(in_field).
                             append(" length should be between ").
                             append(std::to_string(min_len)).append(" and ").
                             append(std::to_string(max_len)),
                             SDK_RET_INVALID_ARG);

    }
}

// API to print "bytes" field of protos
NBB_VOID
pds_ms_print_byte_array (string in_str, string in_msg, string in_field)
{
    // print only if the field has some input
    if (in_str.length())
    {
        int len = 3*in_str.length();
        int buf_idx = 0;
        char buffer [len];
        auto str = (const unsigned char*) in_str.c_str();
        for(int j = 0; j < in_str.length(); j++) {
            buf_idx += snprintf(buffer+buf_idx, len-buf_idx, "%02X:", str[j]);
        }
        PDS_TRACE_DEBUG ("%s.%s: \"%s\"",
                          in_msg.c_str(), in_field.c_str(), buffer);
    }
}

string
pds_ms_get_string_in_byte_array_with_len(NBB_BYTE *in_str, NBB_ULONG len)
{
    std::string ret((const char*)in_str, len);
    return ret;
}

NBB_VOID
pds_ms_get_string_in_byte_array_with_len_oid(NBB_ULONG *oid, string in_str, NBB_LONG getKeyOidIdx, NBB_LONG getKeyOidLenIdx)
{
    oid[getKeyOidLenIdx] = (NBB_ULONG)in_str.length();
    auto str = (const unsigned char* ) in_str.c_str();
    for (NBB_ULONG i=0; i<in_str.length(); i++) {
        oid[getKeyOidIdx + i] = (NBB_ULONG)str[i];
    }
}

NBB_VOID
pds_ms_get_uuid(pds_obj_key_t *uuid, const string& in_str)
{
    NBB_MEMCPY(uuid->id, in_str.c_str(), in_str.length());
}

NBB_VOID
pds_ms_set_string_in_byte_array(NBB_BYTE *field, string in_str)
{
    NBB_MEMCPY(field, in_str.c_str(), in_str.length());
}

NBB_VOID
pds_ms_set_string_in_byte_array_oid(NBB_ULONG *oid, string in_str, NBB_LONG setKeyOidIdx)
{
    auto str = (const unsigned char*) in_str.c_str();
    for (NBB_ULONG i=0; i<in_str.length(); i++) {
        oid[setKeyOidIdx + i] = (NBB_ULONG)str[i];
    }
}

string
pds_ms_get_string_in_byte_array(NBB_BYTE *val, NBB_ULONG len)
{
    std::string ret(val, val + len);
    return ret;
}

NBB_VOID
pds_ms_get_string_in_byte_array_oid(NBB_ULONG *oid, string in_str, NBB_LONG getKeyOidIdx)
{
    auto str = (const unsigned char*) in_str.c_str();
    for (NBB_ULONG i=0; i<in_str.length(); i++) {
        oid[getKeyOidIdx + i] = (NBB_ULONG)str[i];
    }
}

sdk_ret_t
pds_ms_api_to_sdk_ret (ApiStatus api_err)
{
    switch (api_err) {
    case API_STATUS_OK: {return SDK_RET_OK;}
    case API_STATUS_ERR: {return SDK_RET_ERR;}
    case API_STATUS_INVALID_ARG: {return SDK_RET_INVALID_ARG;}
    case API_STATUS_EXISTS_ALREADY: {return SDK_RET_ENTRY_EXISTS;}
    case API_STATUS_OUT_OF_MEM: {return SDK_RET_OOM;}
    case API_STATUS_NOT_FOUND: {return SDK_RET_ENTRY_NOT_FOUND;}
    case API_STATUS_OUT_OF_RESOURCE: {return SDK_RET_NO_RESOURCE;}
    case API_STATUS_OPERATION_NOT_ALLOWED: {return SDK_RET_INVALID_OP;}
    default: {return SDK_RET_ERR;} // catch undefined errors
    }
    return SDK_RET_ERR;
}

ApiStatus
pds_ms_sdk_ret_to_api_status (sdk_ret_t sdk_ret)
{
    switch (sdk_ret) {
    case SDK_RET_OK: {return API_STATUS_OK;}
    case SDK_RET_ERR: {return API_STATUS_ERR;}
    case SDK_RET_INVALID_ARG: {return API_STATUS_INVALID_ARG;}
    case SDK_RET_ENTRY_EXISTS: {return API_STATUS_EXISTS_ALREADY;}
    case SDK_RET_OOM: {return API_STATUS_OUT_OF_MEM;}
    case SDK_RET_ENTRY_NOT_FOUND: {return API_STATUS_NOT_FOUND;}
    case SDK_RET_NO_RESOURCE: {return API_STATUS_OUT_OF_RESOURCE;}
    case SDK_RET_INVALID_OP: {return API_STATUS_OPERATION_NOT_ALLOWED;}
    default: {return API_STATUS_ERR;} // catch undefined errors
    }
}

const char *
pds_ms_api_ret_str (ApiStatus api_err)
{
    switch (api_err) {
    case API_STATUS_OK: {return "API_STATUS_OK";}
    case API_STATUS_ERR: {return "API_STATUS_ERR";}
    case API_STATUS_INVALID_ARG: {return "API_STATUS_INVALID_ARG";}
    case API_STATUS_EXISTS_ALREADY: {return "API_STATUS_EXISTS_ALREADY";}
    case API_STATUS_OUT_OF_MEM: {return "API_STATUS_OUT_OF_MEM";}
    case API_STATUS_NOT_FOUND: {return "API_STATUS_NOT_FOUND";}
    case API_STATUS_OUT_OF_RESOURCE: {return "API_STATUS_OUT_OF_RESOURCE";}
    case API_STATUS_OPERATION_NOT_ALLOWED: {return "API_STATUS_OPERATION_NOT_ALLOWED";}
    default: {return "Unknown error";} // catch undefined errors
    }
    return "Unknown error";
}

namespace pds_ms {

ApiStatus fill_epoch_get_response(const EpochGetRequest *req,
                                  EpochGetResponse *resp)
{
    auto mgmt_ctxt = mgmt_state_t::thread_context();
    resp->set_apistatus(API_STATUS_OK);
    resp->set_epoch(mgmt_ctxt.state()->epoch());
    return API_STATUS_OK;
}

} // namespace pds_ms
