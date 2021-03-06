#! /usr/bin/python3

import infra.common.defs        as defs
import infra.common.objects     as objects
import infra.config.base        as base

from iris.config.store          import Store
import iris.config.resmgr       as resmgr

from infra.common.logging   import logger
from infra.factory.store    import FactoryStore
import types_pb2            as types_pb2
from iris.config.objects.crypto_keys import CryptoKeyHelper

import iris.config.hal.api           as halapi

import model_sim.src.model_wrap as model_wrap

from infra.common.glopts import GlobalOptions

from scapy.all import *

class NvmeNscb(Packet):
    name = "NvmeNscb"
    fields_desc = [
        LongField("ns_size", 0),
        IntField("key_index", 0),
        IntField("sec_key_index", 0),
        BitField("ns_valid", 0, 1),
        BitField("ns_active", 0, 1),
        BitField("rsvd0", 0, 1),
        BitField("log_lba_size", 0, 5),

        IntField("backend_ns_id", 0),

        BitField("num_sessions", 0, 11),    #1-based
        BitField("rr_session_id_to_be_served", 0, 10), #0-based

        BitField("num_outstanding_req", 0, 11), #1-based

        ShortField("sess_prodcb_start", 0),

        BitField("pad", 0, 40),

        BitField("valid_session_bitmap", 0, 256),
    ]

class NvmeNsObject(base.ConfigObjectBase):
    def __init__(self, lif, ns_id, size, lba_size, key_type, key_size, key):
        super().__init__()
        self.Clone(Store.templates.Get('NVME_NS'))
        self.lif = lif  
        self.id = ns_id
        self.GID("NvmeNS%02d" % self.id)
        self.size = size
        self.lba_size = lba_size
        #for now assume that ns has max of 16 sessions
        self.max_sess = 16
        self.nscb_addr = None
        self.backend_nsid = None
        self.session_list = []
        self.crypto_key = CryptoKeyHelper.main()
        self.crypto_key.Update(key_type, key_size, key)
        return

    def Show(self):
        logger.info("- NVME NS : %s" % self.GID())
        logger.info("   - nsid: %d backend_nsid: %d hw_lif_id: %d" \
                     %(self.id, self.backend_nsid, self.lif.lif.hw_lif_id))
        logger.info("   - size : %dlbas  lba_size: %d max_sess: %d nscb_addr: 0x%x" \
                     %(self.size, self.lba_size, self.max_sess,
                       self.nscb_addr if self.nscb_addr is not None else 0))
        logger.info("   - key_type : %d key_size: %d keyindex: %d" \
                     %(self.crypto_key.key_type, self.crypto_key.key_size, self.crypto_key.keyindex))
        return
    
    def PrepareHALRequestSpec(self, req_spec):
        if (GlobalOptions.dryrun):  return

        #by this time hw_lif_id should be valid to generate backend_nsid
        #backend_nsid = (lif << 12) | id
        self.backend_nsid = (self.lif.lif.hw_lif_id << 12) | self.id

        req_spec.nsid = self.id
        req_spec.hw_lif_id = self.lif.lif.hw_lif_id
        req_spec.backend_nsid = self.backend_nsid
        req_spec.size = self.size
        req_spec.lba_size = self.lba_size
        req_spec.max_sess = self.max_sess
        req_spec.key_index = self.crypto_key.keyindex
        return

    def ProcessHALResponse(self, req_spec, resp_spec):
        self.nscb_addr = resp_spec.nscb_addr
        self.Show()
        return

    def SessionAttach(self, nvme_sess):
        logger.info("Attaching nvme_sess: %s to ns: %s" \
                     %(nvme_sess.GID(), self.GID()))
        self.session_list.append(nvme_sess)

    def SessionGet(self, ns_local_sessid):
        return self.session_list[ns_local_sessid]

    def NscbRead(self, debug=True):
        nscb_size = len(NvmeNscb())
        if (GlobalOptions.dryrun):
            return NvmeNscb(bytes(nscb_size))
               
        if debug is True:
            logger.info("Read Nscb @0x%x size: %d" % (self.nscb_addr, nscb_size))

        nscb = NvmeNscb(model_wrap.read_mem(self.nscb_addr, nscb_size))
        logger.ShowScapyObject(nscb)
        return nscb
        
    def IsFilterMatch(self, selectors):

        #TBD: add nvme specific filters here
        logger.info('selectors: %s' %(str(selectors)))

        match = super().IsFilterMatch(selectors.nvmens.filters)
        logger.info('Matching Nvme NS: %s match: %s' % (self.GID(), match))
        return match

    def SetupTestcaseConfig(self, obj):
        obj.nvmens = self;
        return

    def ShowTestcaseConfig(self, obj):
        logger.info("Config Objects for %s" % self.GID())
        return

class NsObjectHelper:
    def __init__(self):
        self.ns_list = []

    
    def Generate(self, lif, max_ns):
    
        for ns_id in range(1, max_ns+1):
            # size is ns_id * 128 number of LBAs
            size = ns_id * 128

            #lba_size is 512B for odd ns_id and 4096B for even ns_id
            #256 bit crypto key for odd ns_id and 128 bit for even ns_id
            if (ns_id%2 == 0):
                lba_size = 4096
                key_type = types_pb2.CRYPTO_KEY_TYPE_AES128
                key_size = 16
            else:
                lba_size = 512
                key_type = types_pb2.CRYPTO_KEY_TYPE_AES256
                key_size = 32

            key = os.urandom(key_size)
            ns = NvmeNsObject(lif, ns_id, size, lba_size, key_type, key_size, key)
            self.ns_list.append(ns)

    def Configure(self):
        if (GlobalOptions.dryrun):  return
        halapi.NvmeNsCreate(self.ns_list)
        return

    def SessionAttach(self, nsid, nvme_sess):
        self.ns_list[nsid-1].SessionAttach(nvme_sess)

    def GetNs(self, nsid):
        return (self.ns_list[nsid-1])

def GetMatchingObjects(selectors):
    ns_store = Store.objects.GetAllByClass(NvmeNsObject)
    ns_list = []
    for ns in ns_store:
        if ns.IsFilterMatch(selectors): 
            ns_list.append(ns)
    return ns_list
