# {C} Copyright 2019 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET       = _artemis_libp4plus_rxdma_p4pdcli.lib
MODULE_PIPELINE	    = artemis
MODULE_PREREQS      = artemis_commonrxdma_p4pd.swigcli
MODULE_SRC_DIR      = ${BLD_P4GEN_DIR}/p4plus_rxdma
MODULE_SRCS         = $(wildcard ${MODULE_SRC_DIR}/src/*p4pd_cli_backend.cc) \
                      $(wildcard ${MODULE_SRC_DIR}/src/*entry_packing.cc) \
                      $(wildcard ${MODULE_SRC_DIR}/cli/*.cc)
MODULE_INCS         = ${BLD_P4GEN_DIR}/p4plus_rxdma/cli \
                      ${CLI_P4PD_INCS}
MODULE_FLAGS        = ${CLI_P4PD_FLAGS}
MODULE_LDLIBS       = ${CLI_${PIPELINE}_P4PD_LDLIBS}
MODULE_SOLIBS       = ${CLI_${PIPELINE}_P4PD_SOLIBS}
include ${MKDEFS}/post.mk
