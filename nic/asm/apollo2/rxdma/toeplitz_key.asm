#include "apollo_rxdma.h"
#include "../../../p4/apollo2/include/defines.h"
#include "INGRESS_p.h"
#include "ingress.h"
#include "INGRESS_toeplitz_key_k.h"

#include "capri-macros.h"
#include "capri_common.h"

struct toeplitz_key_k_      k;
struct phv_                 p;

%%

toeplitz_key_init:
    // copy key fields from packet
    // key (320bits) is constructed from msb to lsb as -
    // key0: flow_src
    // key1: flow_dst
    // key2: sport, dport, proto (whichever order provided by key-maker)
    phvwr       p.toeplitz_key0_data, k.p4_to_rxdma_header_flow_src
    phvwr       p.toeplitz_key1_data[127:112], k.p4_to_rxdma_header_flow_dst_s0_e15
    phvwr.e     p.toeplitz_key1_data[111:0], k.p4_to_rxdma_header_flow_dst_s16_e127
    // XXX assert to ensure that flow_proto to s_port are contiguous in k
    phvwr       p.toeplitz_key2_data[63:24], k.{p4_to_rxdma_header_flow_proto...p4_to_rxdma_header_flow_sport}
