/******************************************************************************
 * Capri Intrinsic header definitions
 *****************************************************************************/
header cap_phv_intr_global_t capri_intrinsic;
header cap_phv_intr_p4_t capri_p4_intrinsic;
header cap_phv_intr_rxdma_t capri_rxdma_intrinsic;
header cap_phv_intr_txdma_t capri_txdma_intrinsic;

/******************************************************************************
 * Headers
 *****************************************************************************/

header ingress_recirc_header_t ingress_recirc_header;
header p4i_to_p4e_header_t p4i_to_p4e_header;
header p4_to_p4plus_classic_nic_header_t p4e_to_p4plus_classic_nic;
header p4_to_p4plus_ip_addr_t p4e_to_p4plus_classic_nic_ip;
header p4plus_to_p4_s1_t p4plus_to_p4;
@pragma pa_header_union ingress ctag_1
header p4plus_to_p4_s2_t p4plus_to_p4_vlan;
header predicate_header_t predicate_header; /* TODO: To remove */

// layer 0
header ethernet_t ethernet_0;
header vlan_tag_t ctag_0;
@pragma pa_header_union egress ipv4_0 ipv6_0
header ipv4_t ipv4_0;
header ipv6_t ipv6_0;
@pragma pa_header_union egress udp_0 gre_0
header udp_t udp_0;
header gre_t gre_0;
header mpls_t mpls_label1_0;
header mpls_t mpls_label2_0;
header mpls_t mpls_label3_0;

// layer 1
header ethernet_t ethernet_1;
header vlan_tag_t ctag_1;
@pragma pa_header_union xgress ipv4_1 ipv6_1
header ipv4_t ipv4_1;
header ipv6_t ipv6_1;
@pragma pa_header_union xgress udp_1 gre_1
header udp_t udp_1;
header gre_t gre_1;
header mpls_t mpls_label1_1;
header mpls_t mpls_label2_1;
header mpls_t mpls_label3_1;

// layer 2
header ethernet_t ethernet_2;
header vlan_tag_t ctag_2;
@pragma pa_header_union xgress ipv4_2 ipv6_2
header ipv4_t ipv4_2;
header ipv6_t ipv6_2;
@pragma pa_header_union xgress udp_2 tcp icmp
header udp_t udp_2;
@pragma no_ohi ingress
header tcp_t tcp;
header icmp_t icmp;

// TCP options
@pragma hdr_len parser_metadata.parse_tcp_counter
header tcp_options_blob_t tcp_options_blob;
header tcp_option_unknown_t tcp_option_unknown;
header tcp_option_eol_t tcp_option_eol;
header tcp_option_nop_t tcp_option_nop;
header tcp_option_nop_t tcp_option_nop_1;
header tcp_option_mss_t tcp_option_mss;
header tcp_option_ws_t tcp_option_ws;
header tcp_option_sack_perm_t tcp_option_sack_perm;
@pragma pa_header_union ingress tcp_option_two_sack tcp_option_three_sack tcp_option_four_sack
@pragma no_ohi ingress
header tcp_option_one_sack_t tcp_option_one_sack;
@pragma no_ohi ingress
header tcp_option_two_sack_t tcp_option_two_sack;
@pragma no_ohi ingress
header tcp_option_three_sack_t tcp_option_three_sack;
@pragma no_ohi ingress
header tcp_option_four_sack_t tcp_option_four_sack;
@pragma no_ohi ingress
header tcp_option_timestamp_t tcp_option_timestamp;

/******************************************************************************
 * Parser metadata
 *****************************************************************************/
header_type parser_metadata_t {
    fields {
        parse_tcp_counter   : 8;
    }
}
@pragma pa_parser_local
metadata parser_metadata_t parser_metadata;

/******************************************************************************
 * Parser OHI
 *****************************************************************************/
