#! /usr/bin/python3
import time
import iota.harness.api as api
import iota.protos.pygen.topo_svc_pb2 as topo_svc_pb2
import iota.test.iris.testcases.penctl.penctldefs as penctldefs
import iota.test.iris.testcases.penctl.common as common

def Setup(tc):
    tc.Nodes = api.GetNaplesHostnames()
    return api.types.status.SUCCESS

def Trigger(tc):
    if len(tc.Nodes) > 0:
        n = tc.Nodes[0]

        for i in range(0, 3):
            common.SetNaplesModeOOB_Static(n, "1.1.1.1", "2.2.2.2/24")
            common.SetNaplesModeOOB_Dynamic(n)
            common.SetNaplesModeInband_Static(n, "1.1.1.1", "2.2.2.2/24")
            common.SetNaplesModeHost(n)
            common.SetNaplesModeInband_Dynamic(n)

    return api.types.status.SUCCESS

def Verify(tc):
    return api.types.status.SUCCESS

def Teardown(tc):
    return api.types.status.SUCCESS
