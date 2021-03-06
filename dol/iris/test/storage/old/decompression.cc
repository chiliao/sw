// Compression DOLs.
#include "dol/test/storage/compression.hpp"
#include "dol/test/storage/compression_test.hpp"
#include "dol/test/storage/utils.hpp"
#include "dol/test/storage/tests.hpp"
#include "third-party/asic/capri/design/common/cap_addr_define.h"
#include "third-party/asic/capri/model/cap_he/readonly/cap_hens_csr_define.h"

#include <assert.h>
#include <stdint.h>
#include <string.h>
#include <strings.h>
#include <unistd.h>
#include <stdio.h>
#include <byteswap.h>
#include "nic/utils/host_mem/c_if.h"
#include "nic/sdk/model_sim/include/lib_model_client.h"
#include "gflags/gflags.h"

DECLARE_uint64(long_poll_interval);

namespace tests {

static const uint64_t cfg_glob = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_GLB_BYTE_ADDRESS;

static const uint64_t cfg_dist = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_DIST_BYTE_ADDRESS;

static const uint64_t cfg_ueng = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_UENG_W0_BYTE_ADDRESS;

static const uint64_t cfg_q_base = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_Q_BASE_ADR_W0_BYTE_ADDRESS;

static const uint64_t cfg_hotq_base = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_HOTQ_BASE_ADR_W0_BYTE_ADDRESS;

static const uint64_t cfg_q_pd_idx = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_Q_PD_IDX_BYTE_ADDRESS;

static const uint64_t cfg_hotq_pd_idx = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_HOTQ_PD_IDX_BYTE_ADDRESS;

static const uint64_t cfg_host = CAP_ADDR_BASE_MD_HENS_OFFSET +
    CAP_HENS_CSR_DHS_CRYPTO_CTL_DC_CFG_HOST_BYTE_ADDRESS;

static const uint32_t kNumSubqEntries = 4096;
static const uint32_t kQueueMemSize = sizeof(cp_desc_t) * kNumSubqEntries;
static void *queue_mem;
static uint64_t queue_mem_pa;
static uint16_t queue_index = 0;

static const uint32_t kStatusBufSize = 4096;
static void *status_buf;
static uint64_t status_buf_pa;

static const uint32_t kDatainBufSize = 4096;
static void *datain_buf;
static uint64_t datain_buf_pa;

// Size defined in compression_test.hpp
static void *dataout_buf;
static uint64_t dataout_buf_pa;

static const uint32_t kSGLBufSize = 4096;
static uint8_t *sgl_buf;
static uint64_t sgl_buf_pa;

static uint64_t hbm_datain_buf_pa;
static uint64_t hbm_dataout_buf_pa;
static uint64_t hbm_status_buf_pa;
static uint64_t hbm_sgl_buf_pa;

static bool status_poll(bool in_hbm) {
  auto func = [in_hbm] () -> int {
    cp_status_sha512_t *s = (cp_status_sha512_t *)status_buf;
    if (in_hbm) {
      read_mem(hbm_status_buf_pa, (uint8_t *)status_buf, kStatusBufSize);
    }
    if (s->valid) {
      if (in_hbm) {
        usleep(100);
        // Do it once more as status is copied 1st and DB is rung next.
        read_mem(hbm_status_buf_pa, (uint8_t *)status_buf, kStatusBufSize);
      }
      printf("Got status %llx\n", *((unsigned long long *)s));
      return 0;
    }
    return 1;
  };
  tests::Poller poll;
  if (poll(func) == 0)
    return true;
  return false;
}

void
decompression_init()
{
  queue_mem = alloc_page_aligned_host_mem(kQueueMemSize);
  assert(queue_mem != nullptr);
  queue_mem_pa = host_mem_v2p(queue_mem);

  status_buf = alloc_page_aligned_host_mem(kStatusBufSize);
  assert(status_buf != nullptr);
  status_buf_pa = host_mem_v2p(status_buf);

  datain_buf = alloc_page_aligned_host_mem(kDatainBufSize);
  assert(datain_buf != nullptr);
  datain_buf_pa = host_mem_v2p(datain_buf);

  dataout_buf = alloc_page_aligned_host_mem(kDataoutBufSize);
  assert(dataout_buf != nullptr);
  dataout_buf_pa = host_mem_v2p(dataout_buf);

  sgl_buf = (uint8_t *)alloc_page_aligned_host_mem(kSGLBufSize);
  assert(sgl_buf != nullptr);
  sgl_buf_pa = host_mem_v2p(sgl_buf);

  assert(utils::hbm_addr_alloc_page_aligned(kDatainBufSize, &hbm_datain_buf_pa) == 0);
  assert(utils::hbm_addr_alloc_page_aligned(kDataoutBufSize, &hbm_dataout_buf_pa) == 0);
  assert(utils::hbm_addr_alloc_page_aligned(kStatusBufSize, &hbm_status_buf_pa) == 0);
  assert(utils::hbm_addr_alloc_page_aligned(kSGLBufSize, &hbm_sgl_buf_pa) == 0);

  // Pre-fill input buffers.
  bcopy(compressed_data, ((uint8_t *)datain_buf)+8, kCompressedDataSize);
  write_mem(hbm_datain_buf_pa+8, ((uint8_t *)datain_buf)+8, kCompressedDataSize);

  // Write queue base.
  uint32_t lo_reg, hi_reg;
  read_reg(cfg_glob, lo_reg);
  write_reg(cfg_glob, (lo_reg & 0xFFFF0000u) | kCPVersion);
  write_reg(cfg_q_base, queue_mem_pa & 0xFFFFFFFFu);
  write_reg(cfg_q_base + 4, (queue_mem_pa >> 32) & 0xFFFFFFFFu);
  // Enable all 16 engines.
  read_reg(cfg_ueng, lo_reg);
  read_reg(cfg_ueng+4, hi_reg);
  lo_reg |= 0x3;
  hi_reg &= ~(1u << (54 - 32));
  hi_reg &= ~(1u << (53 - 32));
  write_reg(cfg_ueng, lo_reg);
  write_reg(cfg_ueng+4, hi_reg);
  // Enable cold/warm queue.
  read_reg(cfg_dist, lo_reg);
  lo_reg |= 1;
  write_reg(cfg_dist, lo_reg);

  queue_index = 0;

  printf("Decompression init done\n");
}

static void populate_sgls(uint16_t data_len, uint16_t num_entries,
                          void *sgl_va_in, uint64_t sgl_pa, uint64_t src_pa) {
  bool sgl_in_host_mem = sgl_pa & 0x8000000000000000ull;
  uint16_t chunk_size = data_len / num_entries;
  uint16_t last_chunk_size = chunk_size + (data_len - (chunk_size * num_entries));
  assert(chunk_size > 0);
  uint16_t count = num_entries;
  cp_sgl_t *sgl_va = (cp_sgl_t *)sgl_va_in;
  cp_sgl_t sgl;
  while (count > 0) {
    bzero(&sgl, sizeof(sgl));
    sgl.link = sgl_pa + sizeof(cp_sgl_t);
    sgl.addr0 = src_pa;
    sgl.len0 = (count == 1) ? last_chunk_size : chunk_size;
    src_pa += sgl.len0;
    count--;
    if (count == 0)
      break;

    sgl.addr1 = src_pa;
    sgl.len1 = (count == 1) ? last_chunk_size : chunk_size;
    src_pa += sgl.len1;
    count--;
    if (count == 0)
      break;

    sgl.addr2 = src_pa;
    sgl.len2 = (count == 1) ? last_chunk_size : chunk_size;
    src_pa += sgl.len2;
    count--;
    if (count == 0)
      break;
    if (sgl_in_host_mem) {
      bcopy(&sgl, sgl_va, sizeof(sgl));
      sgl_va++;
    } else {
      write_mem(sgl_pa, (uint8_t *)&sgl, sizeof(sgl));
    }
    sgl_pa = sgl.link;
  }
  sgl.link = 0;
  if (sgl_in_host_mem) {
    bcopy(&sgl, sgl_va, sizeof(sgl));
  } else {
    write_mem(sgl_pa, (uint8_t *)&sgl, sizeof(sgl));
  }
}

static int run_dc_test(comp_test_t *params) {
  static uint32_t counter = 0;
  cp_hdr_t cp_hdr = {0};
  counter++;
  cp_desc_t d;
  bzero(&d, sizeof(d));
  if (params->cmd_bits.cksum_en) {
    if (params->cmd_bits.cksum_adler)
      cp_hdr.cksum = kADLER32Sum;
    else
      cp_hdr.cksum = kCRC32Sum;
  }
  cp_hdr.data_len = kCompressedDataSize;
  cp_hdr.version = kCPVersion;
  uint64_t *h = (uint64_t *)&cp_hdr;
  *((uint64_t *)datain_buf) = *h;
  write_mem(hbm_datain_buf_pa, (uint8_t *)h, sizeof(cp_hdr_t));

  printf("Starting testcase %s\n", params->test_name.c_str());
  d.cmd = params->cmd;
  uint64_t bufin_pa = params->src_is_hbm ? hbm_datain_buf_pa : datain_buf_pa;
  if (params->num_src_sgls == 1) {
    d.src = bufin_pa;
  } else {
    if (params->src_sgl_is_hbm) {
      populate_sgls(params->datain_len, params->num_src_sgls, nullptr,
                    hbm_sgl_buf_pa, bufin_pa);
      d.src = hbm_sgl_buf_pa;
    } else {
      populate_sgls(params->datain_len, params->num_src_sgls, sgl_buf,
                    sgl_buf_pa, bufin_pa);
      d.src = sgl_buf_pa;
    }
    d.cmd_bits.src_is_list = 1;
  }

  uint64_t bufout_pa = params->dst_is_hbm ? hbm_dataout_buf_pa : dataout_buf_pa;
  bzero(dataout_buf, kDataoutBufSize);
  if (params->dst_is_hbm) {
    write_mem(hbm_dataout_buf_pa, all_zeros, kDataoutBufSize);
  }
  if (params->num_dst_sgls == 1) {
    d.dst = bufout_pa;
  } else {
    if (params->dst_sgl_is_hbm) {
      populate_sgls(params->dataout_len, params->num_dst_sgls, nullptr,
                    hbm_sgl_buf_pa + 2048, bufout_pa);
      d.dst = hbm_sgl_buf_pa + 2048;
    } else {
      populate_sgls(params->dataout_len, params->num_dst_sgls, sgl_buf + 2048,
                    sgl_buf_pa + 2048, bufout_pa);
      d.dst = sgl_buf_pa + 2048;
    }
    d.cmd_bits.dst_is_list = 1;
  }

  bzero(status_buf, kStatusBufSize);
  if (params->status_is_hbm) {
    write_mem(hbm_status_buf_pa, all_zeros, kStatusBufSize);
  }
  const uint64_t kDBData = 0x11223344556677ull;
  const uint32_t kTagData = 0x8899aabbu;
  d.input_len = params->datain_len;
  d.expected_len = params->dataout_len;
  d.status_addr = params->status_is_hbm ? hbm_status_buf_pa : status_buf_pa;
  d.doorbell_addr = d.status_addr + 1024;
  d.doorbell_data = kDBData;
  d.opaque_tag_addr = d.status_addr + 2048;
  d.opaque_tag_data = kTagData;
  d.status_data = counter;

  cp_desc_t *dst_d = (cp_desc_t *)queue_mem;
  bcopy(&d, &dst_d[queue_index], sizeof(d));
  queue_index++;
  if (queue_index == 4096)
    queue_index = 0;
  write_reg(cfg_q_pd_idx, queue_index);

  cp_status_sha512_t *st = (cp_status_sha512_t *)status_buf;
  if (!status_poll(params->status_is_hbm)) {
    printf("ERROR: decompression status never came\n");
    return -1;
  }
  if (!st->valid) {
    printf("ERROR: status valid bit not set\n");
    return -1;
  }
  if (st->err) {
    printf("ERROR: decompression generated err = 0x%x\n", st->err);
    return -1;
  }
  if (st->output_data_len != kUncompressedDataSize) {
    printf("ERROR: output data len mismatch, expected %u, received %u\n",
           kUncompressedDataSize, st->output_data_len);
    return -1;
  }
  if (params->dst_is_hbm) {
    read_mem(hbm_dataout_buf_pa, (uint8_t *)dataout_buf, kUncompressedDataSize);
  }
  if (st->partial_data != counter) {
    printf("ERROR: partial data in status does not match the expected value\n");
    return -1;
  }
  if (bcmp(dataout_buf, uncompressed_data, kUncompressedDataSize) != 0) {
    printf("ERROR: uncompressed data does not match with expected output.\n");
    return -1;
  }
  uint64_t *db_data = (uint64_t *)(((uint8_t *)status_buf) + 1024);
  if (params->cmd_bits.doorbell_on) {
    auto func = [db_data] () ->int {
      if (*db_data == kDBData)
        return 0;
      return 1;
    };
    tests::Poller poll;
    if (poll(func) != 0) {
      printf("ERROR: doorbell is not rung\n");
      return -1;
    }
  } else {
    if (*db_data == kDBData) {
      printf("ERROR: doorbell is rung when not requested\n");
      return -1;
    }
  }
  uint32_t *o = (uint32_t *)(((uint8_t *)status_buf) + 2048);
  if (params->cmd_bits.opaque_tag_on) {
    auto func = [o] () -> int {
      if (*o == kTagData)
        return 0;
      return 1;
    };
    tests::Poller poll;
    if (poll(func) != 0) {
      printf("ERROR: Opaque tag not written\n");
      return -1;
    }
  } else {
    if (*o == kTagData) {
      printf("ERROR: Opaque tag is written when not requested\n");
      return -1;
    }
  }

  printf("Testcase %s passed\n", params->test_name.c_str());
  return 0;
}

int decompress_host_flat() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;

