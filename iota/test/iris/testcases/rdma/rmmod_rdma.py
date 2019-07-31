#! /usr/bin/python3
import iota.harness.api as api
import iota.protos.pygen.topo_svc_pb2 as topo_svc_pb2
import iota.test.iris.verif.utils.rdma_utils as rdma
import iota.test.iris.utils.naples_host as host

def Setup(tc):
    tc.iota_path = api.GetTestsuiteAttr("driver_path")
    tc.nodes = api.GetNaplesHostnames()
    tc.os = api.GetNodeOs(tc.nodes[0])
    tc.insmod_opts = api.GetTestsuiteAttr("insmod_opts")
    return api.types.status.SUCCESS

def Trigger(tc):
    pairs = api.GetRemoteWorkloadPairs()

    w1 = pairs[0][0]
    w2 = pairs[0][1]

    #===============================================================
    # Uninstall RDMA module, so next tests would not have dependency
    #===============================================================
    req = api.Trigger_CreateExecuteCommandsRequest(serial = True)
    api.Logger.info("Uninstalling RDMA driver on the following nodes: {0}".format(tc.nodes))

    repeat = int(getattr(tc.args, 'reload', 0))
    if repeat:
        api.Logger.info("Repeating unload + reload {repeat} times"
                .format(repeat=repeat))
    for _ in range(repeat):
        for n in tc.nodes:
            if tc.os == host.OS_TYPE_LINUX:
                api.Trigger_AddHostCommand(req, n, "rmmod ionic_rdma")
                api.Trigger_AddHostCommand(req, n,
                        "insmod {path}drivers/rdma/drv/ionic/ionic_rdma.ko {opts}"
                        .format(path=tc.iota_path, opts=tc.insmod_opts))
            else:
                api.Trigger_AddHostCommand(req, n, "kldunload ionic_rdma")
                api.Trigger_AddHostCommand(req, n,
                        "kldload {path}sys/modules/ionic_rdma/ionic_rdma.ko"
                        .format(path=tc.iota_path))
            # allow device to register before proceeding
            api.Trigger_AddHostCommand(req, n, "sleep 2")

    api.Logger.info("Final unload")
    for n in tc.nodes:
        if tc.os == host.OS_TYPE_LINUX:
            api.Trigger_AddHostCommand(req, n, "rmmod ionic_rdma")
        else:
            api.Trigger_AddHostCommand(req, n, "kldunload ionic_rdma")
            #Reset kenv options that may have been configured (ignore errors here)
            api.Trigger_AddHostCommand(req, n, "kenv -u compat.linuxkpi.ionic_rdma_spec; true")

    tc.resp = api.Trigger(req)

    return api.types.status.SUCCESS

def Verify(tc):
    if tc.resp is None:
        return api.types.status.FAILURE

    api.Logger.info("rmmod_rdma results")

    for cmd in tc.resp.commands:
        api.PrintCommandResults(cmd)
        if cmd.exit_code != 0 and not api.Trigger_IsBackgroundCommand(cmd):
            return api.types.status.FAILURE

    return api.types.status.SUCCESS

def Teardown(tc):
    return api.types.status.SUCCESS
