# {C} Copyright 2018 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
ifeq "${P4VER}" "P4_16"
MODULE_PIPELINE   = athena_dontuse
MODULE_TARGET     = athena_p4.dontuse
else
MODULE_TARGET   = athena_p4.p4bin
MODULE_SRCS     = ${MODULE_SRC_DIR}/athena.p4
MODULE_PIPELINE = athena
MODULE_NCC_OPTS = --pipeline=athena --asm-out --gen-dir ${BLD_P4GEN_DIR} \
                  --cfg-dir ${BLD_PGMBIN_DIR}/athena_p4 \
                  --i2e-user --fe-flags="-I${TOPDIR} -I${SDKDIR}"
MODULE_DEPS     = $(shell find ${MODULE_DIR} -name '*.p4' -o -name '*.h') \
                  $(shell find ${MODULE_DIR}/../include -name '*.p4' -o -name '*.h') \
                  $(shell find ${TOPDIR}/nic/p4/include -name '*')
endif
MODULE_POSTGEN_MK = module_p4pd.mk
include ${MKDEFS}/post.mk

