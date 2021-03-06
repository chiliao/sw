/******************************************************************************/
/* Tunnel lookup 1                                                            */
/******************************************************************************/
action tep1_rx_info(decap_next, src_vpc_id) {
    // if table lookup is a miss, drop
    ingress_drop(P4I_DROP_TEP1_RX_MISS);

    modify_field(tunnel_metadata.decap_next, decap_next);
    modify_field(vnic_metadata.src_vpc_id, src_vpc_id);
}

@pragma stage 0
table tep1_rx {
    reads {
        vxlan_1.vni : ternary;
    }
    actions {
        tep1_rx_info;
    }
    size : TEP1_RX_TABLE_SIZE;
}

/******************************************************************************/
/* Tunnel lookup 2                                                            */
/******************************************************************************/
action tep2_rx_info(src_vpc_id) {
    // if table lookup is a miss, drop
    ingress_drop(P4I_DROP_TEP2_RX_MISS);

    modify_field(vnic_metadata.src_vpc_id, src_vpc_id);
}

@pragma stage 1
table tep2_rx {
    reads {
        vxlan_2.vni                 : ternary;
        tunnel_metadata.tep2_dst    : ternary;
    }
    actions {
        tep2_rx_info;
    }
    size : TEP2_RX_TABLE_SIZE;
}

control tunnel_rx {
    if (vxlan_1.valid == TRUE) {
        apply(tep1_rx);
    }
    if (vxlan_2.valid == TRUE) {
        if (tunnel_metadata.decap_next == TRUE) {
            apply(tep2_rx);
        }
    }
}
