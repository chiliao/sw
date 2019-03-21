// {C} Copyright 2019 Pensando Systems Inc. All rights reserved

#ifndef __HAL_PD_IRIS_FLOW_PD_HPP__
#define __HAL_PD_IRIS_FLOW_PD_HPP__

#include "nic/include/base.hpp"
#include "nic/sdk/lib/table/memhash/mem_hash.hpp"
#include "gen/proto/system.pb.h"
#include "gen/proto/table.pb.h"

using sdk::table::mem_hash;
using table::TableResponse;

namespace hal {
namespace pd {

class flow_table_pd {
public:
    static flow_table_pd *factory();
    static void destroy(flow_table_pd *ftpd);

    flow_table_pd() {}
    ~flow_table_pd() {}

    hal_ret_t init();
    hal_ret_t insert(void *key, void *data,
                     uint32_t *hash_value, bool hash_valid);
    hal_ret_t remove(void *key);
    
    hal_ret_t meta_get(table::TableMetadataResponseMsg *rsp_msg);
    hal_ret_t stats_get(sys::TableStatsEntry *stats_entry);
    hal_ret_t dump(TableResponse *rsp);

private:
    mem_hash *table_;
    std::string table_name_;
    uint32_t table_size_;
    uint32_t oflow_table_size_;

};

} // namespace pd
} // namespace hal

#endif // __HAL_PD_IRIS_FLOW_PD_HPP__
