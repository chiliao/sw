//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
//
#include "nic/apollo/test/api/utils/vpc.hpp"
#include "nic/apollo/test/api/utils/svc_mapping.hpp"

namespace test {
namespace api {

//----------------------------------------------------------------------------
// Svc maping feeder class routines
//----------------------------------------------------------------------------

void
svc_mapping_feeder::init(pds_obj_key_t key, std::string vip_str,
                         uint16_t svc_port, pds_obj_key_t backend_vpc,
                         std::string backend_ip_str, uint16_t backend_port,
                         std::string backend_pip_str,
                         uint32_t num_svc_mapping) {

    memset(&spec, 0, sizeof(pds_svc_mapping_spec_t));
    spec.key = key;
    spec.skey.vpc = backend_vpc;
    test::extract_ip_addr(backend_ip_str.c_str(), &spec.skey.backend_ip);
    spec.skey.backend_port = backend_port;
    test::extract_ip_addr(vip_str.c_str(), &spec.vip);
    spec.svc_port = svc_port;
    test::extract_ip_addr(backend_pip_str.c_str(), &spec.backend_provider_ip);
    num_obj = num_svc_mapping;
}

void
svc_mapping_feeder::iter_next(int width) {
    spec.key = int2pdsobjkey(pdsobjkey2int(spec.key) + width);
    spec.skey.vpc =
        int2pdsobjkey((pdsobjkey2int(spec.skey.vpc) + width) % PDS_MAX_VPC + 1);
    spec.skey.backend_ip.addr.v6_addr.addr64[1] += width;
    spec.skey.backend_port += width;
    spec.vip.addr.v6_addr.addr64[1] += width;
    spec.svc_port += width;
    if (!ip_addr_is_zero(&spec.backend_provider_ip))
        spec.backend_provider_ip.addr.v6_addr.addr64[1] += width;
    cur_iter_pos++;
}

void
svc_mapping_feeder::key_build(pds_obj_key_t *key) const {
    memcpy(key , &spec.key, sizeof(pds_obj_key_t));
}

void
svc_mapping_feeder::spec_build(pds_svc_mapping_spec_t *spec) const {
    memset(spec, 0, sizeof(pds_svc_mapping_spec_t));
    this->key_build(&spec->key);
    memcpy(&spec->skey, &this->spec.skey, sizeof(pds_svc_mapping_key_t));
    spec->vip = this->spec.vip;
    spec->svc_port = this->spec.svc_port;
    spec->backend_provider_ip = this->spec.backend_provider_ip;
}

bool
svc_mapping_feeder::key_compare(const pds_obj_key_t *key) const {
    if (this->spec.key != *key)
        return false;
    return true;
}

bool
svc_mapping_feeder::spec_compare(const pds_svc_mapping_spec_t *spec) const {
    // todo @njose remove this once read() is complete
    return true;

    if (memcmp(&spec->skey, &this->spec.skey,
               sizeof(pds_svc_mapping_key_t))) {
        return false;
    }

    if (!IPADDR_EQ(&spec->vip, &this->spec.vip))
        return false;

    if (spec->svc_port != this->spec.svc_port)
        return false;

    if (!IPADDR_EQ(&spec->backend_provider_ip, &this->spec.backend_provider_ip))
        return false;

    return true;
}

bool
svc_mapping_feeder::status_compare(
    const pds_svc_mapping_status_t *status1,
    const pds_svc_mapping_status_t *status2) const {
    return true;
}

}    // namespace api
}    // namespace test
