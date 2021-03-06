//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// security policy datapath database handling
///
//----------------------------------------------------------------------------

#include "nic/apollo/api/include/pds_policy.hpp"
#include "nic/apollo/api/impl/artemis/security_policy_impl_state.hpp"

namespace api {
namespace impl {

/// \defgroup PDS_SECURITY_POLICY_IMPL_STATE - security policy database
///                                            functionality
/// \ingroup PDS_SECURITY_POLICY
/// @{

security_policy_impl_state::security_policy_impl_state(pds_state *state) {
    v4_idxr_ =
        rte_indexer::factory(state->mempartition()->block_count("sacl_v4"),
                             false, false);
    SDK_ASSERT(v4_idxr_ != NULL);

    v4_region_addr_ = state->mempartition()->start_addr("sacl_v4");
    SDK_ASSERT(v4_region_addr_ != INVALID_MEM_ADDRESS);
    v4_table_size_ = state->mempartition()->block_size("sacl_v4");

    v6_idxr_ =
        rte_indexer::factory(state->mempartition()->block_count("sacl_v6"),
                             false, false);
    SDK_ASSERT(v6_idxr_ != NULL);

    v6_region_addr_ = state->mempartition()->start_addr("sacl_v6");
    SDK_ASSERT(v6_region_addr_ != INVALID_MEM_ADDRESS);
    v6_table_size_ = state->mempartition()->block_size("sacl_v6");
}

security_policy_impl_state::~security_policy_impl_state() {
    rte_indexer::destroy(v4_idxr_);
    rte_indexer::destroy(v6_idxr_);
}

sdk_ret_t
security_policy_impl_state::table_transaction_begin(void) {
    return SDK_RET_OK;
}

sdk_ret_t
security_policy_impl_state::table_transaction_end(void) {
    return SDK_RET_OK;
}

/// \@}

}    // namespace impl
}    // namespace api
