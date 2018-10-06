#! /usr/bin/python3
import ipaddress

class IpAddressStep(object):
    def __init__(self, subnet, step, count = 1):
        super().__init__()
        self.start = ipaddress.IPv4Address(subnet)
        self.step = ipaddress.IPv4Address(step)
        self.curr = self.start
        self.count = count
        return

    def __getone(self):
        out = self.curr
        self.curr = ipaddress.IPv4Address(int(self.curr) + int(self.step))
        return out

    def Alloc(self):
        return self.__getone()

    def GetCount(self):
        return self.count

    def GetLast(self):
        #Assuming /24
        return ipaddress.IPv4Address(int(self.start) + 254)

ControlIpAllocator = IpAddressStep("64.0.0.0", "0.0.0.1")
