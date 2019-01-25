// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package clusterCliUtilsBackend is a auto generated package.
Input file: cluster.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateClusterFlags specifies flags for Cluster create operation
var CreateClusterFlags = []gen.CliFlag{
	{
		ID:     "auto-admit-nics",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "certs",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "key",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "ntp-servers",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "quorum-nodes",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "virtual-ip",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeClusterOper(obj interface{}) error {
	if v, ok := obj.(*cluster.Cluster); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = cluster.ClusterStatus{}
	}
	return nil
}

// CreateHostFlags specifies flags for Host create operation
var CreateHostFlags = []gen.CliFlag{
	{
		ID:     "mac-address",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeHostOper(obj interface{}) error {
	if v, ok := obj.(*cluster.Host); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = cluster.HostStatus{}
	}
	return nil
}

// CreateNodeFlags specifies flags for Node create operation
var CreateNodeFlags = []gen.CliFlag{}

func removeNodeOper(obj interface{}) error {
	if v, ok := obj.(*cluster.Node); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = cluster.NodeStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("cluster.Cluster", "create", CreateClusterFlags)
	cl.AddRemoveObjOperFunc("cluster.Cluster", removeClusterOper)

	cl.AddCliInfo("cluster.Host", "create", CreateHostFlags)
	cl.AddRemoveObjOperFunc("cluster.Host", removeHostOper)

	cl.AddRemoveObjOperFunc("cluster.Node", removeNodeOper)

}
