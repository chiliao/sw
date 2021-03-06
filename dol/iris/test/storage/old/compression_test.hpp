// Configuration of compression test.

#ifndef _COMPRESSION_TEST_HPP_
#define _COMPRESSION_TEST_HPP_

#include "dol/test/storage/compression.hpp"

#include <stdint.h>
#include <string>

namespace tests {

static constexpr uint16_t kCPVersion = 0x1234;
static constexpr uint32_t kADLER32Sum = 0x6802be13;
static constexpr uint32_t kCRC32Sum = 0xb12e7b4a;

static constexpr uint16_t kUncompressedDataSize = 4096;
extern uint8_t uncompressed_data[kUncompressedDataSize];

static constexpr uint16_t kCompressedDataSize = 2208;
extern uint8_t compression_hdr[8];
extern uint8_t compressed_data[kCompressedDataSize];

extern uint8_t sha256_post[32];
extern uint8_t sha512_post[64];
static const uint32_t kDataoutBufSize = 4096;
extern uint8_t all_zeros[kDataoutBufSize];

// Testcase spec.
// Src and dst are always 4K buffers. The test code will create
// SGLs (or use plain lists) based on datain_len and dataout_len.
typedef struct comp_test {
  std::string test_name;
  union {
    ccmd_t   cmd_bits;
    uint16_t cmd;
  };
  uint8_t num_src_sgls;
  uint8_t num_dst_sgls;
  uint32_t datain_size;
  uint32_t dataout_size;
  uint16_t comp_threshold;
  uint16_t src_is_hbm:1,
           src_sgl_is_hbm:1,
           dst_is_hbm:1,
           dst_sgl_is_hbm:1,
           status_is_hbm:1,
           output_same_as_input:1,
           unused:1;
} comp_test_t;

void compression_init();
void decompression_init();
int compress_host_flat();
int compress_hbm_flat();
int compress_host_to_hbm_flat();
int compress_hbm_to_host_flat();
int compress_host_sgl();
int compress_hbm_sgl();
int compress_host_nested_sgl();
int compress_hbm_nested_sgl();
int compress_nested_sgl_in_hbm();
int compress_return_through_hbm();
int compress_adler_sha256();
int compress_crc_sha512();
int compress_doorbell_odata();
int compress_max_features();
int compress_output_through_sequencer();
int compress_flat_host_buf_performance();

int decompress_host_flat();
int decompress_hbm_flat();
int decompress_host_to_hbm_flat();
int decompress_hbm_to_host_flat();
int decompress_host_sgl();
int decompress_hbm_sgl();
int decompress_host_nested_sgl();
int decompress_hbm_nested_sgl();
int decompress_nested_sgl_in_hbm();
int decompress_return_through_hbm();
int decompress_adler();
int decompress_crc();
int decompress_doorbell_odata();
int decompress_flat_host_buf_performance();

}  // namespace tests

#endif  // _COMPRESSION_TEST_HPP_
