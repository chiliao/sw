#! /usr/bin/python3

import apollo.config.objects.lmapping as lmapping
import apollo.config.objects.policy as policy
import apollo.config.utils as utils

class PolicyObjectHelper:
    def __init__(self):
        return

    def __get_policyID_from_subnet(self, subnet, af, direction):
        if af == 'IPV6':
            return subnet.IngV6SecurityPolicyId if direction == 'ingress' else subnet.EgV6SecurityPolicyId
        else:
            return subnet.IngV4SecurityPolicyId if direction == 'ingress' else subnet.EgV4SecurityPolicyId

    def __is_lmapping_match(self, policyobj, lobj):
        if lobj.VNIC.SUBNET.VPC.VPCId != policyobj.VPCId:
            return False
        if lobj.AddrFamily == policyobj.AddrFamily:
            return (policyobj.PolicyId == self.__get_policyID_from_subnet(lobj.VNIC.SUBNET, policyobj.AddrFamily, policyobj.Direction))
        return False

    def GetMatchingConfigObjects(self, selectors):
        objs = []
        policyobjs = filter(lambda x: x.IsFilterMatch(selectors), policy.client.Objects())
        for policyObj in policyobjs:
            for lobj in lmapping.client.Objects():
                if self.__is_lmapping_match(policyObj, lobj):
                    policyObj.l_obj = lobj
                    objs.append(policyObj)
                    break
        return utils.GetFilteredObjects(objs, selectors.maxlimits, False)

PolicyHelper = PolicyObjectHelper()

def GetMatchingObjects(selectors):
    return PolicyHelper.GetMatchingConfigObjects(selectors)
