/*****************************************************************************/
/* Inter pipe : ingress pipeline                                             */
/*****************************************************************************/
action tunnel_decap() {
    remove_header(ethernet_1);
    remove_header(ctag_1);
    remove_header(ipv4_1);
    remove_header(ipv6_1);
    remove_header(udp_1);
    remove_header(vxlan_1);
    subtract(capri_p4_intrinsic.packet_len, capri_p4_intrinsic.frame_size,
             offset_metadata.l2_2);
}

action ingress_to_egress() {
    add_header(capri_p4_intrinsic);
    add_header(txdma_to_p4e);
    add_header(p4i_i2e);
    remove_header(capri_txdma_intrinsic);
    remove_header(p4plus_to_p4);
    remove_header(p4plus_to_p4_vlan);
    remove_header(ingress_recirc);
    modify_field(capri_intrinsic.tm_oport, TM_PORT_EGRESS);
}

action ingress_to_rxdma() {
    add_header(capri_p4_intrinsic);
    add_header(capri_rxdma_intrinsic);
    add_header(p4i_to_rxdma);
    remove_header(capri_txdma_intrinsic);
    remove_header(p4plus_to_p4);
    remove_header(p4plus_to_p4_vlan);
    remove_header(ingress_recirc);
    modify_field(capri_intrinsic.tm_oport, TM_PORT_DMA);
    modify_field(capri_intrinsic.lif, APULU_SERVICE_LIF);
    modify_field(capri_rxdma_intrinsic.rx_splitter_offset,
                 (CAPRI_GLOBAL_INTRINSIC_HDR_SZ +
                  CAPRI_RXDMA_INTRINSIC_HDR_SZ + APULU_P4I_TO_RXDMA_HDR_SZ));
}

action ingress_recirc() {
    add_header(capri_p4_intrinsic);
    add_header(ingress_recirc);
    remove_header(capri_txdma_intrinsic);
    remove_header(p4plus_to_p4);
    remove_header(p4plus_to_p4_vlan);
    modify_field(capri_intrinsic.tm_oport, TM_PORT_INGRESS);
}

action p4i_inter_pipe() {
    if ((ingress_recirc.flow_done == FALSE) or
        (ingress_recirc.local_mapping_done == FALSE)) {
        ingress_recirc();
        // return
    }

    if (control_metadata.tunneled_packet == TRUE) {
        tunnel_decap();
    }

    if (control_metadata.flow_miss == TRUE) {
        ingress_to_rxdma();
    } else {
        ingress_to_egress();
    }
}

@pragma stage 5
table p4i_inter_pipe {
    actions {
        p4i_inter_pipe;
    }
}

control ingress_inter_pipe {
    if (capri_intrinsic.drop == 0) {
        apply(p4i_inter_pipe);
    }
}

/*****************************************************************************/
/* Inter pipe : egress pipeline                                              */
/*****************************************************************************/
action egress_to_rxdma() {
    add_header(capri_rxdma_intrinsic);
    add_header(p4e_to_p4plus_classic_nic);
    add_header(p4e_to_p4plus_classic_nic_ip);

    modify_field(p4e_to_p4plus_classic_nic.packet_len,
                 capri_p4_intrinsic.packet_len);
    modify_field(p4e_to_p4plus_classic_nic.p4plus_app_id,
                 P4PLUS_APPTYPE_CLASSIC_NIC);
    modify_field(capri_rxdma_intrinsic.rx_splitter_offset,
                 (CAPRI_GLOBAL_INTRINSIC_HDR_SZ + CAPRI_RXDMA_INTRINSIC_HDR_SZ +
                  P4PLUS_CLASSIC_NIC_HDR_SZ));

    if (ctag_1.valid == TRUE) {
    }

    modify_field(key_metadata.sport, key_metadata.sport);
    modify_field(key_metadata.dport, key_metadata.dport);
    if (ipv4_1.valid == TRUE) {
        modify_field(key_metadata.src, ipv4_1.srcAddr);
        modify_field(key_metadata.dst, ipv4_1.dstAddr);
        if (ipv4_1.protocol == IP_PROTO_TCP) {
            modify_field(p4e_to_p4plus_classic_nic.pkt_type,
                         CLASSIC_NIC_PKT_TYPE_IPV4_TCP);
        } else {
            if (ipv4_1.protocol == IP_PROTO_UDP) {
                modify_field(p4e_to_p4plus_classic_nic.pkt_type,
                             CLASSIC_NIC_PKT_TYPE_IPV4_UDP);
            } else {
                modify_field(p4e_to_p4plus_classic_nic.pkt_type,
                             CLASSIC_NIC_PKT_TYPE_IPV4);
            }
        }
    }
    if (ipv6_1.valid == TRUE) {
        if (ipv6_1.nextHdr == IP_PROTO_TCP) {
            modify_field(p4e_to_p4plus_classic_nic.pkt_type,
                         CLASSIC_NIC_PKT_TYPE_IPV6_TCP);
        } else {
            if (ipv6_1.nextHdr == IP_PROTO_UDP) {
                modify_field(p4e_to_p4plus_classic_nic.pkt_type,
                             CLASSIC_NIC_PKT_TYPE_IPV6_UDP);
            } else {
                modify_field(p4e_to_p4plus_classic_nic.pkt_type,
                             CLASSIC_NIC_PKT_TYPE_IPV6);
            }
        }
    }
}

action egress_recirc() {
    add_header(egress_recirc);
    modify_field(capri_intrinsic.tm_oport, TM_PORT_EGRESS);
}

action p4e_inter_pipe() {
    if (capri_intrinsic.tm_oq != TM_P4_RECIRC_QUEUE) {
        modify_field(capri_intrinsic.tm_iq, capri_intrinsic.tm_oq);
    } else {
        modify_field(capri_intrinsic.tm_oq, capri_intrinsic.tm_iq);
    }

    if (egress_recirc.mapping_done == FALSE) {
        egress_recirc();
        //return
    }

    remove_header(p4e_i2e);
    remove_header(txdma_to_p4e);
    remove_header(egress_recirc);
    if (capri_intrinsic.tm_oport == TM_PORT_DMA) {
        egress_to_rxdma();
    }
}

@pragma stage 5
table p4e_inter_pipe {
    actions {
        p4e_inter_pipe;
    }
}

control egress_inter_pipe {
    if (capri_intrinsic.drop == 0) {
        apply(p4e_inter_pipe);
    }
}
