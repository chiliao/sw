//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// DHCP state handling
///
//----------------------------------------------------------------------------

#ifndef __API_DHCP_STATE_HPP__
#define __API_DHCP_STATE_HPP__

#include "nic/sdk/lib/ht/ht.hpp"
#include "nic/sdk/lib/slab/slab.hpp"
#include "nic/apollo/framework/state_base.hpp"
#include "nic/apollo/api/dhcp.hpp"

namespace api {

/// \defgroup PDS_DHCP_STATE - DHCP state functionality
/// \ingroup PDS_DHCP
/// @{

/// \brief    state maintained for DHCP configuration
class dhcp_state : public state_base {
public:
    /// \brief constructor
    dhcp_state();

    /// \brief destructor
    ~dhcp_state();

    /// \brief      allocate memory required for a DHCP policy entry
    /// \return     pointer to the allocated DHCP policy entry, NULL if no memory
    dhcp_policy *alloc(void);

    /// \brief    insert given DHCP policy entry instance into the db
    /// \param[in] policy DHCP policy entry to be added to the db
    /// \return   SDK_RET_OK on success, failure status code on error
    sdk_ret_t insert(dhcp_policy *policy);

    /// \brief     remove the DHCP policy entry object from db
    /// \param[in] policy DHCP policy entry to be deleted from the db
    /// \return    pointer to the removed DHCP policy entry or NULL, if not found
    dhcp_policy *remove(dhcp_policy *policy);

    /// \brief      free DHCP policy entry back to slab
    /// \param[in]  policy pointer to the allocated DHCP policy entry
    void free(dhcp_policy *policy);

    /// \brief      lookup a DHCP policy entry in database given the key
    /// \param[in]  key DHCP policy entry key
    /// \return     pointer to the DHCP policy entry found or NULL
    dhcp_policy *find(pds_obj_key_t *key) const;

    /// \brief API to walk all the db elements
    /// \param[in] walk_cb    callback to be invoked for every node
    /// \param[in] ctxt       opaque context passed back to the callback
    /// \return   SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t walk(state_walk_cb_t walk_cb, void *ctxt) override;

    /// \brief API to walk all the slabs
    /// \brief API to walk all the slabs
    /// \param[in] walk_cb    callback to be invoked for every slab
    /// \param[in] ctxt       opaque context passed back to the callback
    /// \return   SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t slab_walk(state_walk_cb_t walk_cb, void *ctxt) override;

    friend void slab_delay_delete_cb(void *timer, uint32_t slab_id, void *elem);

private:
    ht *dhcp_policy_ht(void) const { return dhcp_policy_ht_; }
    slab *dhcp_policy_slab(void) const { return dhcp_policy_slab_; }

    /// dhcp_policy class is friend of dhcp_state
    friend class dhcp_policy;

private:
    ht *dhcp_policy_ht_;        ///< DHCP policy entry hash table root
    slab *dhcp_policy_slab_;    ///< slab for allocating DHCP policy entry
};

static inline dhcp_policy *
dhcp_policy_find (pds_obj_key_t *key)
{
    return (dhcp_policy *)api_base::find_obj(OBJ_ID_DHCP_POLICY, key);
}

/// \@}    // end of PDS_DHCP

}    // namespace api

using api::dhcp_state;

#endif    // __API_DHCP_STATE_HPP__
