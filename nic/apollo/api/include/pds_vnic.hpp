//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This module defines VNIC API
///
//----------------------------------------------------------------------------

#ifndef __INCLUDE_API_PDS_VNIC_HPP__
#define __INCLUDE_API_PDS_VNIC_HPP__

#include "nic/sdk/include/sdk/eth.hpp"
#include "nic/apollo/api/include/pds.hpp"
#include "nic/apollo/api/include/pds_vcn.hpp"
#include "nic/apollo/api/include/pds_subnet.hpp"

/// \defgroup PDS_VNIC VNIC API
/// @{

#define PDS_MAX_VNIC 1024

/// \brief VNIC key
typedef struct pds_vnic_key_s {
    pds_vnic_id_t id;    ///< Unique VNIC ID (in the range 0 to 1024)
} __PACK__ pds_vnic_key_t;

/// \brief VNIC specification
typedef struct pds_vnic_spec_s {
    pds_vnic_key_t key;                ///< vnic's key
    pds_vcn_key_t vcn;                 ///< vcn of this vnic
    pds_subnet_key_t subnet;           ///< subnet of this vnic
    pds_encap_t vnic_encap;            ///< vnic encap for this vnic
    pds_encap_t fabric_encap;          ///< fabric encap for this vnic
    mac_addr_t mac_addr;               ///< vnic's overlay mac mac address
    pds_rsc_pool_id_t rsc_pool_id;     ///< resource pool
    bool src_dst_check;                ///< tRUE if src/dst check is enabled
    uint8_t tx_mirror_session_bmap;    ///< Tx mirror sessions, if any
    uint8_t rx_mirror_session_bmap;    ///< Rx mirror sessions, if any
} __PACK__ pds_vnic_spec_t;

/// \brief VNIC status
typedef struct pds_vnic_status_s {
} pds_vnic_status_t;

/// \brief VNIC statistics
typedef struct pds_vnic_stats_s {
    uint64_t rx_pkts;     ///< received packet count
    uint64_t rx_bytes;    ///< received bytes
    uint64_t tx_pkts;     ///< transmit packet count
    uint64_t tx_bytes;    ///< transmit bytes
} pds_vnic_stats_t;

/// \brief VNIC information
typedef struct pds_vnic_info_s {
    pds_vnic_spec_t spec;        ///< vnic specification
    pds_vnic_status_t status;    ///< vnic status
    pds_vnic_stats_t stats;      ///< vnic stats
} pds_vnic_info_t;

/// \brief Create VNIC
///
/// \param[in] spec Specification
/// \return #SDK_RET_OK on success, failure status code on error
sdk_ret_t pds_vnic_create(pds_vnic_spec_t *spec);

/// \brief Read VNIC information
///
/// \param[in] key Key
/// \param[out] info Information
/// \return #SDK_RET_OK on success, failure status code on error
sdk_ret_t pds_vnic_read(pds_vnic_key_t *key, pds_vnic_info_t *info);

/// \brief Update VNIC specification
///
/// \param[in] spec Specififcation
/// \return #SDK_RET_OK on success, failure status code on error
sdk_ret_t pds_vnic_update(pds_vnic_spec_t *spec);

/// \brief Delete VNIC
///
/// \param[in] key Key
/// \return #SDK_RET_OK on success, failure status code on error
sdk_ret_t pds_vnic_delete(pds_vnic_key_t *key);

/// @}

#endif    ///  __INCLUDE_API_PDS_VNIC_HPP__
