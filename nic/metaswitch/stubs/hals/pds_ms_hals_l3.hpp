//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// PDS Metaswitch L3 HALS stub integration subcomponent APIs
//---------------------------------------------------------------

#ifndef __PDS_MS_HALS_L3_HPP__
#define __PDS_MS_HALS_L3_HPP__

#include <hals_integ_subcomp.hpp>

namespace pds_ms {

class hals_l3_integ_subcomp_t : public hals::IntegrationSubcomponent {
public:
    //-------------------------------------------------------------------------
    // NHPI
    //-------------------------------------------------------------------------
    NBB_BYTE nhpi_add_update_ecmp (ATG_NHPI_ADD_UPDATE_ECMP *add_update_ecmp_ips); 

    NBB_BYTE nhpi_delete_ecmp(ATG_NHPI_DELETE_ECMP *delete_ecmp_ips);

    NBB_BYTE nhpi_destroy_ecmp(NBB_CORRELATOR ecmp_corr);

    NBB_BYTE nhpi_add_ecmp_nh(NBB_CORRELATOR ecmp_corr,
                              ATG_NHPI_APPENDED_NEXT_HOP *next_hop,
                              NBB_BYTE cascaded);

    NBB_BYTE nhpi_delete_ecmp_nh(NBB_CORRELATOR ecmp_corr,
                                 NBB_CORRELATOR nh_corr);
};

} // End namespace

#endif