#! /usr/bin/python3
import os
import time
import threading
import iota.harness.api as api
from iota.test.iris.testcases.alg.dns.dns_utils import *
from iota.test.iris.testcases.alg.alg_utils import *
import iota.test.iris.testcases.vmotion.vm_utils as vm_utils 
import pdb

def DNSCleanup(server, client):
    # Cleanup DNS Server
    SetupDNSServer(server, True)

    # Cleanup DNS Client
    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    api.Trigger_AddCommand(req, client.node_name, client.workload_name, "sudo rm /etc/resolv.conf")

    trig_resp = api.Trigger(req)
    term_resp = api.Trigger_TerminateAllCommands(trig_resp)

    for cmd in trig_resp.commands:
        api.PrintCommandResults(cmd)

    return api.types.status.SUCCESS

def Setup(tc):
    tc.move_info       = []
    tc.uuidMap         = api.GetNaplesNodeUuidMap()
    tc.num_moves       = int(getattr(tc.args, "num_moves", 1))
    tc.node_to_move    = getattr(tc.args,"","client")
    tc.local_or_remote = getattr(tc.args,"","local")
    update_sgpolicy('dns', True)
    return api.types.status.SUCCESS

def Trigger(tc):
    if tc.local_or_remote == "local":
        pairs = api.GetLocalWorkloadPairs()
    else:
        pairs = api.GetRemoteWorkloadPairs()

    if not len(pairs):
        return api.types.status.FAILURE

    server = pairs[0][0]
    client = pairs[0][1]

    if tc.node_to_move == "client":
        tc.vm_node = client
    else:
        tc.vm_node = server

    vm_utils.get_vm_dbg_stats(tc)

    tc.cmd_cookies = []

    serverReq = None
    clientReq = None

    serverReq = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    clientReq = api.Trigger_CreateExecuteCommandsRequest(serial = True)

    tc.cmd_descr = "Server: %s(%s) <--> Client: %s(%s)" %\
                   (server.workload_name, server.ip_address, client.workload_name, client.ip_address)
    api.Logger.info("Starting DNS test from %s" % (tc.cmd_descr))

    SetupDNSServer(server)

    api.Trigger_AddCommand(serverReq, server.node_name, server.workload_name, "sudo systemctl start named")
    tc.cmd_cookies.append("Before vMotion - Start Named")
    api.Trigger_AddCommand(serverReq, server.node_name, server.workload_name, "sudo systemctl enable named")
    tc.cmd_cookies.append("Before vMotion - Enable Named")
    api.Trigger_AddCommand(clientReq, client.node_name, client.workload_name, "sudo rm /etc/resolv.conf")
    tc.cmd_cookies.append("Before vMotion - Remove resolv.conf")

    api.Trigger_AddCommand(clientReq, client.node_name, client.workload_name,
                           "sudo echo \'nameserver %s\' | sudo tee -a /etc/resolv.conf"%(server.ip_address))
    tc.cmd_cookies.append("Before vMotion - Setup resolv conf")

    api.Trigger_AddCommand(clientReq, client.node_name, client.workload_name, "nslookup test3.example.com")
    tc.cmd_cookies.append("Before vMotion - Query DNS server") 
    api.Trigger_AddNaplesCommand(clientReq, tc.vm_node.node_name,
                                 "/nic/bin/halctl show session --dstport 53 --dstip {}".format(server.ip_address))
    tc.cmd_cookies.append("Before vMotion - show session")

    # Trigger the commands
    tc.server_resp = api.Trigger(serverReq)
    tc.client_resp = api.Trigger(clientReq)
    tc.resp        = api.Trigger_AggregateCommandsResponse(tc.server_resp, tc.client_resp)

    # Trigger vMotion
    new_node = vm_utils.find_new_node_to_move_to(tc, tc.vm_node)
    vm_utils.update_move_info(tc,[tc.vm_node],False,new_node)
    vm_utils.do_vmotion(tc, True)

    vm_utils.update_node_info(tc, tc.client_resp)

    # After vMotion - Wait for Commands to end
    tc.client_resp = api.Trigger_WaitForAllCommands(tc.client_resp)
    api.Trigger_TerminateAllCommands(tc.server_resp)

    # After vMotion - Show sessions dump after vMotion
    tc.cmd_cookies = []
    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)

    api.Trigger_AddCommand(req, client.node_name, client.workload_name, "nslookup test3.example.com")
    tc.cmd_cookies.append("After vMotion - Query DNS server") 

    trig_resp = api.Trigger(req)
    term_resp = api.Trigger_TerminateAllCommands(trig_resp)
    tc.resp   = api.Trigger_AggregateCommandsResponse(trig_resp, term_resp)

    DNSCleanup(server, client)

    return api.types.status.SUCCESS

def Verify(tc):
    if tc.resp is None:
        return api.types.status.FAILURE

    result = api.types.status.SUCCESS
    cookie_idx = 0
    api.Logger.info("Results for %s" % (tc.cmd_descr))
    for cmd in tc.resp.commands:
        api.PrintCommandResults(cmd)
        if cmd.exit_code != 0 and not api.Trigger_IsBackgroundCommand(cmd):
            if tc.cmd_cookies[cookie_idx].find("Before") != -1:
                result = api.types.status.SUCCESS
            else:
                result = api.types.status.FAILURE
        if tc.cmd_cookies[cookie_idx].find("show session") != -1 and \
           cmd.stdout != '':
           result = api.types.status.FAILURE
        cookie_idx += 1       
    return result

def Teardown(tc):
    if tc.GetStatus() != api.types.status.SUCCESS:
        api.Logger.info("verify failed, returning without teardown")
        return tc.GetStatus()

    vm_utils.move_back_vms(tc)
    return vm_utils.verify_vm_dbg_stats(tc)

