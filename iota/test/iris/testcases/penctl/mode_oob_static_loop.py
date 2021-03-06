#! /usr/bin/python3
import time
import iota.harness.api as api
import iota.protos.pygen.topo_svc_pb2 as topo_svc_pb2
import iota.test.iris.testcases.penctl.penctldefs as penctldefs
import iota.test.iris.testcases.penctl.common as common


def Setup(tc):
    tc.Nodes = api.GetNaplesHostnames()
    tc.venice_ips = ["1.1.1.1", "2.2.2.2", "3.3.3.3"]
    tc.mgmt_ips = ["1.1.1.11/24", "2.2.2.2/24", "3.3.3.3/24"]
    tc.mgmt_ip_res = []
    tc.controller_ip_pentcl = []
    return api.types.status.SUCCESS

def Trigger(tc):
    for n in tc.Nodes:
        for i in range(0, len(tc.venice_ips)):
            common.SetNaplesModeOOB_Static(n, tc.venice_ips[i], tc.mgmt_ips[i])
            tc.mgmt_ip_res.append(common.PenctlGetNaplesMgtmIp(n))
            tc.controller_ip_pentcl.append(common.PenctlGetControllers(n)[0])

    return api.types.status.SUCCESS

def Verify(tc):
    for i in range(0, len(tc.venice_ips)):
        api.Logger.info("MGMT IP : {} PENCTL_CTRL : {}".format(tc.mgmt_ip_res[i], tc.controller_ip_pentcl[i]))
        if tc.mgmt_ip_res[i]  != tc.mgmt_ips[i] or tc.controller_ip_pentcl[i] != tc.venice_ips[i]:
            return api.types.status.FAILURE

    return api.types.status.SUCCESS

def Teardown(tc):
    for n in tc.Nodes:
        common.DeleteNMDDb(n)
    return api.types.status.SUCCESS
