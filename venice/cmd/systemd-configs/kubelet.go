package configs

import (
	"fmt"
	"net"
	"path"

	"github.com/pkg/errors"

	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/certs"
	"github.com/pensando/sw/venice/utils/systemd"
)

const (
	// Config files
	kubeletSystemdCfgFile = "kubelet.conf"
	kubeletCgroupDriver   = "cgroupfs"
)

// generateSystemdKubeletConfig generates the systemd configuration file for
// starting kubelet with the right arguments.
func generateSystemdKubeletConfig(nodeID string) error {
	const (
		// Environment variables
		kubeletAddressVar    = "KUBELET_ADDRESS"
		allowPrivilegedVar   = "ALLOW_PRIVILEGED"
		requireKubeconfigVar = "REQUIRE_KUBE_CONFIG"
		cgroupDriverVar      = "CGROUP_DRIVER"
		nodeIPVar            = "NODEIP"
		hostNameOverrideVar  = "HOSTNAME_OVERRIDE"
		readOnlyPortVar      = "READONLY_PORT"

		// Parameters
		kubeletAddressParam    = "--address"
		allowPrivilegedParam   = "--allow-privileged"
		requireKubeconfigParam = "--require-kubeconfig"
		cgroupDriverParam      = "--cgroup-driver"
		nodeIPParam            = "--node-ip"
		hostNameOverrideParam  = "--hostname-override"
		readOnlyPortParam      = "--read-only-port"
	)

	// Kubelet gets a single set of credentials that it uses to:
	// - authenticate itself as a client when opening a connection to ApiServer
	// - authenticate itself as a server when it receives incoming connections from ApiServer or other clients
	kubeletCertFile, kubeletKeyFile, kubeletCACertFile := certs.GetTLSCredentialsPaths(globals.KubeletPKIDir)

	cfgMap := make(map[string]string)
	// TODO: Bind to node IP.
	cfgMap[kubeletAddressVar] = fmt.Sprintf("%s 0.0.0.0", kubeletAddressParam)
	cfgMap[allowPrivilegedVar] = allowPrivilegedParam
	cfgMap[kubeconfigVar] = fmt.Sprintf("%s=%s", kubeconfigParam, path.Join(globals.KubeletConfigDir, kubeconfigFileName))
	cfgMap[requireKubeconfigVar] = requireKubeconfigParam
	cfgMap[cgroupDriverVar] = fmt.Sprintf("%s %s", cgroupDriverParam, kubeletCgroupDriver)
	cfgMap[hostNameOverrideVar] = fmt.Sprintf("%s %s", hostNameOverrideParam, nodeID)
	// The read-only port is unauthenticated, so it's better to disable it
	cfgMap[readOnlyPortVar] = fmt.Sprintf("%s %s", readOnlyPortParam, "0") // disabled
	ip := net.ParseIP(nodeID)
	if ip != nil {
		cfgMap[nodeIPVar] = fmt.Sprintf("%s %s", nodeIPParam, nodeID)
	}

	cfgMap[anonymousAuthVar] = fmt.Sprintf("%s=%s", anonymousAuthParam, "false")
	cfgMap[authzModeVar] = fmt.Sprintf("%s=%s", authzModeParam, "Webhook")
	cfgMap[clientCAFileVar] = fmt.Sprintf("%s %s", clientCAFileParam, kubeletCACertFile)
	cfgMap[tlsKeyFileVar] = fmt.Sprintf("%s %s", tlsKeyFileParam, kubeletKeyFile)
	cfgMap[tlsCertFileVar] = fmt.Sprintf("%s %s", tlsCertFileParam, kubeletCertFile)

	return systemd.WriteCfgMapToFile(cfgMap, path.Join(globals.KubeletConfigDir, kubeletSystemdCfgFile))
}

// GenerateKubeletConfig generates kubelet kubeconfig and systemd configuration files.
func GenerateKubeletConfig(nodeID, kubeAPIServerAddr string, kubeAPIServerPort string) error {
	// Kubeconfig contains the parameters needed to connect to Kubernetes API Server
	err := generateKubeconfig(kubeAPIServerAddr, kubeAPIServerPort, "kubelet", globals.KubeletConfigDir, globals.KubeletPKIDir)
	if err != nil {
		return errors.Wrapf(err, "Error generating kubeconfig for Kubelet")
	}
	err = generateSystemdKubeletConfig(nodeID)
	if err != nil {
		return errors.Wrapf(err, "Error generating systemd config for kubelet")
	}
	return nil
}

// RemoveKubeletConfig removes the config files associated with kubelet.
func RemoveKubeletConfig() {
	removeFiles([]string{kubeletSystemdCfgFile})
	removeKubeconfig(globals.KubeletConfigDir)
}
