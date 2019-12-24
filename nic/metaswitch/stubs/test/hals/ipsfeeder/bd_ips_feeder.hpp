//------------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//------------------------------------------------------------------------------
#ifndef __BD_IPS_FEEDER_HPP__
#define __BD_IPS_FEEDER_HPP__

#include "nic/metaswitch/stubs/test/hals/bd_test_params.hpp"
#include "nic/metaswitch/stubs/common/pdsa_util.hpp"
#include "nic/metaswitch/stubs/common/pds_ms_ifindex.hpp"
#include "nic/metaswitch/stubs/common/pdsa_ms_defs.hpp"
#include <l2f_c_includes.hpp>
#include "nic/metaswitch/stubs/hals/pds_ms_l2f.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_l2f_bd.hpp"

namespace pdsa_test {
using pdsa_stub::ms_ifindex_t;

class bd_ips_feeder_t final : public bd_input_params_t {
public:
    pds_ms::l2f_integ_subcomp_t  l2f_is; 
    ms_ifindex_t   prev_if_bind = 0;
   void init() override {
       bd_input_params_t::init();
    }

    ATG_BDPI_UPDATE_BD generate_add_upd_ips(void) {
        ATG_BDPI_UPDATE_BD add_upd {0};
      // generate_ips_header (add_upd); 
        add_upd.bd_id.bd_id = subnet_spec.key.id;
        add_upd.bd_properties.vni = subnet_spec.fabric_encap.val.vnid;
        return add_upd;
    }

    void trigger_create(void) override {
        auto add_upd = generate_add_upd_ips();
        l2f_is.add_upd_bd(&add_upd);
    }

    void trigger_delete(void) override {
        bd_input_params_t::trigger_delete();
        ATG_L2_BD_ID bd_id = {ATG_L2_BRIDGE_DOMAIN_EVPN, subnet_spec.key.id, 0};
        l2f_is.delete_bd(&bd_id, NBB_CORRELATOR{0});
    }

    void trigger_update(void) override {
        if (test_if_bind) {
            pds_ms::l2f_bd_t  bd;
            auto ms_ifindex = pds_ms::pds_to_ms_ifindex(subnet_spec.host_ifindex,
                                                        IF_TYPE_LIF);
            bd.handle_add_if(subnet_spec.key.id, ms_ifindex);
            prev_if_bind = ms_ifindex;
            test_if_bind = false;
            return;
        }
        if (test_if_unbind) {
            ATG_BDPI_INTERFACE_BIND if_bind;
            ATG_L2_BD_ID bd_id; bd_id.bd_id = subnet_spec.key.id;
            if_bind.if_index = prev_if_bind;
            l2f_is.delete_bd_if(&bd_id, NBB_CORRELATOR(), &if_bind);
            test_if_unbind = false;
            return;
        }
        auto add_upd = generate_add_upd_ips();
        l2f_is.add_upd_bd(&add_upd);
    }

    void trigger_if_bind(void) {
    }
    void trigger_if_unbind(void) {
    }
    bool ips_mock() override {return true;}
};

} // End Namespace

#endif