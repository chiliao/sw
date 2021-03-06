#! /usr/bin/python3
import json
import time
import iota.harness.api as api
import iota.protos.pygen.cfg_svc_pb2 as cfg_svc_pb2

json_object = {
    "kind": "Cluster",
    "api-version": "v1",
    "meta": {
        "name": "e2eCluster"
     },
    "spec": {
        "auto-admit-dscs": True,
        "quorum-nodes": [],
    }
}

def Main(step):
    req = cfg_svc_pb2.MakeClusterMsg()
    venice_ips = api.GetVeniceMgmtIpAddresses()
    req.endpoint = "%s:9001/api/v1/cluster" % venice_ips[0]
    venice_names = api.GetVeniceHostnames()
    for vn in venice_names:
        json_object["spec"]["quorum-nodes"].append(vn)
    req.config = json.dumps(json_object)

    for i in range(1, 4):
        time.sleep(3)
        resp = api.MakeCluster(req)
        if resp != None:
            return api.types.status.SUCCESS

    return api.types.status.FAILURE
