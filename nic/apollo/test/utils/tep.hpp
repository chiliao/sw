//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains the tep test utility routines
///
//----------------------------------------------------------------------------

#ifndef __TEST_UTILS_TEP_HPP__
#define __TEST_UTILS_TEP_HPP__

#include "nic/sdk/include/sdk/base.hpp"
#include "nic/apollo/include/api/oci_tep.hpp"

namespace api_test {

/// TEP test utility class
class tep_util {
public:
    // Test parameters
    std::string ip_str;       // TEP IP
    oci_encap_type_t type;    // TEP type

    /// \brief default constructor
    tep_util();

    /// \brief parameterized constructor
    tep_util(std::string ip_str);

    /// \brief parameterized constructor
    tep_util(std::string ip_str, oci_encap_type_t type);

    /// \brief destructor
    ~tep_util();

    /// \brief Create TEP
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    sdk_ret_t create(void);

    /// \brief Read TEP
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    sdk_ret_t read(oci_tep_info_t *info);

    /// \brief Update TEP
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    sdk_ret_t update(void);

    /// \brief Delete TEP
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    sdk_ret_t del(void);

    /// \brief Validate TEP
    ///
    /// validates the test parameters of this object against info
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    sdk_ret_t validate(oci_tep_info_t *info);

    /// \brief Create multiple TEPs
    ///
    /// Create "num_tep" TEPs of type "tep_type" with IPs in range
    /// ip_str...ip_str+num_tep
    ///
    /// \param[in] num_tep number of TEPs to be created
    ///
    /// \param[in] ip_str starting IP address
    ///
    /// \param[in] type TEP encap type
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    static sdk_ret_t many_create(uint32_t num_tep, std::string ip_str,
                                 oci_encap_type_t type = OCI_ENCAP_TYPE_VNIC);

    /// \brief Get all TEPs configured on the NAPLES
    ///
    /// \param[out] count number of TEP objects
    /// \param[out] teps_info list of TEP objects
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    static sdk_ret_t all_get(int *count, oci_tep_info_t *teps_info);

    /// \brief Update a TEP given its IP and updated spec
    ///
    /// \param[in] ip_str TEP IP
    ///
    /// \param[in] new_tep_spec modified TEP spec
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    static sdk_ret_t update(std::string ip_str, oci_tep_spec_t *new_tep_spec);

    /// \brief Delete a TEP given its IP
    ///
    /// \param[in] ip_str TEP IP
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    static sdk_ret_t del(std::string ip_str);

    /// \brief Delete all TEPs configured
    ///
    /// \returns #SDK_RET_OK on success, failure status code on error
    static sdk_ret_t all_delete(void);
};

}    // namespace api_test

#endif    // __TEST_UTILS_TEP_HPP__
