// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#ifndef __HAL_CONTROL_HPP__
#define __HAL_CONTROL_HPP__

#include "nic/include/hal_cfg.hpp"

namespace hal {
namespace pd {

#define HAL_CONTROL_Q_SIZE                           128
#define HAL_CONTROL_OPERATION_PORT_TIMER             0
#define HAL_CONTROL_OPERATION_PORT_ENABLE            1
#define HAL_CONTROL_OPERATION_PORT_DISABLE           2

// starting point for control thread
void *hal_control_start(void *ctxt);

//------------------------------------------------------------------------------
// hal-control thread notification by other threads
//------------------------------------------------------------------------------
hal_ret_t
hal_control_notify (uint8_t operation, void *ctxt);

//------------------------------------------------------------------------------
// hal control thread operation entry.
// one such entry is added to the queue for every operation
//------------------------------------------------------------------------------
typedef struct hal_ctrl_entry_ {
    uint8_t           opn;     // operation requested to perform
    std::atomic<bool> done;    // TRUE if thread performed operation
    hal_ret_t         status;  // result status of operation requested
    void              *data;   // data passed by called
} hal_ctrl_entry_t;

//------------------------------------------------------------------------------
// hal control thread maintains one queue per HAL thread to serve
// operations by HAL thread, thus avoiding locking altogether
//------------------------------------------------------------------------------
typedef struct hal_ctrl_queue_s {
    std::atomic<uint32_t> nentries;    // no. of entries in the queue
    uint16_t              pindx;       // producer index
    uint16_t              cindx;       // consumer index
    hal_ctrl_entry_t      entries[HAL_CONTROL_Q_SIZE];    // entries
} hal_ctrl_queue_t;

// per producer read/write request queues
extern hal_ctrl_queue_t g_hal_ctrl_workq[HAL_THREAD_ID_MAX];

}    // namespace pd
}    // namespace hal

#endif  // __HAL_CONTROL_HPP__
