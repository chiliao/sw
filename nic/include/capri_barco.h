#ifndef __CAPRI_BARCO_H__
#define __CAPRI_BARCO_H__

#define CAPRI_BARCO_MD_HENS_REG_BASE                    (0x6580000)

#define CAPRI_BARCO_MD_HENS_REG_GCM0_PRODUCER_IDX       (CAPRI_BARCO_MD_HENS_REG_BASE + 0x20c)
#define CAPRI_BARCO_MD_HENS_REG_GCM0_CONSUMER_IDX       (CAPRI_BARCO_MD_HENS_REG_BASE + 0x280)

#define CAPRI_BARCO_MP_MPNS_REG_BASE                    (0x6560000) // CAP_ADDR_BASE_MP_MPNS_OFFSET

#define CAPRI_BARCO_MP_MPNS_REG_MPP1_PRODUCER_IDX       (CAPRI_BARCO_MP_MPNS_REG_BASE + 0x10c)
#define CAPRI_BARCO_MP_MPNS_REG_MPP1_CONSUMER_IDX       (CAPRI_BARCO_MP_MPNS_REG_BASE + 0x180)

#endif  /* __CAPRI_BARCO_H__ */
