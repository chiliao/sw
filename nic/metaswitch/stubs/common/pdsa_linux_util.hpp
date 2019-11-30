//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// Utilities to interface with Linux OS
//---------------------------------------------------------------

#ifndef __PDSA_LINUX_UTIL_HPP__
#define __PDSA_LINUX_UTIL_HPP__

#include "nic/sdk/include/sdk/eth.hpp"
#include "nic/sdk/include/sdk/ip.hpp"
#include <string>

namespace pdsa_stub {

// Fetch Linux parameters (IfIndex, MAC) for given interface
bool get_linux_intf_params(const char* ifname,      // In
                           uint32_t*   lnx_ifindex, // Out
                           mac_addr_t& mac);        // Out

void config_linux_loopback_ip (const ip_addr_t& ip, uint32_t prefix_len, 
                               bool del = false);

} // End namespace
#endif
