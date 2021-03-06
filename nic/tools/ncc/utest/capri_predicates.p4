// Test a small parse graph, no loops, no tlv parsing
// used capri intrinsic headers

#include "includes/capri_intrinsic.p4"

header_type ethernet_t {
    fields {
        dstAddr : 48;
        srcAddr : 48;
        etherType : 16;
    }
}

header_type ipv4_t {
    fields {
        version : 4;
        ihl : 4;
        diffserv : 8;
        totalLen : 16;
        identification : 16;
        flags : 3;
        fragOffset : 13;
        ttl : 8;
        protocol : 8;
        hdrChecksum : 16;
        srcAddr : 32;
        dstAddr: 32;
    }
}

header_type tcp_t {
    fields {
        srcPort : 16;
        dstPort : 16;
        seqNo : 32;
        ackNo : 32;
        dataOffset : 4;
        res : 4;
        flags : 8;
        window : 16;
        checksum : 16;
        urgentPtr : 16;
    }
}

header_type udp_t {
    fields {
        srcPort : 16;
        dstPort : 16;
        len : 16;
        checksum : 16;
    }
}

header_type ipv6_t {
    fields {
        version : 4;
        trafficClass : 8;
        flowLabel : 20;
        payloadLen : 16;
        nextHdr : 8;
        hopLimit : 8;
        srcAddr : 128;
        dstAddr : 128;
    }
}

header_type control_meta_t {
    fields {
        en_B : 2;
        en_C : 1;
        en_D : 1;
    }
}

// Capri specific headers
header cap_phv_intr_global_t capri_intrinsic;                                                       
metadata cap_phv_intr_p4_t capri_p4_intrinsic;                                                      

metadata control_meta_t control_meta;

header ethernet_t ethernet;
//@pragma pa_header_union xgress ipv6
header ipv4_t ipv4;
header ipv6_t ipv6;
header tcp_t tcp;
header udp_t udp;

#define ETHERTYPE_VLAN         0x8100
#define ETHERTYPE_IPV4         0x0800
#define ETHERTYPE_IPV6         0x86dd

parser start {
    extract(capri_intrinsic);
    return select(current(0, 4)) {
        default : parse_ethernet;
    }
}

parser parse_ethernet {
    extract(ethernet);
    return select(latest.etherType) {
        ETHERTYPE_IPV4: parse_ipv4;
        ETHERTYPE_IPV6: parse_ipv6;
        default : ingress;
    }
}

#define IP_PROTOCOLS_TCP               6
#define IP_PROTOCOLS_UDP               17

parser parse_ipv4 {
    extract(ipv4);
    return select(latest.fragOffset, latest.protocol) {
        IP_PROTOCOLS_TCP : parse_tcp;
        IP_PROTOCOLS_UDP : parse_udp;
        default: ingress;
    }
}

parser parse_ipv6 {
    extract(ipv6);
    return select(latest.nextHdr) {
        IP_PROTOCOLS_TCP : parse_tcp;
        IP_PROTOCOLS_UDP : parse_udp;
        default: ingress;
    }
}

parser parse_tcp {
    extract(tcp);
    return ingress;
}

parser parse_udp {
    extract(udp);
    return ingress;
}

action nop() {
    modify_field(capri_p4_intrinsic.frame_size, 1024);    // dummy instr to force
}

table A_hash {
    reads {
        //ethernet.valid      : exact;
        ethernet.dstAddr    : exact;
        //udp.valid           : exact;
        udp.dstPort         : exact;
    }
    actions {
        nop;
    }
    size : 64;
}

action update_ipv4() {
    modify_field(ipv4.identification, 0x900d);
}

action update_ipv6() {
    modify_field(ipv6.flowLabel, 0x0cafe);
}

table B_tcam {
    reads {
        ipv4.valid      : ternary;
        ipv6.valid      : ternary;
        ipv4.dstAddr    : ternary;
        ipv6.dstAddr    : ternary;
    }
    actions {
        update_ipv4;
        update_ipv6;
    }
    size : 16;
}

action update_ttl(ttl_val) {
    modify_field(ipv4.ttl, ttl_val);
}

table C_idx {
    reads {
        ipv4.ttl     : exact;
    }
    actions {
        update_ttl;
    }
    size : 256;
}

action mpu_only_action() {
    modify_field(control_meta.en_C, control_meta.en_D); // causes 1 bit I
}

table D_mpu_only {
    actions {
        mpu_only_action;
    }
}    

control ingress {
    if (valid(ipv4)) {
        apply(D_mpu_only);
    }
    if (control_meta.en_C == 1 and control_meta.en_D == 0) {
        apply(C_idx);
    }
    if ((control_meta.en_B == 0x3) and valid(ipv4)) {
        apply(B_tcam);
    } else {
        apply(A_hash);
    }
}

control egress {
}
