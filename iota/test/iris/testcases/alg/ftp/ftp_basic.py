#! /usr/bin/python3
import time
from iota.test.iris.testcases.alg.ftp.ftp_utils import *
from iota.test.iris.testcases.alg.alg_utils import *
import pdb

TCP_TICKLE_GAP = 15
NUM_TICKLES = 5

dir_path = os.path.dirname(os.path.realpath(__file__))

def Setup(tc):
    update_app('ftp', tc.iterators.timeout)
    update_sgpolicy('ftp')
    return api.types.status.SUCCESS

def SetupFTPClient(node, workload, server, mode):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    ftpdata = dir_path + '/' + "ftp_client.txt"
    api.Logger.info("fullpath %s" % (ftpdata))
    resp = api.CopyToWorkload(node, workload, [ftpdata], 'ftpdir')
    if resp is None:
       return None

    ftpcommand = dir_path + '/' + "ftp.sh" 
    f = open(ftpcommand, "w")
    f.write("#!/bin/bash\n")
    f.write("ftp -inv %s <<EOF\n"%(server.ip_address))
    f.write("user admin linuxpassword\n")
    f.write("cd /home/admin/ftp\n")
    f.write("pwd\n")
    f.write("dir\n")
    if (mode == 'active'):
        f.write("pass\n")
    f.write("mget ftp_server.txt\n")
    f.write("mput ftp_client.txt\n")
    f.close()

    resp = api.CopyToWorkload(node, workload, [ftpcommand], 'ftpdir')
    if resp is None:
       return None

    return api.types.status.SUCCESS 

def Cleanup(server, client):
    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    api.Trigger_AddCommand(req, server.node_name, server.workload_name,
                           "rm -rf ftpdir")
    api.Trigger_AddCommand(req, client.node_name, client.workload_name,
                           "rm -rf ftpdir")
    api.Trigger_AddCommand(req, server.node_name, server.workload_name,
                           "rm -rf /home/admin/ftp")
    api.Trigger_AddCommand(req, client.node_name, client.workload_name,
                           "rm -rf /home/admin/ftp")
    ftpfile = os.path.dirname(os.path.realpath(__file__)) + '/' + "ftp.sh"
    os.remove(ftpfile)
    trig_resp = api.Trigger(req)
    term_resp = api.Trigger_TerminateAllCommands(trig_resp)
    return api.types.status.SUCCESS

def Trigger(tc):
    pairs = api.GetLocalWorkloadPairs()
    w1 = pairs[0][0]
    w2 = pairs[0][1]
    tc.cmd_cookies = []
    timeout = timetoseconds(tc.iterators.timeout)

    naples = w1
    if not w1.IsNaples():
       naples = w2
       if not w2.IsNaples():
          return api.types.status.FAILURE

    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    tc.cmd_descr = "Server: %s(%s) <--> Client: %s(%s)" %\
                   (w1.workload_name, w1.ip_address, w2.workload_name, w2.ip_address)
    api.Logger.info("Starting FTP test from %s" % (tc.cmd_descr))

    SetupFTPServer(w1.node_name, w1.workload_name)
    SetupFTPClient(w2.node_name, w2.workload_name, w1, tc.iterators.mode)

    api.Trigger_AddNaplesCommand(req, naples.node_name,
                                "/nic/bin/halctl clear session")
    tc.cmd_cookies.append("clear session")

    api.Trigger_AddCommand(req, w1.node_name, w1.workload_name, 
                           "tcpdump -i {} -nn > out.txt".format(w1.interface), background=True)
    tc.cmd_cookies.append("tcpdump");

    api.Trigger_AddCommand(req, w2.node_name, w2.workload_name,
                           "cat ftpdir/ftp_server.txt | grep \"I am FTP server\"")
    tc.cmd_cookies.append("Before get on client")

    api.Trigger_AddCommand(req, w1.node_name, w1.workload_name,
                           "cat /home/admin/ftp/ftp_client.txt | grep \"I am FTP client\"")
    tc.cmd_cookies.append("Before put on server")

    api.Trigger_AddCommand(req, w2.node_name, w2.workload_name,
                           "cd ftpdir && chmod +x ftp.sh && ./ftp.sh", timeout=300)
    tc.cmd_cookies.append("Run FTP")

    ## Add Naples command validation
    api.Trigger_AddNaplesCommand(req, naples.node_name,
                                "/nic/bin/halctl show session --alg ftp")
    tc.cmd_cookies.append("show session FTP established")
 
    api.Trigger_AddCommand(req, w2.node_name, w2.workload_name,
                           "cat ftpdir/ftp_server.txt | grep \"I am FTP server\"")
    tc.cmd_cookies.append("After get on client")
    
    api.Trigger_AddCommand(req, w1.node_name, w1.workload_name,
                           "cat /home/admin/ftp/ftp_client.txt | grep \"I am FTP client\"")
    tc.cmd_cookies.append("After put on server")

    # Get the timeout
    timeout += get_timeout('tcp-close') + (TCP_TICKLE_GAP * NUM_TICKLES)
    api.Trigger_AddNaplesCommand(req, naples.node_name,
                                "sleep %s"%(timeout), timeout=300)
    tc.cmd_cookies.append("sleep")

    api.Trigger_AddNaplesCommand(req, naples.node_name,
                                "/nic/bin/halctl show session --alg ftp --yaml")
    tc.cmd_cookies.append("After timeout show session")

    trig_resp = api.Trigger(req)
    term_resp = api.Trigger_TerminateAllCommands(trig_resp)
    tc.resp = api.Trigger_AggregateCommandsResponse(trig_resp, term_resp)

    api.CopyFromWorkload(w1.node_name, w1.workload_name, ['out.txt'], dir_path)

    Cleanup(w1, w2)
    return api.types.status.SUCCESS

def Verify(tc):
    if tc.resp is None:
        return api.types.status.FAILURE

    result = api.types.status.SUCCESS
    api.Logger.info("Results for %s" % (tc.cmd_descr))
    cookie_idx = 0
    for cmd in tc.resp.commands:
        api.PrintCommandResults(cmd)
        if cmd.exit_code != 0 and not api.Trigger_IsBackgroundCommand(cmd):
            if tc.cmd_cookies[cookie_idx].find("Before") != -1 or \
               tc.cmd_cookies[cookie_idx].find("After timeout") != -1:
                result = api.types.status.SUCCESS
            else:
                result = api.types.status.FAILURE
        if tc.cmd_cookies[cookie_idx].find("After timeout") != -1 and \
           cmd.stdout != '':
           result = api.types.status.FAILURE
       
        cookie_idx += 1

    return result

def Teardown(tc):
    return api.types.status.SUCCESS
