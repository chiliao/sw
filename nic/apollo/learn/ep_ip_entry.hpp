
//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// IP Learning entry handling
///
//----------------------------------------------------------------------------

#ifndef __LEARN_EP_IP_ENTRY_HPP__
#define __LEARN_EP_IP_ENTRY_HPP__

#include "nic/apollo/learn/learn.hpp"

namespace learn {

/// \defgroup EP_LEARN - Endpoint IP Learning Entry functionality
/// @{

class ep_mac_entry;

/// \brief    IP Learning entry
class ep_ip_entry {
public:
    /// \brief          factory method to create an EP IP entry
    /// \param[in]      key         key info for IP address learnt
    /// \param[in]      vnic_obj_id object id of associated vnic
    /// \return         new instance of IP entry or NULL, in case of error
    static ep_ip_entry *factory(ep_ip_key_t *key, uint32_t vnic_obj_id);

    /// \brief          free memory allocated to IP entry
    /// \param[in]      ep_ip    pointer to IP entry
    static void destroy(ep_ip_entry *ep_ip);

    /// \brief          add this entry to ep database
    /// \return         SDK_RET_OK on sucess, failure status code on error
    sdk_ret_t add_to_db(void);

    /// \brief          del this entry from ep database
    /// \return         SDK_RET_OK on success, SDK_RET_ENTRY_NOT_FOUND
    ///                 if entry not found in db
    sdk_ret_t del_from_db(void);

    /// \brief          initiate delay deletion of this object
    sdk_ret_t delay_delete(void);

    /// \brief          get the state of this entry
    /// \return         state of this entry
    ep_state_t state(void) const { return state_; }

    /// \brief          set the state of this entry
    /// \param[in]      state    state to be set on this entry
    void set_state(ep_state_t state);

    /// \brief          helper function to get key given ep IP entry
    /// \param[in]      entry    pointer to ep IP instance
    /// \return         pointer to the ep IP instance's key
    static void *ep_ip_key_func_get(void *entry) {
        ep_ip_entry *ep_ip = (ep_ip_entry *)entry;
        return (void *)&(ep_ip->key_);
    }

    /// \brief          get mac entry associated with this IP entry
    /// \return         pointer to mac entry, which this IP entry associates to
    const ep_mac_entry *mac_entry(void);

    /// \brief          set the vnic object id for this entry
    /// \param[in]      vnic object id
    void set_vnic_obj_id(uint32_t vnic_obj_id) { vnic_obj_id_ = vnic_obj_id; }

    /// \brief          compare vnic obj id with the one assoc with this entry
    /// \return         true if vnic object ids are equal else false
    bool vnic_compare(uint32_t vnic_obj_id_);

    /// \brief          test if ep is busy with active state transition
    /// \return         true if ep is in transition, false otherwise
    bool active(void) const { return (state_ != EP_STATE_CREATED &&
                                      state_ != EP_STATE_PROBING);
    }

private:
    /// \brief          constructor
    ep_ip_entry();

    /// \brief          destructor
    ~ep_ip_entry();

private:
    ep_ip_key_t    key_;            ///< IP learning entry key
    uint32_t       vnic_obj_id_;    ///< key for vnic associated
    ep_state_t     state_;          ///< state of this entry
    ht_ctxt_t      ht_ctxt_;        ///< hash table context

    ///< ep IP state class is a friend of IP entry
    friend class ep_ip_state;
};

/// \@}    // end of EP_IP_ENTRY

}    // namespace learn

using learn::ep_ip_entry;

#endif    // __LEARN_EP_IP_ENTRY_HPP__