header_type parser_ohi_t {
    fields {
        ipv4_1_len          : 16;
        ipv4_2_len          : 16;
        l4_1_len            : 16;
        l4_2_len            : 16;
        ipv4_1___start_off  : 16;
        ipv6_1___start_off  : 16;
        ipv4_2___start_off  : 16;
        ipv6_2___start_off  : 16;
        tcp___start_off     : 16;
        udp_1___start_off   : 16;
        udp_2___start_off   : 16;
    }
}
@pragma parser_write_only
@pragma parser_share_ohi ipv4_1___start_off ipv6_1___start_off
@pragma parser_share_ohi ipv4_2___start_off ipv6_2___start_off
metadata parser_ohi_t ohi;


/******************************************************************************
 * Parser start
 *****************************************************************************/
parser start {
    extract(capri_intrinsic);
    return select(capri_intrinsic.tm_iport) {
        TM_PORT_DMA     : parse_txdma_to_ingress;
        TM_PORT_INGRESS : parse_ingress_recirc_header;
        default : parse_uplink;
        0x1 mask 0 : deparse_ingress;
        0x2 mask 0 : egress_start;
    }
}

@pragma xgress ingress
parser parse_txdma_to_ingress {
    extract(capri_txdma_intrinsic);
    return select(current(72, 4)) { /* p4plus_to_p4_classic_header_t: p4plus_app_id*/
        0 : parse_ingress_pass2; /* FIXME: Not a scenario ? */
        default : parse_txdma_app;
    }
}

@pragma xgress ingress
parser parse_txdma_app {
    extract(p4plus_to_p4);
    extract(p4plus_to_p4_vlan);
    return select(p4plus_to_p4.p4plus_app_id) {
        P4PLUS_APPTYPE_CPU : parse_cpu_packet;
        default : parse_packet_from_host; /* TODO: Not a scenario for OCI */
    }
}

@pragma xgress ingress
parser parse_cpu_packet {
    return parse_packet; /* TODO: Assumption made that the direction information is derived from the NACL */
}

@pragma xgress ingress
parser parse_uplink {
    return select(capri_intrinsic.tm_iport) {
        TM_PORT_UPLINK_0 : parse_packet_from_host;
        TM_PORT_UPLINK_1 : parse_packet_from_switch;
        default : parse_packet;
    }
}

@pragma xgress ingress
parser parse_ingress_recirc_header {
    extract(capri_p4_intrinsic);
    extract(ingress_recirc_header);
    return parse_ingress_pass2;
}

@pragma xgress ingress
parser parse_ingress_pass2 {
    extract(predicate_header);
    return select(predicate_header.direction) {
        RX_FROM_SWITCH : parse_packet_from_switch;
        TX_FROM_HOST : parse_packet_from_host;
        default : ingress;
    }
}


@pragma xgress ingress
parser parse_packet_from_host {
    set_metadata(control_metadata.direction, TX_FROM_HOST);
    return parse_packet;
}

@pragma xgress ingress
parser parse_packet_from_switch {
    set_metadata(control_metadata.direction, RX_FROM_SWITCH);
    return parse_packet;
}

/******************************************************************************
 * Layer 1
 *****************************************************************************/
@pragma allow_set_meta offset_metadata.l2_1
parser parse_packet {
    extract(ethernet_1);
    set_metadata(offset_metadata.l2_1, current + 0);
    return select(latest.etherType) {
        ETHERTYPE_CTAG : parse_ctag_1;
        ETHERTYPE_IPV4 : parse_ipv4_1;
        ETHERTYPE_IPV6 : parse_ipv6_1;
        default: ingress;
    }
}

parser parse_ctag_1 {
    extract(ctag_1);
    return select(latest.etherType) {
        ETHERTYPE_IPV4 : parse_ipv4_1;
        ETHERTYPE_IPV6 : parse_ipv6_1;
        default: ingress;
    }
}

parser parse_ipv4_1 {
    set_metadata(offset_metadata.l3_1, current + 0);
    return parse_ipv4_1_split;
}

