/******************************************************************************/
/* Local IP mappings                                                          */
/******************************************************************************/
action local_ip_mapping_info(vnic_id, vpc_id, service_tag, pa_or_ca_xlate_idx,
                             public_xlate_idx, ca6_xlate_idx) {
    // save the original vpc id
    modify_field(vnic_metadata.vpc_id2, vnic_metadata.vpc_id);
    // if table lookup is a miss, return

    if (vnic_id != 0) {
        modify_field(vnic_metadata.vnic_id, vnic_id);
        modify_field(vnic_metadata.vpc_id, vpc_id);
    }
    modify_field(p4_to_rxdma.service_tag, service_tag);
    modify_field(nat_metadata.pa_or_ca_xlate_idx, pa_or_ca_xlate_idx);
    modify_field(nat_metadata.public_xlate_idx, public_xlate_idx);
    modify_field(nat_metadata.ca6_xlate_idx, ca6_xlate_idx);
}

@pragma stage 1
table local_ip_mapping {
    reads {
        vnic_metadata.vpc_id    : exact;
        key_metadata.mapping_ip : exact;
    }
    actions {
        local_ip_mapping_info;
    }
    size : LOCAL_IP_MAPPING_HASH_TABLE_SIZE;
}

@pragma stage 1
@pragma overflow_table local_ip_mapping
table local_ip_mapping_otcam {
    reads {
        vnic_metadata.vpc_id    : ternary;
        key_metadata.mapping_ip : ternary;
    }
    actions {
        local_ip_mapping_info;
    }
    size : LOCAL_IP_MAPPING_OTCAM_TABLE_SIZE;
}

/******************************************************************************/
/* Service mappings                                                           */
/******************************************************************************/
action service_mapping_info(service_xlate_idx, service_xlate_port) {
    // if table lookup is a miss, return

    modify_field(nat_metadata.service_xlate_idx, service_xlate_idx);
    modify_field(nat_metadata.service_xlate_port, service_xlate_port);
}

@pragma stage 4
table service_mapping {
    reads {
        vnic_metadata.vpc_id2       : exact;
        key_metadata.mapping_ip     : exact;
        key_metadata.mapping_ip2    : exact;
        key_metadata.mapping_port   : exact;
    }
    actions {
        service_mapping_info;
    }
    size : SERVICE_MAPPING_HASH_TABLE_SIZE;
}

@pragma stage 4
@pragma overflow_table service_mapping
table service_mapping_otcam {
    reads {
        vnic_metadata.vpc_id2       : ternary;
        key_metadata.mapping_ip     : ternary;
        key_metadata.mapping_ip2    : ternary;
        key_metadata.mapping_port   : ternary;
    }
    actions {
        service_mapping_info;
    }
    size : SERVICE_MAPPING_OTCAM_TABLE_SIZE;
}

control ingress_mappings {
    apply(local_ip_mapping);
    apply(local_ip_mapping_otcam);
    apply(service_mapping);
    apply(service_mapping_otcam);
}
