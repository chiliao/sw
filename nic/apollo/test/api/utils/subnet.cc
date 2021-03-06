//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------

#include "nic/sdk/include/sdk/ip.hpp"
#include "nic/apollo/test/api/utils/batch.hpp"
#include "nic/apollo/test/api/utils/subnet.hpp"
#include "nic/apollo/test/api/utils/vpc.hpp"

namespace test {
namespace api {

//----------------------------------------------------------------------------
// Subnet feeder class routines
//----------------------------------------------------------------------------

void
subnet_feeder::init(pds_obj_key_t key, pds_obj_key_t vpc_key,
                    std::string cidr_str, std::string vrmac_str,
                    int num_subnet, int num_policies_per_subnet,
                    int start_policy_index) {
    ip_prefix_t pfx = {0};
    int i = 0;
    int policy_start = TEST_POLICY_ID_BASE + start_policy_index; // 0 or 1

    memset(&spec, 0, sizeof(pds_subnet_spec_t));
    spec.key = key;
    spec.vpc = vpc_key;
    str2ipv4pfx((char *)cidr_str.c_str(), &pfx);
    spec.v4_prefix.len = pfx.len;
    spec.v4_prefix.v4_addr = pfx.addr.addr.v4_addr;
    mac_str_to_addr((char *)vrmac_str.c_str(), spec.vr_mac);
    spec.v4_route_table = int2pdsobjkey(pdsobjkey2int(key));
    // Unique id, 1-1024 reserved
    spec.v6_route_table = int2pdsobjkey(pdsobjkey2int(key) + 1024);
                                            // for IPv4 rt table
    // TODO: fix for multiple policies
    // num_policies can be either of 1 or 2 as of now.
    spec.num_ing_v4_policy = num_policies_per_subnet;
    spec.num_ing_v6_policy = num_policies_per_subnet;
    spec.num_egr_v4_policy = num_policies_per_subnet;
    spec.num_egr_v6_policy = num_policies_per_subnet;
    while(i < num_policies_per_subnet) {
        spec.ing_v4_policy[i] = int2pdsobjkey(policy_start + i + 1);
        spec.ing_v6_policy[i] = int2pdsobjkey(policy_start + i + 6);
        spec.egr_v4_policy[i] = int2pdsobjkey(policy_start + i + 11);
        spec.egr_v6_policy[i] = int2pdsobjkey(policy_start + i + 16);
        i++;
    }
    spec.fabric_encap.val.vnid = pdsobjkey2int(key) + 512;
    spec.fabric_encap.type = PDS_ENCAP_TYPE_VXLAN;
    //spec.host_if = int2pdsobjkey(1);

    num_obj = num_subnet;
}

subnet_feeder::subnet_feeder(const subnet_feeder& feeder) {
    memcpy(&this->spec, &feeder.spec, sizeof(pds_subnet_spec_t));
    num_obj = feeder.num_obj;
}

void
subnet_feeder::iter_next(int width) {
    spec.key = int2pdsobjkey(pdsobjkey2int(spec.key) + width);
    spec.v4_prefix.v4_addr += (1 << spec.v4_prefix.len);
    spec.v4_route_table = int2pdsobjkey(pdsobjkey2int(spec.v4_route_table) + width);
    spec.v6_route_table = int2pdsobjkey(pdsobjkey2int(spec.v6_route_table) + width);
    spec.fabric_encap.val.vnid += width;

    cur_iter_pos++;
}

void
subnet_feeder::key_build(pds_obj_key_t *key) const {
    memset(key, 0, sizeof(pds_obj_key_t));
    *key = spec.key;
}

void
subnet_feeder::spec_build(pds_subnet_spec_t *spec) const {
    memcpy(spec, &this->spec, sizeof(pds_subnet_spec_t));
}

bool
subnet_feeder::key_compare(const pds_obj_key_t *key) const {
    return (memcmp(key, &this->spec.key, sizeof(pds_obj_key_t)) == 0);
}

// TODO: fix for multiple policies
bool
subnet_feeder::spec_compare(const pds_subnet_spec_t *spec) const {
    if (spec->vpc != this->spec.vpc) {
        return false;
    }

    if (spec->v4_route_table != this->spec.v4_route_table) {
        return false;
    }

    if (spec->v6_route_table != this->spec.v6_route_table) {
        return false;
    }

    if (spec->num_ing_v4_policy == this->spec.num_ing_v4_policy) {
        for (uint32_t i = 0; i < spec->num_ing_v4_policy; i++) {
            if (spec->ing_v4_policy[i] != this->spec.ing_v4_policy[i]) {
                return false;
            }
        }
    } else {
        return false;
    }

    if (spec->num_ing_v6_policy == this->spec.num_ing_v6_policy) {
        for (uint32_t i = 0; i < spec->num_ing_v6_policy; i++) {
            if (spec->ing_v6_policy[i] != this->spec.ing_v6_policy[i]) {
                return false;
            }
        }
    } else {
        return false;
    }

    if (spec->num_egr_v4_policy == this->spec.num_egr_v4_policy) {
        for (uint32_t i = 0; i < spec->num_egr_v4_policy; i++) {
            if (spec->egr_v4_policy[i] != this->spec.egr_v4_policy[i]) {
                return false;
            }
        }
    } else {
        return false;
    }

    if (spec->num_egr_v6_policy == this->spec.num_egr_v6_policy) {
        for (uint32_t i = 0; i < spec->num_egr_v6_policy; i++) {
            if (spec->egr_v6_policy[i] != this->spec.egr_v6_policy[i]) {
                return false;
            }
        }
    } else {
        return false;
    }

    if (memcmp(&spec->vr_mac, this->spec.vr_mac, sizeof(mac_addr_t))) {
        return false;
    }

    if (spec->v4_vr_ip != this->spec.v4_vr_ip) {
        return false;
    }

    if (memcmp(&spec->v6_vr_ip, &this->spec.v6_vr_ip, sizeof(ip_addr_t))) {
        return false;
    }

    if (memcmp(&spec->fabric_encap, &this->spec.fabric_encap,
               sizeof(pds_encap_t))) {
        return false;
    }
    return true;
}

bool
subnet_feeder::status_compare(const pds_subnet_status_t *status1,
                              const pds_subnet_status_t *status2) const {
    return true;
}

//----------------------------------------------------------------------------
// Subnet CRUD helper routines
//----------------------------------------------------------------------------

void
subnet_create (subnet_feeder& feeder)
{
    pds_batch_ctxt_t bctxt = batch_start();

    SDK_ASSERT_RETURN_VOID(
        (SDK_RET_OK == many_create<subnet_feeder>(bctxt, feeder)));
    batch_commit(bctxt);
}

void
subnet_read (subnet_feeder& feeder, sdk_ret_t exp_result)
{
    SDK_ASSERT_RETURN_VOID(
        (SDK_RET_OK == many_read<subnet_feeder>(feeder, exp_result)));
}

static void
subnet_attr_update (subnet_feeder& feeder, pds_subnet_spec_t *spec,
                    uint64_t chg_bmap)
{
    if (bit_isset(chg_bmap, SUBNET_ATTR_VPC)) {
        feeder.spec.vpc = spec->vpc;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V4_PREFIX)) {
        feeder.spec.v4_prefix = spec->v4_prefix;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V6_PREFIX)) {
        feeder.spec.v6_prefix = spec->v6_prefix;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V4_VRIP)) {
        feeder.spec.v4_vr_ip = spec->v4_vr_ip;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V6_VRIP)) {
        feeder.spec.v6_vr_ip = spec->v6_vr_ip;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_VR_MAC)) {
        memcpy(&feeder.spec.vr_mac, &spec->vr_mac, sizeof(spec->vr_mac));
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V4_RTTBL)) {
        feeder.spec.v4_route_table = spec->v4_route_table;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V6_RTTBL)) {
        feeder.spec.v6_route_table = spec->v6_route_table;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V4_INGPOL)) {
        feeder.spec.num_ing_v4_policy = spec->num_ing_v4_policy;
        memcpy(&feeder.spec.ing_v4_policy, &spec->ing_v4_policy,
               sizeof(spec->ing_v4_policy[0]) * spec->num_ing_v4_policy);
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V6_INGPOL)) {
        feeder.spec.num_ing_v6_policy = spec->num_ing_v6_policy;
        memcpy(&feeder.spec.ing_v6_policy, &spec->ing_v6_policy,
               sizeof(spec->ing_v6_policy[0]) * spec->num_ing_v6_policy);
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V4_EGRPOL)) {
        feeder.spec.num_egr_v4_policy = spec->num_egr_v4_policy;
        memcpy(&feeder.spec.egr_v4_policy, &spec->egr_v4_policy,
               sizeof(spec->egr_v4_policy[0]) * spec->num_egr_v4_policy);
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_V6_EGRPOL)) {
        feeder.spec.num_egr_v6_policy = spec->num_egr_v6_policy;
        memcpy(&feeder.spec.egr_v6_policy, &spec->egr_v6_policy,
               sizeof(spec->egr_v6_policy[0]) * spec->num_egr_v6_policy);
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_FAB_ENCAP)) {
        feeder.spec.fabric_encap = spec->fabric_encap;
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_HOST_IF)) {
        feeder.spec.num_host_if = 1;
        feeder.spec.host_if[0] = spec->host_if[0];
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_DHCP_POL)) {
        feeder.spec.num_dhcp_policy = spec->num_dhcp_policy;
        memcpy(&feeder.spec.dhcp_policy, &spec->dhcp_policy,
               sizeof(spec->dhcp_policy));
    }
    if (bit_isset(chg_bmap, SUBNET_ATTR_TOS)) {
        feeder.spec.tos = spec->tos;
    }
}