  return run_dc_test(&spec);
}

int decompress_hbm_flat() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.src_is_hbm   = 1;
  spec.dst_is_hbm   = 1;

  return run_dc_test(&spec);
}

int decompress_host_to_hbm_flat() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.dst_is_hbm   = 1;

  return run_dc_test(&spec);
}

int decompress_hbm_to_host_flat() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.src_is_hbm   = 1;

  return run_dc_test(&spec);
}

int decompress_host_sgl() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 3;
  spec.num_dst_sgls = 2;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;

  return run_dc_test(&spec);
}

int decompress_hbm_sgl() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 3;
  spec.num_dst_sgls = 2;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.src_is_hbm   = 1;
  spec.dst_is_hbm   = 1;

  return run_dc_test(&spec);
}

int decompress_host_nested_sgl() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 4;
  spec.num_dst_sgls = 4;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;

  return run_dc_test(&spec);
}

int decompress_hbm_nested_sgl() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 4;
  spec.num_dst_sgls = 4;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.src_is_hbm   = 1;
  spec.dst_is_hbm   = 1;

  return run_dc_test(&spec);
}

int decompress_nested_sgl_in_hbm() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 4;
  spec.num_dst_sgls = 4;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.src_is_hbm   = 1;
  spec.src_sgl_is_hbm   = 1;
  spec.dst_is_hbm   = 1;
  spec.dst_sgl_is_hbm   = 1;

  return run_dc_test(&spec);
}

