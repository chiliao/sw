#! /usr/bin/python3
import iota.harness.api as api

def Setup(tc):
 
    tc.desc = '''
    Test  :   ib_*_bw
	test  :   send, write, read
    Opcode:   Only
    Num QP:   1, 2
    Pad   :   No
    Inline:   No
    modes :   bidirectional
    rdma_cm:  yes, no
    flip  :   yes, no
    '''

    tc.iota_path = api.GetTestsuiteAttr("driver_path")
    tc.vlan_idx = api.GetTestsuiteAttr("vlan_idx")

    pairs = api.GetRemoteWorkloadPairs()
    # get workloads from each node
    tc.w = []

    if getattr(tc.iterators, 'flip', None) == 'no':
        tc.w.append(pairs[tc.vlan_idx][0])
        tc.w.append(pairs[tc.vlan_idx][1])
    else:
        tc.w.append(pairs[tc.vlan_idx][1])
        tc.w.append(pairs[tc.vlan_idx][0])
    tc.devices = []
    tc.gid = []
    tc.ib_prefix = []
    for i in range(2):
        tc.devices.append(api.GetTestsuiteAttr(tc.w[i].ip_address+'_device'))
        tc.gid.append(api.GetTestsuiteAttr(tc.w[i].ip_address+'_gid'))
        if tc.w[i].IsNaples():
            tc.ib_prefix.append('cd ' + tc.iota_path + ' && ./run_rdma.sh  ')
        else:
            tc.ib_prefix.append('')

    return api.types.status.SUCCESS

def Trigger(tc):

    #==============================================================
    # trigger the commands
    #==============================================================
    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)

    if tc.iterators.rdma_cm == 'yes':
        cm_opt = " -R "
    else:
        cm_opt = " "

    i = 0
    j = i + 1
    w1 = tc.w[i]
    w2 = tc.w[j]

    tc.cmd_descr = "Server: %s(%s) <--> Client: %s(%s)" %\
                    (w1.workload_name, w1.ip_address, w2.workload_name, w2.ip_address)

    api.Logger.info("Starting ib_%s_bw test from %s" % (tc.iterators.test, tc.cmd_descr))

    # cmd for server
    cmd = "ib_" + tc.iterators.test + "_bw -d " + tc.devices[i] + " -n 10 -F -x " + tc.gid[i] + " -s 1024 -b -q " + str(tc.iterators.num_qp) + cm_opt + " --report_gbits"
    api.Trigger_AddCommand(req, 
                           w1.node_name, 
                           w1.workload_name,
                           tc.ib_prefix[i] + cmd,
                           background = True)

    # On Naples-Mellanox setups, with Mellanox as server, it takes a few seconds before the server
    # starts listening. So sleep for a few seconds before trying to start the client
    cmd = 'sleep 2'
    api.Trigger_AddCommand(req,
                           w1.node_name,
                           w1.workload_name,
                           cmd)

    # cmd for client
    cmd = "ib_" + tc.iterators.test + "_bw -d " + tc.devices[j] + " -n 10 -F -x " + tc.gid[j] + " -s 1024 -b -q " + str(tc.iterators.num_qp) + cm_opt + " --report_gbits " + w1.ip_address
    api.Trigger_AddCommand(req, 
                           w2.node_name, 
                           w2.workload_name,
                           tc.ib_prefix[j] + cmd)

    # trigger the request
    trig_resp = api.Trigger(req)
    term_resp = api.Trigger_TerminateAllCommands(trig_resp)

    tc.resp = api.Trigger_AggregateCommandsResponse(trig_resp, term_resp)
    return api.types.status.SUCCESS

def Verify(tc):
    if tc.resp is None:
        return api.types.status.FAILURE

    result = api.types.status.SUCCESS

    api.Logger.info("ib_send_bw results for %s" % (tc.cmd_descr))
    for cmd in tc.resp.commands:
        api.PrintCommandResults(cmd)
        if cmd.exit_code != 0 and not api.Trigger_IsBackgroundCommand(cmd):
            result = api.types.status.FAILURE
    return result

def Teardown(tc):
    return api.types.status.SUCCESS
