# /usr/bin/python3
import pdb

import infra.common.defs        as defs
import infra.common.objects     as objects
import infra.config.base        as base
import config.resmgr            as resmgr

from config.store               import Store
from infra.common.logging       import cfglogger
from config.objects.tls_proxy_cb     import TlsCbHelper

import config.hal.defs          as haldefs
import config.hal.api           as halapi

class TcpCbObject(base.ConfigObjectBase):
    def __init__(self):
        super().__init__()
        self.Clone(Store.templates.Get('TCPCB'))
        return
        
    #def Init(self, spec_obj):
    def Init(self, queue):
        self.id = resmgr.TcpCbIdAllocator.get()
        gid = "TcpCb%04d" % self.id
        self.GID(gid)
        # self.spec = spec_obj
        # cfglogger.info("  - %s" % self)

        # self.uplinks = objects.ObjectDatabase()
        # for uplink_spec in self.spec.uplinks:
            # uplink_obj = uplink_spec.Get(Store)
            # self.uplinks.Set(uplink_obj.GID(), uplink_obj)

        # assert(len(self.uplinks) > 0)
        cfglogger.info("  - %s" % self)
        self.queue = queue
        self.tlscb = TlsCbHelper.main(self)

        return


    def PrepareHALRequestSpec(self, req_spec):
        #req_spec.meta.tcpcb_id             = self.id
        req_spec.key_or_handle.tcpcb_id    = self.id
        req_spec.rcv_nxt                   = self.rcv_nxt
        return

    def ProcessHALResponse(self, req_spec, resp_spec):
        cfglogger.info("  - TcpCb %s = %s" %\
                       (self.id, \
                        haldefs.common.ApiStatus.Name(resp_spec.api_status)))
        return

    def IsFilterMatch(self, spec):
        return super().IsFilterMatch(spec.filters)

    def Read(self):
        return

    def Write(self):
        return



# Helper Class to Generate/Configure/Manage TcpCb Objects.
class TcpCbObjectHelper:
    def __init__(self):
        self.objlist = []
        return

    def Configure(self, objlist = None):
        if objlist == None:
            objlist = Store.objects.GetAllByClass(TcpCbObject)
        cfglogger.info("Configuring %d TcpCbs." % len(objlist)) 
        halapi.ConfigureTcpCbs(objlist)
        return
        
    def __gen_one(self, queue):
        cfglogger.info("Creating TcpCb")
        tcpcb_obj = TcpCbObject()
        tcpcb_obj.Init(queue)
        Store.objects.Add(tcpcb_obj)
        return tcpcb_obj

    def Generate(self, queue):
        self.objlist.append(self.__gen_one(queue))
        return self.objlist 

    def main(self, queue):
        objlist = self.Generate(queue)
        #self.Configure(self.tcpcbs)
        #Store.objects.SetAll(objlist)
        return objlist

TcpCbHelper = TcpCbObjectHelper()
