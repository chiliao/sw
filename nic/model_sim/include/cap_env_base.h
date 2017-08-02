// -*- C++ -*-
//************************************************************
// (C)  Copyright 2017 Pensando Systems Inc. All rights reserved.
//************************************************************
#ifndef _CAP_ENV_BASE_H_
#define _CAP_ENV_BASE_H_
#include "cap_model_base.h"
#include "cap_blk_env_base.h"
#include <vector>
#include <string>
class cap_top_csr_t;
class cap_env_base : public cap_blk_env_base {
protected:
    cap_model_base * cap_mod;
    int my_id;
 public:
    cap_top_csr_t * cap0_ptr;
    cap_env_base(int id);
    // Call this function if you want to run just StandAlone Model without RTL.
    virtual void SAM_mode(void);

    virtual void load_cfg();
    virtual void load_prog();
    virtual void load_debug();

    // Push the network packet to capri
    virtual void step_network_pkt(const std::vector<uint8_t> & pkt, uint32_t port);
    // Get next packet from one of the ethernet ports. It returns the
    // port number and COS.
    virtual bool get_next_pkt(std::vector<uint8_t> &pkt, uint32_t &port, uint32_t& cos);

    // Read capri register
    virtual bool read_reg (uint64_t addr, uint32_t& data);
    // Write capri register
    virtual bool write_reg(uint64_t addr, uint32_t  data);

    // Memory access:
    // Address map will define if it goes to HBM model or host
    // memory. Currently MSB of address indicates it going to "host
    // memory"
    // Range 0x80000000 to 0x80000000+0xffffffff is HBM
    // The size and address should be such that the data being written
    // or read does not cross 4KB boundary. Data buffer is allocated
    // by the caller.
    virtual bool read_mem(uint64_t addr, uint8_t * data, uint32_t size);
    virtual bool write_mem(uint64_t addr, uint8_t * data, uint32_t size);
    virtual void step_doorbell(uint64_t addr, uint64_t data);
    virtual void init(void);
    virtual int get_chip_id(void)
    {
        return my_id;
    }

    virtual ~cap_env_base();
};
#endif
