# Segment Configuration Spec
meta:
    id: SEGMENT_INFRA_PROXY

type        : infra
native      : False
fabencap    : vlan
broadcast   : flood 
multicast   : drop
l4lb        : False
endpoints   :
    useg    : 0
    pvlan   : 0
    direct  : 0
    remote  : 2 # Remote TEPs
