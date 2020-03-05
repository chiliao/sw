//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains athena_app test utility API
///
//----------------------------------------------------------------------------

#include "app_test_utils.hpp"
#include "nic/apollo/api/impl/athena/ftl_pollers_client.hpp"

namespace test {
namespace athena_app {

const static std::map<string,bool> truefalse2hool_map =
{
    {"true",            true},
    {"false",           false},
    {"TRUE",            true},
    {"FALSE",           false},
    {"True",            true},
    {"False",           false},
};

const static std::map<string,pds_flow_type_t> flowtype2num_map =
{
    {"tcp",             PDS_FLOW_TYPE_TCP},
    {"TCP",             PDS_FLOW_TYPE_TCP},
    {"udp",             PDS_FLOW_TYPE_UDP},
    {"UDP",             PDS_FLOW_TYPE_UDP},
    {"icmp",            PDS_FLOW_TYPE_ICMP},
    {"ICMP",            PDS_FLOW_TYPE_ICMP},
    {"others",          PDS_FLOW_TYPE_OTHERS},
    {"OTHERS",          PDS_FLOW_TYPE_OTHERS},
};

const static std::map<string,pds_flow_state_t> flowstate2num_map =
{
    {"unestablished",   UNESTABLISHED},
    {"UNESTABLISHED",   UNESTABLISHED},
    {"syn_sent",        PDS_FLOW_STATE_SYN_SENT},
    {"SYN_SENT",        PDS_FLOW_STATE_SYN_SENT},
    {"syn_recv",        PDS_FLOW_STATE_SYN_RECV},
    {"SYN_RECV",        PDS_FLOW_STATE_SYN_RECV},
    {"synack_sent",     PDS_FLOW_STATE_SYNACK_SENT},
    {"SYNACK_SENT",     PDS_FLOW_STATE_SYNACK_SENT},
    {"synack_recv",     PDS_FLOW_STATE_SYNACK_RECV},
    {"SYNACK_RECV",     PDS_FLOW_STATE_SYNACK_RECV},
    {"established",     ESTABLISHED},
    {"ESTABLISHED",     ESTABLISHED},
    {"fin_sent",        FIN_SENT},
    {"FIN_SENT",        FIN_SENT},
    {"fin_recv",        FIN_RECV},
    {"FIN_RECV",        FIN_RECV},
    {"time_wait",       TIME_WAIT},
    {"TIME_WAIT",       TIME_WAIT},
    {"rst_close",       RST_CLOSE},
    {"RST_CLOSE",       RST_CLOSE},
    {"removed",         REMOVED},
    {"REMOVED",         REMOVED},
};

static const bool *
truefalse2hool_find(const std::string& token)
{
    auto iter = truefalse2hool_map.find(token);
    if (iter != truefalse2hool_map.end()) {
        return &iter->second;
    }
    return nullptr;
}

static const pds_flow_type_t *
flowtype2num_find(const std::string& token)
{
    auto iter = flowtype2num_map.find(token);
    if (iter != flowtype2num_map.end()) {
        return &iter->second;
    }
    return nullptr;
}

static const pds_flow_state_t *
flowstate2num_find(const std::string& token)
{
    auto iter = flowstate2num_map.find(token);
    if (iter != flowstate2num_map.end()) {
        return &iter->second;
    }
    return nullptr;
}

/**
 * Test parameters
 */
sdk_ret_t 
test_param_t::num(uint32_t *ret_num,
                  bool suppress_err_log) const
{
    if (is_num()) {
        *ret_num = num_;
        return SDK_RET_OK;
    }

    *ret_num = 0;
    if (!suppress_err_log) {
        TEST_LOG_INFO("Numeric token not found\n");
    }
    return SDK_RET_ERR;
}

sdk_ret_t 
test_param_t::str(std::string *ret_str,
                  bool suppress_err_log) const
{
    if (is_str()) {
        ret_str->assign(str_);
        return SDK_RET_OK;
    }

    ret_str->clear();
    if (!suppress_err_log) {
        TEST_LOG_INFO("String token not found\n");
    }
    return SDK_RET_ERR;
}

sdk_ret_t
test_param_t::bool_val(bool *ret_bool,
                       bool suppress_err_log) const
{
    switch (type) {

    case TOKEN_TYPE_NUM:
        *ret_bool = !!num_;
        break;

    case TOKEN_TYPE_STR: {
        const bool *find_val = truefalse2hool_find(str_);
        if (find_val) {
            *ret_bool = *find_val;
            break;
        }

        // Fall through!!!
    }

    default:
        *ret_bool = false;
        if (!suppress_err_log) {
            TEST_LOG_INFO("Bool token not found\n");
        }
        return SDK_RET_ERR;

    }
    return SDK_RET_OK;
}

sdk_ret_t
test_param_t::flowtype(pds_flow_type_t *ret_flowtype,
                       bool suppress_err_log) const
{
    *ret_flowtype = PDS_FLOW_TYPE_TCP;
    switch (type) {

    case TOKEN_TYPE_NUM:
        if (num_ > (uint32_t)PDS_FLOW_TYPE_OTHERS) {
            if (!suppress_err_log) {
                TEST_LOG_INFO("Invalid flow type %u\n", num_);
            }
            return SDK_RET_ERR;
        }
        *ret_flowtype = (pds_flow_type_t)num_;
        break;

    case TOKEN_TYPE_STR: {
        const pds_flow_type_t *find_val = flowtype2num_find(str_);
        if (!find_val) {
            if (!suppress_err_log) {
                TEST_LOG_INFO("Invalid flow type %s\n", str_.c_str());
            }
            return SDK_RET_ERR;
        }
        *ret_flowtype = *find_val;
        break;
    }

    default:
        if (!suppress_err_log) {
            TEST_LOG_INFO("Flow type token not found\n");
        }
        return SDK_RET_ERR;

    }
    return SDK_RET_OK;
}

sdk_ret_t
test_param_t::flowstate(pds_flow_state_t *ret_flowstate,
                        bool suppress_err_log) const
{
    *ret_flowstate = UNESTABLISHED;
    switch (type) {

    case TOKEN_TYPE_NUM:
        if (num_ > (uint32_t)REMOVED) {
            if (!suppress_err_log) {
                TEST_LOG_INFO("Invalid flow state %u\n", num_);
            }
            return SDK_RET_ERR;
        }
        *ret_flowstate = (pds_flow_state_t)num_;
        break;

    case TOKEN_TYPE_STR: {
        const pds_flow_state_t *find_val = flowstate2num_find(str_);
        if (!find_val) {
            if (!suppress_err_log) {
                TEST_LOG_INFO("Invalid flow state %s\n", str_.c_str());
            }
            return SDK_RET_ERR;
        }
        *ret_flowstate = *find_val;
        break;
    }

    default:
        if (!suppress_err_log) {
            TEST_LOG_INFO("Flow state token not found\n");
        }
        return SDK_RET_ERR;
    }
    return SDK_RET_OK;
}

sdk_ret_t 
test_param_t::tuple(test_param_tuple_t *ret_tuple,
                    bool suppress_err_log) const
{
    if (is_tuple()) {
        *ret_tuple = tuple_;
        return SDK_RET_OK;
    }

    ret_tuple->clear();
    if (!suppress_err_log) {
        TEST_LOG_INFO("Tuple not found\n");
    }
    return SDK_RET_ERR;
}

/**
 * Vector parameters
 */
sdk_ret_t 
test_vparam_t::num(uint32_t idx,
                   uint32_t *ret_num,
                   bool suppress_err_log) const
{
    if (idx < vparam.size()) {
        const test_param_t& param = vparam.at(idx);
        return param.num(ret_num);
    }

    *ret_num = 0;
    if (!suppress_err_log) {
        TEST_LOG_INFO("Numeric token not found at index %u\n", idx);
    }
    return SDK_RET_ERR;
}

sdk_ret_t 
test_vparam_t::str(uint32_t idx,
                   std::string *ret_str,
                   bool suppress_err_log) const
{
    if (idx < vparam.size()) {
        const test_param_t& param = vparam.at(idx);
        return param.str(ret_str);
    }

    ret_str->clear();
    if (!suppress_err_log) {
        TEST_LOG_INFO("String token not found at index %u\n", idx);
    }
    return SDK_RET_ERR;
}

sdk_ret_t
test_vparam_t::flowtype(uint32_t idx,
                        pds_flow_type_t *ret_flowtype,
                        bool suppress_err_log) const
{
    if (idx < vparam.size()) {
        const test_param_t& param = vparam.at(idx);
        return param.flowtype(ret_flowtype);
    }

    *ret_flowtype = PDS_FLOW_TYPE_TCP;
    if (!suppress_err_log) {
        TEST_LOG_INFO("Flow type token not found at index %u\n", idx);
    }
    return SDK_RET_ERR;
}

sdk_ret_t
test_vparam_t::flowstate(uint32_t idx,
                        pds_flow_state_t *ret_flowstate,
                        bool suppress_err_log) const
{
    if (idx < vparam.size()) {
        const test_param_t& param = vparam.at(idx);
        return param.flowstate(ret_flowstate);
    }

    *ret_flowstate = UNESTABLISHED;
    if (!suppress_err_log) {
        TEST_LOG_INFO("Flow state token not found at index %u\n", idx);
    }
    return SDK_RET_ERR;
}

sdk_ret_t 
test_vparam_t::tuple(uint32_t idx,
                     test_param_tuple_t *ret_tuple,
                     bool suppress_err_log) const
{
    if (idx < vparam.size()) {
        const test_param_t& param = vparam.at(idx);
        return param.tuple(ret_tuple);
    }

    ret_tuple->clear();
    if (!suppress_err_log) {
        TEST_LOG_INFO("Tuple not found at index %u\n", idx);
    }
    return SDK_RET_ERR;
}

bool
test_vparam_t::expected_bool(bool dflt) const
{
    if (vparam.size()) {
        const test_param_t& param = vparam.at(0);
        bool  ret_val;
        if (param.bool_val(&ret_val) == SDK_RET_OK) {
            return ret_val;
        }
    }
    return dflt;
}

uint32_t
test_vparam_t::expected_num(uint32_t dflt) const
{
    if (vparam.size()) {
        const test_param_t& param = vparam.at(0);
        if (param.is_num()) {
            return param.num_;
        }
    }
    return dflt;
}

const std::string&
test_vparam_t::expected_str(void) const
{
    static std::string empty_str;

    if (vparam.size()) {
        const test_param_t& param = vparam.at(0);
        if (param.is_str()) {
            return param.str_;
        }
    }
    return empty_str;
}

/*
 * Tuple evaluation helper
 */
void
tuple_eval_t::reset(test_vparam_ref_t vparam,
                    uint32_t vparam_idx)
{
    fail_count = 0;
    if (vparam.tuple(vparam_idx, &tuple) != SDK_RET_OK) {
        fail_count++;
    }
}

uint32_t
tuple_eval_t::num(uint32_t idx)
{
    uint32_t ret_val = 0;

    if ((idx >= tuple.size()) ||
        (tuple.at(idx).num(&ret_val) != SDK_RET_OK)) {
        fail_count++;
    }
    return ret_val;
}

std::string
tuple_eval_t::str(uint32_t idx)
{
    std::string ret_str;

    if ((idx >= tuple.size()) ||
        (tuple.at(idx).str(&ret_str) != SDK_RET_OK)) {
        fail_count++;
    }
    return ret_str;
}

pds_flow_type_t
tuple_eval_t::flowtype(uint32_t idx)
{
    pds_flow_type_t ret_val = PDS_FLOW_TYPE_TCP;

    if ((idx >= tuple.size()) ||
        (tuple.at(idx).flowtype(&ret_val) != SDK_RET_OK)) {
        fail_count++;
    }
    return ret_val;
}

pds_flow_state_t
tuple_eval_t::flowstate(uint32_t idx)
{
    pds_flow_state_t ret_val = UNESTABLISHED;

    if ((idx >= tuple.size()) ||
        (tuple.at(idx).flowstate(&ret_val) != SDK_RET_OK)) {
        fail_count++;
    }
    return ret_val;
}

/**
 * Metrics
 */
sdk_ret_t
aging_metrics_t::baseline(void)
{
    return refresh(base);
}

uint64_t
aging_metrics_t::delta_expired_entries(void) const
{
    ftl_dev_if::lif_attr_metrics_t  curr;
    uint64_t                        delta = 0;

    switch (qtype) {

    case ftl_dev_if::FTL_QTYPE_SCANNER_SESSION:
    case ftl_dev_if::FTL_QTYPE_SCANNER_CONNTRACK:

        if (refresh(curr) == SDK_RET_OK) {
            delta = curr.scanners.total_expired_entries -
                    base.scanners.total_expired_entries;
        }
        break;

    default:
        break;
    }

    return delta;
}

uint64_t
aging_metrics_t::delta_num_qfulls(void) const
{
    ftl_dev_if::lif_attr_metrics_t  curr;
    uint64_t                        delta = 0;

    switch (qtype) {

    case ftl_dev_if::FTL_QTYPE_POLLER:

        if (refresh(curr) == SDK_RET_OK) {
            delta = curr.pollers.total_num_qfulls -
                    base.pollers.total_num_qfulls;
        }
        break;

    default:
        break;
    }

    return delta;
}

void
aging_metrics_t::show(bool latest) const
{
    ftl_dev_if::lif_attr_metrics_t  curr;

    if (latest) {
        refresh(curr);
    } else {
        curr = base;
    }

    switch (qtype) {

    case ftl_dev_if::FTL_QTYPE_SCANNER_SESSION:
    case ftl_dev_if::FTL_QTYPE_SCANNER_CONNTRACK:

        TEST_LOG_INFO("total_cb_cfg_discards   : %" PRIu64 "\n", 
                      curr.scanners.total_cb_cfg_discards);
        TEST_LOG_INFO("total_scan_invocations  : %" PRIu64 "\n",
                      curr.scanners.total_scan_invocations);
        TEST_LOG_INFO("total_expired_entries   : %" PRIu64 "\n",
                      curr.scanners.total_expired_entries);
        TEST_LOG_INFO("min_range_elapsed_ns    : %" PRIu64 "\n",
                      curr.scanners.min_range_elapsed_ns);
        TEST_LOG_INFO("avg_min_range_elapsed_ns: %" PRIu64 "\n",
                      curr.scanners.avg_min_range_elapsed_ns);
        TEST_LOG_INFO("max_range_elapsed_ns    : %" PRIu64 "\n",
                      curr.scanners.max_range_elapsed_ns);
        TEST_LOG_INFO("avg_max_range_elapsed_ns: %" PRIu64 "\n",
                      curr.scanners.avg_max_range_elapsed_ns);
        break;

    case ftl_dev_if::FTL_QTYPE_POLLER:
        TEST_LOG_INFO("total_num_qfulls        : %" PRIu64 "\n", 
                      curr.pollers.total_num_qfulls);
        break;

    default:
        break;
    }
}

sdk_ret_t
aging_metrics_t::refresh(ftl_dev_if::lif_attr_metrics_t& m) const
{
    sdk_ret_t   ret = SDK_RET_OK;

    switch (qtype) {

    case ftl_dev_if::FTL_QTYPE_SCANNER_SESSION:
        ret = ftl_pollers_client::session_scanners_metrics_get(&m);
        break;

    case ftl_dev_if::FTL_QTYPE_SCANNER_CONNTRACK:
        ret = ftl_pollers_client::conntrack_scanners_metrics_get(&m);
        break;

    case ftl_dev_if::FTL_QTYPE_POLLER:
        ret = ftl_pollers_client::pollers_metrics_get(&m);
        break;

    default:
        memset(&m, 0, sizeof(m));
        break;
    }

    return ret;
}

/*
 * Aging timeout configs
 */
void 
aging_tmo_cfg_t::reset(void)
{
    sdk_ret_t   ret;

    failures.clear();
    ret = is_accel_tmo ?
          pds_flow_age_accel_timeouts_set(&tmo_rec) :
          pds_flow_age_normal_timeouts_set(&tmo_rec);

    if (ret != SDK_RET_OK) {
        TEST_LOG_INFO("Failed to set %s timeouts\n",
                      is_accel_tmo ? "accelerated" : "normal");
        failures.counters.set++;
    }
}

/*
 * Return configured inactivity timeout for a given flowtype and flowstate
 */
void
aging_tmo_cfg_t::conntrack_tmo_set(pds_flow_type_t flowtype,
                                   pds_flow_state_t flowstate,
                                   uint32_t tmo_val)
{
    uint32_t    errors = 0;

    switch (flowtype) {

    case PDS_FLOW_TYPE_TCP:

        switch (flowstate) {

        case UNESTABLISHED:
        case PDS_FLOW_STATE_SYN_SENT:
        case PDS_FLOW_STATE_SYN_RECV:
        case PDS_FLOW_STATE_SYNACK_SENT:
        case PDS_FLOW_STATE_SYNACK_RECV:
            tmo_rec.tcp_syn_tmo = tmo_val;
            break;
        case ESTABLISHED:
            tmo_rec.tcp_est_tmo = tmo_val;
            break;
        case FIN_SENT:
        case FIN_RECV:
            tmo_rec.tcp_fin_tmo = tmo_val;
            break;
        case TIME_WAIT:
            tmo_rec.tcp_timewait_tmo = tmo_val;
            break;
        case RST_CLOSE:
            tmo_rec.tcp_rst_tmo = tmo_val;
            break;
        default:
            TEST_LOG_INFO("Invalid TCP flowstate %u\n", flowstate);
            errors++;
            break;
        }
        break;

    case PDS_FLOW_TYPE_UDP:

        switch (flowstate) {
        case UNESTABLISHED:
            tmo_rec.udp_tmo = tmo_val;
            break;
        case ESTABLISHED:
            tmo_rec.udp_est_tmo = tmo_val;
            break;
        default:
            TEST_LOG_INFO("Invalid UDP flowstate %u\n", flowstate);
            errors++;
            break;
        }
        break;

    case PDS_FLOW_TYPE_ICMP:
        tmo_rec.icmp_tmo = tmo_val;
        break;

    case PDS_FLOW_TYPE_OTHERS:
        tmo_rec.others_tmo = tmo_val;
        break;

    default:
        TEST_LOG_INFO("Invalid flowtype %u\n", flowtype);
        errors++;
        break;
    }

    failures.counters.set += errors;
    if (!errors) {
        tmo_set();
    }
}

/*
 * Return configured inactivity timeout for a given flowtype and flowstate
 */
uint32_t
aging_tmo_cfg_t::conntrack_tmo_get(pds_flow_type_t flowtype,
                                   pds_flow_state_t flowstate)
{
    uint32_t    ret_tmo = 0;

    switch (flowtype) {

    case PDS_FLOW_TYPE_TCP:

        switch (flowstate) {

        case UNESTABLISHED:
        case PDS_FLOW_STATE_SYN_SENT:
        case PDS_FLOW_STATE_SYN_RECV:
        case PDS_FLOW_STATE_SYNACK_SENT:
        case PDS_FLOW_STATE_SYNACK_RECV:
            ret_tmo = tmo_rec.tcp_syn_tmo;
            break;
        case ESTABLISHED:
            ret_tmo = tmo_rec.tcp_est_tmo;
            break;
        case FIN_SENT:
        case FIN_RECV:
            ret_tmo = tmo_rec.tcp_fin_tmo;
            break;
        case TIME_WAIT:
            ret_tmo = tmo_rec.tcp_timewait_tmo;
            break;
        case RST_CLOSE:
            ret_tmo = tmo_rec.tcp_rst_tmo;
            break;
        default:
            break;
        }
        break;

    case PDS_FLOW_TYPE_UDP:

        switch (flowstate) {
        case UNESTABLISHED:
            ret_tmo = tmo_rec.udp_tmo;
            break;
        case ESTABLISHED:
            ret_tmo = tmo_rec.udp_est_tmo;
            break;
        default:
            break;
        }
        break;

    case PDS_FLOW_TYPE_ICMP:
        ret_tmo = tmo_rec.icmp_tmo;
        break;

    case PDS_FLOW_TYPE_OTHERS:
        ret_tmo = tmo_rec.others_tmo;
        break;

    default:
        break;
    }
    return ret_tmo;
}

void 
aging_tmo_cfg_t::tmo_set(void)
{
    sdk_ret_t   ret;

    ret = is_accel_tmo ?
          pds_flow_age_accel_timeouts_set(&tmo_rec) :
          pds_flow_age_normal_timeouts_set(&tmo_rec);

    if (ret != SDK_RET_OK) {
        TEST_LOG_INFO("Failed to set %s timeouts\n",
                      is_accel_tmo ? "accelerated" : "normal");
        failures.counters.set++;
    }
}

/*
 * Aging results, with tolerance
 */
void 
aging_tolerance_t::reset(create_id_map_mode_t mode)
{
    create_id_map_mode = mode;
    create_count_ = 0;
    expiry_count_ = 0;
    create_id_map.clear();
    failures.clear();
    normal_tmo.reset();
    accel_tmo.reset();

    /*
     * Note: we don't reset accel_control to allow caller to test
     * with whichever control is currently active.
     */
}

void 
aging_tolerance_t::age_accel_control(bool enable_sense)
{
    if (pds_flow_age_accel_control(enable_sense) == SDK_RET_OK) {
        curr_tmo = enable_sense ? accel_tmo : normal_tmo;
    } else {
        TEST_LOG_INFO("Failed to set accelerated aging control %u\n",
                      enable_sense);
        failures.counters.accel_control++;
    }
}

void 
aging_tolerance_t::create_id_map_insert(uint32_t id)
{
    if (create_id_map_mode == CREATE_ID_MAP_WITH_IDS) {
         if (!create_id_map.insert(id)) {

             /*
              * We can tolerate id_map insertion failure for 2 reasons:
              * 1) A test might deliberately try to create the same flow ID
              *    multiple times, or
              * 2) A test using random IDs may end up getting same IDs
              *    randomly generated
              *
             TEST_LOG_INFO("entry_id %u could not be added to id_map\n", id);
             failures.counters.create_add++;
             */
         }
    } else {

        /*
         * Count-only mode should not have been used for the 2 types of
         * test stated above or the test could fail.
         */
        create_count_++;
    }
}

void 
aging_tolerance_t::create_id_map_find_erase(uint32_t id)
{
    if (create_id_map_mode == CREATE_ID_MAP_WITH_IDS) {
        if (!create_id_map.find_erase(id)) {
            TEST_LOG_INFO("entry_id %u was not created for this test\n", id);
            failures.counters.create_erase++;
        }
    } else if (create_count_) {
        create_count_--;
    } else {
        TEST_LOG_INFO("entry_id %u not removable (empty id_map)\n", id);
        failures.counters.create_erase++;
    }
}

uint32_t 
aging_tolerance_t::create_id_map_size(void)
{
    return create_id_map_mode == CREATE_ID_MAP_WITH_IDS ?
           create_id_map.size() : create_count_;
}

void 
aging_tolerance_t::create_id_map_empty_check(void)
{
    if (create_id_map_size()) {
        TEST_LOG_INFO("Not all entries were aged out, remaining count: %u\n",
                      create_id_map_size());
        failures.counters.create_empty++;
    }
}

/*
 * Validate that HW declared the session timeout within a certain tolerance.
 */
void 
aging_tolerance_t::session_tmo_tolerance_check(uint32_t id)
{
    pds_flow_session_key_t  key;
    pds_flow_session_info_t info;

    /*
     * Only HW has real timestamp applicable for tolerance check
     */
    if (hw()) {
        flow_session_key_init(&key);
        key.session_info_id = id;
        if (pds_flow_session_info_read(&key, &info) == SDK_RET_OK) {
            tmo_tolerance_check(id, info.status.timestamp,
                                curr_tmo.session_tmo_get());
        } else {
            failures.counters.info_read++;
        }
    }
}

/*
 * Validate that HW declared the conntrack entry timeout within a certain tolerance.
 */
void 
aging_tolerance_t::conntrack_tmo_tolerance_check(uint32_t id)
{
    pds_conntrack_key_t     key;
    pds_conntrack_info_t    info;
    uint32_t                applic_tmo;

    /*
     * Only HW has real timestamp applicable for tolerance check
     */
    if (hw()) {
        flow_conntrack_key_init(&key);
        key.conntrack_id = id;
        if (pds_conntrack_state_read(&key, &info) == SDK_RET_OK) {
            applic_tmo = curr_tmo.conntrack_tmo_get(info.spec.data.flow_type,
                                                    info.spec.data.flow_state);
            tmo_tolerance_check(id, info.status.timestamp, applic_tmo);
        } else {
            failures.counters.info_read++;
        }
    }
}

void 
aging_tolerance_t::tmo_tolerance_check(uint32_t id,
                                       uint32_t entry_ts,
                                       uint32_t applic_tmo_secs)
{
    uint32_t    curr_ts = mpu_timestamp();
    uint32_t    delta_secs;

    delta_secs = mpu_timestamp2secs(curr_ts - entry_ts);
    if (delta_secs < applic_tmo_secs) {
        TEST_LOG_INFO("entry_id %u aged out in less than timeout of %u seconds "
                      "(actual: %u)\n", id, applic_tmo_secs, delta_secs);
        failures.counters.ts_tolerance++;
    } else if ((delta_secs - applic_tmo_secs) > tolerance_secs) {
        TEST_LOG_INFO("entry_id %u took extra %u seconds to age out "
                      "(tolerance is %u seconds)\n",
                      id, delta_secs - applic_tmo_secs, tolerance_secs);
        failures.counters.ts_tolerance++;
    }
}

/*
 * On Capri, P4 updates flow timestamp using bits 47:23 of the MPU tick.
 * With a clock speed of 833MHz, this gives interval of 1.01E-02 (which is
 * very close to 1/100 s, i.e., 10ms). 
 * 
 */
uint32_t
mpu_timestamp(void)
{
    return 0; /* temporary */
}

uint32_t
mpu_timestamp2secs(uint32_t mpu_timestamp)
{
    return mpu_timestamp / 100;
}


}    // namespace athena_app
}    // namespace test

