#! /usr/bin/python3
import time
import iota.harness.api as api
from iota.test.iris.testcases.aging.aging_utils import *
import pdb
import ipaddress

IP_PROTO_ICMP = 1

INVALID_ICMP_TYPE_4 = 4
INVALID_ICMP_TYPE_6 = 6

INSTANCES_CREATE = 1
INSTANCES_REUSE  = 2

NUMBER_OF_CLIENT_SERVER_OBJECTS       = 4
NUMBER_OF_VALIDATIONS_PER_OBJECT      = 8
MIN_NUMBER_OF_DROP_PACKETS_PER_OBJECT = 1
MIN_NUMBER_OF_DROP_BYTES_PER_PACKET   = 60

DROP_MALFORMED_PKT_TTL_0 = 0x00000001

def add_command(tc, req, cmd_cookie, cmd, host, bg):
    tc.cmd_cookies.append(cmd_cookie)
    api.Trigger_AddCommand(req, host.node_name, host.workload_name, cmd,
                           background = bg)

def add_naples_command(tc, req, cmd_cookie, cmd, naples):
    tc.cmd_cookies.append(cmd_cookie)
    api.Trigger_AddNaplesCommand(req, naples.node_name, cmd)

def establishNaplesWorkloads(tc):
    tc.workloads = api.GetWorkloads()
    if len(tc.workloads) == 0:
        api.Logger.info("ERROR: No workloads")
        return api.types.status.FAILURE

    for wl in tc.workloads:
        if wl.IsNaples():
            tc.naples = wl
            break
    if tc.naples is None:
        api.Logger.info("ERROR: No Naples workload")
        return api.types.status.FAILURE

    for wl in tc.workloads:
        if tc.iterators.peer == 'local' and\
           wl.node_name == tc.naples.node_name and\
           wl != tc.naples and wl.uplink_vlan == tc.naples.uplink_vlan:
            tc.naples_peer = wl
            break
        elif tc.iterators.peer == 'remote' and\
             wl.node_name != tc.naples.node_name and\
             wl.uplink_vlan == tc.naples.uplink_vlan:
            tc.naples_peer = wl
            break

    #
    # Handle Single-Host Testbed case
    #
    if tc.naples_peer is None:
        for wl in tc.workloads:
            if wl != tc.naples and wl.uplink_vlan == tc.naples.uplink_vlan:
                tc.naples_peer = wl
                break
    if tc.naples_peer is None:
        api.Logger.info("ERROR: No Naples-peer workload")
        return api.types.status.FAILURE

    return api.types.status.SUCCESS

def Setup(tc):
    #
    # Establish Client/Server Workloads
    #
    tc.skip = False
    tc.naples = None
    tc.naples_peer = None

    result = establishNaplesWorkloads(tc)
    if result != api.types.status.SUCCESS:
        tc.skip = True
        return result

    if tc.iterators.naples == 'client':
        tc.client = tc.naples
        tc.server = tc.naples_peer
    else:
        tc.server = tc.naples
        tc.client = tc.naples_peer

    #
    # Preserve current Time-out configs and
    # Set-up ICMP-timeout per testcase specification
    #
    tc.icmp_timeout_val = get_timeout_val('icmp-timeout')
    update_timeout('icmp-timeout', tc.iterators.timeout)

    return api.types.status.SUCCESS

