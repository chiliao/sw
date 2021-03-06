//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This module defines protobuf API for session object
///
//----------------------------------------------------------------------------

#ifndef __AGENT_SVC_SESSION_SVC_HPP__
#define __AGENT_SVC_SESSION_SVC_HPP__

#include "nic/apollo/agent/svc/specs.hpp"
#include "nic/apollo/agent/svc/session.hpp"

static inline void
pds_session_to_proto (void *ctxt)
{
}

static inline void
pds_ipv4_flow_to_proto (ftlite::internal::ipv4_entry_t *ipv4_entry,
                        void *ctxt)
{
    flow_get_t *fget = (flow_get_t *)ctxt;
    auto flow = fget->msg.add_flow();
    auto ipflowkey = flow->mutable_key()->mutable_ipflowkey();
    auto srcaddr = ipflowkey->mutable_srcip();
    auto dstaddr = ipflowkey->mutable_dstip();
    auto tcpudpinfo = ipflowkey->mutable_l4info()->mutable_tcpudpinfo();

    srcaddr->set_af(types::IP_AF_INET);
    srcaddr->set_v4addr(ipv4_entry->src);
    dstaddr->set_af(types::IP_AF_INET);
    dstaddr->set_v4addr(ipv4_entry->dst);
    ipflowkey->set_ipprotocol(ipv4_entry->proto);
    tcpudpinfo->set_srcport(ipv4_entry->sport);
    tcpudpinfo->set_dstport(ipv4_entry->dport);

    flow->set_vpc(ipv4_entry->vpc_id);
    flow->set_flowrole(ipv4_entry->flow_role);
    flow->set_sessionidx(ipv4_entry->session_index);
    flow->set_epoch(ipv4_entry->epoch);

    fget->count ++;
    if (fget->count == FLOW_GET_STREAM_COUNT) {
        fget->msg.set_apistatus(types::ApiStatus::API_STATUS_OK);
        fget->writer->Write(fget->msg);
        fget->msg.Clear();
        fget->count = 0;
    }
}

static inline void
pds_ipv6_flow_to_proto (ftlite::internal::ipv6_entry_t *ipv6_entry,
                        void *ctxt)
{
    flow_get_t *fget = (flow_get_t *)ctxt;
    auto flow = fget->msg.add_flow();
    auto ipflowkey = flow->mutable_key()->mutable_ipflowkey();
    auto srcaddr = ipflowkey->mutable_srcip();
    auto dstaddr = ipflowkey->mutable_dstip();
    auto tcpudpinfo = ipflowkey->mutable_l4info()->mutable_tcpudpinfo();

    srcaddr->set_af(types::IP_AF_INET6);
    srcaddr->set_v6addr(std::string((const char *)(ipv6_entry->src),
                        IP6_ADDR8_LEN));
    dstaddr->set_af(types::IP_AF_INET6);
    srcaddr->set_v6addr(std::string((const char *)(ipv6_entry->dst),
                        IP6_ADDR8_LEN));
    ipflowkey->set_ipprotocol(ipv6_entry->proto);
    tcpudpinfo->set_srcport(ipv6_entry->sport);
    tcpudpinfo->set_dstport(ipv6_entry->dport);

    flow->set_vpc(ipv6_entry->vpc_id);
    flow->set_flowrole(ipv6_entry->flow_role);
    flow->set_sessionidx(ipv6_entry->session_index);
    flow->set_epoch(ipv6_entry->epoch);

    fget->count ++;
    if (fget->count == FLOW_GET_STREAM_COUNT) {
        fget->msg.set_apistatus(types::ApiStatus::API_STATUS_OK);
        fget->writer->Write(fget->msg);
        fget->msg.Clear();
        fget->count = 0;
    }
}

static inline void
pds_flow_to_proto (ftlite::internal::ipv4_entry_t *ipv4_entry,
                   ftlite::internal::ipv6_entry_t *ipv6_entry,
                   void *ctxt)
{
    if (ipv4_entry) {
        pds_ipv4_flow_to_proto(ipv4_entry, ctxt);
    } else if (ipv6_entry) {
        pds_ipv6_flow_to_proto(ipv6_entry, ctxt);
    } else {
        flow_get_t *fget = (flow_get_t *)ctxt;
        if (fget->count) {
            fget->msg.set_apistatus(types::ApiStatus::API_STATUS_OK);
            fget->writer->Write(fget->msg);
            fget->msg.Clear();
            fget->count = 0;
        }
    }
}

static inline sdk_ret_t
pds_flow_proto_to_flow_key (pds_flow_key_t *key,
                            const pds::FlowFilter &flow_filter)
{
    key->lookup_id = flow_filter.vpc();
    key->proto = flow_filter.ipproto();
    key->sport = flow_filter.srcport();
    key->dport = flow_filter.dstport();

    if (flow_filter.srcaddr().af() == types::IPAF::IP_AF_INET) {
        key->src_ip.af = IP_AF_IPV4;
        key->src_ip.addr.v4_addr = flow_filter.srcaddr().v4addr();
        key->dst_ip.af = IP_AF_IPV4;
        key->dst_ip.addr.v4_addr = flow_filter.dstaddr().v4addr();
    } else {
        key->src_ip.af = IP_AF_IPV6;
        memcpy(key->src_ip.addr.v6_addr.addr8,
               flow_filter.srcaddr().v6addr().c_str(),
               IP6_ADDR8_LEN);
        key->dst_ip.af = IP_AF_IPV6;
        memcpy(key->dst_ip.addr.v6_addr.addr8,
               flow_filter.dstaddr().v6addr().c_str(),
               IP6_ADDR8_LEN);
    }
    return SDK_RET_OK;
}

#endif    //__AGENT_SVC_SESSION_SVC_HPP__
