
control dnat_lookup(inout cap_phv_intr_global_h capri_intrinsic,
            inout cap_phv_intr_p4_h intr_p4,
            inout headers hdr,
            inout metadata_t metadata) {


  
    @appdatafields ("addr_type", "addr")
    @hwfields_access_api  
    @name(".dnat")
    action dnat_a( 
			bit<22>    pad1,
			//bit<3>     pad2,		
			bit<2>     addr_type,
			bit<128>   addr,
                        bit<16>    epoch,
			bit<18>    hash1,
			bit<12>    hint1,
			bit<18>    hash2,
			bit<12>    hint2,
			bit<18>    hash3,
			bit<12>    hint3,
			bit<18>    hash4,
			bit<12>    hint4,
			bit<18>    hash5,
			bit<12>    hint5,
			bit<1>     more_hashes,
			bit<12>    more_hints,
			bit<1>     entry_valid
			//                      bit<5>     pad2
		  ) {

      bit<32>  hardware_hash = __hash_value();

       if (entry_valid == FALSE) {
	 hdr.ingress_recirc_header.dnat_done = TRUE;
	 
	 //JUST TO TEST
	 /*
	 if(hdr.ingress_recirc_header.isValid()) {
	   hdr.ingress_recirc_header.dnat_done = TRUE;
	 } else {
	    hdr.ingress_recirc_header.dnat_done = FALSE;
	    hdr.ingress_recirc_header.dnat_ohash_lkp = TRUE;
	    hdr.ingress_recirc_header.dnat_ohash = 
	     POS_OVERFLOW_HASH_BIT | 0x3456;

	 }
	 */

          return;
        }

       if (__table_hit()) {
	    hdr.ingress_recirc_header.dnat_done = TRUE;
	    metadata.key.src= addr;
	    metadata.key.ktype= addr_type;
	    hdr.p4i_to_p4e_header.dnat_epoch = epoch;
	    hdr.ingress_recirc_header.dnat_ohash_lkp = metadata.cntrl.dnat_ohash_lkp;
	    //          metadata.scratch.flow_hash = hardware_hash[31:14];
            // metadata.scratch.hint_valid = TRUE;
	    
       } else {
	 bit<1> hint_valid = FALSE;
	 if ((hint1 != 0 )  && (hash1 == hardware_hash[31:14])) {
	   hint_valid = TRUE;
	   metadata.scratch.flow_hint = (bit<20>)hint1;
	 } else if ((hint2 != 0 )  && (hash2 == hardware_hash[31:14])) {
	   hint_valid = TRUE;
	   metadata.scratch.flow_hint = (bit<20>)hint2;
	 } else if ((hint3 != 0 ) &&  (hash3 == hardware_hash[31:14])) {
	   hint_valid = TRUE;
	   metadata.scratch.flow_hint = (bit<20>)hint3;
	 } else if ((hint4 != 0 ) && (hash4 == hardware_hash[31:14])) {
	   hint_valid = TRUE;
	   metadata.scratch.flow_hint = (bit<20>)hint4;
	 } else if ((hint5 != 0 )  && (hash5 == hardware_hash[31:14])) {
	   hint_valid = TRUE;
	   metadata.scratch.flow_hint = (bit<20>)hint5;
	 } else if ((more_hashes == 1) ) {
	   hint_valid = TRUE;
	   metadata.scratch.flow_hint = (bit<20>)more_hints;
	 }
	 
	 if( hint_valid == TRUE) {
	   metadata.cntrl.dnat_ohash_lkp = TRUE;
	   hdr.ingress_recirc_header.dnat_ohash_lkp = TRUE;
	   hdr.ingress_recirc_header.dnat_ohash = 
	     POS_OVERFLOW_HASH_BIT | (bit<32>)metadata.scratch.flow_hint;
	 } else {
	   hdr.ingress_recirc_header.dnat_done = TRUE;
	 }
	 
       }
    }

    @hbm_table
    @name(".dnat_ohash")
     table dnat_ohash {
      key =  {
        hdr.ingress_recirc_header.dnat_ohash : exact;
      }
	
      actions = {
        dnat_a;	
      }
      size = DNAT_OHASH_TABLE_SIZE;
      placement = HBM;
      default_action = dnat_a;
      stage = 2;


    }


    @capi_bitfields_struct
    @name(".dnat")
    table dnat {
        key = {
            metadata.key.vnic_id : exact;
	    metadata.key.ktype7 : exact @name(".key.metadata.ktype");
	    metadata.key.src : exact;
	}
        actions  = {
           dnat_a;
        }
        size =  DNAT_TABLE_SIZE;
        placement = HBM;
        stage = 1;
        default_action = dnat_a;
        hash_collision_table = dnat_ohash;
	
    }

    /*    
    @name(".key_fix")
      action key_fix_a() {
        metadata.key.vnic_id = metadata.cntrl.vnic_id;
	if(metadata.cntrl.l2_fwd == FALSE) {
	  metadata.key.smac = (bit<48>)0;
	  metadata.key.dmac = (bit<48>)0;
	}
     }
      
    @name(".key_fix")
    table key_fix {
        actions  = {
           key_fix_a;
        }
        default_action = key_fix_a;
        stage = 1;
    }
    */

    apply {
      //    if(!hdr.ingress_recirc_header.isValid()) {
	if(metadata.cntrl.skip_dnat_lkp == FALSE) {
	  dnat.apply();
	}
	//      }

      if(metadata.cntrl.dnat_ohash_lkp == TRUE) {
	dnat_ohash.apply();
      }
      
      //      if(metadata.cntrl.direction == RX_FROM_SWITCH) {
      //	key_fix.apply();
      // }
    }
}