def Trigger(tc):
    if tc.skip == True:
        return api.types.status.FAILURE

    #
    # Set-up Test Environment
    #
    tc.cmd_cookies = []
    cmd_cookie = "%s(%s) --> %s(%s)" %\
                 (tc.server.workload_name, tc.server.ip_address,
                  tc.client.workload_name, tc.client.ip_address)
    api.Logger.info("Starting Multiple-IPv4-ICMP-Flow-Drops test from %s" %\
                   (cmd_cookie))

    #
    # Start TCPDUMP in background on Server/Client
    #
    req1 = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    #cmd_cookie = "start tcpdump on Server"
    #cmd = "sudo tcpdump -nnSXi {} > out.txt".format(tc.server.interface)
    #add_command(tc, req1, cmd_cookie, cmd, tc.server, True)

    #cmd_cookie = "start tcpdump on Client"
    #cmd = "sudo tcpdump -nnSXi {} > out.txt".format(tc.client.interface)
    #add_command(tc, req1, cmd_cookie, cmd, tc.client, True)

    #
    # Start with a clean slate by clearing all sessions/flows
    #
    cmd_cookie = "clear session"
    cmd = "/nic/bin/halctl clear session"
    add_naples_command(tc, req1, cmd_cookie, cmd, tc.naples)

    #
    # Make sure that Client<=>Server Forwarding is set up
    #
    #cmd_cookie = "trigger ping: Create case"
    #cmd = "ping -c1 %s -I %s" %\
    #      (tc.server.ip_address, tc.client.interface)
    #add_command(tc, req1, cmd_cookie, cmd, tc.client, False)

    idx = 0
    while (idx < tc.iterators.types):
        #
        # Send Good-Data from Client / Server Non-zero TTL
        #
        cmd_cookie = "send good data from Client: Create case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --count 1 {}"\
        .format(idx, tc.server.ip_address)
        add_command(tc, req1, cmd_cookie, cmd, tc.client, False)

        cmd_cookie = "send good data from Server: Create case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --count 1 {}"\
        .format(idx, tc.client.ip_address)
        add_command(tc, req1, cmd_cookie, cmd, tc.server, False)

        #
        # Send Bad-Data (TTL=0) from Client / Server
        #
        cmd_cookie = "send bad data from Client TTL=0: Create case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --ttl 0 --count 1 {}"\
        .format(idx, tc.server.ip_address)
        add_command(tc, req1, cmd_cookie, cmd, tc.client, False)

        cmd_cookie = "send bad data from Server TTL=0: Create case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --ttl 0 --count 1 {}"\
        .format(idx, tc.client.ip_address)
        add_command(tc, req1, cmd_cookie, cmd, tc.server, False)
        idx += 1

    #
    # Do "show session" command
    #
    cmd_cookie = "show session"
    cmd = "/nic/bin/halctl show session"
    add_naples_command(tc, req1, cmd_cookie, cmd, tc.naples)

    #
    # Trigger "metrics get IPv4FlowDropMetrics" output
    #
    cmd_cookie = "show flow-drop: Create case"
    cmd = "PATH=$PATH:/platform/bin/;\
           LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/platform/lib/:/nic/lib/;\
           export PATH; export LD_LIBRARY_PATH;\
           /nic/bin/delphictl metrics get IPv4FlowDropMetrics"
    add_naples_command(tc, req1, cmd_cookie, cmd, tc.naples)

    tc.resp1 = api.Trigger(req1)
    for command in tc.resp1.commands:
        api.PrintCommandResults(command)

    #
    # Clearing all sessions/flows
    #
    req2 = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    cmd_cookie = "clear session"
    cmd = "/nic/bin/halctl clear session"
    add_naples_command(tc, req2, cmd_cookie, cmd, tc.naples)

    #
    # Make sure that Client<=>Server Forwarding is set up
    #
    #cmd_cookie = "trigger ping: Re-use case"
    #cmd = "ping -c1 %s -I %s" %\
    #      (tc.server.ip_address, tc.client.interface)
    #add_command(tc, req2, cmd_cookie, cmd, tc.client, False)

    idx = 0
    while (idx < tc.iterators.types):
        #
        # Re-send Good-Data from Client / Server Non-zero TTL
        #
        cmd_cookie = "re-send good data from Client: Re-use case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --count 1 {}"\
        .format(idx, tc.server.ip_address)
        add_command(tc, req2, cmd_cookie, cmd, tc.client, False)

        cmd_cookie = "re-send good data from Server: Re-use case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --count 1 {}"\
        .format(idx, tc.client.ip_address)
        add_command(tc, req2, cmd_cookie, cmd, tc.server, False)

        #
        # Re-send Bad-Data (TTL=0) from Client / Server
        #
        cmd_cookie = "re-send bad data from Client TTL=0: Re-use case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --ttl 0 --count 1 {}"\
        .format(idx, tc.server.ip_address)
        add_command(tc, req2, cmd_cookie, cmd, tc.client, False)

        cmd_cookie = "re-send bad data from Server TTL=0: Re-use case"
        cmd = "hping3 --icmp --icmptype {} --force-icmp --ttl 0 --count 1 {}"\
        .format(idx, tc.client.ip_address)
        add_command(tc, req2, cmd_cookie, cmd, tc.server, False)
        idx += 1

    #
    # Do "show session" command
    #
    cmd_cookie = "show session"
    cmd = "/nic/bin/halctl show session"
    add_naples_command(tc, req2, cmd_cookie, cmd, tc.naples)

    #
    # Trigger "metrics get IPv4FlowDropMetrics" output
    #
    cmd_cookie = "show flow-drop: Re-use case"
    cmd = "PATH=$PATH:/platform/bin/;\
           LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/platform/lib/:/nic/lib/;\
           export PATH; export LD_LIBRARY_PATH;\
           /nic/bin/delphictl metrics get IPv4FlowDropMetrics"
    add_naples_command(tc, req2, cmd_cookie, cmd, tc.naples)

    tc.resp2 = api.Trigger(req2)
    for command in tc.resp2.commands:
        api.PrintCommandResults(command)

    #
    # Clearing all sessions/flows and Sleep for 45secs
    #
    req3 = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    cmd_cookie = "clear session"
    cmd = "/nic/bin/halctl clear session; sleep 45"
    add_naples_command(tc, req3, cmd_cookie, cmd, tc.naples)

    #
    # Do "show session" command
    #
    cmd_cookie = "show session"
    cmd = "/nic/bin/halctl show session"
    add_naples_command(tc, req3, cmd_cookie, cmd, tc.naples)

    #
    # Trigger "metrics get IPv4FlowDropMetrics" output
    #
    cmd_cookie = "show flow-drop: Delete case"
    cmd = "PATH=$PATH:/platform/bin/;\
           LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/platform/lib/:/nic/lib/;\
           export PATH; export LD_LIBRARY_PATH;\
           /nic/bin/delphictl metrics get IPv4FlowDropMetrics"
    add_naples_command(tc, req3, cmd_cookie, cmd, tc.naples)

    tc.resp3 = api.Trigger(req3)
    for command in tc.resp3.commands:
        api.PrintCommandResults(command)

    return api.types.status.SUCCESS