void
subnet_update (subnet_feeder& feeder, pds_subnet_spec_t *spec,
               uint64_t chg_bmap, sdk_ret_t exp_result)
{
    pds_batch_ctxt_t bctxt = batch_start();

    subnet_attr_update(feeder, spec, chg_bmap);
    SDK_ASSERT_RETURN_VOID(
        (SDK_RET_OK == many_update<subnet_feeder>(bctxt, feeder)));

    // if expected result is err, batch commit should fail
    if (exp_result == SDK_RET_ERR)
        batch_commit_fail(bctxt);
    else
        batch_commit(bctxt);
}

void
subnet_delete (subnet_feeder& feeder)
{
    pds_batch_ctxt_t bctxt = batch_start();

    SDK_ASSERT_RETURN_VOID(
        (SDK_RET_OK == many_delete<subnet_feeder>(bctxt, feeder)));
    batch_commit(bctxt);
}

//----------------------------------------------------------------------------
// Misc routines
//----------------------------------------------------------------------------

// do not modify these sample values as rest of system is sync with these
pds_obj_key_t k_subnet_key = int2pdsobjkey(1);
static subnet_feeder k_subnet_feeder;

void sample_subnet_setup(pds_batch_ctxt_t bctxt) {
    // setup and teardown parameters should be in sync
    k_subnet_feeder.init(k_subnet_key, k_vpc_key, "10.1.0.0/16",
                         "00:02:01:00:00:01");
    create(bctxt, k_subnet_feeder);
}

