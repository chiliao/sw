// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package auditApiServer is a auto generated package.
Input file: svc_audit.proto
*/
package audit

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapSvc_audit = map[string]*api.Struct{}

var keyMapSvc_audit = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSvc_audit)
	schema.AddPaths(keyMapSvc_audit)
}
