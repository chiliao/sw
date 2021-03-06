//
// Copyright (c) 2019 Pensando Systems, Inc.
//----------------------------------------------------------------------------
///
/// \file
/// service mapping entry handling
///
//----------------------------------------------------------------------------

#include "nic/apollo/core/mem.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/api/service.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/apollo/framework/api_params.hpp"
#include "nic/apollo/framework/api_base.hpp"
#include "nic/apollo/framework/api_engine.hpp"
#include "nic/apollo/framework/api_params.hpp"

namespace api {

svc_mapping::svc_mapping() {
    ht_ctxt_.reset();
    skey_ht_ctxt_.reset();
    impl_ = NULL;
}

svc_mapping *
svc_mapping::factory(pds_svc_mapping_spec_t *spec) {
    svc_mapping *mapping;

    // create service mapping entry with defaults, if any
    mapping = svc_mapping_db()->alloc();
    if (mapping) {
        new (mapping) svc_mapping();
        mapping->impl_ =
            impl_base::factory(impl::IMPL_OBJ_ID_SVC_MAPPING, spec);
        if (mapping->impl_ == NULL) {
            svc_mapping::destroy(mapping);
            return NULL;
        }
    }
    return mapping;
}

svc_mapping::~svc_mapping() {
}

void
svc_mapping::destroy(svc_mapping *mapping) {
    mapping->nuke_resources_();
    if (mapping->impl_) {
        impl_base::destroy(impl::IMPL_OBJ_ID_SVC_MAPPING, mapping->impl_);
    }
    mapping->~svc_mapping();
    svc_mapping_db()->free(mapping);
}

api_base *
svc_mapping::clone(api_ctxt_t *api_ctxt) {
    svc_mapping *cloned_mapping;

    cloned_mapping = svc_mapping_db()->alloc();
    if (cloned_mapping) {
        new (cloned_mapping) svc_mapping();
        if (cloned_mapping->init_config(api_ctxt) != SDK_RET_OK) {
            goto error;
        }
        cloned_mapping->impl_ = impl_->clone();
        if (unlikely(cloned_mapping->impl_ == NULL)) {
            PDS_TRACE_ERR("Failed to clone mapping %s impl", key2str().c_str());
            goto error;
        }
    }
    return cloned_mapping;

error:

    cloned_mapping->~svc_mapping();
    svc_mapping_db()->free(cloned_mapping);
    return NULL;
}

sdk_ret_t
svc_mapping::free(svc_mapping *mapping) {
    if (mapping->impl_) {
        impl_base::free(impl::IMPL_OBJ_ID_SVC_MAPPING, mapping->impl_);
    }
    mapping->~svc_mapping();
    svc_mapping_db()->free(mapping);
    return SDK_RET_OK;
}

svc_mapping *
svc_mapping::build(pds_svc_mapping_key_t *skey) {
    svc_mapping *mapping;

    // create service mapping entry with defaults, if any
    mapping = svc_mapping_db()->alloc();
    if (mapping) {
        new (mapping) svc_mapping();
        memcpy(&mapping->skey_, skey, sizeof(*skey));
        mapping->impl_ = impl_base::build(impl::IMPL_OBJ_ID_SVC_MAPPING,
                                          skey, mapping);
        if (mapping->impl_ == NULL) {
            svc_mapping::destroy(mapping);
            return NULL;
        }
    }
    return mapping;
}

svc_mapping *
svc_mapping::build(pds_obj_key_t *key) {
    pds_svc_mapping_key_t skey;
    svc_mapping *mapping = NULL;

    // find the 2nd-ary key corresponding to this primary key
    if (svc_mapping_db()->skey(key, &skey) == SDK_RET_OK) {
        // and then build the object
        mapping = svc_mapping::build(&skey);
        if (mapping) {
            memcpy(&mapping->key_, key, sizeof(*key));
        }
    }
    return mapping;
}

void
svc_mapping::soft_delete(svc_mapping *mapping) {
    if (mapping->impl_) {
        impl_base::soft_delete(impl::IMPL_OBJ_ID_SVC_MAPPING, mapping->impl_);
    }
    mapping->del_from_db();
    mapping->~svc_mapping();
    svc_mapping_db()->free(mapping);
}

 sdk_ret_t
svc_mapping::reserve_resources(api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    return impl_->reserve_resources(this, orig_obj, obj_ctxt);
}

sdk_ret_t
svc_mapping::release_resources(void) {
    return impl_->release_resources(this);
}

sdk_ret_t
svc_mapping::nuke_resources_(void) {
    if (impl_ == NULL) {
        return SDK_RET_OK;
    }
    return impl_->nuke_resources(this);
}

sdk_ret_t
svc_mapping::init_config(api_ctxt_t *api_ctxt) {
    pds_svc_mapping_spec_t *spec = &api_ctxt->api_params->svc_mapping_spec;

    memcpy(&key_, &spec->key, sizeof(key_));
    memcpy(&skey_, &spec->skey, sizeof(skey_));
    return SDK_RET_OK;
}

sdk_ret_t
svc_mapping::program_create(api_obj_ctxt_t *obj_ctxt) {
    return impl_->program_hw(this, obj_ctxt);
}

sdk_ret_t
svc_mapping::cleanup_config(api_obj_ctxt_t *obj_ctxt) {
    return impl_->cleanup_hw(this, obj_ctxt);
}

sdk_ret_t
svc_mapping::program_update(api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    return impl_->update_hw(orig_obj, this, obj_ctxt);
}

sdk_ret_t
svc_mapping::activate_config(pds_epoch_t epoch, api_op_t api_op,
                             api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    PDS_TRACE_DEBUG("Activating %s", key_.str());
    return impl_->activate_hw(this, orig_obj, epoch, api_op, obj_ctxt);
}

sdk_ret_t
svc_mapping::read(pds_obj_key_t *key, pds_svc_mapping_info_t *info) {
    pds_svc_mapping_key_t skey;

    PDS_TRACE_DEBUG("Reading %s", key->str());
    // find the 2nd-ary key corresponding to this primary key
    if (svc_mapping_db()->skey(key, &skey) == SDK_RET_OK) {
        // and then read from h/w
        memcpy(&info->spec.key, key, sizeof(*key));
        memcpy(&info->spec.skey, &skey, sizeof(skey));
        return impl_->read_hw(this, (impl::obj_key_t *)&skey,
                              (impl::obj_info_t *)info);
    }
    return SDK_RET_ENTRY_NOT_FOUND;
}

sdk_ret_t
svc_mapping::read(pds_svc_mapping_key_t *skey, pds_svc_mapping_info_t *info) {
    return impl_->read_hw(this, (impl::obj_key_t *)skey,
                          (impl::obj_info_t *)info);
}

// even though mapping object is stateless, we need to temporarily insert
// into the db as back-to-back operations on the same object can be issued
// in same batch
sdk_ret_t
svc_mapping::add_to_db(void) {
    return svc_mapping_db()->insert(this);
}

sdk_ret_t
svc_mapping::del_from_db(void) {
    if (svc_mapping_db()->remove(this)) {
        return SDK_RET_OK;
    }
    return SDK_RET_ENTRY_NOT_FOUND;
}

sdk_ret_t
svc_mapping::update_db(api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    if (svc_mapping_db()->remove((svc_mapping *)orig_obj)) {
        return svc_mapping_db()->insert(this);
    }
    return SDK_RET_ENTRY_NOT_FOUND;
}

sdk_ret_t
svc_mapping::delay_delete(void) {
    return delay_delete_to_slab(PDS_SLAB_ID_SVC_MAPPING, this);
}

}    // namespace api
