//------------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//------------------------------------------------------------------------------
#ifndef __PDS_MS_TEST_PHY_PORT_PDS_MOCK_HPP__
#define __PDS_MS_TEST_PHY_PORT_PDS_MOCK_HPP__

#include "nic/metaswitch/stubs/test/hals/phy_port_test_params.hpp"
#include "nic/metaswitch/stubs/test/pdsmock/pds_api_mock.hpp"

namespace pds_ms_test {

class phy_port_pds_mock_t final : public pds_mock_t {
public:
    void validate() override {
        validate_();
        reset_mock_fail();
    }
    void expect_create() override {
        op_create_ = true; op_delete_ = false;
        clear_batches();
        auto phy_port_input = dynamic_cast<phy_port_input_params_t*>
                                 (test_params()->test_input);
        generate_addupd_specs(*phy_port_input, expected_pds);
        ++num_if_objs_;
    }
    void expect_update() override {
        op_create_ = false; op_delete_ = false;
        clear_batches();
        auto phy_port_input = dynamic_cast<phy_port_input_params_t*>
                                 (test_params()->test_input);
        generate_addupd_specs(*phy_port_input, expected_pds);
    }
    void expect_pds_spec_op_fail(void) override {
        pds_mock_t::expect_pds_spec_op_fail();
        ++num_if_objs_;
    }
    void expect_pds_batch_commit_fail(void) override {
        pds_mock_t::expect_pds_batch_commit_fail();
        ++num_if_objs_;
    }

    void expect_create_pds_async_fail() override {
        mock_pds_batch_async_fail_ = true;
        expect_create();
    }
    void cleanup(void) override {
        pds_mock_t::cleanup();
        num_if_objs_ = 0;
    }
private:
    uint32_t           num_if_objs_ = 0;
    void generate_addupd_specs(const phy_port_input_params_t& input,
                               batch_spec_t& pds_batch);
    void generate_del_specs(const phy_port_input_params_t& input,
                            batch_spec_t& pds_batch);
    void validate_();
    pds_obj_key_t make_l3if_key_(const phy_port_input_params_t& input);
};

} // End namespace pds_ms_test

#endif
