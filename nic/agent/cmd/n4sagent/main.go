// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package main

import (
	"flag"

	"github.com/pensando/sw/nic/agent/netagent"
	"github.com/pensando/sw/nic/agent/netagent/datapath"
	"github.com/pensando/sw/nic/agent/nmd"
	"github.com/pensando/sw/nic/agent/nmd/platform"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/netutils"
)

// Main function
func main() {
	// command line flags
	var (
		hostIf       = flag.String("hostif", "ntrunk0", "Host facing interface")
		uplinkIf     = flag.String("uplink", "eth2", "Uplink interface")
		agentDbPath  = flag.String("agentdb", "/tmp/n4sagent.db", "Agent Database file")
		nmdDbPath    = flag.String("nmddb", "/tmp/nmd.db", "NMD Database file")
		npmURL       = flag.String("npm", "master.local:"+globals.NpmRPCPort, "NPM RPC server URL")
		cmdURL       = flag.String("cmd", "master.local:"+globals.CMDGRPCPort, "CMD RPC server URL")
		mode         = flag.String("mode", "classic", "Naples mode, \"classic\" or \"managed\" ")
		debugflag    = flag.Bool("debug", false, "Enable debug mode")
		logToFile    = flag.String("logtofile", "/var/log/pensando/n4sagent.log", "Redirect logs to file")
		resolverURLs = flag.String("resolver-urls", ":"+globals.CMDGRPCPort, "comma separated list of resolver URLs <IP:Port>")
	)
	flag.Parse()

	// Fill logger config params
	logConfig := &log.Config{
		Module:      "N4sAgent",
		Format:      log.JSONFmt,
		Filter:      log.AllowInfoFilter,
		Debug:       *debugflag,
		CtxSelector: log.ContextAll,
		LogToStdout: true,
		LogToFile:   true,
		FileCfg: log.FileConfig{
			Filename:   *logToFile,
			MaxSize:    10, // TODO: These needs to be part of Service Config Object
			MaxBackups: 3,  // TODO: These needs to be part of Service Config Object
			MaxAge:     7,  // TODO: These needs to be part of Service Config Object
		},
	}

	// Initialize logger config
	log.SetConfig(logConfig)

	// create a dummy channel to wait forver
	waitCh := make(chan bool)

	// read the mac address of the host interface
	macAddr, err := netutils.GetIntfMac(*hostIf)
	if err != nil {
		log.Fatalf("Error getting host interface's mac addr. Err: %v", err)
	}

	// create a network datapath
	dp, err := datapath.NewNaplesDatapath(*hostIf, *uplinkIf)
	if err != nil {
		log.Fatalf("Error creating fake datapath. Err: %v", err)
	}

	// create a platform agent
	pa, err := platform.NewNaplesPlatformAgent()
	if err != nil {
		log.Fatalf("Error creating platform agent. Err: %v", err)
	}

	// create the new NetAgent
	ag, err := netagent.NewAgent(dp, *agentDbPath, macAddr.String(), *npmURL, *resolverURLs, ":"+globals.AgentRESTPort)
	if err != nil {
		log.Fatalf("Error creating NetAgent. Err: %v", err)
	}
	log.Printf("NetAgent {%+v} is running", ag)

	// create the new NMD
	nm, err := nmd.NewAgent(pa, *nmdDbPath, macAddr.String(), *cmdURL, *resolverURLs, ":"+globals.NmdRESTPort, *mode)
	if err != nil {
		log.Fatalf("Error creating NMD. Err: %v", err)
	}

	log.Printf("NMD {%+v} is running", nm)

	// wait forever
	<-waitCh
}
