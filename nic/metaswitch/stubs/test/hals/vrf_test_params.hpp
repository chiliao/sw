//------------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//------------------------------------------------------------------------------
#ifndef __PDSA_TEST_VRF_PARAMS_HPP__
#define __PDSA_TEST_VRF_PARAMS_HPP__

#include "nic/metaswitch/stubs/test/hals/test_params.hpp"
#include "nic/metaswitch/stubs/mgmt/pds_ms_vpc.hpp"
#include "nic/metaswitch/stubs/common/pdsa_state.hpp"
#include "nic/apollo/api/include/pds_vpc.hpp"
#include "nic/sdk/include/sdk/if.hpp"

namespace pdsa_test{

class vrf_input_params_t : public test_input_base_t {
public:
    uint32_t           vrf_id;
    pds_vpc_spec_t  vpc_spec = {0};

   // These inputs are used to generate feeder inputs 
   // as well as output verifications 
   virtual void init() {
    vrf_id = 1;
    vpc_spec.key.id = vrf_id;
    vpc_spec.type = PDS_VPC_TYPE_TENANT;
    vpc_spec.v4_prefix.len = 24; 
    str2ipv4addr("23.1.10.1", &vpc_spec.v4_prefix.v4_addr);
    mac_str_to_addr((char*) "04:06:03:09:00:03", vpc_spec.vr_mac);
    vpc_spec.v4_route_table.id = vrf_id;
    vpc_spec.fabric_encap.type = PDS_ENCAP_TYPE_VXLAN;
    vpc_spec.fabric_encap.val.vnid  = 100;
    vpc_spec.tos = 5;

    auto state_ctxt = pdsa_stub::state_t::thread_context(); 
    state_ctxt.state()->vpc_store().add_upd (vrf_id, new pdsa_stub::vpc_obj_t(vpc_spec));
   }
   void modify(void) override {
       vpc_spec.fabric_encap.val.vnid  += 100;
   };
   void next(void) override {
       // Set an initial subnet spec in the BD store 
       vrf_id = vrf_id+1;
       vpc_spec.key.id = vrf_id;
       str2ipv4addr("25.2.10.1", &vpc_spec.v4_prefix.v4_addr);
       mac_str_to_addr((char*) "04:26:23:29:20:03", vpc_spec.vr_mac);
       vpc_spec.v4_route_table.id = vrf_id;
       vpc_spec.fabric_encap.val.vnid  += 100;
       auto state_ctxt = pdsa_stub::state_t::thread_context(); 
       state_ctxt.state()->vpc_store().add_upd (vrf_id, new pdsa_stub::vpc_obj_t(vpc_spec));
   }; 
   void trigger_delete(void) override { 
       auto state_ctxt = pdsa_stub::state_t::thread_context(); 
       state_ctxt.state()->vpc_store().erase (vrf_id);
   }
   virtual ~vrf_input_params_t(void) {};
   virtual void init_direct_update() {
       // Set an initial subnet spec in the BD store 
       vrf_id = vrf_id+1;
       vpc_spec.key.id = vrf_id;
       str2ipv4addr("33.3.10.1", &vpc_spec.v4_prefix.v4_addr);
       mac_str_to_addr((char*) "04:66:63:69:60:03", vpc_spec.vr_mac);
       vpc_spec.v4_route_table.id = vrf_id;
       vpc_spec.fabric_encap.val.vnid  += 100;
       auto state_ctxt = pdsa_stub::state_t::thread_context(); 
       state_ctxt.state()->vpc_store().add_upd (vrf_id, new pdsa_stub::vpc_obj_t(vpc_spec));

       // And then change it to simulate Direct Update
       vpc_spec.fabric_encap.val.vnid  += 100;
       str2ipv4addr("43.3.10.1", &vpc_spec.v4_prefix.v4_addr);
   }
   virtual void send_direct_update() {
       pds_ms::vpc_update(&vpc_spec, 0);
   }
   void modify_direct_update(void) {
       vpc_spec.fabric_encap.val.vnid  += 100;
       str2ipv4addr("53.3.10.1", &vpc_spec.v4_prefix.v4_addr);
   }
};

void load_vrf_test_input(void);
void load_vrf_test_output(void);

static inline void
load_vrf_test (void) 
{
    load_vrf_test_input();
    load_vrf_test_output();
}

} // End namespace

#endif
