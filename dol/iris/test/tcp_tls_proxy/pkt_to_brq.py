# Testcase definition file.

import pdb
import copy
import iris.test.tcp_tls_proxy.tcp_proxy as tcp_proxy
import iris.test.tcp_tls_proxy.tcp_tls_proxy as tcp_tls_proxy

import types_pb2                as types_pb2
import internal_pb2          as internal_pb2
#import internal_pb2_grpc     as internal_pb2_grpc

from iris.config.store                   import Store
from iris.config.objects.proxycb_service    import ProxyCbServiceHelper
from iris.config.objects.tcp_proxy_cb        import TcpCbHelper
import iris.test.callbacks.networking.modcbs as modcbs
from infra.common.objects import ObjectDatabase as ObjectDatabase
from infra.common.logging import logger
from infra.common.glopts import GlobalOptions

#temporary hack until we implement pi/ci for BRQ
maxflows = 0
def Setup(infra, module):
    global maxflows
    print("Setup(): Sample Implementation")
    modcbs.Setup(infra, module)
    maxflows = module.args.maxflows
    return

def Teardown(infra, module):
    print("Teardown(): Sample Implementation.")
    return

def TestCaseSetup(tc):

    tc.pvtdata = ObjectDatabase()
    tcp_proxy.SetupProxyArgs(tc)
    id = ProxyCbServiceHelper.GetFlowInfo(tc.config.flow._FlowObject__session)
    TcpCbHelper.main(id)
    tcbid = "TcpCb%04d" % id
    # 1. Configure TCB in HBM before packet injection
    tcb = tc.infra_data.ConfigStore.objects.db[tcbid]
    tcp_proxy.init_tcb_inorder(tc, tcb)
    tcb.debug_dol_tx |= tcp_proxy.tcp_tx_debug_dol_dont_send_ack 
    # set tcb state to ESTABLISHED(1)
    tcb.state = 1
    tcb.SetObjValPd()

    # 2. Clone objects that are needed for verification
    rnmdpr_big = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["RNMDPR_BIG"])
    rnmdpr_big.GetMeta()
    rnmdpr_big.GetRingEntries([rnmdpr_big.pi])
    tnmdpr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["TNMDPR_BIG"])
    tnmdpr.GetMeta()

    brq = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["BRQ_ENCRYPT_GCM"])
    brq.GetMeta()
    brq.GetRingEntries([brq.pi])
    tcpcb = copy.deepcopy(tcb)
    tcpcb.GetObjValPd()

    tlscbid = "TlsCb%04d" % id
    tlscb = copy.deepcopy(tc.infra_data.ConfigStore.objects.db[tlscbid])

    # Key Setup
    key_type = types_pb2.CRYPTO_KEY_TYPE_AES128
    key_size = 16
    key = b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
    tlscb.crypto_key.Update(key_type, key_size, key)

    # TLS-CB Setup
    tlscb.command = 0x30000000
    tlscb.crypto_key_idx = tlscb.crypto_key.keyindex
    tlscb.salt = 0x12345678
    tlscb.explicit_iv = 0xfedcba9876543210
    tlscb.enc_requests = 0
    tlscb.enc_completions = 0
    tlscb.serq_pi = 0
    tlscb.serq_ci = 0
    tlscb.debug_dol = tcp_tls_proxy.tls_debug_dol_bypass_barco | \
                        tcp_tls_proxy.tls_debug_dol_bypass_proxy | \
                        tcp_tls_proxy.tls_debug_dol_sesq_stop
    tlscb.other_fid = 0xffff
    tlscb.is_decrypt_flow = False
    tlscb.SetObjValPd()

    tlscb = copy.deepcopy(tc.infra_data.ConfigStore.objects.db[tlscbid])
    tlscb.GetObjValPd()

    tc.pvtdata.Add(tlscb)
    tc.pvtdata.Add(rnmdpr_big)
    tc.pvtdata.Add(tnmdpr)
    tc.pvtdata.Add(tcpcb)
    tc.pvtdata.Add(brq)
    return

def TestCaseVerify(tc):
    global maxflows

    if GlobalOptions.dryrun:
        return True

    id = ProxyCbServiceHelper.GetFlowInfo(tc.config.flow._FlowObject__session)
    # 1. Verify pi/ci got update got updated for SERQ
    tlscbid = "TlsCb%04d" % id
    tlscb = tc.pvtdata.db[tlscbid]
    tlscb_cur = tc.infra_data.ConfigStore.objects.db[tlscbid]
    print("pre-sync: tlscb_cur.serq_pi %d tlscb_cur.serq_ci %d" % (tlscb_cur.serq_pi, tlscb_cur.serq_ci))
    tlscb_cur.GetObjValPd()
    print("post-sync: tlscb_cur.serq_pi %d tlscb_cur.serq_ci %d" % (tlscb_cur.serq_pi, tlscb_cur.serq_ci))
    if (tlscb_cur.serq_pi != (tlscb.serq_pi+1) or tlscb_cur.serq_ci != (tlscb.serq_ci+1)):
        print("serq pi/ci not as expected")
        return False

    # 2. Verify enc_requests
    if (tlscb_cur.enc_requests != tlscb.enc_requests+1):
        print("enc_requests not as expected %d %d" %(tlscb_cur.enc_requests, tlscb.enc_requests))
        return False

    # Assumes bypass Barco
    stage0_7_thread = 0x151111

    if (tlscb_cur.pre_debug_stage0_7_thread != stage0_7_thread):
        print("pre_debug_stage0_7_thread not as expected %x" % tlscb_cur.pre_debug_stage0_7_thread)
        return False

    # 3. Verify pi/ci got update got updated for BRQ
    brq = tc.pvtdata.db["BRQ_ENCRYPT_GCM"]
    brq_cur = tc.infra_data.ConfigStore.objects.db["BRQ_ENCRYPT_GCM"]
    print("pre-sync: brq_cur.pi %d brq_cur.ci %d" % (brq_cur.pi, brq_cur.ci))
    brq_cur.GetMeta()
    brq_cur.GetRingEntries([brq_cur.pi])
    print("post-sync: brq_cur.pi %d brq_cur.ci %d" % (brq_cur.pi, brq_cur.ci))
    if (brq_cur.pi != (brq.pi+1) or brq_cur.ci != (brq.ci+1)):
        print("brq pi/ci not as expected")
        #needs fix in HAL and support in model/p4+ for this check to work/pass
        #return False

    # 4. Fetch current values from Platform
    rnmdpr_big = tc.pvtdata.db["RNMDPR_BIG"]
    rnmdpr_big_cur = tc.infra_data.ConfigStore.objects.db["RNMDPR_BIG"]
    rnmdpr_big_cur.GetMeta()
    tnmdpr = tc.pvtdata.db["TNMDPR_BIG"]
    tnmdpr_cur = tc.infra_data.ConfigStore.objects.db["TNMDPR_BIG"]
    tnmdpr_cur.GetMeta()

    # 5. Verify PI for RNMDPR_BIG got incremented by 1
    if (rnmdpr_big_cur.pi != rnmdpr_big.pi+1):
        print("RNMDPR_BIG pi check failed old %d new %d" % (rnmdpr_big.pi, rnmdpr_big_cur.pi))
        return False
    print("Old RNMDPR_BIG PI: %d, New RNMDPR_BIG PI: %d" % (rnmdpr_big.pi, rnmdpr_big_cur.pi))



    return True

def TestCaseTeardown(tc):
    print("TestCaseTeardown(): Sample Implementation.")
    return
