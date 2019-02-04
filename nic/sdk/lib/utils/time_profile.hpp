// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
#include <time.h>
#include <sys/time.h>
#include <unistd.h>
#include <sys/resource.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>

namespace sdk {
namespace utils {
namespace time_profile {

//#define RUSAGE_PROFILING_ENABLE

#ifdef RUSAGE_PROFILING_ENABLE

class time_profile_info {
private:
    struct timespec before_;
    struct timespec after_;
    uint64_t total_;

public:
    time_profile_info() {
        total_ = 0;
    }
    void start();
    void stop();
    uint64_t total();
};

#define ENUM_ENTRY_LIST(ENUM_ENTRY) \
        ENUM_ENTRY(TABLE_LIB_MEMHASH_INSERT) \
        ENUM_ENTRY(P4PD_ENTRY_READ) \
        ENUM_ENTRY(P4PD_ENTRY_INSTALL) \
        ENUM_ENTRY(P4PD_HWKEY_HWMASK_BUILD) \
        ENUM_ENTRY(ASICPD_HBM_TABLE_ENTRY_READ) \
        ENUM_ENTRY(ASICPD_HBM_TABLE_ENTRY_WRITE) \
        ENUM_ENTRY(COMPUTE_CRC) \
        ENUM_ENTRY(PAL_MEM_RD) \
        ENUM_ENTRY(PAL_MEM_WR) \
        ENUM_ENTRY(TIME_PROFILE_ID_MAX)

#define GENERATE_ENUM(ENUM) ENUM,
#define GENERATE_STRING(STRING) #STRING,

typedef enum time_profile_id_s {
    ENUM_ENTRY_LIST(GENERATE_ENUM)
} time_profile_id_t;

#if 0
typedef enum time_profile_id_s {
    TABLE_LIB_MEMHASH_INSERT,
    P4PD_ENTRY_READ,
    P4PD_ENTRY_INSTALL,
    P4PD_HWKEY_BUILD,
    P4PD_ASICPD_HBM_ENTRY_READ,
    P4PD_ASICPD_HBM_ENTRY_WRITE,
    TIME_PROFILE_ID_MAX,
} time_profile_id_t;
#endif

extern time_profile_info time_profile_db[];

void print();

#define time_profile_begin(_id) \
        sdk::utils::time_profile::time_profile_db[_id].start();
#define time_profile_end(_id) \
        sdk::utils::time_profile::time_profile_db[_id].stop();
#define time_profile_print() \
        sdk::utils::time_profile::print();
#else
#define time_profile_begin(_id)
#define time_profile_end(_id)
#define time_profile_print()
#endif

} // namespace time_profile
} // namespace utils
} // namespace sdk
