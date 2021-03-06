// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityCliUtilsBackend is a auto generated package.
Input file: networkencryption.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/security"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateTrafficEncryptionPolicyFlags specifies flags for TrafficEncryptionPolicy create operation
var CreateTrafficEncryptionPolicyFlags = []gen.CliFlag{
	{
		ID:     "cipher-suite",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "encryption-transform",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "integrity-transform",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "key-rotation-interval-secs",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "mode",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "version",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeTrafficEncryptionPolicyOper(obj interface{}) error {
	if v, ok := obj.(*security.TrafficEncryptionPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = security.TrafficEncryptionPolicyStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("security.TrafficEncryptionPolicy", "create", CreateTrafficEncryptionPolicyFlags)
	cl.AddRemoveObjOperFunc("security.TrafficEncryptionPolicy", removeTrafficEncryptionPolicyOper)

}
