header_type apulu_i2e_metadata_t {
    fields {
        mapping_lkp_addr    : 128;
        entropy_hash        : 32;
        mapping_lkp_type    : 2;
        flow_role           : 1;
        rx_packet           : 1;
        session_id          : 20;
        vnic_id             : 16;
        bd_id               : 16;
        vpc_id              : 16;
        pad0                : 7;
        update_checksum     : 1;
    }
}
