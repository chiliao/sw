#! /bin/bash

set -e
export NIC_DIR=/nic/
export CONFIG_PATH=$NIC_DIR/conf/
export NON_PERSISTENT_LOG_DIR=/var/log/pensando/
export HAL_LIBRARY_PATH=$NIC_DIR/lib:$NIC_DIR/../platform/lib:/usr/local/lib:/usr/lib/aarch64-linux-gnu:$LD_LIBRARY_PATH
export HAL_PBC_INIT_CONFIG="2x100_hbm"
export HAL_PLATFORM_HW=1
$NIC_DIR/bin/artemis_scale_test -d -c hal_hw.json -i /nic/conf/scale_cfg.json > /data/logs.txt &
