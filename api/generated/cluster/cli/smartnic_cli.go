// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package clusterCliUtilsBackend is a auto generated package.
Input file: smartnic.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateDistributedServiceCardFlags specifies flags for DistributedServiceCard create operation
var CreateDistributedServiceCardFlags = []gen.CliFlag{
	{
		ID:     "admit",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "controllers",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "default-gw",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "dns-servers",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "dscprofile",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "id",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "ip-address",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "mgmt-mode",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "mgmt-vlan",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "network-mode",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "routing-config",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeDistributedServiceCardOper(obj interface{}) error {
	if v, ok := obj.(*cluster.DistributedServiceCard); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = cluster.DistributedServiceCardStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("cluster.DistributedServiceCard", "create", CreateDistributedServiceCardFlags)
	cl.AddRemoveObjOperFunc("cluster.DistributedServiceCard", removeDistributedServiceCardOper)

}
