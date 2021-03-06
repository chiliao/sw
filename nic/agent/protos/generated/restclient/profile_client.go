// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: profile.proto
*/
package restclient

import (
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
)

// ProfileList lists all Profile objects
func (cl *AgentClient) ProfileList() ([]netproto.Profile, error) {
	var profileList []netproto.Profile

	err := netutils.HTTPGet("http://"+cl.agentURL+"/api/profiles/", &profileList)

	return profileList, err
}

// ProfilePost creates Profile object
func (cl *AgentClient) ProfileCreate(postData netproto.Profile) error {
	var resp Response

	err := netutils.HTTPPost("http://"+cl.agentURL+"/api/profiles/", &postData, &resp)

	return err

}

// ProfileDelete deletes Profile object
func (cl *AgentClient) ProfileDelete(deleteData netproto.Profile) error {
	var resp Response

	err := netutils.HTTPDelete("http://"+cl.agentURL+"/api/profiles/default/default/testDeleteProfile", &deleteData, &resp)

	return err
}

// ProfilePut updates Profile object
func (cl *AgentClient) ProfileUpdate(putData netproto.Profile) error {
	var resp Response

	err := netutils.HTTPPut("http://"+cl.agentURL+"/api/profiles/default/default/preCreatedProfile", &putData, &resp)

	return err
}
