//------------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//------------------------------------------------------------------------------
#ifndef __PDS_MS_TEST_ROUTE_PDS_MOCK_HPP__
#define __PDS_MS_TEST_ROUTE_PDS_MOCK_HPP__

#include "nic/metaswitch/stubs/test/hals/route_test_params.hpp"
#include "nic/metaswitch/stubs/test/pdsmock/pds_api_mock.hpp"

namespace pds_ms_test {

class route_pds_mock_t final : public pds_mock_t {
public:
    void validate() override {
        validate_();
        reset_mock_fail();
    }
    void expect_create() override {
        op_create_ = true; op_delete_ = false;
        clear_batches();
        auto route_input = dynamic_cast<route_input_params_t*>(test_params()->test_input);
        generate_addupd_specs(*route_input, expected_pds);
    }
    void expect_update() override {
        op_create_ = false; op_delete_ = false;
        clear_batches();
        auto route_input = dynamic_cast<route_input_params_t*>(test_params()->test_input);
        generate_addupd_specs(*route_input, expected_pds);
    }
    void expect_delete() override {
        op_create_ = false; op_delete_ = false;
        op_delete_update_ = true;
        clear_batches();
        auto route_input = dynamic_cast<route_input_params_t*>(test_params()->test_input);
        generate_del_specs(*route_input, expected_pds);
    }
    void expect_create_pds_async_fail() override {
        mock_pds_batch_async_fail_ = true;
        expect_create();
    }
    void expect_pds_spec_op_fail(void) override {
        pds_mock_t::expect_pds_spec_op_fail();
    }
    void expect_pds_batch_commit_fail(void) override {
        pds_mock_t::expect_pds_batch_commit_fail();
    }

    void cleanup(void) override {
        pds_mock_t::cleanup();
    }
private:
    bool op_delete_update_ = false;
    void generate_addupd_specs(const route_input_params_t& input,
                               batch_spec_t& pds_batch);
    void generate_del_specs(const route_input_params_t& input,
                            batch_spec_t& pds_batch);
    void validate_();
};

} // End namespace pds_ms_test

#endif