int decompress_return_through_hbm() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.num_src_sgls = 4;
  spec.num_dst_sgls = 4;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;
  spec.dst_is_hbm   = 1;
  spec.dst_sgl_is_hbm   = 1;
  spec.status_is_hbm = 1;

  return run_dc_test(&spec);
}

int decompress_adler() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.cmd_bits.cksum_en = 1;
  spec.cmd_bits.cksum_adler = 1;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;

  return run_dc_test(&spec);
}

int decompress_crc() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.cmd_bits.cksum_en = 1;
  spec.cmd_bits.cksum_adler = 0;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;

  return run_dc_test(&spec);
}

int decompress_doorbell_odata() {
  comp_test_t spec;
  bzero(&spec, sizeof(spec));
  spec.test_name    = __func__;
  spec.cmd          = 2;
  spec.cmd_bits.doorbell_on = 1;
  spec.cmd_bits.opaque_tag_on = 1;
  spec.num_src_sgls = 1;
  spec.num_dst_sgls = 1;
  spec.datain_len   = kCompressedDataSize + 8;
  spec.dataout_len  = 4096;

  return run_dc_test(&spec);
}

// Run decompression continuously.
int decompress_flat_host_buf_performance() {
  cp_desc_t d;
  bzero(&d, sizeof(d));
  uint32_t tag_terminal_value = 63; // total 64 commands

  cp_hdr_t cp_hdr = {0};
  cp_hdr.cksum = kCRC32Sum;
  cp_hdr.data_len = kCompressedDataSize;
  cp_hdr.version = kCPVersion;
  uint64_t *h = (uint64_t *)&cp_hdr;
  *((uint64_t *)datain_buf) = *h;

  printf("Starting testcase decompress_flat_host_buf_performance\n");
  d.cmd_bits.header_present = 1;
  d.cmd_bits.cksum_en = 1;  // CRC32
  d.cmd_bits.opaque_tag_on = 1;
  d.src = datain_buf_pa;
  d.dst = dataout_buf_pa;
  d.input_len = kCompressedDataSize + 8;
  d.expected_len = 4096;
  d.status_addr = status_buf_pa;
  d.opaque_tag_addr = d.status_addr + 2048;

  cp_desc_t *dst_d = (cp_desc_t *)queue_mem;
  for (unsigned i = 0; i <= tag_terminal_value; i++) {
    d.opaque_tag_data = i;
    bcopy(&d, &dst_d[queue_index], sizeof(d));
    queue_index++;
    if (queue_index == 4096)
      queue_index = 0;
    write_reg(cfg_q_pd_idx, queue_index);
  }

  // Wait for opaque tag to reach terminal value.
  unsigned *tag_addr = (unsigned *)(((uint8_t *)status_buf) + 2048);
  auto func = [tag_addr, tag_terminal_value] () -> int {
    if (*tag_addr == tag_terminal_value)
      return 0;
    return 1;
  };
  tests::Poller poll(FLAGS_long_poll_interval);
  if (poll(func) != 0) {
    printf("testcase decompress_flat_host_buf_performance failed : poll timeout\n");
    return -1;
  }
  // Do status validation.
  cp_status_sha512_t *st = (cp_status_sha512_t *)status_buf;
  if (!st->valid) {
    printf("ERROR: status valid bit not set\n");
    return -1;
  }
  if (st->err) {
    printf("ERROR: decompression generated err = 0x%x\n", st->err);
    return -1;
  }
  if (st->output_data_len != kUncompressedDataSize) {
    printf("ERROR: output data len mismatch, expected %u, received %u\n",
           kUncompressedDataSize, st->output_data_len);
    return -1;
  }
  printf("testcase decompress_flat_host_buf_performance passed\n");
  return 0;
}

}  // namespace tests
