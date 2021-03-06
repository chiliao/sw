//
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This module defines resource pool API
///
//----------------------------------------------------------------------------

#ifndef __INCLUDE_API_PDS_RESOURCE_POOL_HPP__
#define __INCLUDE_API_PDS_RESOURCE_POOL_HPP__

#include "nic/apollo/api/include/pds.hpp"

/// \defgroup PDS_RESOURCE_POOL Resource pool API
/// @{

/// \brief Traffic class type
typedef enum pds_traffic_class_type_s {
    PDS_TRAFFIC_CLASS_AGGR_TX    = 0,    ///< all traffic from vnic
    PDS_TRAFFIC_CLASS_AGGR_RX    = 1,    ///< all traffic to vnic
    PDS_TRAFFIC_CLASS_IGW_TX     = 2,    ///< traffic to internet gateway
    PDS_TRAFFIC_CLASS_IGW_RX     = 3,    ///< traffic from internet gateway
    PDS_TRAFFIC_CLASS_INT_SVC_TX = 4,    ///< traffic to local public service
    PDS_TRAFFIC_CLASS_INT_SVC_RX = 5,    ///< traffic from local public service
    PDS_TRAFFIC_CLASS_MAX        = 6,
} pds_traffic_class_type_t;

/// \brief Class limits
typedef struct pds_class_limits_s {
    uint32_t bytes_per_second;    ///< bps limit for given traffic class
} __PACK__ pds_class_limits_t;

/// \brief Connection tracking config
typedef struct pds_conntrack_cfg_s {
    uint32_t session_limit;    ///< Uppoer limit; 0 indicates unlimited sessions
} __PACK__ pds_conntrack_cfg_t;

/// \brief Resource limits
typedef struct pds_rsc_limits_s {
    /// Limits per class
    pds_class_limits_t class_limits[PDS_TRAFFIC_CLASS_MAX];
    /// Connection tracking limits
    pds_conntrack_cfg_t conn_track;
} __PACK__ pds_rsc_limits_t;

/// Resource pool
typedef struct pds_rsc_pool_s {
    pds_rsc_pool_id_t key;      ///< Key
    pds_rsc_limits_t limits;    ///< Limits
} __PACK__ pds_rsc_pool_t;

/// \brief Create resource pool
/// \param[in] rsc_pool Resource pool information
/// \param[in] bctxt batch context if API is invoked in a batch
/// \return #SDK_RET_OK on success, failure status code on error
pds_status_t pds_rsc_pool_create(pds_rsc_pool_t *rsc_pool,
                                 pds_batch_ctxt_t bctxt = PDS_BATCH_CTXT_INVALID);

/// \brief Delete resource pool
/// \param[in] key Key
/// \param[in] bctxt batch context if API is invoked in a batch
/// \return #SDK_RET_OK on success, failure status code on error
pds_status_t pds_rsc_pool_delete(pds_rsc_pool_key_t *key,
                                 pds_batch_ctxt_t bctxt = PDS_BATCH_CTXT_INVALID);


/// @}

#endif    // __INCLUDE_API_PDS_RESOURCE_POOL_HPP_