def Verify(tc):
    if tc.skip == True:
        return api.types.status.FAILURE

    if tc.resp1 is None or tc.resp2 is None or tc.resp3 is None:
        return api.types.status.FAILURE

    #
    # Start with SUCCESS result assumption
    #
    result = api.types.status.SUCCESS

    #
    # Parse and Verify Results for Create case
    #
    objects = 0
    cookie_idx = 0
    for cmd in tc.resp1.commands:
        if "show flow-drop: Create case" in tc.cmd_cookies[cookie_idx]:
            for line in cmd.stdout.split('\n'):
                newline = line.replace(',', '')
                if "IPv4FlowDropMetrics" in newline:
                    objects += 1
                    sip = 0
                    dip = 0
                elif "Sip" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            sip = int(s)
                elif "Dip" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            dip = int(s)
                            if sip == int(ipaddress.ip_address(tc.client.
                                          ip_address)):
                                if dip != int(ipaddress.ip_address(tc.server.
                                              ip_address)):
                                    print("************ERROR************")
                                    print(line)
                                    result = api.types.status.FAILURE
                            elif sip == int(ipaddress.ip_address(tc.server.
                                          ip_address)):
                                if dip != int(ipaddress.ip_address(tc.client.
                                              ip_address)):
                                    print("************ERROR************")
                                    print(line)
                                    result = api.types.status.FAILURE
                            else:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "Sport" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if int(s) >= tc.iterators.types:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
