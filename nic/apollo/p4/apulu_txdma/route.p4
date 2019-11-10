/*****************************************************************************/
/* IPv4 Route LPM lookup                                                     */
/*****************************************************************************/

#include "txlpm1.p4"

action route_res_handler() {
    // Disable route LPM lookup in the subsequent recircs
    modify_field(txdma_predicate.lpm1_enable, FALSE);
    //Write the result of the route LPM to PHV
    modify_field(txdma_to_p4e.meter_id,\
                 scratch_metadata.field32 >> ROUTE_RESULT_METERID_SHIFT);
    modify_field(txdma_to_p4e.nexthop_id,\
                 scratch_metadata.field32 >> ROUTE_RESULT_NEXTHOP_SHIFT);
    modify_field(txdma_to_p4e.nexthop_type,\
                 scratch_metadata.field32 >> ROUTE_RESULT_NHTYPE_SHIFT);
}

control route_lookup {
    apply(txlpm1_2);
    apply(txlpm1_3);
    apply(txlpm1_4);
    apply(txlpm1_5);
    apply(txlpm1_6);
    apply(txlpm1);
}
