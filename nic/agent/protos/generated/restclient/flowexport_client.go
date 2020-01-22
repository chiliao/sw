// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: flowexport.proto
*/
package restclient

import (
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
)

// FlowExportPolicyList lists all FlowExportPolicy objects
func (cl *AgentClient) FlowExportPolicyList() ([]netproto.FlowExportPolicy, error) {
	var flowexportpolicyList []netproto.FlowExportPolicy

	err := netutils.HTTPGet("http://"+cl.agentURL+"/api/telemetry/flowexports/", &flowexportpolicyList)

	return flowexportpolicyList, err
}

// FlowExportPolicyPost creates FlowExportPolicy object
func (cl *AgentClient) FlowExportPolicyCreate(postData netproto.FlowExportPolicy) error {
	var resp Response

	err := netutils.HTTPPost("http://"+cl.agentURL+"/api/telemetry/flowexports/", &postData, &resp)

	return err

}

// FlowExportPolicyDelete deletes FlowExportPolicy object
func (cl *AgentClient) FlowExportPolicyDelete(deleteData netproto.FlowExportPolicy) error {
	var resp Response

	err := netutils.HTTPDelete("http://"+cl.agentURL+"/api/telemetry/flowexports/default/default/testDeleteFlowExportPolicy", &deleteData, &resp)

	return err
}

// FlowExportPolicyPut updates FlowExportPolicy object
func (cl *AgentClient) FlowExportPolicyUpdate(putData netproto.FlowExportPolicy) error {
	var resp Response

	err := netutils.HTTPPut("http://"+cl.agentURL+"/api/telemetry/flowexports/default/default/preCreatedFlowExportPolicy", &putData, &resp)

	return err
}