#               elif "Dport" in newline:
#                   for s in newline.split():
#                       if s.isdigit():
#                           if int(s) >= tc.iterators.types:
#                               print("************ERROR************")
#                               print(line)
#                               result = api.types.status.FAILURE
                elif "Ip_proto" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if int(s) != IP_PROTO_ICMP:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "instances" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            instances = int(s)
                            if instances != INSTANCES_CREATE:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "drop_packets" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            drop_packets = int(s)
                            if drop_packets <\
                               (MIN_NUMBER_OF_DROP_PACKETS_PER_OBJECT *\
                                instances):
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "drop_bytes" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if int(s) < (MIN_NUMBER_OF_DROP_BYTES_PER_PACKET *\
                                         drop_packets):
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "drop_reason" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if (int(s) & DROP_MALFORMED_PKT_TTL_0) == 0:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
        cookie_idx += 1

    if objects == 0:
        print("************ NULL-objects in create phase ************")
        result = api.types.status.FAILURE
    else:
        print("**** Number of objects create phase ****", objects)

                    
    #
    # Parse and Verify Results for Re-use case
    #
    objects = 0
    instances_create = 0
    instances_reuse = 0
    for cmd in tc.resp2.commands:
        if "show flow-drop: Re-use case" in tc.cmd_cookies[cookie_idx]:
            objects = 0
            for line in cmd.stdout.split('\n'):
                newline = line.replace(',', '')
                if "IPv4FlowDropMetrics" in newline:
                    objects += 1
                    sip = 0
                    dip = 0
                elif "Sip" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            sip = int(s)
                elif "Dip" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            dip = int(s)
                            if sip == int(ipaddress.ip_address(tc.client.
                                          ip_address)):
                                if dip != int(ipaddress.ip_address(tc.server.
                                              ip_address)):
                                    print("************ERROR************")
                                    print(line)
                                    result = api.types.status.FAILURE
                            elif sip == int(ipaddress.ip_address(tc.server.
                                          ip_address)):
                                if dip != int(ipaddress.ip_address(tc.client.
                                              ip_address)):
                                    print("************ERROR************")
                                    print(line)
                                    result = api.types.status.FAILURE
                            else:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "Sport" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if int(s) >= tc.iterators.types:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
#               elif "Dport" in newline:
#                   for s in newline.split():
#                       if s.isdigit():
#                           if int(s) >= tc.iterators.types:
#                               print("************ERROR************")
#                               print(line)
#                               result = api.types.status.FAILURE
                elif "Ip_proto" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if int(s) != IP_PROTO_ICMP:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "instances" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            instances = int(s)
                            if instances == INSTANCES_CREATE:
                                instances_create += 1
                            elif instances == INSTANCES_REUSE:
                                instances_reuse += 1
                            else:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "drop_packets" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            drop_packets = int(s)
                            if drop_packets <\
                               (MIN_NUMBER_OF_DROP_PACKETS_PER_OBJECT *\
                                instances):
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "drop_bytes" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if int(s) < (MIN_NUMBER_OF_DROP_BYTES_PER_PACKET *\
                                         drop_packets):
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
                elif "drop_reason" in newline:
                    for s in newline.split():
                        if s.isdigit():
                            if (int(s) & DROP_MALFORMED_PKT_TTL_0) == 0:
                                print("************ERROR************")
                                print(line)
                                result = api.types.status.FAILURE
        cookie_idx += 1

    if objects == 0:
        print("************ NULL-objects in re-use phase ************")
        result = api.types.status.FAILURE
    else:
        print("**** Number of objects re-use phase ****", objects)
        print("**** Number of instances_create     ****", instances_create)
        print("**** Number of instances_reuse      ****", instances_reuse)

    #
    # Parse and Verify Results for Delete case
    #
    delete_success = False
    for cmd in tc.resp3.commands:
        if "show flow-drop: Delete case" in tc.cmd_cookies[cookie_idx]:
            if cmd.stdout != '':
                print("************ERROR************")
                api.PrintCommandResults(cmd)
                result = api.types.status.FAILURE
            else:
                delete_success = True
        cookie_idx += 1

    if delete_success == False:
        return api.types.status.FAILURE

    return result

def Teardown(tc):
    if tc.skip == True:
        return api.types.status.FAILURE

    #
    # Restore current Time-out configs
    #
    update_timeout('icmp-timeout', tc.icmp_timeout_val)

    #
    # Terminate all commands
    #
    api.Trigger_TerminateAllCommands(tc.resp1)
    api.Trigger_TerminateAllCommands(tc.resp2)
    api.Trigger_TerminateAllCommands(tc.resp3)

    tc.SetTestCount(1)
    return api.types.status.SUCCESS

