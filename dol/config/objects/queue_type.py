#! /usr/bin/python3
import pdb
import math

import infra.common.defs        as defs
import infra.config.base        as base
import infra.common.objects     as objects
import config.resmgr            as resmgr
import config.objects.queue     as queue
import config.objects.eth.queue as eth_queue
import config.objects.rdma.queue as rdma_queue

import config.hal.api           as halapi
import config.hal.defs          as haldefs

from config.store               import Store
from infra.common.logging       import cfglogger

eth_queue_type_ids = {'RX', 'TX', 'ADMIN'}
rdma_queue_type_ids = {'RDMA_SQ', 'RDMA_RQ', 'RDMA_CQ', 'RDMA_EQ'}

class QueueTypeObject(base.ConfigObjectBase):
    def __init__(self):
        super().__init__()
        return

    def Init(self, lif, spec):
        self.GID(spec.id)
        self.id         = spec.id
        self.lif        = lif
        self.type       = spec.type
        self.purpose    = spec.purpose.upper()
        self.size       = spec.size
        self.count      = spec.count
        self.entries    = None

        self.queueid_allocator = objects.TemplateFieldObject("range/0/16384")

        self.queues = objects.ObjectDatabase(cfglogger)
        self.need_type_specific_configure = True
        if spec.id in eth_queue_type_ids:
            self.obj_helper_q = eth_queue.EthQueueObjectHelper()
        elif spec.id in rdma_queue_type_ids:
            self.obj_helper_q = rdma_queue.RdmaQueueObjectHelper()
        else:
            self.need_type_specific_configure = False
            return
        self.obj_helper_q.Generate(self, spec)
        if len(self.obj_helper_q.queues) > 0:
            self.queues.SetAll(self.obj_helper_q.queues)

        self.Show()


    def __copy__(self):
        q_type = QueueTypeObject()
        q_type.id = self.id
        q_type.lif = self.lif
        q_type.type = self.type
        q_type.purpose = self.purpose
        q_type.size = self.size
        q_type.count = self.count
        q_type.entries = self.entries
        
        return q_type

    
    def Equals(self, other, lgh):
        if not isinstance(other, self.__class__):
            return False
        fields = ["id", "type", "purpose", "entries"]
        if not self.CompareObjectFields(other, fields, lgh):
            return False
       
        return True
        
    def PrepareHALRequestSpec(self, req_spec):
        req_spec.type_num   = self.type
        req_spec.size       = int(math.log(self.size, 2)) - 5
        req_spec.entries    = self.entries = int(math.log(self.count, 2))
        req_spec.purpose = haldefs.interface.LifQPurpose.Value(self.purpose)
        if self.lif.cosA:
            req_spec.cos_a.qos_class_handle = self.lif.cosA.hal_handle
        if self.lif.cosB:
            req_spec.cos_b.qos_class_handle = self.lif.cosB.hal_handle

    def PrepareHALGetRequestSpec(self, get_req_spec):
        #Should never be called.
        assert 0

    def ProcessHALGetResponse(self, get_req_spec, get_resp):
        self.type = get_resp.type_num
        self.entries = get_resp.entries
        self.purpose = haldefs.interface.LifQPurpose.Name(get_resp.purpose)

    def GetQid(self):
        return self.queueid_allocator.get()

    def GetQstateAddr(self):
        return self.lif.GetQstateAddr(self.type)

    def ConfigureQueues(self):
        if self.need_type_specific_configure:
            self.obj_helper_q.Configure()

    def Show(self):
        cfglogger.info('Queue Type: %s' % self.GID())
        cfglogger.info('- lif       : %s' % self.lif.GID())
        cfglogger.info('- type      : %s' % self.type)
        cfglogger.info('- purpose   : %s' % self.purpose)
        cfglogger.info('- size      : %s' % self.size)
        cfglogger.info('- count     : %d' % len(self.obj_helper_q.queues))


class QueueTypeObjectHelper:
    def __init__(self):
        self.queue_types = []

    def Generate(self, lif, lifspec):
        for espec in lifspec.queue_types:
            queue_type = QueueTypeObject()
            queue_type.Init(lif, espec.queue_type)
            self.queue_types.append(queue_type)

    def Configure(self):
        for queue_type in self.queue_types:
            queue_type.ConfigureQueues()
