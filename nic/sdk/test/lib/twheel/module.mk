# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MAKEDEFS}/pre.mk
MODULE_TARGET = twheel_test.gtest
MODULE_SOLIBS = twheel slab shmmgr logger
MODULE_LDLIBS = rt
include ${MAKEDEFS}/post.mk
