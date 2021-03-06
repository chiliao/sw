//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//
// FTP Utility APIs and data structures
//-----------------------------------------------------------------------------

namespace hal {
namespace plugins {
namespace alg_ftp {

/*
 * Forward Declaration
 */
typedef struct ftp_info_ ftp_info_t;

/*
 * Constants
 */

#define FTP_MAX_REQ  4

#ifdef LOGIN_ERR_NEEDED
#define FTP_ERR_RSP  10
#define FTP_MAX_RSP  13 
#else
#define FTP_ERR_RSP  5
#define FTP_MAX_RSP  8
#endif

/*
 * FTP Proto States
 */
#define FTP_STATE(ENTRY)                                  \
    ENTRY(FTP_INIT,           0,  "FTP_INIT")             \
    ENTRY(FTP_PORT,           1,  "FTP_PORT")             \
    ENTRY(FTP_EPRT,           2,  "FTP_EPRT")             \
    ENTRY(FTP_PASV,           3,  "FTP_PASV")             \
    ENTRY(FTP_EPSV,           4,  "FTP_EPSV")             \
    ENTRY(FTP_USER,           5,  "FTP_USER")             \
    ENTRY(FTP_MORE_INFO_PASS, 6,  "FTP_MORE_INFO_PASS")   \
    ENTRY(FTP_PASS,           7,  "FTP_PASS")             \
    ENTRY(FTP_MORE_INFO_ACCT, 8,  "FTP_MORE_INFO_ACCT")   \
    ENTRY(FTP_ACCT,           9,  "FTP_ACCT")             \
    ENTRY(FTP_SYNTAX_ERR,     10, "FTP_SYNTAX_ERR")       \
    ENTRY(FTP_TRANS_ERR,      11, "FTP_TRANS_ERR")        \
    ENTRY(FTP_ERROR_RSP,      12, "FTP_ERR_RSP")          \

DEFINE_ENUM(ftp_state_t, FTP_STATE)
#undef FTP_STATE

#define min_t(type, x, y) ({      \
    type __min1 = (x);            \
    type __min2 = (y);            \
    __min1 < __min2 ? __min1: __min2; })

typedef size_t (*ftp_callback_t) (void *ctx, uint8_t *payload, size_t len);

/*
 * Data Structures
 */

typedef struct ftp_info_ {
     ftp_state_t     state;
     uint8_t         isIPv6;
     ipvx_addr_t     sip;
     ipvx_addr_t     dip;
     uint16_t        sport;
     uint16_t        dport;
     bool            add_exp_flow;
     ftp_callback_t  callback;
     bool            skip_sfw;
     bool            allow_mismatch_ip_address;
     uint32_t        parse_errors;
     uint32_t        login_errors;
} ftp_info_t;


typedef int (*parse_cb_t)(const char *, uint32_t dlen, char,
                          uint32_t *offset, ftp_info_t *ftp_info);

typedef struct ftp_search_ {
    const char *pattern;
    uint32_t plen;
    char skip;
    char term;
    ftp_state_t state;
    parse_cb_t  cb;
} ftp_search_t;

#define __FTP_CMD(__pattern, __skip, __term, __state, __cb)   \
{                                                             \
       .pattern = (__pattern),                                \
       .plen    = sizeof(__pattern) - 1,                      \
       .skip    = (__skip),                                   \
       .term    = (__term),                                   \
       .state   = (__state),                                  \
       .cb      = (__cb),                                     \
}

/*
 * Function declarations
 */
size_t __parse_ftp_req(void *ctx, uint8_t *payload, size_t data_len);
size_t __parse_ftp_rsp(void *ctx, uint8_t *payload, size_t data_len);

}  // namespace alg_ftp
}  // namespace plugins
}  // namespace hal
