# {C} Copyright 2019 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
include $(TOPDIR)/nic/metaswitch/pre.mk
MODULE_TARGET   = libpdsapimock.lib
MODULE_PREREQS  = metaswitch.submake
MODULE_PIPELINE = apollo artemis apulu
MODULE_INCS 	= $(TOPDIR)/nic/metaswitch/stubs/hals \
		  $(addprefix $(MS_ROOT)/,$(MS_INCLPATH))
MODULE_FLAGS	= $(addprefix -D,$(MS_COMPILATION_SWITCH))
include ${MKDEFS}/post.mk
