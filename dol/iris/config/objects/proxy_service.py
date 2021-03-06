#! /usr/bin

import pdb

import infra.common.defs        as defs
import infra.common.objects     as objects
import infra.config.base        as base

import iris.config.hal.defs          as haldefs
import iris.config.hal.api           as halapi

from iris.config.store               import Store
from infra.common.logging       import logger

class ProxyServiceObject(base.ConfigObjectBase):
    def __init__(self, name, type):
        super().__init__()
        self.Clone(Store.templates.Get("SERVICE_LIF"))
        self.type = type
        self.GID("%s_%s" % (name, type))
        return

    def Show(self):
        logger.info("Service List for %s" % self.GID())
        return

    def PrepareHALRequestSpec(self, reqspec):
        if self.type == "TCP_PROXY":
            reqspec.key_or_handle.proxy_id = 0
            reqspec.proxy_type = 1
         
        if self.type == "TLS_PROXY":
            reqspec.key_or_handle.proxy_id = 1
            reqspec.proxy_type = 2

        if self.type == "IPSEC_PROXY":
            reqspec.key_or_handle.proxy_id = 2
            reqspec.proxy_type = 3
 
        if self.type == "CPU_PROXY":
            reqspec.key_or_handle.proxy_id = 3
            reqspec.proxy_type = 4

        if self.type == "APP_REDIR":
            reqspec.key_or_handle.proxy_id = 4
            reqspec.proxy_type = 7

        if self.type == "P4PT":
            reqspec.key_or_handle.proxy_id = 5
            reqspec.proxy_type = 8

        if self.type == "APP_REDIR_PROXY_TCP":
            reqspec.key_or_handle.proxy_id = 6
            reqspec.proxy_type = 9

        if self.type == "APP_REDIR_SPAN":
            reqspec.key_or_handle.proxy_id = 7
            reqspec.proxy_type = 10

        if self.type == "APP_REDIR_PROXY_TCP_SPAN":
            reqspec.key_or_handle.proxy_id = 8
            reqspec.proxy_type = 11

        return

    def ProcessHALResponse(self, req_spec, resp_spec):
        logger.info("[%s]Received response %s for type: %d" % (self.GID(), haldefs.common.ApiStatus.Name(resp_spec.api_status),  resp_spec.proxy_status.proxy_handle))
        return

class ProxyServiceObjectHelper:
    def __init__(self):
        self.proxy_service_list = []
        return

    def Configure(self):
        halapi.ConfigureProxyService(self.proxy_service_list)
        return

    def Generate(self, type):
        spec = Store.specs.Get(type)
        for srvct in spec.entries:
            srvct = ProxyServiceObject(srvct.entry.name, srvct.entry.type)
            self.proxy_service_list.append(srvct)
            Store.objects.Add(srvct)
        return

    def Show(self):
        for srvc in self.proxy_service_list:
            srvc.Show()
        return

    def main(self):
        self.Generate("PROXY_SERVICE")
        self.Configure()
        self.Show()

ProxyServiceHelper = ProxyServiceObjectHelper()

