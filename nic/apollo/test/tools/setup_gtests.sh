#! /bin/bash

set -e

export ASIC="${ASIC:-capri}"
CUR_DIR=$( readlink -f $( dirname $0 ) )
source $CUR_DIR/../../../tools/setup_env_mock.sh $PIPELINE

export SKIP_VERIFY=1
export GEN_TEST_RESULTS_DIR=${BUILD_DIR}/gtest_results
export VAL_CMD=valgrind
export VPP_IPC_MOCK_MODE=1

function finish () {
    # auto invoked on any exit
    ${PDSPKG_TOPDIR}/tools/print-cores.sh
    echo "===== Collecting logs ====="
    ${PDSPKG_TOPDIR}/apollo/test/tools/savelogs.sh
    rm -f ${PDSPKG_TOPDIR}/conf/pipeline.json
    sudo rm -f /tmp/*.db /tmp/pen_* /dev/shm/pds_* /dev/shm/ipc_*
    sudo rm -f /dev/shm/metrics_*
    if [ $PIPELINE == 'apulu' ]; then
        if [ -z "$VPP_IPC_MOCK_MODE" ]; then
            sudo pkill -9 vpp
        fi
        sudo pkill -9 dhcpd
    fi
}
trap finish EXIT

function setup () {
    sudo rm -rf ${PDSPKG_TOPDIR}/*log* ${PDSPKG_TOPDIR}/core*
    sudo rm -rf ${PDSPKG_TOPDIR}/conf/pipeline.json
    sudo rm -rf /tmp/pen_* /dev/shm/pds_* /dev/shm/ipc_*
    ln -s ${PDSPKG_TOPDIR}/conf/${PIPELINE}/pipeline.json ${PDSPKG_TOPDIR}/conf/pipeline.json

    echo "xxx"
    if [ $PIPELINE == 'apulu' ]; then
        if [ -z "$VPP_IPC_MOCK_MODE" ]; then
            echo "Starting VPP"
            sudo ${PDSPKG_TOPDIR}/vpp/tools/start-vpp-mock.sh --pipeline apulu
            if [[ $? != 0 ]]; then
                echo "Failed to bring up VPP"
                exit -1
            fi
        fi

        echo "Starting dhcpd"
        sudo ${PDSPKG_TOPDIR}/apollo/tools/apulu/start-dhcpd-sim.sh -p apulu
        if [[ $? != 0 ]]; then
            echo "Failed to bring up dhcpd"
            exit -1
        fi
    fi
}

function run_gtest () {
    TEST_OBJECT=$1
    TEST_NAME=${PIPELINE}_${TEST_OBJECT}_test
    TEST_LOG=${TEST_NAME}_log.txt
    for cmdargs in "$@"
    do
        arg=$(echo $cmdargs | cut -f1 -d=)
        val=$(echo $cmdargs | cut -f2 -d=)
        case "$arg" in
            LOG) TEST_LOG=${val};;
            CFG) TEST_CFG=${val};;
            *)
        esac
    done
    echo "`date +%x_%H:%M:%S:%N` : Running ${TEST_NAME}  > ${TEST_LOG} "
    if [ -z "$GDB" ]; then
        ${TEST_NAME} -c hal.json ${TEST_CFG} --gtest_output="xml:${GEN_TEST_RESULTS_DIR}/${TEST_NAME}.xml" > ${TEST_LOG};
    else
        $GDB ${TEST_NAME} -c hal.json ${TEST_CFG} --gtest_output="xml:${GEN_TEST_RESULTS_DIR}/${TEST_NAME}.xml";
    fi
    [[ $? -ne 0 ]] && echo "${TEST_NAME} failed!" && exit 1
    return 0
}

function run_sdk_gtest () {
    TEST_NAME=${PIPELINE}_sdk_test
    echo "`date +%x_%H:%M:%S:%N` : Running ${TEST_NAME}"
    ${PDSPKG_TOPDIR}/sdk/tools/run_sdk_gtests.sh
    [[ $? -ne 0 ]] && echo "${TEST_NAME} failed!" && exit 1
    return 0
}

function run_valgrind_gtest () {
    TEST_OBJECT=$1
    TEST_NAME=${PIPELINE}_${TEST_OBJECT}_test
    TEST_LOG=${PDSPKG_TOPDIR}/valgrind_${TEST_NAME}_log.txt
    # TODO: check if function arg parsing can be re-used
    for cmdargs in "$@"
    do
        arg=$(echo $cmdargs | cut -f1 -d=)
        val=$(echo $cmdargs | cut -f2 -d=)
        case "$arg" in
            CFG) TEST_CFG=${val};;
            *)
        esac
    done
    echo "`date +%x_%H:%M:%S:%N` : Running ${TEST_NAME} > ${TEST_LOG}.stdout "
    ${VAL_CMD} --track-origins=yes --error-limit=no --leak-check=summary --show-leak-kinds=definite -v --log-file=${TEST_LOG} --gen-suppressions=all --suppressions=${PDSPKG_TOPDIR}/apollo/test/tools/valgrind_suppression.txt ${PDSPKG_TOPDIR}/build/x86_64/${PIPELINE}/${ASIC}/bin/${TEST_NAME} -c hal.json ${TEST_CFG} --gtest_output="xml:${GEN_TEST_RESULTS_DIR}/${TEST_NAME}.xml" > ${TEST_LOG}.stdout
    # $GDB ${PDSPKG_TOPDIR}/build/x86_64/${PIPELINE}/${ASIC}/bin/${TEST_NAME} -c hal.json ${TEST_CFG} --gtest_output="xml:${GEN_TEST_RESULTS_DIR}/${TEST_NAME}.xml"

    # check valgrind log for leaks
    python ${PDSPKG_TOPDIR}/apollo/test/tools/parse_valgrind_log.py ${TEST_LOG}
    [[ $? -ne 0 ]] && echo "${TEST_NAME} failed!" && exit 1
    return 0
}

function clean_exit () {
    # to be invoked ONLY for successful run
    echo "Success"
    exit 0
}
