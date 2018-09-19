# env
export NIC_DIR='/nic'
export HAL_CONFIG_PATH=$NIC_DIR/conf/
export LD_LIBRARY_PATH=$NIC_DIR/lib:$NIC_DIR/conf/sdk:$NIC_DIR/conf/linkmgr:$NIC_DIR/conf/sdk/external:/usr/local/lib:/usr/lib/aarch64-linux-gnu:$LD_LIBRARY_PATH

# start linkmgr
ulimit -c unlimited
$GDB $NIC_DIR/bin/linkmgr 2>&1 > /run.log &

sleep 10

# start AACS server
kill -STOP `pidof hal`
/nic/bin/port_client -g localhost:50053 -o 26 -w 9000
kill -CONT `pidof hal`
