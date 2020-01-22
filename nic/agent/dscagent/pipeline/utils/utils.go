package utils

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/dscagent/types"
	"github.com/pensando/sw/venice/utils/log"
)

// CreateNewGRPCClient creates a new RPC Client to the pipeline
func CreateNewGRPCClient(portEnvVar string, defaultPort string) (rpcClient *grpc.ClientConn, err error) {
	return waitForHAL(portEnvVar, defaultPort)
}

func isHalConnected(portEnvVar string, defaultPort string) (*grpc.ClientConn, error) {
	halPort := os.Getenv(portEnvVar)
	if halPort == "" {
		halPort = defaultPort
	}
	halURL := fmt.Sprintf("%s:%s", types.HalGRPCDefaultBaseURL, halPort)
	log.Infof("HAL URL: %s", halURL)
	return grpc.Dial(halURL, grpc.WithMaxMsgSize(math.MaxInt32-1), grpc.WithInsecure(), grpc.WithBlock())
}

func waitForHAL(portEnvVar string, defaultPort string) (rpcClient *grpc.ClientConn, err error) {
	halUP := make(chan bool, 1)
	ticker := time.NewTicker(types.HalGRPCTickerDuration)
	timeout := time.After(types.HalGRPCWaitTimeout)
	rpcClient, err = isHalConnected(portEnvVar, defaultPort)
	if err == nil {
		log.Infof("1st TickStaus: %s", types.InfoConnectedToHAL)
		return
	}

	for {
		select {
		case <-ticker.C:
			rpcClient, err = isHalConnected(portEnvVar, defaultPort)
			if err != nil {
				halUP <- true
			}
		case <-halUP:
			log.Infof("Agent HAL Status: %v", types.InfoConnectedToHAL)
			return
		case <-timeout:
			log.Errorf("Agent could not connect to HAL. | Err: %v", err)
			return nil, errors.Wrapf(types.ErrPipelineNotAvailabe, "Agent could not connect to HAL. Err: %v | Err: %v", types.ErrPipelineTimeout, err)
		}
	}
}

// ValidateMeta validates object keys based on kind.
func ValidateMeta(oper types.Operation, kind string, meta api.ObjectMeta) error {
	if oper == types.List {
		if len(kind) == 0 {
			return errors.Wrapf(types.ErrBadRequest, "Empty Kind %v", types.ErrEmptyFields)
		}
		return nil
	}

	switch strings.ToLower(kind) {
	case "tenant":
		if len(meta.Name) == 0 {
			return errors.Wrapf(types.ErrBadRequest, "Kind: %v | Meta: %v | Err: %v", kind, meta, types.ErrEmptyFields)
		}
	case "namespace":
		if len(meta.Tenant) == 0 || len(meta.Name) == 0 {
			return errors.Wrapf(types.ErrBadRequest, "Kind: %v | Meta: %v | Err: %v", kind, meta, types.ErrEmptyFields)
		}
	default:
		if len(meta.Tenant) == 0 || len(meta.Namespace) == 0 || len(meta.Name) == 0 {
			return errors.Wrapf(types.ErrBadRequest, "Kind: %v | Meta: %v | Err: %v", kind, meta, types.ErrEmptyFields)
		}
	}
	return nil
}

// ValidateIPAddresses ensures that IP Address string is a valid v4 address. TODO v6 support
func ValidateIPAddresses(ipAddresses ...string) (err error) {
	for _, a := range ipAddresses {
		ip := net.ParseIP(strings.TrimSpace(a))
		if len(ip) == 0 {
			err = errors.Wrapf(types.ErrInvalidIP, "IP Address: %s | Err: %v", a, types.ErrBadRequest)
			return
		}
	}
	return
}

// ValidateIPAddressesPrefix ensures that IP Address string is a valid v4 address prefix in CIDR format. TODO v6 support
func ValidateIPAddressesPrefix(ipAddressPrefixes ...string) error {
	for _, p := range ipAddressPrefixes {
		_, _, err := net.ParseCIDR(strings.TrimSpace(p))
		if err != nil {
			return errors.Wrapf(types.ErrInvalidIPPrefix, "CIDR Block: %s | Err: %v", p, types.ErrBadRequest)
		}
	}
	return nil
}

// ValidateIPAddressRange ensures that IP Address range is a valid v4 address separated by a hyphen. TODO v6 support
func ValidateIPAddressRange(ipAddressRanges ...string) error {
	for _, r := range ipAddressRanges {
		components := strings.Split(r, "-")
		if len(components) != 2 {
			return errors.Wrapf(types.ErrInvalidIPRange, "Range: %s | Err: %v", r, types.ErrBadRequest)

		}
		err := ValidateIPAddresses(components[0])
		if err != nil {
			return err
		}

		err = ValidateIPAddresses(components[1])
		if err != nil {
			return err
		}
	}
	return nil
}