parser parse_ipv4_1_split {
    extract(ipv4_1);
    set_metadata(ohi.ipv4_1_len, ipv4_1.ihl << 2);
    set_metadata(ohi.l4_1_len, ipv4_1.totalLen - (ipv4_1.ihl << 2));
    return select(latest.fragOffset, latest.protocol) {
        IP_PROTO_ICMP : parse_icmp;
        IP_PROTO_TCP : parse_tcp;
        IP_PROTO_UDP : parse_udp_1;
        default: ingress;
    }
}

parser parse_ipv4_in_ip_1 {
    return parse_ipv4_2;
}

parser parse_ipv6_in_ip_1 {
    return parse_ipv6_2;
}

parser parse_ipv6_1 {
    extract(ipv6_1);
    set_metadata(offset_metadata.l3_1, current + 0);
    set_metadata(ohi.l4_1_len, ipv6_1.payloadLen + 0);
    return select(latest.nextHdr) {
        IP_PROTO_ICMPV6 : parse_icmp;
        IP_PROTO_TCP : parse_tcp;
        IP_PROTO_UDP : parse_udp_1;
        IP_PROTO_IPV4 : parse_ipv4_in_ip_1;
        IP_PROTO_IPV6 : parse_ipv6_in_ip_1;
        IP_PROTO_GRE  : parse_gre_1;
        default : ingress;
    }
}

parser parse_icmp {
    extract(icmp);
    set_metadata(offset_metadata.l4_2, current + 0);
    return ingress;
}

parser parse_tcp {
    extract(tcp);
    set_metadata(offset_metadata.l4_2, current + 0);
    set_metadata(key_metadata.sport, latest.srcPort);
    set_metadata(key_metadata.dport, latest.dstPort);
    set_metadata(parser_metadata.parse_tcp_counter, (tcp.dataOffset << 2) - 20);
    return select(parser_metadata.parse_tcp_counter) {
        0 : ingress;
        0x80 mask 0x80: parse_tcp_option_error;
        default : parse_tcp_options_blob;
    }
}

@pragma dont_advance_packet
@pragma capture_payload_offset
parser parse_tcp_options_blob {
    set_metadata(parser_metadata.parse_tcp_counter, parser_metadata.parse_tcp_counter + 0);
    extract(tcp_options_blob);
    return parse_tcp_options;
}

parser parse_tcp_option_one_sack {
    extract(tcp_option_one_sack);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 10);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_two_sack {
    extract(tcp_option_two_sack);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 18);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_three_sack {
    extract(tcp_option_three_sack);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 26);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_four_sack {
    extract(tcp_option_four_sack);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 34);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_EOL {
    extract(tcp_option_eol);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 1);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_NOP {
    extract(tcp_option_nop);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 1);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_mss {
    extract(tcp_option_mss);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 4);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_ws {
    extract(tcp_option_ws);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 3);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_option_sack_perm {
    extract(tcp_option_sack_perm);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 2);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_timestamp {
    extract(tcp_option_timestamp);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - 10);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

parser parse_tcp_unknown_option {
    extract(tcp_option_unknown);
    set_metadata(parser_metadata.parse_tcp_counter,
                 parser_metadata.parse_tcp_counter - tcp_option_unknown.optLength);
    return select(parser_metadata.parse_tcp_counter) {
        0x00 : ingress;
        default : parse_tcp_options;
    }
}

@pragma dont_capture_payload_offset
parser parse_tcp_option_error {
    set_metadata(control_metadata.parse_tcp_option_error, 1);
    return ingress;
}

@pragma dont_capture_payload_offset
parser parse_tcp_option_error2 {
    set_metadata(control_metadata.parse_tcp_option_error, 1);
    return ingress;
}

@pragma header_ordering tcp_option_mss tcp_option_ws tcp_option_sack_perm tcp_option_one_sack tcp_option_two_sack tcp_option_three_sack tcp_option_four_sack tcp_option_timestamp tcp_option_unknown tcp_option_nop tcp_option_nop_1 tcp_option_eol
parser parse_tcp_options {
    return select(parser_metadata.parse_tcp_counter, current(0, 8)) {
        0x0000 mask 0xff00 : ingress;
        0x8000 mask 0x8000 : parse_tcp_option_error;
        0x0000 mask 0x00ff : parse_tcp_option_EOL;
        0x0001 mask 0x00ff : parse_tcp_option_NOP;
        default : parse_tcp_multibyte_options;
    }
}

