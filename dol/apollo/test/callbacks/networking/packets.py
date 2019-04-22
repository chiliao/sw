# /usr/bin/python3
import pdb
import ipaddress
import random
from infra.common.logging import logger
import infra.api.api as infra_api
import apollo.config.resmgr as resmgr
import apollo.config.utils as utils

IPV4_HOST = ipaddress.IPv4Address(0xbadee1ba)
IPV6_HOST = ipaddress.IPv6Address('e1ba:aced:a11:face:b00c:bade:da75:900d')

def __get_packet_template_impl(obj, args):
    template = 'ETH'
    template += "_%s" % (obj.AddrFamily)

    if args is not None:
        template += "_%s" % (args.proto)
    return infra_api.GetPacketTemplate(template)

def GetPacketTemplateFromMapping(testcase, packet, args=None):
    return __get_packet_template_impl(testcase.config.localmapping, args)

def GetPacketTemplateFromPolicy(testcase, packet, args=None):
    return __get_packet_template_impl(testcase.config.policy, args)

def __get_non_default_random_route(routes):
    numroutes = len(routes)
    if numroutes == 0:
        return None
    elif numroutes == 1:
        route = None
        if not utils.isDefaultRoute(routes[0]):
            route = routes[0]
        return route
    while True:
        route = random.choice(routes)
        if not utils.isDefaultRoute(route):
            break
    return route

def __get_host_from_route_impl(route, af):
    """
        returns a random usable host from non-default route
        if no route / only default route exists, returns IPV4_HOST / IPV6_HOST
    """
    if route is None or utils.isDefaultRoute(route):
        if af == "IPV4":
            return str(IPV4_HOST)
        elif af == "IPV6":
            return str(IPV6_HOST)
        else:
            logger.error("ERROR: invalid AddrFamily ", af)
            sys.exit(1)
            return None
    total_hosts = route.num_addresses
    host = route.network_address
    if total_hosts > 1:
        host += random.randint(0, total_hosts-1)
    return str(host)

def GetUsableHostFromRoute(testcase, packet, args=None):
    route = __get_non_default_random_route(testcase.config.route.routes)
    return __get_host_from_route_impl(route, testcase.config.route.AddrFamily)

def __get_non_default_random_route_from_rule(rules):
    numrules = len(rules)
    if numrules == 0:
        return None
    elif numrules == 1:
        route = None
        if not utils.isDefaultRoute(rules[0].Prefix):
            route = rules[0].Prefix
        return route
    while True:
        route = random.choice(rules).Prefix
        if not utils.isDefaultRoute(route):
            break
    return route

def GetUsableHostFromPolicy(testcase, packet, args=None):
    route = __get_non_default_random_route_from_rule(testcase.config.policy.rules)
    return __get_host_from_route_impl(route, testcase.config.policy.AddrFamily)

def GetInvalidMPLSTag(testcase, packet, args=None):
    return next(resmgr.InvalidMplsSlotIdAllocator)

def GetInvalidVnid(testcase, packet, args=None):
    return next(resmgr.InvalidVxlanIdAllocator)

def __get_packet_encap_type_impl(obj, args):
    if obj.IsWorkload():
        return 'ENCAP_MPLS2'
    elif obj.IsIgw():
        return 'ENCAP_MPLS'
    else:
        return None

def __get_packet_encap_impl(obj, tunnel, args):
    if obj.IsEncapTypeVXLAN():
        encap = 'ENCAP_VXLAN'
    elif obj.IsEncapTypeMPLS():
        encap = __get_packet_encap_type_impl(tunnel, args)
    else:
        assert 0
    return infra_api.GetPacketTemplate(encap)

# This can be called for packets to switch or from switch
def GetPacketEncapFromMapping(testcase, packet, args=None):
    encaps = []
    encaps.append(__get_packet_encap_impl(testcase.config.devicecfg, testcase.config.tunnel, args))
    return encaps

def __get_packet_srcmac_impl(fwdmode, dobj, robj, lobj, args):
    if dobj.IsEncapTypeMPLS():
        if fwdmode == 'L2':
            return robj.MACAddr
        else:
            return lobj.VNIC.SUBNET.VirtualRouterMACAddr
    elif dobj.IsEncapTypeVXLAN():
        return lobj.VNIC.SUBNET.VirtualRouterMACAddr
    else:
        assert 0

# This can be called for packets to host from switch
def GetPacketSrcMacAddrFromMapping(testcase, packet, args=None):
    return __get_packet_srcmac_impl(testcase.config.root.FwdMode,
            testcase.config.devicecfg, testcase.config.remotemapping,
            testcase.config.localmapping, args)

def __get_ip_localmapping_impl(localmapping, tunnel):
    if tunnel is not None:
        if tunnel.Nat is True:
            if hasattr(localmapping, "PublicIP"):
                return localmapping.PublicIP
        else:
            # If TEP does not have nat flag set, then no translation occurs.
            return localmapping.IP

def GetIPFromLocalMapping(testcase, packet, args=None):
    if testcase.config.route is not None and testcase.config.route.Tunnel is not None:
        tunnel = testcase.config.route.Tunnel
    else:
        tunnel = None
    return __get_ip_localmapping_impl(testcase.config.localmapping, tunnel)