// ValidateMacAddresses ensures that MAC Address is in a valid OUI format
func ValidateMacAddresses(macAddresses ...string) error {
	for _, m := range macAddresses {
		_, err := net.ParseMAC(strings.TrimSpace(m))
		if err != nil {
			return errors.Wrapf(types.ErrInvalidMACAddress, "MAC Address: %s | Err: %v", m, types.ErrBadRequest)
		}
	}
	return nil
}

//func isHalConnected() (*grpc.ClientConn, error) {
//	halPort := os.Getenv("HAL_GRPC_PORT")
//	if halPort == "" {
//		halPort = types.HalGRPCDefaultPort
//	}
//	halURL := fmt.Sprintf("%s:%s", types.HalGRPCDefaultBaseURL, halPort)
//	log.Infof("HAL URL: %s", halURL)
//	return grpc.Dial(halURL, grpc.WithMaxMsgSize(math.MaxInt32-1), grpc.WithInsecure(), grpc.WithBlock())
//}

// ResolveIPAddress resolves IPAddresses and returns its ARP Cache.
//func ResolveIPAddress(mgmtIPAddress, mgmtIntf string, ipAddresses ...string) (map[string]string, error) {
//	arpCache := make(map[string]string)
//
//	mgmtLink, err := net.InterfaceByName(mgmtIntf)
//	if err != nil {
//		log.Error(errors.Wrapf(types.ErrARPManagementInterfaceNotFound, "Err: %v", err))
//		return arpCache, errors.Wrapf(types.ErrARPManagementInterfaceNotFound, "Err: %v", err)
//	}
//	arpClient, err := arp.Dial(mgmtLink)
//	if err != nil {
//		log.Error(errors.Wrapf(types.ErrARPClientDialFailure, "Err: %v", err))
//		return arpCache, errors.Wrapf(types.ErrARPClientDialFailure, "Err: %v", err)
//	}
//	defer arpClient.Close()
//
//	for _, a := range ipAddresses {
//		addr := net.ParseIP(a)
//		macAddress, err := resolveARPWithTimeout(mgmtIPAddress, addr, arpClient)
//		if err != nil || macAddress == nil {
//			log.Error(errors.Wrapf(types.ErrARPResolution, "%s | Err: %v", a, err))
//			return nil, errors.Wrapf(types.ErrARPResolution, "%s | Err: %v", a, err)
//		}
//		arpCache[a] = macAddress.String()
//	}
//
//	return arpCache, nil
//}

// mgmtIPAddress will be in CIDR format
//func resolveARPWithTimeout(mgmtIP string, addr net.IP, arpClient *arp.Client) (net.HardwareAddr, error) {
//
//	arpChan := make(chan net.HardwareAddr, 1)
//
//	go func() {
//		var macAddr net.HardwareAddr
//		var err error
//		// Do subnet checks here
//		_, mgmtNet, _ := net.ParseCIDR(mgmtIP)
//		if !mgmtNet.Contains(addr) {
//			macAddr, err = resolveARPForDefaultGateway(addr, arpClient)
//
//		} else {
//			log.Infof("Pipeline Utils Handler: %s", types.InfoARPingForSameSubnetIP)
//			macAddr, err = arpClient.Resolve(addr)
//			if err != nil {
//				log.Error(errors.Wrapf(types.ErrARPEntryMissingForSameSubnetIP, "Same Subnet IP: %s | Err: %v", addr.String(), err))
//			}
//		}
//		arpChan <- macAddr
//	}()
//
//	select {
//	case macAddr := <-arpChan:
//		return macAddr, nil
//	case <-time.After(types.ARPResolutionTimeout):
//		return nil, types.ErrARPResolutionTimeoutExceeded
//	}
//}

//func waitForHAL() (rpcClient *grpc.ClientConn, err error) {
//	halUP := make(chan bool, 1)
//	ticker := time.NewTicker(types.HalGRPCTickerDuration)
//	timeout := time.After(types.HalGRPCWaitTimeout)
//	rpcClient, err = isHalConnected()
//	if err == nil {
//		log.Infof("1st TickStaus: %s", types.InfoConnectedToHAL)
//		return
//	}
//
//	for {
//		select {
//		case <-ticker.C:
//			rpcClient, err = isHalConnected()
//			if err != nil {
//				halUP <- true
//			}
//		case <-halUP:
//			log.Infof("Agent HAL Status: %v", types.InfoConnectedToHAL)
//			return
//		case <-timeout:
//			log.Errorf("Agent could not connect to HAL. | Err: %v", err)
//			return nil, errors.Wrapf(types.ErrPipelineNotAvailabe, "Agent could not connect to HAL. Err: %v | Err: %v", types.ErrPipelineTimeout, err)
//		}
//	}
//}

// Ipv4Touint32 converts net.IP to 32 bit integer
func Ipv4Touint32(ip net.IP) uint32 {
	if ip == nil {
		return 0
	}
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}