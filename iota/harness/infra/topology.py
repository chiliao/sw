#! /usr/bin/python3
import os
import pdb

from iota.harness.infra.utils.logger import Logger as Logger
from iota.harness.infra.glopts import GlobalOptions as GlobalOptions

import iota.harness.infra.store as store
import iota.harness.infra.resmgr as resmgr
import iota.harness.infra.types as types
import iota.harness.infra.utils.parser as parser
import iota.harness.infra.utils.loader as loader
import iota.harness.api as api
import iota.harness.infra.testcase as testcase

import iota.protos.pygen.topo_svc_pb2 as topo_pb2
import iota.protos.pygen.types_pb2 as types_pb2

class Node(object):
    def __init__(self, spec):
        self.__spec = spec
        self.__name = spec.name
        self.__inst = store.GetTestbed().AllocateInstance()

        self.__ip_address = self.__inst.NodeMgmtIP
        self.__os = self.__inst.NodeOs

        self.__role = self.__get_instance_role(spec.role)

        self.__control_ip = resmgr.ControlIpAllocator.Alloc()
        self.__control_intf = "eth1"

        self.__data_intfs = [ "eth2", "eth3" ]
        self.__host_intfs = []
        self.__host_if_alloc_idx = 0

        if self.IsWorkloadNode():
            self.__workload_type = topo_pb2.WorkloadType.Value(spec.workloads.type)
            self.__workload_image = spec.workloads.image
        Logger.info("- New Node: %s: %s (%s)" %\
                    (spec.name, self.__ip_address, topo_pb2.PersonalityType.Name(self.__role)))
        return

    def __get_instance_role(self, role):
        if role != 'auto':
            return topo_pb2.PersonalityType.Value(role)

        nic_type = self.__inst.Resource.NICType
        if nic_type == 'pensando':
            return topo_pb2.PERSONALITY_NAPLES
        elif nic_type == 'mellanox':
            return topo_pb2.PERSONALITY_MELLANOX
        elif nic_type == 'broadcom':
            return topo_pb2.PERSONALITY_BROADCOM
        elif nic_type == 'intel':
            return topo_pb2.PERSONALITY_INTEL
        else:
            Logger.error("Unknown NIC Type : %s" % nic_type)
            sys.exit(1)

    def GetNicType(self):
        return self.__inst.Resource.NICType


    def Name(self):
        return self.__name
    def Role(self):
        return self.__role
    def GetOs(self):
        return self.__os
    def IsNaplesSim(self):
        return self.__role == topo_pb2.PERSONALITY_NAPLES_SIM
    def IsNaplesHw(self):
        return self.__role == topo_pb2.PERSONALITY_NAPLES
    def IsNaples(self):
        return self.IsNaplesSim() or self.IsNaplesHw()
    def IsMellanox(self):
        return self.__role == topo_pb2.PERSONALITY_MELLANOX
    def IsBroadcom(self):
        return self.__role == topo_pb2.PERSONALITY_BROADCOM
    def IsIntel(self):
        return self.__role == topo_pb2.PERSONALITY_INTEL
    def IsThirdParty(self):
        return self.IsMellanox() or self.IsBroadcom() or self.IsIntel()
    def IsWorkloadNode(self):
        return self.__role != topo_pb2.PERSONALITY_VENICE

    def UUID(self):
        if self.IsThirdParty():
            return self.Name()
        return self.__uuid

    def HostInterfaces(self):
        return self.__host_intfs

    def AllocateHostInterface(self):
        host_if = self.__host_intfs[self.__host_if_alloc_idx]
        self.__host_if_alloc_idx = (self.__host_if_alloc_idx + 1) % len(self.__host_intfs)
        return host_if

    def ControlIpAddress(self):
        return self.__control_ip

    def MgmtIpAddress(self):
        return self.__ip_address

    def WorkloadType(self):
        return self.__workload_type

    def WorkloadImage(self):
        return self.__workload_image

    def AddToNodeMsg(self, msg, topology, testsuite):
        msg.type = self.__role
        msg.image = ""
        msg.ip_address = self.__ip_address
        msg.name = self.__name

        if self.Role() == topo_pb2.PERSONALITY_VENICE:
            msg.venice_config.control_intf = self.__control_intf
            msg.venice_config.control_ip = str(self.__control_ip)
            msg.image = os.path.basename(testsuite.GetImages().venice)
            for n in topology.Nodes():
                if n.Role() != topo_pb2.PERSONALITY_VENICE: continue
                peer_msg = msg.venice_config.venice_peers.add()
                peer_msg.host_name = n.Name()
                peer_msg.ip_address = str(n.ControlIpAddress())
        else:
            msg.naples_config.control_intf = self.__control_intf
            msg.naples_config.control_ip = str(self.__control_ip)
            if not self.IsNaplesHw() and not self.IsThirdParty():
                msg.image = os.path.basename(testsuite.GetImages().naples)
            for data_intf in self.__data_intfs:
                msg.naples_config.data_intfs.append(data_intf)
            for n in topology.Nodes():
                if n.Role() != topo_pb2.PERSONALITY_VENICE: continue
                msg.naples_config.venice_ips.append(str(n.ControlIpAddress()))

            # TBD: Fix these hard-code values and use it from testbed json.
            msg.naples_config.naples_ip_address = "1.0.0.2"
            msg.naples_config.naples_username = "root"
            msg.naples_config.naples_password = "pen123"

            host_entity = msg.entities.add()
            host_entity.type = topo_pb2.ENTITY_TYPE_HOST
            host_entity.name = self.__name + "_host"
            if self.IsNaples():
                nic_entity = msg.entities.add()
                nic_entity.type = topo_pb2.ENTITY_TYPE_NAPLES
                nic_entity.name = self.__name + "_naples"

        script = self.GetStartUpScript()
        if script != None:
            msg.startup_script = script

        return types.status.SUCCESS

    def ProcessResponse(self, resp):
        self.__uuid = resp.node_uuid
        Logger.info("Node: %s UUID: %s" % (self.__name, self.__uuid))
        if self.IsMellanox():
            self.__host_intfs = resp.mellanox_config.host_intfs
        elif self.IsBroadcom():
            self.__host_intfs = resp.broadcom_config.host_intfs
        elif self.IsNaples():
            self.__host_intfs = resp.naples_config.host_intfs
        elif self.IsIntel():
            self.__host_intfs = resp.intel_config.host_intfs
        Logger.info("Node: %s Host Interfaces: %s" % (self.__name, self.__host_intfs))

        if len(self.__host_intfs) == 0:
            if GlobalOptions.dryrun:
                self.__host_intfs = ["dummy_intf0", "dummy_intf1"]
            else:
                Logger.error("Interfaces not found on Host: ", self.MgmtIpAddress())
                if self.IsNaples():
                    Logger.error("Check if IONIC driver is installed.")
                sys.exit(1)
        return

    def GetStartUpScript(self):
        if self.IsNaplesHw():
            return api.HOST_NAPLES_DIR + "/" + "nodeinit.sh"
        return None

