# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MAKEDEFS}/pre.mk
MODULE_TARGET   = ipsec_p4plus_esp_v4_tun_ntoh.p4bin
MODULE_SRCS     = esp_v4_tunnel_n2h_rxdma.p4 esp_v4_tunnel_n2h_txdma1.p4 esp_v4_tunnel_n2h_txdma2.p4
MODULE_NCC_OPTS = --p4-plus --pd-gen --asm-out --no-ohi \
                  --two-byte-profile --fe-flags="-I${TOPDIR}" \
				  --gen-dir ${BLD_GEN_DIR}
MODULE_DEPS     = $(shell find ${MODULE_DIR}/ -name '*.p4' -o -name '*.h') \
                  $(shell find ${MODULE_DIR}/../ -name '*.p4' -o -name '*.h') \
                  $(shell find ${MODULE_DIR}/../../include -name '*') \
				  $(shell find ${MODULE_DIR}/../../common-p4+ -name '*.p4')
include ${MAKEDEFS}/post.mk
