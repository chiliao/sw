# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MAKEDEFS}/pre.mk
MODULE_TARGET = tcam_test.gtest
MODULE_SOLIBS = tcam indexer ht logger shmmgr p4pd_mock
MODULE_LDLIBS = rt
include ${MAKEDEFS}/post.mk
