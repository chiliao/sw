//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#ifndef HAL_PLUGINS_NETWORK_EP_LEARN_COMMON_TRANS_HPP_
#define HAL_PLUGINS_NETWORK_EP_LEARN_COMMON_TRANS_HPP_
#include "nic/include/base.hpp"
#include "nic/hal/iris/include/hal_state.hpp"
#include <netinet/if_ether.h>
#include "nic/utils/fsm/fsm.hpp"
#include "lib/periodic/periodic.hpp"

using hal::utils::fsm_state_machine_t;
using hal::utils::fsm_state_machine_def_t;
using hal::utils::fsm_event_data;
using hal::utils::fsm_timer_t;
using hal::utils::fsm_state_timer_ctx;
using namespace hal;

namespace hal {
namespace eplearn {

#define TRANS_NOP_EVENT UINT32_MAX

typedef struct trans_ip_entry_key_s {
    vrf_id_t vrf_id;    // VRF id
    ip_addr_t ip_addr;  // IP address of the endpoint
} __PACK__ trans_ip_entry_key_t;


class trans_t {
private:
    sdk_spinlock_t slock_;  // lock to protect this structure
    bool marked_for_delete_;
    trans_ip_entry_key_t ip_entry_key_;

public:
    fsm_state_machine_t *sm_;
    trans_t() {
        SDK_SPINLOCK_INIT(&this->slock_, PTHREAD_PROCESS_PRIVATE);
        marked_for_delete_ = false;
        sm_= nullptr;
        memset(&ip_entry_key_, 0, sizeof(trans_ip_entry_key_t));
    }
    virtual ~trans_t() {
        SDK_SPINLOCK_DESTROY(&this->slock_);
    }

    virtual void log_error(const char *message)  {}
    virtual void log_info(const char *message)  {}

    class trans_timer_t : public fsm_timer_t {
       public:
        trans_timer_t(uint32_t timer_id) : fsm_timer_t(timer_id) {
        }

        virtual fsm_state_timer_ctx add_timer(uint64_t timeout,
                                              fsm_state_machine_t *ctx,
                                              bool periodic = false) {
            void *timer = sdk::lib::timer_schedule(
                this->get_timer_id(), timeout, ctx, timeout_handler, periodic);
            return reinterpret_cast<fsm_state_timer_ctx>(timer);
        }
        virtual void delete_timer(fsm_state_timer_ctx timer) {
            sdk::lib::timer_delete(timer);
        }

        fsm_state_timer_ctx add_timer_with_custom_handler(uint64_t timeout,
                                              fsm_state_machine_t *ctx,
                                              sdk::lib::twheel_cb_t cb,
                                              bool periodic = false) {
            void *timer = sdk::lib::timer_schedule(
                this->get_timer_id(), timeout, ctx, cb, periodic);
            return reinterpret_cast<fsm_state_timer_ctx>(timer);
        }

        virtual uint64_t get_timeout_remaining(fsm_state_timer_ctx timer) {
            return sdk::lib::get_timeout_remaining(timer) / TIME_MSECS_PER_SEC ;
        }
        static void timeout_handler(void *timer, uint32_t timer_id, void *ctxt) {
            fsm_state_machine_t* sm_ = reinterpret_cast<fsm_state_machine_t*>(ctxt);
            sm_->reset_timer();
            trans_t* trans =
                reinterpret_cast<trans_t*>(sm_->get_ctx());
            trans_t::process_transaction(trans, sm_->get_timeout_event(), NULL);
        }

        ~trans_timer_t() {}
    };

    static void process_transaction(trans_t *trans, uint32_t event,
                fsm_event_data data) {
        trans->log_info("Received event..");
        if (trans->marked_for_delete_) {
            return;
        }
        SDK_SPINLOCK_LOCK(&trans->slock_);
        if (!trans->marked_for_delete_) {
            trans->sm_->process_event(event, data);
            if (trans->sm_->state_machine_competed()) {
                trans->marked_for_delete_ = true;
                /* Assuming delayed delete is setup initiate the transaction delete. */
                trans->log_info("Transaction completed, deleting...");
                delete trans;
            }
        } else {
            trans->log_error("Transaction in marked for delete state, skipping processing");
        }
        SDK_SPINLOCK_UNLOCK(&trans->slock_);
    }

