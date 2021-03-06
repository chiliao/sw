# {C} Copyright 2019 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET   = artemis_txdma.p4bin
MODULE_SRCS     = ${MODULE_SRC_DIR}/p4plus_txdma.p4
MODULE_PIPELINE = artemis
MODULE_NCC_OPTS = --pipeline=artemis --asm-out --p4-plus --no-ohi --two-byte-profile \
                  --pd-gen --p4-plus-module txdma \
                  --gen-dir ${BLD_P4GEN_DIR} \
                  --cfg-dir ${BLD_PGMBIN_DIR}/artemis_txdma \
                  --fe-flags="-I${TOPDIR} -I${SDKDIR}"
MODULE_DEPS     = $(shell find ${MODULE_DIR} -name '*.p4' -o -name '*.h') \
                  $(shell find ${MODULE_DIR}/../include -name '*.p4' -o -name '*.h') \
                  $(shell find ${TOPDIR}/nic/p4/include -name '*')
MODULE_POSTGEN_MK = module_p4pd.mk
include ${MKDEFS}/post.mk
