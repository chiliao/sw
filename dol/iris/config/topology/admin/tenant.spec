# Configuration Template.
meta:
    id: TENANT_ADMIN

type    : tenant
overlay : vlan

segments:
    - spec  : ref://store/specs/id=SEGMENT_ADMIN
      count : 1

# NFV Endpoints:
# - They will attach to a 'TRUNK' Enic.
# - All the segments will be enabled on these Enics              
nfveps: 0

security_profile: ref://store/objects/id=SEC_PROF_ACTIVE

sessions:
    unidest:
        ipv4:
            - ref://store/specs/id=SESSION_UDP_ADMIN
        ipv6:
        mac:


lif: ref://store/specs/id=LIF_ETH
