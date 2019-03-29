// {C} Copyright 2019 Pensando Systems Inc. All rights reserved

#include <iostream>
#include "nic/sdk/platform/fru/fru.hpp"

static int usage(int argc, char* argv[]) {
    fprintf(stderr, "Usage: %s <mode>\n",argv[0]);
    fprintf(stderr, "Possible modes are:\n");
    fprintf(stderr, "   %-16s\n","read");

    return 1;
}

static int printfru() {
    std::string value;

    if (readKey(MANUFACTURERDATE_KEY, value))
        return -1;
    printf("Manufacturing Date: %s\n", value.c_str());

    if (readKey(MANUFACTURER_KEY, value))
        return -1;
    printf("Manufacturer: %s\n", value.c_str());

    if (readKey(PRODUCTNAME_KEY, value))
        return -1;
    printf("Product Name: %s\n", value.c_str());

    if (readKey(SERIALNUMBER_KEY, value))
        return -1;
    printf("Serial Number: %s\n", value.c_str());

    if (readKey(PARTNUM_KEY, value))
        return -1;
    printf("Part Number: %s\n", value.c_str());

    if (readKey(BOARDID_KEY, value))
        return -1;
    printf("Board Id: %s\n", value.c_str());

    if (readKey(ENGCHANGELEVEL_KEY, value))
        return -1;
    printf("Eng Change Level: %s\n", value.c_str());

    if (readKey(NUMMACADDR_KEY, value))
        return -1;
    printf("Num of Macaddr: %s\n", value.c_str());
    
    if (readKey(MACADDRESS_KEY, value))
        return -1;
    printf("Base Macaddr: %s\n", value.c_str());

    return 0;

}

int main(int argc, char* argv[])
{
    std::string value;

    if (argc < 2)
        return usage(argc, argv);

    if (!strcmp(argv[1], "read")) {
        if(printfru()) {
            fprintf(stderr, "%s", "Error Reading FRU!\n");
            return -1;
        }
    } else {
        return usage(argc, argv);
    }

    return 0;
}
