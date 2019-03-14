//
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// Base object definition for all impl objects
///
//----------------------------------------------------------------------------

#ifndef __FRAMEWORK_IMPL_BASE_HPP__
#define __FRAMEWORK_IMPL_BASE_HPP__


#include "nic/sdk/include/sdk/base.hpp"
#include "nic/apollo/framework/api.hpp"
#include "nic/apollo/framework/obj_base.hpp"
#include "nic/apollo/framework/api_base.hpp"
#include "nic/apollo/framework/impl.hpp"
#include "nic/apollo/framework/asic_impl_base.hpp"
#include "nic/apollo/framework/pipeline_impl_base.hpp"
#include "nic/sdk/asic/asic.hpp"

namespace api {
namespace impl {

/// \brief Base class for all impl objects
class impl_base : public obj_base {
public:
    /// \brief Constructor
    impl_base() {}

    /// \brief Destructor */
    ~impl_base() {}

    /// \brief One time init function that must be called during bring up
    ///
    /// \param[in] params Initialization parameters passed by application
    /// \param[in] asic_cfg ASIC configuration parameters
    /// \return #SDK_RET_OK on success, failure status code on error
    static sdk_ret_t init(pds_init_params_t *params, asic_cfg_t *asic_cfg);

    /// \brief Dump all the debug information to given file
    ///
    /// \param[in] fp File handle
    static void debug_dump(FILE *fp);

    /// \brief Factory method to instantiate an impl object
    ///
    /// \param[in] impl Object id
    /// \param[in] args Args (not interpreted by this class)
    static impl_base *factory(impl_obj_id_t obj_id, void *args);

    /// \brief Release all the resources associated with this object
    ///
    /// \param[in] obj_id Object id
    // \param[in] impl_obj Impl instance to be freed
    static void destroy(impl_obj_id_t obj_id, impl_base *impl_obj);

    /// \brief Instantiate an impl object based on current state (sw and/or hw
    //         state) given its key
    /// \param[in] obj_id Object id
    /// \param[in] args Args (not interpreted by this class)
    static impl_base *build(impl_obj_id_t obj_id, void *args);

    /// \brief Allocate/reserve h/w resources for this object
    ///
    /// \param[in] orig_obj Old version of the unmodified object
    /// \param[in] obj_ctxt Transient state associated with this API
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t reserve_resources(api_base *orig_obj,
                                        obj_ctxt_t *obj_ctxt) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Free h/w resources used by this object, if any
    /// This API is invoked during object deletes
    ///
    /// \param[in] api_obj API object holding the resources
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t nuke_resources(api_base *api_obj) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Release h/w resources reserved for this object, if any
    /// This API is invoked during the rollback stage
    ///
    /// \param[in] api_obj API object holding the resources
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t release_resources(api_base *api_obj) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Program hardware
    /// Program all h/w tables relevant to this object except stage 0 table(s),
    /// if any
    ///
    /// \param[in] api_obj API object being programmed
    /// \param[in] obj_ctxt Transient state associated with this API
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t program_hw(api_base *api_obj,
                                 obj_ctxt_t *obj_ctxt) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Cleanup hardware
    /// Cleanup all h/w tables relevant to this object except stage 0 table(s),
    /// if any, by updating packed entries with latest epoch
    ///
    /// \param[in] api_obj API object being cleaned up
    /// \param[in] obj_ctxt Transient state associated with this API
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t cleanup_hw(api_base *api_obj,
                                 obj_ctxt_t *obj_ctxt) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Update hardware
    /// Update all h/w tables relevant to this object except stage 0 table(s),
    /// if any, by updating packed entries with latest epoch
    ///
    /// \param[in] orig_obj Old version of the unmodified object
    /// \param[in] curr_obj Cloned and updated version of the object
    /// \param[in] obj_ctxt Transient state associated with this API
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t update_hw(api_base *orig_obj, api_base *curr_obj,
                                obj_ctxt_t *obj_ctxt) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Activate hardware
    /// Activate the epoch in the dataplane by programming stage 0 tables,
    /// if any
    ///
    /// \param[in] epoch Epoch being activated
    /// \param[in] api_op API operation
    /// \param[in] obj_ctxt Transient state associated with this API
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t activate_hw(api_base *api_obj,
                                  pds_epoch_t epoch,
                                  api_op_t api_op,
                                  obj_ctxt_t *obj_ctxt) {
        return sdk::SDK_RET_INVALID_OP;
    }

    /// \brief Return ASIC impl class instance
    static asic_impl_base *asic_impl(void) { return asic_impl_; }

    /// \brief Return pipeline impl class instance
    static pipeline_impl_base *pipeline_impl(void) { return pipeline_impl_; }

private:
    static asic_impl_base        *asic_impl_;
    static pipeline_impl_base    *pipeline_impl_;
};

}    // namespace impl
}    // namespace api

using api::impl::impl_base;
 
#endif    // __FRAMEWORK_IMPL_BASE_HPP__
