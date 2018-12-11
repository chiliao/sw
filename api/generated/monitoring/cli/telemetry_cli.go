// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringCliUtilsBackend is a auto generated package.
Input file: telemetry.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/monitoring"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateFlowExportPolicyFlags specifies flags for FlowExportPolicy create operation
var CreateFlowExportPolicyFlags = []gen.CliFlag{
	{
		ID:     "format",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "interval",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeFlowExportPolicyOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.FlowExportPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.FlowExportPolicyStatus{}
	}
	return nil
}

// CreateFwlogPolicyFlags specifies flags for FwlogPolicy create operation
var CreateFwlogPolicyFlags = []gen.CliFlag{
	{
		ID:     "filter",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "format",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeFwlogPolicyOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.FwlogPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.FwlogPolicyStatus{}
	}
	return nil
}

// CreateStatsPolicyFlags specifies flags for StatsPolicy create operation
var CreateStatsPolicyFlags = []gen.CliFlag{
	{
		ID:     "downsample-retention-time",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "retention-time",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeStatsPolicyOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.StatsPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.StatsPolicyStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("monitoring.FlowExportPolicy", "create", CreateFlowExportPolicyFlags)
	cl.AddRemoveObjOperFunc("monitoring.FlowExportPolicy", removeFlowExportPolicyOper)

	cl.AddCliInfo("monitoring.FwlogPolicy", "create", CreateFwlogPolicyFlags)
	cl.AddRemoveObjOperFunc("monitoring.FwlogPolicy", removeFwlogPolicyOper)

	cl.AddCliInfo("monitoring.StatsPolicy", "create", CreateStatsPolicyFlags)
	cl.AddRemoveObjOperFunc("monitoring.StatsPolicy", removeStatsPolicyOper)

}
