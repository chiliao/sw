// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#include <stdio.h>
#include <string>
#include <errno.h>
#include <stdlib.h>

#include "nic/include/base.h"
#include "nic/gen/common_rxdma_actions/include/common_rxdma_actions_p4pd_table.h"
#include "nic/gen/common_txdma_actions/include/common_txdma_actions_p4pd_table.h"
#include "nic/hal/pd/p4pd_api.hpp"

/*
 * Based on table id call appropriate table routine.
 * For now this API is only called for p4pd and rxdma
 * routines. TODO: Add other APIs as well here.
 */
void
p4pd_global_hwentry_query (uint32_t tableid,
                           uint32_t *hwkey_len,
                           uint32_t *hwkeymask_len,
                           uint32_t *hwactiondata_len)
{
    if ((tableid >= p4pd_tableid_min_get()) &&
        (tableid <= p4pd_tableid_max_get())) {
        p4pd_hwentry_query(tableid, hwkey_len, hwkeymask_len, hwactiondata_len);
    } else if ((tableid >= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        p4pd_common_rxdma_actions_hwentry_query(tableid, hwkey_len,
            hwkeymask_len, hwactiondata_len);
    } else if ((tableid >= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        p4pd_common_txdma_actions_hwentry_query(tableid, hwkey_len,
            hwkeymask_len, hwactiondata_len);
    } else {
        HAL_ASSERT(0);
    }
    return;
}

/*
 * Based on table id call appropriate table routine.
 * For now this API is only called for p4pd and rxdma
 * routines. TODO: Add other APIs as well here.
 */
p4pd_error_t
p4pd_global_entry_write (uint32_t tableid,
                         uint32_t  index,
                         uint8_t   *hwkey, 
                         uint8_t   *hwkey_mask,
                         void      *actiondata)
{
    if ((tableid >= p4pd_tableid_min_get()) &&
        (tableid <= p4pd_tableid_max_get())) {
        return (p4pd_entry_write(tableid, index, hwkey, hwkey_mask, actiondata));
    } else if ((tableid >= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_rxdma_actions_entry_write(tableid,
                index, hwkey, hwkey_mask, actiondata));
    } else if ((tableid >= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_txdma_actions_entry_write(tableid,
                index, hwkey, hwkey_mask, actiondata));
    } else {
        HAL_ASSERT(0);
    }
    return HAL_RET_OK;
}

/*
 * Based on table id call appropriate table routine.
 * For now this API is only called for p4pd and rxdma
 * routines. TODO: Add other APIs as well here.
 */
p4pd_error_t
p4pd_global_entry_write_with_datamask (uint32_t tableid,
                                       uint32_t  index,
                                       uint8_t   *hwkey, 
                                       uint8_t   *hwkey_mask,
                                       void      *actiondata,
                                       void      *actiondata_mask)
{
    if ((tableid >= p4pd_tableid_min_get()) &&
        (tableid <= p4pd_tableid_max_get())) {
        return (p4pd_entry_write_with_datamask(tableid, index, hwkey, hwkey_mask, 
                                               actiondata, actiondata_mask));
    } else if ((tableid >= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_rxdma_actions_entry_write_with_datamask(tableid,
                index, hwkey, hwkey_mask, actiondata, actiondata_mask));
    } else if ((tableid >= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_txdma_actions_entry_write_with_datamask(tableid,
                index, hwkey, hwkey_mask, actiondata, actiondata_mask));
    } else {
        HAL_ASSERT(0);
    }
    return HAL_RET_OK;
}

/*
 * Based on table id call appropriate table routine.
 * For now this API is only called for p4pd and rxdma
 * routines. TODO: Add other APIs as well here.
 */
p4pd_error_t
p4pd_global_entry_read (uint32_t tableid,
                        uint32_t   index,
                        void       *swkey,
                        void       *swkey_mask,
                        void       *actiondata)
{
    if ((tableid >= p4pd_tableid_min_get()) &&
        (tableid <= p4pd_tableid_max_get())) {
        return (p4pd_entry_read(tableid, index, swkey, swkey_mask, actiondata));
    } else if ((tableid >= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_rxdma_actions_entry_read(tableid,
                index, swkey, swkey_mask, actiondata));
    } else if ((tableid >= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_txdma_actions_entry_read(tableid,
                index, swkey, swkey_mask, actiondata));
    } else {
        HAL_ASSERT(0);
    }
    return HAL_RET_OK;
}

p4pd_error_t
p4pd_global_table_ds_decoded_string_get (uint32_t   tableid,
                                         uint32_t   index,
                                         void*      sw_key,
                                         /* Valid only in case of TCAM;
                                          * Otherwise can be NULL) */
                                         void*      sw_key_mask,
                                         void*      action_data,
                                         char*      buffer,
                                         uint16_t   buf_len)
{
    if ((tableid >= p4pd_tableid_min_get()) &&
        (tableid <= p4pd_tableid_max_get())) {
        return (p4pd_table_ds_decoded_string_get(tableid,
                index, sw_key, sw_key_mask, action_data, buffer, buf_len));
    } else if ((tableid >= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_rxdma_actions_table_ds_decoded_string_get(tableid,
                index, sw_key, sw_key_mask, action_data, buffer, buf_len));
    } else if ((tableid >= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pd_common_txdma_actions_table_ds_decoded_string_get(tableid,
                index, sw_key, sw_key_mask, action_data, buffer, buf_len));
    } else {
        HAL_ASSERT(0);
    }
    return HAL_RET_OK;
}

/* P4PD API that uses tableID to return table properties that HAL
 * layer can use to construct, initialize P4 tables in local memory.
 *
 * Arguments:
 *
 *  IN  : uint32_t          tableid    : Table Id that identifies
 *                                       P4 table. This id is obtained
 *                                       from p4pd_table_id_enum.
 *  OUT : p4pd_table_ctx_t *table_ctx  : Returns a structure of data
 *                                       that contains table properties.
 * Return Value:
 *  P4PD_SUCCESS                       : Table properties copied into tbl_ctx
 *                                       Memory for tbl_ctx is provided by
 *                                       API caller.
 *
 *  P4PD_FAIL                          : If tableid is not valid
 */
p4pd_error_t
p4pd_global_table_properties_get (uint32_t tableid, void *tbl_ctx)
{
    if ((tableid >= p4pd_tableid_min_get()) &&
        (tableid <= p4pd_tableid_max_get())) {
        return (p4pd_table_properties_get(tableid,
                (p4pd_table_properties_t*)tbl_ctx));
    } else if ((tableid >= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_RXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pluspd_rxdma_table_properties_get(tableid,
               (p4pd_table_properties_t*) tbl_ctx));
    } else if ((tableid >= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMIN) &&
         (tableid <= P4_COMMON_TXDMA_ACTIONS_TBL_ID_TBLMAX)) {
        return (p4pluspd_txdma_table_properties_get(tableid,
               (p4pd_table_properties_t*) tbl_ctx));
    } else {
        HAL_ASSERT(0);
    }
    return P4PD_SUCCESS;
}
