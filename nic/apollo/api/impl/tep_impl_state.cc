//
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains Tunnel EndPoint (TEP) datapath database handling
///
//----------------------------------------------------------------------------


#include "nic/apollo/api/include/pds_tep.hpp"
#include "nic/apollo/api/impl/tep_impl_state.hpp"
#include "gen/p4gen/apollo/include/p4pd.h"
#include "nic/sdk/lib/p4/p4_api.hpp"

// default drop/blackhole nexthop
#define PDS_DROP_NEXTHOP_HW_ID         0

namespace api {
namespace impl {

/// \defgroup PDS_TEP_IMPL_STATE - tep database functionality
/// \ingroup PDS_TEP
/// \@{

tep_impl_state::tep_impl_state(pds_state *state) {
    p4pd_table_properties_t    tinfo;

    // instantiate P4 tables for bookkeeping
    p4pd_table_properties_get(P4TBL_ID_TEP_TX, &tinfo);
    // TODO: table_health_monitor_cb is passed as NULL here !!
    tep_tx_tbl_ = directmap::factory(tinfo.tablename, P4TBL_ID_TEP_TX,
                                     tinfo.tabledepth,
                                     tinfo.actiondata_struct_size,
                                     false, true, NULL);
    SDK_ASSERT(tep_tx_tbl_ != NULL);

    p4pd_table_properties_get(P4TBL_ID_NEXTHOP_TX, &tinfo);
    // TODO: table_health_monitor_cb is passed as NULL here !!
    nh_tx_tbl_ = directmap::factory(tinfo.tablename, P4TBL_ID_NEXTHOP_TX,
                                    tinfo.tabledepth,
                                    tinfo.actiondata_struct_size,
                                    false, true, NULL);
    SDK_ASSERT(nh_tx_tbl_ != NULL);

    // reserve system default blackhole/drop nexthop entry
    SDK_ASSERT(nh_tx_tbl_->reserve_index(PDS_DROP_NEXTHOP_HW_ID) == SDK_RET_OK);
}

tep_impl_state::~tep_impl_state() {
    directmap::destroy(tep_tx_tbl_);
    directmap::destroy(nh_tx_tbl_);
}

sdk_ret_t
tep_impl_state::table_transaction_begin(void) {
    //tep_tx_tbl_->txn_start();
    //nh_tx_tbl_->txn_start();
    return SDK_RET_OK;
}

sdk_ret_t
tep_impl_state::table_transaction_end(void) {
    //tep_tx_tbl_->txn_end();
    //nh_tx_tbl_->txn_end();
    return SDK_RET_OK;
}

/// \@}    // end of PDS_TEP_IMPL_STATE

}    // namespace impl
}    // namespace api
