rx_hdr_transpositions:
    sne             c1, d.rx_hdr_transpositions_d.hdr2_bits, 0
    balcf           r7, [c1], rx_hdr_transpositions_layer2
    crestore        [c7-c1], d.rx_hdr_transpositions_d.hdr2_bits[6:0], 0x7F

    sne             c1, d.rx_hdr_transpositions_d.hdr3_bits, 0
    balcf           r7, [c1], rx_hdr_transpositions_layer3
    crestore        [c7-c1], d.rx_hdr_transpositions_d.hdr3_bits[6:0], 0x7F

    sne             c1, d.rx_hdr_transpositions_d.hdr1_bits, 0
    nop.!c1.e
    crestore        [c7-c1], d.rx_hdr_transpositions_d.hdr1_bits[6:0], 0x7F

rx_hdr_transpositions_layer1:
    bcf             [!c1&!c2&!c3&!c4&!c5&!c6&!c7], rx_hdr_transpositions_layer1_next
    phvwr.c1        p.ethernet_1_valid, FALSE
    phvwr.c2        p.ctag_1_valid, FALSE
    phvwr.c3        p.ipv4_1_valid, FALSE
    phvwr.c4        p.ipv6_1_valid, FALSE
    phvwr.c5        p.ethernet_1_srcAddr, d.rx_hdr_transpositions_d.ethernet_src
    phvwr.c6        p.ethernet_1_dstAddr, d.rx_hdr_transpositions_d.ethernet_dst
    phvwr.c7        p.ethernet_1_etherType, d.rx_hdr_transpositions_d.ethernet_type
rx_hdr_transpositions_layer1_next:
    crestore        [c7-c1], d.rx_hdr_transpositions_d.hdr1_bits[13:7], 0x7F
    phvwr.c1        p.ctag_1_valid, TRUE
    phvwr.c1        p.ethernet_1_etherType, ETHERTYPE_CTAG
    phvwrpair.c1    p.ctag_1_vid, d.rx_hdr_transpositions_d.ctag, \
                        p.ctag_1_etherType, d.rx_hdr_transpositions_d.ethernet_type
    bbeq            k.ipv4_1_valid, TRUE, rx_hdr_transpositions_layer1_ipv4
    bbeq            k.ipv6_1_valid, TRUE, rx_hdr_transpositions_layer1_ipv6
    nop
    nop.e
    nop

rx_hdr_transpositions_layer1_ipv4:
    phvwr.c2        p.ipv4_1_srcAddr, d.rx_hdr_transpositions_d.ip_src
    phvwr.c3        p.ipv4_1_dstAddr, d.rx_hdr_transpositions_d.ip_dst
    phvwr.c4        p.ipv4_1_diffserv, d.rx_hdr_transpositions_d.ip_dscp
    phvwr.c5        p.ipv4_1_protocol, d.rx_hdr_transpositions_d.ip_proto
    phvwr.c6        p.ipv4_1_ttl, d.rx_hdr_transpositions_d.ip_ttl
    setcf.e         c7, [c2|c3|c4|c5|c6]
    phvwr.c7        p.control_metadata_update_checksum_1, TRUE

rx_hdr_transpositions_layer1_ipv6:
    phvwr.c2        p.ipv6_1_srcAddr, d.rx_hdr_transpositions_d.ip_src
    phvwr.c3        p.ipv6_1_dstAddr, d.rx_hdr_transpositions_d.ip_dst
    phvwr.c4        p.ipv6_1_trafficClass, d.rx_hdr_transpositions_d.ip_dscp
    phvwr.c5        p.ipv6_1_nextHdr, d.rx_hdr_transpositions_d.ip_proto
    phvwr.c6        p.ipv6_1_hopLimit, d.rx_hdr_transpositions_d.ip_ttl
    setcf.e         c7, [c2|c3]
    phvwr.c7        p.control_metadata_update_checksum_1, TRUE

rx_hdr_transpositions_layer2:
    bcf             [!c1&!c2&!c3&!c4&!c5&!c6&!c7], rx_hdr_transpositions_layer2_next
    phvwr.c1        p.ethernet_2_valid, FALSE
    phvwr.c2        p.ctag_2_valid, FALSE
    phvwr.c3        p.ipv4_2_valid, FALSE
    phvwr.c4        p.ipv6_2_valid, FALSE
    phvwr.c5        p.ethernet_2_srcAddr, d.rx_hdr_transpositions_d.ethernet_src
    phvwr.c6        p.ethernet_2_dstAddr, d.rx_hdr_transpositions_d.ethernet_dst
    phvwr.c7        p.ethernet_2_etherType, d.rx_hdr_transpositions_d.ethernet_type
rx_hdr_transpositions_layer2_next:
    crestore        [c7-c1], d.rx_hdr_transpositions_d.hdr2_bits[13:7], 0x7F
    phvwr.c1        p.ctag_2_valid, TRUE
    phvwr.c1        p.ethernet_2_etherType, ETHERTYPE_CTAG
    jrcf            [!c2&!c3&!c4&!c5&!c6], r7
    phvwrpair.c1    p.ctag_2_vid, d.rx_hdr_transpositions_d.ctag, \
                        p.ctag_2_etherType, d.rx_hdr_transpositions_d.ethernet_type
    bbeq            k.ipv4_2_valid, TRUE, rx_hdr_transpositions_layer2_ipv4
    bbeq            k.ipv6_2_valid, TRUE, rx_hdr_transpositions_layer2_ipv6
    nop
    jr              r7
    nop

