//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// Main entry point for the Pensando Distributed Services Agent (PDSA)
//---------------------------------------------------------------

#include "nic/apollo/api/include/pds_init.hpp"
#include "nic/metaswitch/pdsa_stub_init.hpp"
#include "nic/metaswitch/stubs/mgmt/pdsa_mgmt_init.hpp"

unsigned int g_node_a_ip;
unsigned int g_node_b_ip;
unsigned int g_node_a_ac_ip;
unsigned int g_node_b_ac_ip;
unsigned int g_evpn_if_index;

int main(void)
{
    if(!pdsa_stub_mgmt_init()) {
        return 1; 
    }

    pdsa_stub::init();
    return 0;
}
