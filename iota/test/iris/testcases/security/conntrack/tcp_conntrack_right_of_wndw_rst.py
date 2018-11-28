#! /usr/bin/python3

import iota.harness.api as api
from iota.test.iris.testcases.security.conntrack.session_info import *

def Setup(tc):
    api.Logger.info("Setup.")
    return api.types.status.SUCCESS

def Trigger(tc):
    api.Logger.info("Trigger.")
    pairs = api.GetLocalWorkloadPairs()
    tc.cmd_cookies = {}
    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    tc.resp = []
    #for w1,w2 in pairs:
    server,client  = pairs[0]
    cmd_cookie = "nc -l 1241"
    api.Trigger_AddCommand(req, server.node_name, server.workload_name, cmd_cookie, background=True)
    tc.cmd_cookies['server'] = cmd_cookie

    cmd_cookie = "nc {} 1241 -p 52259".format(server.ip_address)
    api.Trigger_AddCommand(req, client.node_name, client.workload_name,cmd_cookie, background=True)
    tc.cmd_cookies['client'] = cmd_cookie
       
    cmd_cookie = "/nic/bin/halctl show session --dstport 1241 --dstip {} --yaml".format(server.ip_address)
    api.Trigger_AddNaplesCommand(req, client.node_name, cmd_cookie)
    trig_resp1 = api.Trigger(req)
    cmd = trig_resp1.commands[-1] 
    api.PrintCommandResults(cmd)
    iseq_num, iack_num, iwindosz, iwinscale, rseq_num, rack_num, rwindo_sz, rwinscale = get_conntrackinfo(cmd)
    new_seq_num = iseq_num + rwindo_sz * (2 ** rwinscale)

    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)

    #hping3 -c 1 -s 52259 -p 1241 -M 0 -L 0 --ack --tcp-timestamp 192.168.100.102
    cmd_cookie = "hping3 -c 1 -s 52259 -p 1241 -M {}  -L {} --rst --ack --tcp-timestamp {} -d 10".format(new_seq_num, iack_num, server.ip_address)    
    api.Trigger_AddCommand(req, client.node_name, client.workload_name, cmd_cookie)
    tc.cmd_cookies['fail ping'] = cmd_cookie

    cmd_cookie = "sleep 3 && /nic/bin/halctl show session --dstport 1241 --dstip {} --yaml".format(server.ip_address)
    api.Trigger_AddNaplesCommand(req, client.node_name, cmd_cookie)
    tc.cmd_cookies['show after'] = cmd_cookie
    
    trig_resp = api.Trigger(req)
    term_resp = api.Trigger_TerminateAllCommands(trig_resp)
    term_resp1 = api.Trigger_TerminateAllCommands(trig_resp1)
    tc.resp = api.Trigger_AggregateCommandsResponse(trig_resp, term_resp)
    return api.types.status.SUCCESS    
        
def Verify(tc):
    api.Logger.info("Verify.")
    for cmd in tc.resp.commands:
        #api.PrintCommandResults(cmd)
        if tc.cmd_cookies['show after'] == cmd.command:
            if not cmd.stdout:
                return api.types.status.SUCCESS
            print(cmd.stdout)
            yaml_out = get_yaml(cmd)
            init_flow = get_initflow(yaml_out)
            conn_info = get_conntrack_info(init_flow)
            excep =  get_exceptions(conn_info)
            if (excep['tcpoutofwindow'] == 'false'):
                return api.types.status.FAILURE 
        
    #print(tc.resp)
    return api.types.status.SUCCESS

def Teardown(tc):
    api.Logger.info("Teardown.")
    return api.types.status.SUCCESS