parser parse_tcp_multibyte_options {
    return select(current(0, 16)) {
        0x0204 mask 0xffff : parse_tcp_option_mss;
        0x0200 mask 0xff00 : parse_tcp_option_error2;
        0x0303 mask 0xffff : parse_tcp_option_ws;
        0x0300 mask 0xff00 : parse_tcp_option_error2;
        0x0402 mask 0xffff : parse_tcp_option_sack_perm;
        0x0400 mask 0xff00 : parse_tcp_option_error2;
        0x050a mask 0xffff : parse_tcp_option_one_sack;
        0x0512 mask 0xffff : parse_tcp_option_two_sack;
        0x051a mask 0xffff : parse_tcp_option_three_sack;
        0x0522 mask 0xffff : parse_tcp_option_four_sack;
        0x0500 mask 0xff00 : parse_tcp_option_error2;
        0x080a mask 0xffff : parse_tcp_timestamp;
        0x0800 mask 0xff00 : parse_tcp_option_error2;
        0x0000 mask 0x00ff : parse_tcp_option_error2;
        default: parse_tcp_unknown_option;
    }
}


parser parse_udp_1 {
    extract(udp_1);
    set_metadata(offset_metadata.l4_1, current + 0);
    set_metadata(ohi.l4_1_len, udp_1.len + 0);
    return select(latest.dstPort) {
        UDP_PORT_MPLS : parse_mpls_label1_1;
        default: ingress;
    }
}

parser parse_gre_1 {
    extract(gre_1);
#ifdef IPV6_SUPPORT
    set_metadata(tunnel_metadata.tunnel_type_1, INGRESS_TUNNEL_TYPE_GRE);
#endif
    return select(latest.C, latest.R, latest.K, latest.S, latest.s,
                  latest.recurse, latest.flags, latest.ver, latest.proto) {
        ETHERTYPE_MPLS_UNICAST : parse_mpls_label1_1;
        /* FIXME: Handling of unsupported protocols */
    }
}


parser parse_mpls_label1_1 {
    extract(mpls_label1_1);
    return select(latest.bos) {
        0 : parse_mpls_label2_1;
//      1 : error!
    }
}

parser parse_mpls_label2_1 {
    extract(mpls_label2_1);
    return select(latest.bos) {
        0 : parse_mpls_label3_1;
        1 : parse_mpls_payload;
    }
}

parser parse_mpls_label3_1 {
    extract(mpls_label3_1);
    return select(latest.bos) {
//      0 : error ! ;
        1 : parse_mpls_payload;
    }
}

/*
 *  FIXME: Peeking into IP version to determine the next layer in MPLS.
 *  This should be replaced with Oracle's encoding of the layer type in the
 *  MPLS label
 */
parser parse_mpls_payload {
    return select(current(0, 4)) {
        0x4 : parse_mpls_inner_ipv4;
        0x6 : parse_mpls_inner_ipv6;
        default: parse_eompls;
    }
}

parser parse_eompls {
    return parse_ethernet_2;
}

parser parse_mpls_inner_ipv4 {
    return parse_ipv4_2;
}

parser parse_mpls_inner_ipv6 {
    return parse_ipv6_2;
}

/******************************************************************************
 * Layer 2
 *****************************************************************************/
@pragma allow_set_meta offset_metadata.l2_2
parser parse_ethernet_2 {
    set_metadata(offset_metadata.l2_2, current + 0);
    extract(ethernet_2);
    return select(latest.etherType) {
        ETHERTYPE_CTAG : parse_ctag_2;
        ETHERTYPE_IPV4 : parse_ipv4_2;
        ETHERTYPE_IPV6 : parse_ipv6_2;
        default: ingress;
    }
}

