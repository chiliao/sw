#! /bin/bash

export PIPELINE=apollo
export NICDIR=`pwd`

# initial setup
source ${NICDIR}/apollo/test/tools/setup_gtests.sh
setup

# unset mock mem mode for valgrind
unset ASIC_MOCK_MEMORY_MODE

# run all gtests
run_valgrind_gtest scale CFG=" -i ${NICDIR}/apollo/test/scale/scale_cfg_1vpc.json "

# end of script
clean_exit
