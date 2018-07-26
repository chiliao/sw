/******************************************************************************/
/* Key derivation tables                                                      */
/******************************************************************************/
action native_ipv4_packet() {
    modify_field(key_metadata.ktype, KEY_TYPE_IPV4);
    modify_field(key_metadata.src, ipv4_1.srcAddr);
    modify_field(key_metadata.dst, ipv4_1.dstAddr);
    modify_field(key_metadata.proto, ipv4_1.protocol);
    if (udp_1.valid == TRUE) {
        modify_field(key_metadata.sport, udp_1.srcPort);
        modify_field(key_metadata.dport, udp_1.dstPort);
    }
    modify_field(control_metadata.mapping_lkp_addr, ipv4_1.srcAddr);
    modify_field(slacl_metadata.ip_15_00, ipv4_1.dstAddr);
    modify_field(slacl_metadata.ip_31_16, (ipv4_1.dstAddr >> 16) & 0xFFFF);
}

action native_ipv6_packet() {
    modify_field(key_metadata.ktype, KEY_TYPE_IPV6);
    modify_field(key_metadata.src, ipv6_1.srcAddr);
    modify_field(key_metadata.dst, ipv6_1.dstAddr);
    modify_field(key_metadata.proto, ipv6_1.nextHdr);
    if (udp_1.valid == TRUE) {
        modify_field(key_metadata.sport, udp_1.srcPort);
        modify_field(key_metadata.dport, udp_1.dstPort);
    }
    modify_field(control_metadata.mapping_lkp_addr, ipv6_1.srcAddr);
}

action native_nonip_packet() {
    modify_field(key_metadata.ktype, KEY_TYPE_MAC);
    modify_field(key_metadata.src, ethernet_1.srcAddr);
    modify_field(key_metadata.dst, ethernet_1.dstAddr);
    if (ctag_1.valid == TRUE) {
        modify_field(key_metadata.dport, ctag_1.etherType);
    } else {
        modify_field(key_metadata.dport, ethernet_1.etherType);
    }
}

action tunneled_ipv4_packet() {
    modify_field(key_metadata.ktype, KEY_TYPE_IPV4);
    modify_field(key_metadata.src, ipv4_2.srcAddr);
    modify_field(key_metadata.dst, ipv4_2.dstAddr);
    modify_field(key_metadata.proto, ipv4_2.protocol);
    modify_field(slacl_metadata.ip_15_00, ipv4_2.srcAddr);
    modify_field(slacl_metadata.ip_31_16, (ipv4_2.srcAddr >> 16) & 0xFFFF);
    modify_field(control_metadata.mapping_lkp_addr, ipv4_2.dstAddr);
}

action tunneled_ipv6_packet() {
    modify_field(key_metadata.ktype, KEY_TYPE_IPV6);
    modify_field(key_metadata.src, ipv6_2.srcAddr);
    modify_field(key_metadata.dst, ipv6_2.dstAddr);
    modify_field(key_metadata.proto, ipv6_2.nextHdr);
    modify_field(control_metadata.mapping_lkp_addr, ipv6_2.dstAddr);
}

action tunneled_nonip_packet() {
    modify_field(key_metadata.ktype, KEY_TYPE_MAC);
    modify_field(key_metadata.src, ethernet_2.srcAddr);
    modify_field(key_metadata.dst, ethernet_2.dstAddr);
    modify_field(key_metadata.dport, ethernet_2.etherType);
}

@pragma stage 0
table key_native {
    reads {
        tunnel_metadata.tunnel_type : ternary;
        ipv4_1.valid                : ternary;
        ipv6_1.valid                : ternary;
        ipv4_2.valid                : ternary;
        ipv6_2.valid                : ternary;
    }
    actions {
        nop;
        native_ipv4_packet;
        native_ipv6_packet;
        native_nonip_packet;
    }
    size : KEY_MAPPING_TABLE_SIZE;
}

@pragma stage 0
table key_tunneled {
    reads {
        tunnel_metadata.tunnel_type : ternary;
        ipv4_1.valid                : ternary;
        ipv6_1.valid                : ternary;
        ipv4_2.valid                : ternary;
        ipv6_2.valid                : ternary;
    }
    actions {
        nop;
        tunneled_ipv4_packet;
        tunneled_ipv6_packet;
        tunneled_nonip_packet;
    }
    size : KEY_MAPPING_TABLE_SIZE;
}

action init_config() {
    if (service_header.valid == TRUE) {
        modify_field(control_metadata.local_ip_mapping_ohash_lkp,
                     ~service_header.local_ip_mapping_done);
        modify_field(control_metadata.flow_ohash_lkp,
                     ~service_header.flow_done);
        modify_field(rewrite_metadata.nexthop_index,
                     service_header.nexthop_index);
    }
    modify_field(lpm_metadata.addr, lpm_metadata.addr + key_metadata.dst);

    modify_field(scratch_metadata.addr, (slacl_metadata.ip_31_16 / 51) << 6);
    add(slacl_metadata.addr1, slacl_metadata.addr1, scratch_metadata.addr);
    modify_field(scratch_metadata.addr, (slacl_metadata.ip_15_00 / 51) << 6);
    add(slacl_metadata.addr2, slacl_metadata.addr2, scratch_metadata.addr);
}

@pragma stage 1
table init_config {
    actions {
        init_config;
    }
}

control key_init {
    apply(key_native);
    apply(key_tunneled);
    apply(init_config);
}
