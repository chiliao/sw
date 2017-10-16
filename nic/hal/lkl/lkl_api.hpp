#ifndef __LKL_API_HPP__
#define __LKL_API_HPP__
#include "nic/include/cpupkt_headers.hpp"
#include "nic/hal/src/session.hpp"

namespace hal {
namespace pd {

int lkl_init(void);
void* lkl_alloc_skbuff(const p4_to_p4plus_cpu_pkt_t* rxhdr, const uint8_t* pkt, size_t pkt_len, hal::flow_direction_t);
bool lkl_handle_flow_miss_pkt(void* skb, hal::flow_direction_t dir, uint32_t iqid, uint32_t rqid, const p4_to_p4plus_cpu_pkt_t *rxhdr);
bool lkl_handle_flow_hit_pkt(void* skb, hal::flow_direction_t dir, const p4_to_p4plus_cpu_pkt_t* rxhdr);
uint32_t lkl_get_tcpcb_rcv_nxt(void *tcpcb);
uint32_t lkl_get_tcpcb_snd_nxt(void *tcpcb);
uint32_t lkl_get_tcpcb_snd_una(void *tcpcb);
uint32_t lkl_get_tcpcb_rcv_tsval(void *tcpcb);
uint32_t lkl_get_tcpcb_ts_recent(void *tcpcb);
uint32_t lkl_get_tcpcb_state(void *tcpcb);
}
}
#endif //__LKL_API_HPP__
