// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#ifndef __LINKMGR_UTILS_HPP__
#define __LINKMGR_UTILS_HPP__

#include "gen/proto/types.pb.h"
#include "nic/sdk/include/sdk/ip.hpp"
#include "gen/proto/port.pb.h"
#include "nic/sdk/include/sdk/types.hpp"

using types::ApiStatus;

namespace hal {

#define NUM_DASHES 20
#define __FILENAME__ (strrchr(__FILE__, '/') ? strrchr(__FILE__, '/') + 1 : __FILE__)

hal_ret_t ip_addr_spec_to_ip_addr(ip_addr_t *out_ipaddr,
                                  const types::IPAddress& in_ipaddr);
hal_ret_t ip_addr_to_spec(types::IPAddress *ip_addr_spec,
                          const ip_addr_t *ip_addr);
hal_ret_t ip_pfx_spec_to_pfx(ip_prefix_t *ip_pfx,
                             const types::IPPrefix& in_ippfx);
bool ip_addr_in_ip_pfx(ip_addr_t *ipaddr, ip_prefix_t *ip_pfx);

ApiStatus hal_prepare_rsp (hal_ret_t ret);
void hal_api_trace(const char *trace);

// handle list specific APIs
void hal_print_handles_list(dllist_ctxt_t  *list);
bool hal_handle_in_list(dllist_ctxt_t *handle_list, hal_handle_t handle);
hal_ret_t hal_add_to_handle_list(dllist_ctxt_t *handle_list,
                                 hal_handle_t handle);
void hal_free_handles_list(dllist_ctxt_t *list);
hal_ret_t hal_cleanup_handle_list(dllist_ctxt_t **list);

}    // namespace hal

#endif    // __LINKMGR_UTILS_HPP__

