#! /usr/bin/python3

import math
import time
import infra.config.base        as base
import infra.common.objects     as objects
import iris.config.objects.eth.queue as eth_queue
import iris.config.objects.eth.doorbell as doorbell
import iris.config.objects.rdma.queue as rdma_queue
import iris.config.objects.nvme.queue as nvme_queue
import iris.config.hal.defs          as haldefs
from infra.common.glopts        import GlobalOptions
from infra.common.defs          import status
from infra.common.logging       import logger


eth_queue_type_ids = {'RX', 'TX', 'ADMIN', 'EQ'}
rdma_queue_type_ids = {'RDMA_AQ', 'RDMA_SQ', 'RDMA_RQ', 'RDMA_CQ', 'RDMA_EQ'}
nvme_queue_type_ids = {'NVME_SQ', 'NVME_CQ', 'NVME_ARMQ', 'NVME_SESS'}


class QueueTypeObject(base.ConfigObjectBase):
    def __init__(self):
        super().__init__()
        return

    def Init(self, lif, spec):
        self.GID(spec.id)
        self.lif        = lif
        self.spec       = spec
        self.id         = spec.id
        self.type       = spec.type
        self.purpose    = spec.purpose.upper()
        self.size       = spec.size
        self.count      = spec.count
        self.entries    = None

        self.queueid_allocator = objects.TemplateFieldObject("range/0/16384")

        self.queues = objects.ObjectDatabase()
        self.need_type_specific_configure = True
        if spec.id in eth_queue_type_ids:
            self.obj_helper_q = eth_queue.EthQueueObjectHelper()
            if GlobalOptions.rtl and not GlobalOptions.skipverify:
                if self.purpose == "LIF_QUEUE_PURPOSE_RX":
                    # Ring the doorbell, don't increment the PI and don't set
                    # the scheduler bit.
                    # The PI will be incremented by directly writing to QState.
                    self.upd = 0x0
                elif self.purpose == "LIF_QUEUE_PURPOSE_TX":
                    # Ring the doorbell, don't increment the PI and but set
                    # the scheduler bit.
                    # The PI will be incremented by directly writing to QState.
                    self.upd = 0x1
                else:
                    self.upd = 0x0
            else:
                if self.purpose == "LIF_QUEUE_PURPOSE_RX":
                    # Ring the doorbell, set PI=index and the scheduler bit.
                    self.upd = 0x9
                elif self.purpose == "LIF_QUEUE_PURPOSE_TX":
                    # Ring the doorbell, set PI=index and the scheduler bit.
                    self.upd = 0x9
                else:
                    self.upd = 0x0
            self.doorbell = doorbell.Doorbell()
            self.doorbell.Init(self, self.spec)
        elif spec.id in rdma_queue_type_ids:
            self.obj_helper_q = rdma_queue.RdmaQueueObjectHelper()
        elif spec.id in nvme_queue_type_ids:
            logger.info("nvme queue type %s" %(self.id))
            self.obj_helper_q = nvme_queue.NvmeQueueObjectHelper()
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

    def Equals(self, other, lgh=logger):
        if not isinstance(other, self.__class__):
            return False
        fields = ["id", "type", "purpose", "entries"]
        if not self.CompareObjectFields(other, fields, lgh):
            return False

        return True

    def GetLif(self):
        return self.lif

    def PrepareHALRequestSpec(self, req_spec):
        req_spec.type_num   = self.type
        req_spec.size       = int(math.log(self.size, 2)) - 5
        req_spec.entries    = self.entries = int(math.log(self.count, 2))
        req_spec.purpose    = haldefs.interface.LifQPurpose.Value(self.purpose)

    def ProcessHALResponseSpec(self, req_spec, rsp_spec):
        pass

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

    def Post(self, descriptor, queue_id=0):
        if GlobalOptions.dryrun or GlobalOptions.cfgonly:
            return status.SUCCESS

        if GlobalOptions.eth_mode != "onepkt" and self.purpose == "LIF_QUEUE_PURPOSE_RX":
            return status.SUCCESS

        ring_id = 0
        queue = self.obj_helper_q.queues[queue_id]
        ring = queue.obj_helper_ring.rings[ring_id]
        ret = queue.Post(descriptor)
        if ret == status.SUCCESS:
            if not descriptor._more:    # All descriptors are posted
                if not GlobalOptions.skipverify:
                    queue.qstate.Read()
                if GlobalOptions.rtl and not GlobalOptions.skipverify:
                    queue.qstate.set_pindex(ring_id, ring.pi)
                self.doorbell.RingDB(queue_id, ring_id, ring.pi, queue.pid)
                if not GlobalOptions.skipverify:
                    queue.qstate.Read()
        return ret

    def Consume(self, descriptor, queue_id=0):
        if GlobalOptions.dryrun or GlobalOptions.cfgonly or GlobalOptions.skipverify:
            return status.SUCCESS

        queue = self.obj_helper_q.queues[queue_id]
        return queue.Consume(descriptor)

    def EnableEQ(self, eq, queue_id=0):
        queue = self.obj_helper_q.queues[queue_id]
        queue.EnableEQ(eq)

    def DisableEQ(self, queue_id=0):
        queue = self.obj_helper_q.queues[queue_id]
        queue.DisableEQ()

    def IsQstateArmed(self, queue_id=0):
        queue = self.obj_helper_q.queues[queue_id]
        return queue.IsQstateArmed()

    def __str__(self):
        return ("%s Lif:%s/Id:%s/Purpose:%s/QueueType:%s/Size:%s/Count:%s" %
                (self.__class__.__name__,
                 self.lif.hw_lif_id,
                 self.id,
                 self.purpose,
                 self.type,
                 self.size,
                 self.count))

    def Show(self):
        logger.info(self)


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
