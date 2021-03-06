# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MKDEFS}/pre.mk
ifeq "${P4VER}" "P4_14"
MODULE_TARGET       = iris_p4pd.swigcli
MODULE_PIPELINE     = iris
MODULE_PREREQS      = iris.p4bin
MODULE_SRC_DIR      = ${BLD_P4GEN_DIR}/p4/cli
MODULE_SRCS         = ${MODULE_SRC_DIR}/*.i
MODULE_FLAGS        = -c++ -python
MODULE_POSTGEN_MK   = module_p4pd_cli.mk
else
MODULE_PIPELINE     = iris_dontuse
MODULE_TARGET       = p4pd_swigcli.dontuse
endif
include ${MKDEFS}/post.mk