void sample_subnet_teardown(pds_batch_ctxt_t bctxt) {
    k_subnet_feeder.init(k_subnet_key, k_vpc_key, "10.1.0.0/16",
                         "00:02:01:00:00:01");
    del(bctxt, k_subnet_feeder);
}

void sample1_subnet_setup(pds_batch_ctxt_t bctxt) {
    // setup and teardown parameters should be in sync
    k_subnet_feeder.init(k_subnet_key, k_vpc_key, "10.0.0.0/8",
                         "00:02:01:00:00:01");
    create(bctxt, k_subnet_feeder);
}

void sample1_subnet_teardown(pds_batch_ctxt_t bctxt) {
    k_subnet_feeder.init(k_subnet_key, k_vpc_key, "10.0.0.0/8",
                         "00:02:01:00:00:01");
    del(bctxt, k_subnet_feeder);
}

void
subnet_spec_policy_fill (pds_subnet_spec_t *spec, uint8_t num_policies,
                         uint8_t start_policy_index)
{
    uint8_t policy_start = TEST_POLICY_ID_BASE + start_policy_index;

    spec->num_ing_v4_policy = num_policies;
    spec->num_ing_v6_policy = num_policies;
    spec->num_egr_v4_policy = num_policies;
    spec->num_egr_v6_policy = num_policies;
    for (uint8_t i = 0; i < num_policies; i++) {
        spec->ing_v4_policy[i] = int2pdsobjkey(policy_start + i + 1);
        spec->ing_v6_policy[i] = int2pdsobjkey(policy_start + i + 6);
        spec->egr_v4_policy[i] = int2pdsobjkey(policy_start + i + 11);
        spec->egr_v6_policy[i] = int2pdsobjkey(policy_start + i + 16);
    }
}

void
subnet_spec_route_table_fill (pds_subnet_spec_t *spec, uint16_t index)
{
    spec->v4_route_table = int2pdsobjkey(index);
    spec->v6_route_table = int2pdsobjkey(index + 1024);
}

}    // namespace api
}    // namespace test
