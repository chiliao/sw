import iota.harness.api as api
import iota.test.utils.naples_host as host
import iota.test.iris.testcases.server.verify_pci as verify_pci
import pdb


def Setup (tc):
    api.Logger.info("Validate flow control")

    tc.nodes = api.GetNaplesHostnames()
    tc.os = api.GetNodeOs(tc.nodes[0])

    if  tc.os != host.OS_TYPE_LINUX:
        api.Logger.info(f"This test runs on Linux, not on {tc.os}" )
        return api.types.status.FAILURE

    if len(tc.nodes) > 1:
        api.Logger.info(f"Expecting one node setup, this testbed has {len(tc.nodes)}")
        return api.types.status.FAILURE

    return api.types.status.SUCCESS


def Trigger (tc):
    naples_nodes =[]
    #for every node in the setup
    for node in tc.nodes:
        if api.IsNaplesNode(node):
            naples_nodes.append (node)
            api.Logger.info(f"Found Naples Node: [{node}]")

    if len(naples_nodes) == 0:
        api.Logger.error(f"Failed to find a Naples Node!")
        return api.types.status.FAILURE

    for reboot in range(tc.args.reboots): 
        api.Logger.info (f"Starting Reboot Loop # {reboot}")
        # RestartNodes (reboot, ipmi, apc)
        # api.RestartNodes(node, tc.iterators.reboot)
        api.RestartNodes(naples_nodes)

        for node in naples_nodes:   
            api.Logger.info (f"Verifying PCI on [{node}]")
            if verify_pci.verify_errors_lspci (node) != api.types.status.SUCCESS:
               api.Logger.error (f"PCIe Failure detected on {node}")

    return api.types.status.SUCCESS

def Verify (tc):

    return api.types.status.SUCCESS