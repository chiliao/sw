#include "ingress.h"
#include "INGRESS_p.h"
#include "../../p4/nw/include/defines.h"

struct drop_stats_k k;
struct drop_stats_d d;
struct phv_         p;

%%

nop:
  nop.e
  nop

.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
drop_stats:
  add         r6, d.u.drop_stats_d.drop_pkts, 1
  beqi        r6, 0xFFFF, drop_stats_overflow
  phvwr       p.capri_intrinsic_tm_span_session, d.u.drop_stats_d.mirror_session_id
  tblwr.e     d.u.drop_stats_d.drop_pkts, r6
  nop

drop_stats_overflow:
  add         r5, r5, d.u.drop_stats_d.stats_idx, 3
  memwr.d.e   r5, r6
  tblwr       d.u.drop_stats_d.drop_pkts, r0
