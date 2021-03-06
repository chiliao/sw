#! /usr/bin/python3
import pdb

from infra.common.logging import logger
from infra.common.objects import MacAddressBase

from apollo.config.store import client as EzAccessStoreClient
from apollo.config.store import EzAccessStore

from apollo.config.resmgr import client as ResmgrClient
from apollo.config.resmgr import Resmgr
from apollo.config.agent.api import ObjectTypes as ObjectTypes
import apollo.config.utils as utils
import apollo.config.topo as topo
import apollo.config.objects.base as base

import tunnel_pb2 as tunnel_pb2
import types_pb2 as types_pb2
import ipaddress
import copy
import time

class TunnelStatus(base.StatusObjectBase):
    def __init__(self):
        super().__init__(ObjectTypes.TUNNEL)

class TunnelObject(base.ConfigObjectBase):
    def __init__(self, node, parent, spec, local):
        super().__init__(ObjectTypes.TUNNEL, node)
        if hasattr(spec, 'origin'):
            self.SetOrigin(spec.origin)
        self.__spec = spec
        if (hasattr(spec, 'id')):
            self.Id = spec.id
        else:
            self.Id = next(ResmgrClient[node].TunnelIdAllocator)
        self.GID("Tunnel%d"%self.Id)
        self.UUID = utils.PdsUuid(self.Id, self.ObjType)
        self.DEVICE = parent
        self.__nhtype = topo.NhType.NONE
        if utils.IsDol() and self.DEVICE.OverlayRoutingEn:
            self.SetOrigin('discovered')
        ################# PUBLIC ATTRIBUTES OF TUNNEL OBJECT #####################
        # In IOTA, we'll use the Local and Remote TEP IPs from testbed json file if present.
        # If not, we'll use the ones from cfgyml.
        self.LocalIPAddr = utils.GetNodeLoopbackIp(node)
        if not self.LocalIPAddr:
           if (hasattr(spec, 'srcaddr')):
               self.LocalIPAddr = ipaddress.IPv4Address(spec.srcaddr)
           else:
               self.LocalIPAddr = self.DEVICE.IPAddr
        self.EncapValue = 0
        self.Nat = False
        self.NexthopId = 0
        self.NEXTHOP = None
        self.NexthopGroupId = 0
        self.NEXTHOPGROUP = None
        self.validateFound = False

        if (hasattr(spec, 'nat')):
            self.Nat = spec.nat
        if local == True:
            self.RemoteIPAddr = self.LocalIPAddr
            self.Type = tunnel_pb2.TUNNEL_TYPE_NONE
        else:
            self.Type = utils.GetTunnelType(spec.type)
            if self.Type == tunnel_pb2.TUNNEL_TYPE_WORKLOAD:
                self.RemoteIPAddr = next(ResmgrClient[node].TepIpAddressAllocator)
                self.RemoteVnicMplsSlotIdAllocator = ResmgrClient[node].CreateRemoteVnicMplsSlotAllocator()
                self.RemoteVnicVxlanIdAllocator = ResmgrClient[node].CreateRemoteVnicVxlanIdAllocator()
            elif self.Type == tunnel_pb2.TUNNEL_TYPE_IGW:
                self.RemoteIPAddr = next(ResmgrClient[node].TepIpAddressAllocator)
                if self.DEVICE.IsEncapTypeMPLS():
                    self.EncapValue = next(ResmgrClient[node].IGWMplsSlotIdAllocator)
                else:
                    self.EncapValue = next(ResmgrClient[node].IGWVxlanIdAllocator)
            elif self.Type == tunnel_pb2.TUNNEL_TYPE_SERVICE:
                self.RemoteIPAddr = next(ResmgrClient[node].TepIpv6AddressAllocator)
                if hasattr(spec, "remote") and spec.remote is True:
                    self.Remote = True
                    self.RemoteServicePublicIP = next(ResmgrClient[node].RemoteSvcTunIPv4Addr)
                    self.RemoteServiceEncap = next(ResmgrClient[node].IGWVxlanIdAllocator)
                else:
                    self.Remote = False
                self.EncapValue = next(ResmgrClient[node].IGWVxlanIdAllocator)
            else:
                if utils.IsV4Stack(self.DEVICE.Stack):
                    self.RemoteIPAddr = utils.GetNodeLoopbackRemoteTEP(node)
                    if not self.RemoteIPAddr:
                        if getattr(spec, 'dstaddr', None) != None:
                            self.RemoteIPAddr = ipaddress.IPv4Address(spec.dstaddr)
                        else:
                            self.RemoteIPAddr = next(ResmgrClient[node].TepIpAddressAllocator)
                else:
                    self.RemoteIPAddr = next(ResmgrClient[node].TepIpv6AddressAllocator)
                # nexthop / nh_group association happens later
                if spec.type == 'underlay':
                    self.__nhtype = topo.NhType.UNDERLAY
                    self.NEXTHOP = None
                    if hasattr(spec, 'nhid'):
                        self.NexthopId = spec.nhid
                    else:
                        self.NexthopId = None
                elif spec.type == 'underlay-ecmp':
                    self.__nhtype = topo.NhType.UNDERLAY_ECMP
                    self.NEXTHOPGROUP = None

        self.RemoteIP = str(self.RemoteIPAddr) # for testspec

        if utils.IsDol() and self.DEVICE.OverlayRoutingEn:
            if hasattr(spec, 'encap'):
                self.EncapValue = spec.encap
            if hasattr(spec, 'macaddress'):
                self.MACAddr = spec.macaddress
            else:
                self.MACAddr = MacAddressBase(string='0000.0000.0000')
        else:
            if getattr(spec, 'macaddress', None) != None:
                self.MACAddr = spec.macaddress
            else:
                self.MACAddr = ResmgrClient[node].TepMacAllocator.get()
        self.Mutable = utils.IsUpdateSupported()

        ################# PRIVATE ATTRIBUTES OF TUNNEL OBJECT #####################
        self.Status = TunnelStatus()
        self.DeriveOperInfo()
        self.Show()
        return

    def Dup(self):
        dupObj = copy.copy(self)
        dupObj.Id = next(ResmgrClient[self.Node].TunnelIdAllocator) + Resmgr.BaseDuplicateIdAllocator
        dupObj.GID("DupTunnel%d"%dupObj.Id)
        dupObj.UUID = utils.PdsUuid(dupObj.Id, dupObj.ObjType)
        dupObj.Interim = True
        self.Duplicate = dupObj
        return dupObj

    def __repr__(self):
        remote = ""
        if hasattr(self, "Remote") and self.Remote is True:
            remote = " Remote:%s"% (self.Remote)
        return "TEP: %s |LocalIPAddr:%s|RemoteIPAddr:%s|TunnelType:%s%s|" \
               "EncapValue:%d|Nat:%s|Mac:%s|NhType:%s|NexthopId:%d" % \
               (self.UUID, self.LocalIPAddr, self.RemoteIPAddr,
               utils.GetTunnelTypeString(self.Type), remote, self.EncapValue,
               self.Nat, self.MACAddr, self.__nhtype, (self.NexthopId if self.NexthopId != None else 0))

    def Show(self):
        logger.info("Tunnel Object: %s" % self)
        logger.info("- %s" % repr(self))
        self.Status.Show()
        return

    def UpdateAttributes(self, spec):
        if self.LocalIPAddr != self.RemoteIPAddr:
            if self.Type == tunnel_pb2.TUNNEL_TYPE_WORKLOAD:
                self.RemoteIPAddr = next(ResmgrClient[self.Node].TepIpAddressAllocator)
            elif self.Type == tunnel_pb2.TUNNEL_TYPE_IGW:
                self.RemoteIPAddr = next(ResmgrClient[self.Node].TepIpAddressAllocator)
            else:
                if utils.IsV4Stack(self.DEVICE.Stack):
                    self.RemoteIPAddr = next(ResmgrClient[self.Node].TepIpAddressAllocator)
                else:
                    self.RemoteIPAddr = next(ResmgrClient[self.Node].TepIpv6AddressAllocator)
            if self.IsUnderlay():
                self.NEXTHOP = ResmgrClient[self.Node].UnderlayNHAllocator.rrnext()
                self.NexthopId = self.NEXTHOP.NexthopId
            elif self.IsUnderlayEcmp():
                self.NEXTHOPGROUP = ResmgrClient[self.Node].UnderlayNhGroupAllocator.rrnext()
                self.NexthopGroupId = self.NEXTHOPGROUP.Id
        self.RemoteIP = str(self.RemoteIPAddr) # for testspec
        self.MACAddr = ResmgrClient[self.Node].TepMacAllocator.get()
        return

    def RollbackAttributes(self):
        attrlist = ["RemoteIPAddr", "NEXTHOP", "NexthopId", "NEXTHOPGROUP", "NexthopGroupId", "RemoteIP", "MACAddr"]
        self.RollbackMany(attrlist)
        return

    def PopulateKey(self, grpcmsg):
        grpcmsg.Id.append(self.GetKey())
        return

    def PopulateSpec(self, grpcmsg):
        spec = grpcmsg.Request.add()
        spec.Id = self.GetKey()
        spec.VPCId = utils.PdsUuid.GetUUIDfromId(0, ObjectTypes.VPC) # TODO: Create Underlay VPC
        utils.GetRpcEncap(self.Node, self.EncapValue, self.EncapValue, spec.Encap)
        spec.Type = self.Type
        utils.GetRpcIPAddr(self.LocalIPAddr, spec.LocalIP)
        utils.GetRpcIPAddr(self.RemoteIPAddr, spec.RemoteIP)
        spec.Nat = self.Nat
        # TODO: Fix mac addr in testspec
        if not utils.IsPipelineApollo():
            spec.MACAddress = self.MACAddr.getnum()
        if utils.IsServiceTunnelSupported():
            if self.Type is tunnel_pb2.TUNNEL_TYPE_SERVICE and self.Remote is True:
                spec.RemoteService = self.Remote
                utils.GetRpcIPAddr(self.RemoteServicePublicIP, spec.RemoteServicePublicIP)
                utils.GetRpcEncap(self.Node, self.RemoteServiceEncap, self.RemoteServiceEncap, spec.RemoteServiceEncap)
        # In IOTA, with BGP Underlay support, we do not need to configure NextHop object/info. Instead,
        # NH resolution will be taken care by pds-agent/MS overlay-underlay NH stitching logic.
        if utils.IsDol():
            if self.IsUnderlay():
                spec.NexthopId = utils.PdsUuid.GetUUIDfromId(self.NexthopId, ObjectTypes.NEXTHOP)
            elif self.IsUnderlayEcmp():
                spec.NexthopGroupId = utils.PdsUuid.GetUUIDfromId(self.NexthopGroupId, ObjectTypes.NEXTHOPGROUP)
        return

    def ValidateSpecDiscovered(self, spec):
        if spec.MACAddress != self.MACAddr.getnum():
            logger.error(f"GRPC Get Tunnel {self.RemoteIPAddr} has unexpected MAC address {spec.MACAddress} expected {self.MACAddr} ({self.MACAddr.getnum()})")
            return False
        return True

    def ValidateSpec(self, spec):
        if utils.IsDol() and self.DEVICE.IsOverlayRoutingEnabled:
            return self.ValidateSpecDiscovered(spec)

        # TODO - Enable tunnel validation
        return True

        if spec.Id != self.GetKey():
            return False
        if utils.ValidateTunnelEncap(self.Node, self.EncapValue, spec.Encap) == False:
            return False
        # TODO: LocalIP is optional & unused
        # if utils.ValidateRpcIPAddr(self.LocalIPAddr, spec.LocalIP) == False:
        #     return False
        if utils.ValidateRpcIPAddr(self.RemoteIPAddr, spec.RemoteIP) == False:
            return False
        if spec.Type != self.Type:
            return False
        if utils.IsPipelineApollo():
            if spec.Nat != self.Nat:
                return False
        else:
            if spec.MACAddress != self.MACAddr.getnum():
                return False
        if utils.IsServiceTunnelSupported():
            if self.Type is tunnel_pb2.TUNNEL_TYPE_SERVICE and self.Remote is True:
                if spec.RemoteService != self.Remote:
                    return False
                # TODO: artemis
                # if utils.ValidateRpcIPAddr(self.RemoteServicePublicIP, spec.RemoteServicePublicIP) == False:
                #     return False
                # if utils.ValidateTunnelEncap(self.Node, self.RemoteServiceEncap, spec.RemoteServiceEncap) == False:
                #     return False
        return True

    def ValidateYamlSpec(self, spec):
        # TODO - Enable tunnel validation
        return True
        if utils.GetYamlSpecAttr(spec) != self.GetKey():
            return False

    def GetDependees(self, node):
        """
        depender/dependent - tunnel
        dependee - nexthop, & nexthop_group
        """
        # TODO: nh & nhg will be linked later
        dependees = [ ]
        if self.NEXTHOP:
            dependees.append(self.NEXTHOP)
        if self.NEXTHOPGROUP:
            dependees.append(self.NEXTHOPGROUP)
        return dependees

    def RestoreNotify(self, cObj):
        logger.info("Notify %s for %s creation" % (self, cObj))
        if not self.IsHwHabitant():
            logger.info(" - Skipping notification as %s already deleted" % self)
            return
        logger.info(" - Linking %s to %s " % (cObj, self))
        if cObj.ObjType == ObjectTypes.NEXTHOP:
            self.NexthopId = cObj.NexthopId
        elif cObj.ObjType == ObjectTypes.NEXTHOPGROUP:
            self.NexthopGroupId = cObj.Id
        else:
            logger.error(" - ERROR: %s not handling %s restoration" %\
                         (self.ObjType.name, cObj.ObjType))
            assert(0)
        self.SetDirty(True)
        self.CommitUpdate()
        return

    def DeleteNotify(self, dObj):
        logger.info("Notify %s for %s deletion" % (self, dObj))
        if not self.IsHwHabitant():
            logger.info(" - Skipping notification as %s already deleted" % self)
            return
        logger.info(" - Unlinking %s from %s " % (dObj, self))
        if dObj.ObjType == ObjectTypes.NEXTHOP:
            self.NexthopId = dObj.Duplicate.NexthopId
        elif dObj.ObjType == ObjectTypes.NEXTHOPGROUP:
            self.NexthopGroupId = dObj.Duplicate.Id
        else:
            logger.error(" - ERROR: %s not handling %s deletion" %\
                         (self.ObjType.name, dObj.ObjType))
            assert(0)
        self.SetDirty(True)
        self.CommitUpdate()
        return

    def IsWorkload(self):
        if self.Type == tunnel_pb2.TUNNEL_TYPE_WORKLOAD:
            return True
        return False

    def IsIgw(self):
        if self.Type == tunnel_pb2.TUNNEL_TYPE_IGW:
            return True
        return False

    def IsSvc(self):
        if self.Type == tunnel_pb2.TUNNEL_TYPE_SERVICE:
            return True
        return False

    def IsNat(self):
        if self.Nat is True:
            return True
        return False

    def IsUnderlay(self):
        if self.__nhtype == topo.NhType.UNDERLAY:
            return True
        return False

    def IsUnderlayEcmp(self):
        if self.__nhtype == topo.NhType.UNDERLAY_ECMP:
            return True
        return False