class Topology(object):
    def __init__(self, spec):
        self.__nodes = {}

        assert(spec)
        Logger.info("Parsing Topology: %s" % spec)
        self.__dirname = os.path.dirname(spec)
        self.__spec = parser.YmlParse(spec)
        self.__parse_nodes()
        return

    def GetDirectory(self):
        return self.__dirname

    def __parse_nodes(self):
        for node_spec in self.__spec.nodes:
            node = Node(node_spec)
            self.__nodes[node.Name()] = node
        return

    def Nodes(self):
        return self.__nodes.values()

    def RestartNodes(self, node_names):
        req = topo_pb2.NodeMsg()

        for node_name in node_names:
            if node_name not in self.__nodes:
                Logger.error("Node %s not found" % node_name)
                return types.status.FAILURE
            msg = req.nodes.add()
            msg.name = node_name


        resp = api.ReloadNodes(req)
        if not api.IsApiResponseOk(resp):
            return types.status.FAILURE

        return types.status.SUCCESS


    def Setup(self, testsuite):
        Logger.info("Adding Nodes:")
        req = topo_pb2.NodeMsg()
        req.node_op = topo_pb2.ADD

        for name,node in self.__nodes.items():
            msg = req.nodes.add()
            ret = node.AddToNodeMsg(msg, self, testsuite)
            assert(ret == types.status.SUCCESS)

        resp = api.AddNodes(req)
        if not api.IsApiResponseOk(resp):
            return types.status.FAILURE

        for node_resp in resp.nodes:
            node = self.__nodes[node_resp.name]
            node.ProcessResponse(node_resp)

        return types.status.SUCCESS

    def GetVeniceMgmtIpAddresses(self):
        ips = []
        for n in self.__nodes.values():
            if n.Role() == topo_pb2.PERSONALITY_VENICE:
                ips.append(n.MgmtIpAddress())
        return ips

    def GetNaplesMgmtIpAddresses(self):
        ips = []
        for n in self.__nodes.values():
            if n.IsNaples():
                ips.append(n.MgmtIpAddress())
        return ips

    def GetNaplesUuidMap(self):
        uuid_map = {}
        for n in self.__nodes.values():
            if n.IsWorkloadNode():
                uuid_map[n.Name()] = n.UUID()
        return uuid_map

    def GetVeniceHostnames(self):
        ips = []
        for n in self.__nodes.values():
            if n.Role() == topo_pb2.PERSONALITY_VENICE:
                ips.append(n.Name())
        return ips

    def GetNaplesHostnames(self):
        ips = []
        for n in self.__nodes.values():
            if n.IsNaples():
                ips.append(n.Name())
        return ips

    def GetNaplesHostInterfaces(self, name):
        return self.__nodes[name].HostInterfaces()

    def GetWorkloadNodeHostnames(self):
        ips = []
        for n in self.__nodes.values():
            if n.IsWorkloadNode():
                ips.append(n.Name())
        return ips

    def GetWorkloadNodeHostInterfaces(self, node_name):
        return self.__nodes[node_name].HostInterfaces()

    def GetWorkloadTypeForNode(self, node_name):
        return self.__nodes[node_name].WorkloadType()

    def GetWorkloadImageForNode(self, node_name):
        return self.__nodes[node_name].WorkloadImage()

    def AllocateHostInterfaceForNode(self, node_name):
        return self.__nodes[node_name].AllocateHostInterface()

    def GetNodeOs(self, node_name):
        return self.__nodes[node_name].GetOs()

    def GetNicType(self, node_name):
        return self.__nodes[node_name].GetNicType()

    def GetNodes(self):
        return list(self.__nodes.values())
