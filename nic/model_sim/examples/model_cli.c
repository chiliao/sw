//  Hello World client
#include <zmq.h>
#include <string.h>
#include <stdio.h>
#include <unistd.h>
#include <assert.h>
#include "buf_hdr.h"
#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <cstdlib>
#include <assert.h>
#include <iostream>
#include "lib_model_client.h"

void print_banner ()
{
    printf("*************************************\n");
    printf("1. Read Reg (0xaddr n_words)\n");
    printf("2. Write Reg (0xaddr 0xdata)\n");
    printf("3. Read Mem (0xaddr size)\n");
    printf("4. Dump HBM\n");
    printf("5. Exit\n");
    printf("*************************************\n");
    return;
}


int main (void)
{
    int opt, size;
    uint64_t addr;
    uint32_t data;
    int nw;
    uint8_t dbuff[4096];
    
    lib_model_connect();
    while (1) {
        print_banner();
        scanf("%d", &opt);
        switch (opt) {
            case 1:
                scanf("%lx %d", &addr, &nw);
                for (int i = 0; i < nw; i++) {
                    read_reg (addr + (i*4), data);
                    printf("Data: 0x%x\n", data);
                }
                break;
            case 2:
                scanf("%lx %x", &addr, &data);
                write_reg(addr, data);
                printf("write_reg complete\n");
                break;
            case 3:
                scanf("%lx %d", &addr, &size);
                read_mem(addr, dbuff, size);
                for (int i = 0; i < size; i++) {
                    printf("0x%x ", dbuff[i]);
                }
                printf("\n");
                break;
            case 4:
                dump_hbm();
                break;
            default:
                break;
        }
        if (opt == 5) break;
    }
    lib_model_conn_close();

    return 0;
}
