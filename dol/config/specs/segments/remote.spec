# Segment Configuration Spec
meta:
    id: SEGMENT_REMOTE

type        : tenant
native      : False
broadcast   : drop
multicast   : drop
l4lb        : True
endpoints   :
    useg    : 0
    pvlan   : 1
    direct  : 0
    remote  : 1 # Remote TEPs
    pd      : ref://store/specs/id=RDMA_PD_DEFAULT
    slab    : ref://store/specs/id=RDMA_SLAB_DEFAULT