rx_hdr_transpositions_layer2_ipv4:
    phvwr.c2        p.ipv4_2_srcAddr, d.rx_hdr_transpositions_d.ip_src
    phvwr.c3        p.ipv4_2_dstAddr, d.rx_hdr_transpositions_d.ip_dst
    phvwr.c4        p.ipv4_2_diffserv, d.rx_hdr_transpositions_d.ip_dscp
    phvwr.c5        p.ipv4_2_protocol, d.rx_hdr_transpositions_d.ip_proto
    phvwr.c6        p.ipv4_2_ttl, d.rx_hdr_transpositions_d.ip_ttl
    setcf           c7, [c2|c3|c4|c5|c6]
    jr              r7
    phvwr.c7        p.control_metadata_update_checksum_2, TRUE

rx_hdr_transpositions_layer2_ipv6:
    phvwr.c2        p.ipv6_2_srcAddr, d.rx_hdr_transpositions_d.ip_src
    phvwr.c3        p.ipv6_2_dstAddr, d.rx_hdr_transpositions_d.ip_dst
    phvwr.c4        p.ipv6_2_trafficClass, d.rx_hdr_transpositions_d.ip_dscp
    phvwr.c5        p.ipv6_2_nextHdr, d.rx_hdr_transpositions_d.ip_proto
    phvwr.c6        p.ipv6_2_hopLimit, d.rx_hdr_transpositions_d.ip_ttl
    setcf           c7, [c2|c3]
    jr              r7
    phvwr.c7        p.control_metadata_update_checksum_2, TRUE

rx_hdr_transpositions_layer3:
    bcf             [!c1&!c2&!c3&!c4&!c5&!c6&!c7], rx_hdr_transpositions_layer3_next
    phvwr.c1        p.ethernet_3_valid, FALSE
    phvwr.c2        p.ctag_3_valid, FALSE
    phvwr.c3        p.ipv4_3_valid, FALSE
    phvwr.c4        p.ipv6_3_valid, FALSE
    phvwr.c5        p.ethernet_3_srcAddr, d.rx_hdr_transpositions_d.ethernet_src
    phvwr.c6        p.ethernet_3_dstAddr, d.rx_hdr_transpositions_d.ethernet_dst
    phvwr.c7        p.ethernet_3_etherType, d.rx_hdr_transpositions_d.ethernet_type
rx_hdr_transpositions_layer3_next:
    crestore        [c7-c1], d.rx_hdr_transpositions_d.hdr3_bits[13:7], 0x7F
    phvwr.c1        p.ctag_3_valid, TRUE
    phvwr.c1        p.ethernet_3_etherType, ETHERTYPE_CTAG
    jrcf            [!c2&!c3&!c4&!c5&!c6], r7
    phvwrpair.c1    p.ctag_3_vid, d.rx_hdr_transpositions_d.ctag, \
                        p.ctag_3_etherType, d.rx_hdr_transpositions_d.ethernet_type
    bbeq            k.ipv4_3_valid, TRUE, rx_hdr_transpositions_layer3_ipv4
    bbeq            k.ipv6_3_valid, TRUE, rx_hdr_transpositions_layer3_ipv6
    nop
    jr              r7
    nop

rx_hdr_transpositions_layer3_ipv4:
    phvwr.c2        p.ipv4_3_srcAddr, d.rx_hdr_transpositions_d.ip_src
    phvwr.c3        p.ipv4_3_dstAddr, d.rx_hdr_transpositions_d.ip_dst
    phvwr.c4        p.ipv4_3_diffserv, d.rx_hdr_transpositions_d.ip_dscp
    phvwr.c5        p.ipv4_3_protocol, d.rx_hdr_transpositions_d.ip_proto
    phvwr.c6        p.ipv4_3_ttl, d.rx_hdr_transpositions_d.ip_ttl
    setcf           c7, [c2|c3|c4|c5|c6]
    jr              r7
    phvwr.c7        p.control_metadata_update_checksum_3, TRUE

rx_hdr_transpositions_layer3_ipv6:
    phvwr.c2        p.ipv6_3_srcAddr, d.rx_hdr_transpositions_d.ip_src
    phvwr.c3        p.ipv6_3_dstAddr, d.rx_hdr_transpositions_d.ip_dst
    phvwr.c4        p.ipv6_3_trafficClass, d.rx_hdr_transpositions_d.ip_dscp
    phvwr.c5        p.ipv6_3_nextHdr, d.rx_hdr_transpositions_d.ip_proto
    phvwr.c6        p.ipv6_3_hopLimit, d.rx_hdr_transpositions_d.ip_ttl
    setcf           c7, [c2|c3]
    jr              r7
    phvwr.c7        p.control_metadata_update_checksum_3, TRUE
