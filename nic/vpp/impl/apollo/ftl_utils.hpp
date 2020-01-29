//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//

#ifndef __VPP_IMPL_APOLLO_FTL_UTILS_H__
#define __VPP_IMPL_APOLLO_FTL_UTILS_H__

void
ftlv4_set_nexthop(ipv4_flow_hash_entry_t *entry,
                  uint32_t nexthop,
                  uint32_t nexthop_type,
                  uint8_t nexthop_valid)
{
    return;
}

void
ftlv6_set_nexthop(flow_hash_entry_t *entry,
                  uint32_t nexthop,
                  uint32_t nexthop_type,
                  uint8_t nexthop_valid)
{
    return;
}

void
ftlv4_set_lookup_id(ipv4_flow_hash_entry_t *entry, 
                    uint16_t lookup_id)
{
    entry->set_key_metadata_lkp_id(lookup_id);
}

void
ftlv6_set_lookup_id(flow_hash_entry_t *entry, 
                    uint16_t lookup_id)
{
    entry->set_key_metadata_lkp_id(lookup_id);
}

uint16_t
ftlv4_get_lookup_id(ipv4_flow_hash_entry_t *entry)
{
    return entry->get_key_metadata_lkp_id();
}

uint16_t
ftlv6_get_lookup_id(flow_hash_entry_t *entry)
{
    return entry->get_key_metadata_lkp_id();
}

#endif  // __VPP_IMPL_APOLLO_FTL_UTILS_H__