parser parse_ctag_2 {
    extract(ctag_2);
    return select(latest.etherType) {
        ETHERTYPE_IPV4 : parse_ipv4_2;
        ETHERTYPE_IPV6 : parse_ipv6_2;
        default: ingress;
    }
}

@pragma allow_set_meta offset_metadata.l3_2
parser parse_ipv4_2 {
    set_metadata(offset_metadata.l3_2, current + 0);
    return parse_ipv4_2_split;
}

parser parse_ipv4_2_split {
    extract(ipv4_2);
    set_metadata(ohi.ipv4_2_len, ipv4_2.ihl << 2);
    set_metadata(ohi.l4_2_len, ipv4_2.totalLen - (ipv4_2.ihl << 2));
    return select(latest.fragOffset, latest.protocol) {
        IP_PROTO_ICMP : parse_icmp;
        IP_PROTO_TCP : parse_tcp;
        IP_PROTO_UDP : parse_udp_2;
        default: ingress;
    }
}

@pragma allow_set_meta offset_metadata.l3_2
parser parse_ipv6_2 {
    extract(ipv6_2);
    set_metadata(offset_metadata.l3_2, current + 0);
    //set_metadata(key_metadata.src, latest.srcAddr);
    set_metadata(ohi.l4_2_len, ipv6_2.payloadLen + 0);
    return select(latest.nextHdr) {
        IP_PROTO_ICMPV6 : parse_icmp;
        IP_PROTO_TCP : parse_tcp;
        IP_PROTO_UDP : parse_udp_2;
        default : ingress;
    }
}

parser parse_udp_2 {
    extract(udp_2);
    set_metadata(offset_metadata.l4_2, current + 0);
    set_metadata(ohi.l4_2_len, udp_2.len + 0);
    set_metadata(key_metadata.sport, latest.srcPort);
    set_metadata(key_metadata.dport, latest.dstPort);
    return ingress;
}

/******************************************************************************/
/* Ingress deparser                                                           */
/******************************************************************************/
@pragma deparse_only
@pragma xgress ingress
parser deparse_ingress {
    // intrinsic headers
    extract(capri_intrinsic);
    extract(capri_p4_intrinsic);

    // Recirc only 
    extract(ingress_recirc_header);

    // P4E only
    extract(p4i_to_p4e_header);
    return parse_packet;
}

/******************************************************************************/
/* Egress parser                                                              */
/******************************************************************************/
/*
 * All data pkts entering p4e from p4i
 */
@pragma xgress egress
parser egress_start {
    extract(capri_intrinsic);
    extract(capri_p4_intrinsic);
    return select(capri_intrinsic.tm_iport) {
        default         : parse_egress;
        0x1 mask 0      : deparse_egress;
    }
}

@pragma xgress egress
parser parse_egress {
    extract(p4i_to_p4e_header);
    return select(p4i_to_p4e_header.flow_miss) {
        TRUE: parse_packet_flow_miss;
        FALSE: parse_packet;
    }
}

@pragma xgress egress
parser parse_packet_flow_miss {
    set_metadata(control_metadata.flow_miss, TRUE);
    return parse_packet;
}

/******************************************************************************/
/* Egress deparser                                                            */
/******************************************************************************/

@pragma deparse_only
@pragma xgress egress
parser deparse_egress {
    extract(capri_intrinsic);
    extract(capri_p4_intrinsic);

    // Redirect to ARM
    extract(capri_rxdma_intrinsic);

    // Packet to classic NIC application
    extract(p4e_to_p4plus_classic_nic);
    extract(p4e_to_p4plus_classic_nic_ip);


    // Packet to uplink - "Push" header - layer 0
    extract(ethernet_0);
    extract(ctag_0);
    extract(ipv4_0);
    extract(ipv6_0);
    extract(udp_0);
    extract(gre_0);
    extract(mpls_label1_0);
    extract(mpls_label2_0);

    return parse_packet;
}


/******************************************************************************
 * Checksums : Layer 0 (compute only, no verification)
 *****************************************************************************/