//------------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//------------------------------------------------------------------------------
#ifndef __PHY_PORT_IPS_FEEDER_HPP__
#define __PHY_PORT_IPS_FEEDER_HPP__

#include "nic/metaswitch/stubs/test/hals/phy_port_test_params.hpp"
#include "nic/metaswitch/stubs/common/pdsa_util.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_li.hpp"
#include "nic/apollo/test/base/base.hpp"

namespace pdsa_test {

class phy_port_ips_feeder_t final : public phy_port_input_params_t {
public:
   void init() override {
       phy_port = 1;
       phy_port_ifindex = 0x10000;
       admin_state = false;
    }

    ATG_LIPI_PORT_ADD_UPDATE generate_add_upd_ips(void) {
        ATG_LIPI_PORT_ADD_UPDATE add_upd;
      // generate_ips_header (add_upd); 
        add_upd.id.if_index = phy_port_ifindex;
        strcpy (add_upd.id.if_name, "eth0");
        add_upd.port_settings.port_enabled = (admin_state)? ATG_YES:ATG_NO;
        add_upd.port_settings.port_enabled_updated = ATG_YES;
        add_upd.port_settings.no_switch_port = ATG_YES;
        add_upd.port_settings.no_switch_port_updated = ATG_YES;
        return add_upd;
    }

    void trigger_create(void) override {
        auto add_upd = generate_add_upd_ips();
        pds_ms::li_is()->port_add_update(&add_upd);
    }

    void trigger_delete(void) override {
        pds_ms::li_is()->port_delete(phy_port_ifindex);
    }

    void trigger_update(void) override {
        auto add_upd = generate_add_upd_ips();
        pds_ms::li_is()->port_add_update(&add_upd);
    }
    void modify(void) override {
        admin_state = !admin_state;
    }
    void next(void) override {
        phy_port += 1;
        phy_port_ifindex += 0x10000;
    }

    bool ips_mock() override {return true;}
};

} // End Namespace

#endif
