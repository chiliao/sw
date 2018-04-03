# Flow generation configuration template.

meta:
    id: VXLANTOPO_SESSION_UDP

proto: udp

entries:
    - entry:
        label: networking
        responder:
            port : const/1
        initiator: 
            port : const/4

    - entry:
        label: gft_drop
        tracking: False
        timestamp: False
        responder: 
            port : const/12345
            flow_info:
                action  : drop
        initiator: 
            port : const/54321
            flow_info:
                action  : drop
