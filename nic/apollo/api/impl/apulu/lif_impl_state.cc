//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// lif implementation state maintenance
///
//----------------------------------------------------------------------------

#include "nic/sdk/lib/p4/p4_api.hpp"
#include "nic/sdk/lib/rte_indexer/rte_indexer.hpp"
#include "nic/apollo/api/impl/lif_impl_state.hpp"
#include "gen/p4gen/p4plus_txdma/include/p4plus_txdma_p4pd.h"

/// \defgroup PDS_LIF_IMPL_STATE - lif state functionality
/// \ingroup PDS_LIF
/// \@{

#define PDS_MAX_LIFS            512

namespace api {
namespace impl {

lif_impl_state::lif_impl_state(pds_state *state) {
    p4pd_table_properties_t    tinfo;

    // uuid based lif database
    lif_ht_ = ht::factory(PDS_MAX_LIFS >> 2,
                          lif_impl::lif_key_func_get,
                          sizeof(pds_obj_key_t));
    SDK_ASSERT(lif_ht_ != NULL);

    // lif (internal) id based lif database
    lif_id_ht_ = ht::factory(PDS_MAX_LIFS >> 2,
                          lif_impl::lif_id_func_get,
                          sizeof(pds_lif_id_t));
    SDK_ASSERT(lif_id_ht_ != NULL);

    p4pluspd_txdma_table_properties_get(
                P4_P4PLUS_TXDMA_TBL_ID_TX_TABLE_S5_T4_LIF_RATE_LIMITER_TABLE,
                &tinfo);
    tx_rate_limiter_tbl_ =
        directmap::factory(tinfo.tablename,
            P4_P4PLUS_TXDMA_TBL_ID_TX_TABLE_S5_T4_LIF_RATE_LIMITER_TABLE,
            tinfo.tabledepth, tinfo.actiondata_struct_size,
            false, true, NULL);
    SDK_ASSERT(tx_rate_limiter_tbl_ != NULL);
}

lif_impl_state::~lif_impl_state() {
    ht::destroy(lif_ht_);
    ht::destroy(lif_id_ht_);
    directmap::destroy(tx_rate_limiter_tbl_);
}

/// \@}

}    // namespace impl
}    // namespace api
