# Segment Configuration Spec
meta:
    id: SEGMENT_VXLAN_TELEMETRY

type        : tenant
fabencap    : vxlan
native      : False
broadcast   : drop
multicast   : drop
l4lb        : False
endpoints   :
    sgenable: True
    useg    : 0
    pvlan   : 2
    direct  : 0
    remote  : 4 # Remote TEPs