    static void process_learning_transaction(trans_t *trans, fte::ctx_t &ctx, uint32_t event,
                fsm_event_data data);
    static void trans_completion_handler(fte::ctx_t& ctx, bool status);

    static hal_ret_t process_ip_move(hal_handle_t ep_handle, const ip_addr_t *ip_addr,
             ht *ip_ht) {

        ip_addr_t ep_ip_addr = {0};
        vrf_t *vrf;
        ep_t *ep_entry = find_ep_by_handle(ep_handle);
        trans_ip_entry_key_t ip_entry_key;

        if (ep_entry == nullptr) {
            return HAL_RET_EP_NOT_FOUND;
        }

        vrf = vrf_lookup_by_handle(ep_entry->vrf_handle);
        if (vrf == NULL) {
            HAL_ABORT(0);
        }

        memcpy(&ep_ip_addr, ip_addr, sizeof(ip_addr_t));
        init_ip_entry_key(ip_addr, vrf->vrf_id, &ip_entry_key);

        trans_t *other_trans = reinterpret_cast<trans_t *>(ip_ht->lookup(&ip_entry_key));


        /* Initiate delete of the transaction only if EP is different */
        if (other_trans != nullptr && (other_trans->get_ep_entry() != ep_entry)) {
            other_trans->log_error("Initiating transaction delete as part of IP move.");
            process_transaction(other_trans, other_trans->sm_->get_remove_event(), NULL);
        }

        return HAL_RET_OK;
    }


    bool transaction_completed() { return sm_->state_machine_competed(); }

    uint32_t get_current_state_timeout() {
        return this->sm_->get_current_state_timeout();
    }

    uint64_t get_timeout_remaining() {
        return this->sm_->get_timeout_remaining();
    }

    static fsm_state_machine_def_t *get_sm_def_func() { return nullptr; };

    static void set_state_timeout(fsm_state_machine_def_t *sm_def,
            uint32_t state, uint32_t timeout) {
        auto result = sm_def->find(state);
        /*  Assert, state not defined in SM */
        HAL_ABORT(result != sm_def->end());
        result->second.timeout = timeout;
    }

    ep_t* get_ep_entry() {
        ep_t *other_ep_entry = NULL;
        ep_l3_entry_t *other_ep_l3_entry;
        ep_l3_key_t l3_key = {0};

        l3_key.vrf_id = this->ip_entry_key_.vrf_id;
        memcpy(&l3_key.ip_addr, &ip_entry_key_.ip_addr, sizeof(ip_addr_t));

        /* Find the EP entry for this IP address */
        other_ep_l3_entry = reinterpret_cast<ep_l3_entry_t *>(
            g_hal_state->ep_l3_entry_ht()->lookup(&l3_key));

        if (other_ep_l3_entry != NULL) {
            other_ep_entry = find_ep_by_handle(other_ep_l3_entry->ep_hal_handle);
        }

        return other_ep_entry;
    }

    trans_ip_entry_key_t *ip_entry_key_ptr() { return &ip_entry_key_; }

    static void init_ip_entry_key(const ip_addr_t *ip_addr, vrf_id_t vrf_id,
                                  trans_ip_entry_key_t *ip_entry_key) {
        *ip_entry_key = {0};
        memcpy(&ip_entry_key->ip_addr, ip_addr, sizeof(ip_addr_t));
        ip_entry_key->vrf_id = vrf_id;
    }

    bool trans_marked_for_delete() {
        return this->marked_for_delete_;
    }

};

void *trans_get_ip_entry_key_func(void *entry);
uint32_t trans_ip_entry_key_size(void);

}
}
#endif /* HAL_PLUGINS_NETWORK_EP_LEARN_COMMON_TRANS_HPP_ */