class TunnelObjectClient(base.ConfigClientBase):
    def __init__(self):
        super().__init__(ObjectTypes.TUNNEL, Resmgr.MAX_TUNNEL)
        # Expected discovered tunnel objects in case of overlay routing
        self.DiscoveredObjs = {}
        self.Parent = None
        return

    def AddObject(self, node, obj):
        self.Objs[node].update({obj.Id: obj})
        return

    def GetTunnelObject(self, node, tunnelid):
        return self.GetObjectByKey(node, tunnelid)

    def AssociateObjects(self, node):
        logger.info(f"Filling nexthops on {node}")
        NhClient = EzAccessStore.GetConfigClient(ObjectTypes.NEXTHOP)
        for tun in self.Objects(node):
            if tun.IsUnderlay():
                if tun.NexthopId == None:
                    nhObj = ResmgrClient[node].UnderlayNHAllocator.rrnext()
                    tun.NEXTHOP = nhObj
                    tun.NexthopId = nhObj.NexthopId
                    logger.info("Linking %s - %s" % (tun, nhObj))
                    nhObj.AddDependent(tun)
                elif tun.NEXTHOP == None:
                    logger.info("Linking Tunnel%d with Nexthop%d" %
                                (tun.Id, tun.NexthopId))
                    tun.NEXTHOP = NhClient.GetNexthopObject(node, tun.NexthopId)
                    tun.NEXTHOP.AddDependent(tun)
        return

    def FillUnderlayNhGroups(self, node):
        logger.info("Filling nexthop groups")
        for tun in self.Objects(node):
            if tun.IsUnderlayEcmp():
                nhGroupObj = ResmgrClient[node].UnderlayNhGroupAllocator.rrnext()
                tun.NEXTHOPGROUP = nhGroupObj
                tun.NexthopGroupId = nhGroupObj.Id
                logger.info("Linking %s - %s" % (tun, nhGroupObj))
                nhGroupObj.AddDependent(tun)
        return

    def AddObjToDict(self, obj):
        self.Objs[obj.Node].update({obj.Id: obj})
        return

    def DeleteObjFromDict(self, obj):
        self.Objs[obj.Node].pop(obj.Id, None)
        return

    def GetKeyfromSpec(self, spec, yaml=False):
        if utils.IsDol() and self.Parent.IsOverlayRoutingEnabled():
            # In case of overlay return IP, L3 VNI tuple
            # to look up against expected DiscoveredObjs
            if not yaml:
                vnid = spec.Encap.value.Vnid
                return (spec.RemoteIP, vnid, False)
            # TODO enable YAML lookup
            # spec['remoteip']['af'] = 1
            # spec['remoteip']['v4addr']
            # spec['encap']['value'] = None
            return (0, 0, True)
        return super().GetKeyfromSpec(spec, yaml)

    def GetObjectByKey(self, node, key):
        if utils.IsDol() and self.Parent.IsOverlayRoutingEnabled() and isinstance(key, tuple):
           yaml = key[2]
           if not yaml:
               logger.info(f"GetObjectByKey discovered Tunnel {key}")
               remoteIPKey = key[0]
               vnid = key[1]
               if remoteIPKey.Af == types_pb2.IP_AF_INET:
                   remoteip = remoteIPKey.V4Addr
               else:
                   remoteip = remoteIPKey.V6Addr
               remoteipaddr = ipaddress.ip_address(remoteip)
               if remoteipaddr not in self.DiscoveredObjs[node]:
                   return None
               if vnid not in self.DiscoveredObjs[node][remoteipaddr]:
                   return None
               return self.DiscoveredObjs[node][remoteipaddr][vnid]
           else:
               # TODO Fix YAML validation. For now return the first object
               objDict = self.Objs[node]
               for obj in objDict.values():
                   return obj
        return super().GetObjectByKey(node, key)

    def GenerateObjects(self, node, parent, tunnelspec):
        def __isTunFeatureSupported(tunnel_type):
            if tunnel_type == 'service':
                return utils.IsServiceTunnelSupported()
            elif tunnel_type == 'underlay' or tunnel_type == 'underlay-ecmp':
                return utils.IsUnderlayTunnelSupported()
            elif tunnel_type == 'internet-gateway':
                return utils.IsIGWTunnelSupported()
            elif tunnel_type == 'workload':
                return utils.IsWorkloadTunnelSupported()
            return False

        # Generate Remote Tunnel object
        self.Parent = parent

        if parent.IsOverlayRoutingEnabled():
            if node in self.DiscoveredObjs:
                discoveredObjs = self.DiscoveredObjs[node]
            else:
                self.DiscoveredObjs.update({node: {}})
                discoveredObjs = self.DiscoveredObjs[node]
        for t in tunnelspec:
            if not __isTunFeatureSupported(t.type):
                continue
            for c in range(t.count):
                obj = TunnelObject(node, parent, t, False)
                self.Objs[node].update({obj.Id: obj})
                if parent.IsOverlayRoutingEnabled():
                    if obj.RemoteIPAddr in discoveredObjs:
                        discoveredObjs[obj.RemoteIPAddr].update({obj.EncapValue: obj})
                    else:
                        discoveredObjs.update({obj.RemoteIPAddr: {obj.EncapValue: obj}})
        EzAccessStoreClient[node].SetTunnels(self.Objects(node))
        ResmgrClient[node].CreateInternetTunnels()
        ResmgrClient[node].CreateVnicTunnels()
        ResmgrClient[node].CollectSvcTunnels()
        ResmgrClient[node].CreateUnderlayTunnels()
        return

    def GetNumHwObjects(self, node):
        if utils.IsDol() and self.Parent.IsOverlayRoutingEnabled():
            count = 0
            for tunnels in self.DiscoveredObjs[node].values():
                count = count + len(tunnels)
            logger.info(f"NumDiscoveredHwObjects returned {count} for {self.ObjType.name} in {node}")
            return count
        return super().GetNumHwObjects(node)

client = TunnelObjectClient()

def GetMatchingObjects(selectors):
    return client.Objects()
