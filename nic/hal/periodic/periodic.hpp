//------------------------------------------------------------------------------
// periodic thread manages few timer wheels and can carry our periodic tasks
// as background activities. Few examples include:
// - delay deleting memory resources to slabs or heap
// - flow table scan to age out sessions or detect dead flows
// - periodic stats collection and/or aggregation
//------------------------------------------------------------------------------
#ifndef __HAL_PERIODIC_HPP__
#define __HAL_PERIODIC_HPP__

#include <base.h>
#include <thread.hpp>
#include <twheel.hpp>
#include <hal_mem.hpp>

namespace hal {

using utils::thread;
extern thread_local thread *t_curr_thread;

namespace periodic {

void *periodic_thread_start(void *ctxt);

// API to delay delete any slab objects
hal_ret_t delay_delete_to_slab(hal_slab_t slab_id,             // slab to free back to
                               void *elem);                    // element to free back

// callback invoked for delay deleting slab elements back to their respective
// blocks
void slab_delay_delete_cb(hal_slab_t slab_id, void *ctxt);

}    // namespace periodic
}    // namespace hal

#endif    // __HAL_PERIODIC_HPP__

