//-----------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------
#ifndef __FTLV4_STATS_HPP__
#define __FTLV4_STATS_HPP__

#include "include/sdk/base.hpp"
#include "include/sdk/table.hpp"
#include <string>

#include "ftlv4_utils.hpp"

using namespace std;

namespace sdk {
namespace table {
namespace ftlint_ipv4 {

class ftlv4_api_stats {
private:
    uint32_t insert_;
    uint32_t insert_duplicate_;
    uint32_t insert_fail_;
    uint32_t remove_;
    uint32_t remove_not_found_;
    uint32_t remove_fail_;
    uint32_t update_;
    uint32_t update_fail_;
    uint32_t get_;
    uint32_t get_fail_;
    uint32_t release_;
    uint32_t release_fail_;

public:
    ftlv4_api_stats() {
        insert_ = 0;
        insert_duplicate_ = 0;
        insert_fail_ = 0;
        remove_ = 0;
        remove_not_found_ = 0;
        remove_fail_ = 0;
        update_ = 0;
        update_fail_ = 0;
        get_ = 0;
        get_fail_ = 0;
        release_ = 0;
        release_fail_ = 0;
    }

    ~ftlv4_api_stats() {
    }

    sdk_ret_t insert(sdk_ret_t status) {
        //FTLV4_TRACE_VERBOSE("Updating insert stats, ret:%d", status);
        if (status == SDK_RET_OK) {
            insert_++;
        } else if (status == SDK_RET_ENTRY_EXISTS) {
            insert_duplicate_++;
        } else {
            insert_fail_++;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t update(sdk_ret_t status) {
        //FTLV4_TRACE_VERBOSE("Updating update stats, ret:%d", status);
        if (status == SDK_RET_OK) {
            update_++;
        } else {
            update_fail_++;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t remove(sdk_ret_t status) {
        //FTLV4_TRACE_VERBOSE("Updating remove stats, ret:%d", status);
        if (status == SDK_RET_OK) {
            remove_++;
        } else if (status == SDK_RET_ENTRY_NOT_FOUND) {
            remove_not_found_++;
        } else {
            remove_fail_++;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t release(sdk_ret_t status) {
        //FTLV4_TRACE_VERBOSE("Updating release stats, ret:%d", status);
        if (status == SDK_RET_OK) {
            release_++;
        } else {
            release_fail_++;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t get(sdk_ret_t status) {
        //FTLV4_TRACE_VERBOSE("Updating get stats, ret:%d", status);
        if (status == SDK_RET_OK) {
            get_++;
        } else {
            get_fail_++;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t get(sdk_table_api_stats_t *stats) {
        stats->insert = insert_;
        stats->insert_duplicate = insert_duplicate_;
        stats->insert_fail = insert_fail_;
        stats->remove = remove_;
        stats->remove_not_found = remove_not_found_;
        stats->remove_fail = remove_fail_;
        stats->update = update_;
        stats->update_fail = update_fail_;
        stats->get = get_;
        stats->get_fail = get_fail_;
        stats->release = release_;
        stats->release_fail = release_fail_;
        return SDK_RET_OK;
    }
};

class ftlv4_table_stats {
private:
    uint32_t    entries_;
    uint32_t    hints_;

public:
    ftlv4_table_stats() {
        entries_ = 0;
        hints_ = 0;
    }

    ~ftlv4_table_stats() {
    }

    sdk_ret_t insert(bool is_hint) {
        entries_++;
        if (is_hint) {
            hints_++;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t remove(bool is_hint) {
        SDK_ASSERT(entries_);
        entries_--;
        if (is_hint) {
            SDK_ASSERT(hints_);
            hints_--;
        }
        return SDK_RET_OK;
    }

    sdk_ret_t get(sdk_table_stats_t *stats) {
        stats->entries = entries_;
        stats->collisions = hints_;
        return SDK_RET_OK;
    }
};

} // namespace ftlint_ipv4
} // namespace table
} // namespace sdk

#endif // __FTLV4_HPP__
