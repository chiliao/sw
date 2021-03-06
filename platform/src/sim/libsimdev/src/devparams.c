/*
 * Copyright (c) 2017, Pensando Systems Inc.
 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

#include "nic/sdk/platform/misc/include/misc.h"
#include "nic/sdk/platform/misc/include/bdf.h"
#include "nic/sdk/platform/misc/include/maclib.h"
#include "simdev_impl.h"

/*
 * type=eth,bdf=03:00.0,mac=02:00:00:00:00:01,lif=4,intr_base=2
 */

int
devparam_str(const char *devparams,
             const char *key, char *buf, const size_t bufsz)
{
    const int keylen = strlen(key);
    char *ldevparams, *dp, *p, *v, *sp = NULL;
    int found;

    /*
     * Make a local copy of devparams so strtok() can write to it.
     */
    dp = ldevparams = strdup(devparams);

    found = 0;
    while ((p = strtok_r(dp, ",", &sp)) != NULL && !found) {
        if (strncmp(p, key, keylen) == 0 && p[keylen] == '=') {
            if (buf != NULL && bufsz) {
                v = &p[keylen + 1];
                strncpy0(buf, v, bufsz);
            }
            found = 1;
        }
        if (dp != NULL) dp = NULL;
    }
    free(ldevparams);
    return found ? 0 : -1;
}

int
devparam_int(const char *devparams, const char *key, int *val)
{
    char str[32];

    if (devparam_str(devparams, key, str, sizeof(str)) >= 0) {
        char *ep = NULL;
        int v = strtol(str, &ep, 0);
        if (*ep == '\0') {
            if (val != NULL) {
                *val = v;
            }
            return 0;
        }
    }
    return -1;
}

int
devparam_u64(const char *devparams, const char *key, u_int64_t *val)
{
    char str[32];

    if (devparam_str(devparams, key, str, sizeof(str)) >= 0) {
        char *ep = NULL;
        u_int64_t v = strtoull(str, &ep, 0);
        if (*ep == '\0') {
            if (val != NULL) {
                *val = v;
            }
            return 0;
        }
    }
    return -1;
}

int
devparam_bdf(const char *devparams, const char *key, u_int16_t *bdf)
{
    char str[32];

    if (devparam_str(devparams, key, str, sizeof(str)) >= 0) {
        int b = bdf_from_str(str);
        if (b != -1) {
            if (bdf != NULL) {
                *bdf = b;
            }
            return 0;
        }
    }
    return -1;
}

int
devparam_mac(const char *devparams, const char *key, mac_t *m)
{
    char macstr[32];
    mac_t lmac;

    if (devparam_str(devparams, key, macstr, sizeof(macstr)) >= 0 &&
        mac_from_str(&lmac, macstr) >= 0) {
        if (m != NULL) {
            memcpy(m, &lmac, sizeof(mac_t));
        }
        return 0;
    }
    return -1;
}

void
devparam_get_value(const char *devparams, const char *key,
                   char *buf, const size_t bufsz, char *dval)
{
    if (devparam_str(devparams, key, buf, bufsz) < 0) {
        strncpy0(buf, dval, bufsz);
    }
}

void
devparam_get_int(const char *devparams, const char *key,
                 int *val, int dval)
{
    if (devparam_int(devparams, key, val) < 0) {
        *val = dval;
    }
}

void
devparam_get_bdf(const char *devparams, const char *key,
                 u_int16_t *bdf, u_int16_t dbdf)
{
    if (devparam_bdf(devparams, key, bdf) < 0) {
        *bdf = dbdf;
    }
}

void
devparam_get_mac(const char *devparams, const char *key,
                 mac_t *m, mac_t *dmac)
{
    if (devparam_mac(devparams, key, m) < 0) {
        memcpy(m, dmac, sizeof(mac_t));
    }
}
