// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package security is a auto generated package.
Input file: securitygroup.proto
*/
package security

import (
	"errors"
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// MakeKey generates a KV store key for the object
func (m *SecurityGroup) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "security-groups/", m.Tenant, "/", m.Name)
}

func (m *SecurityGroup) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/security-groups/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *SecurityGroup) Clone(into interface{}) (interface{}, error) {
	var out *SecurityGroup
	var ok bool
	if into == nil {
		out = &SecurityGroup{}
	} else {
		out, ok = into.(*SecurityGroup)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SecurityGroup))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SecurityGroup) Defaults(ver string) bool {
	var ret bool
	m.Kind = "SecurityGroup"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *SecurityGroupSpec) Clone(into interface{}) (interface{}, error) {
	var out *SecurityGroupSpec
	var ok bool
	if into == nil {
		out = &SecurityGroupSpec{}
	} else {
		out, ok = into.(*SecurityGroupSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SecurityGroupSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SecurityGroupSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *SecurityGroupStatus) Clone(into interface{}) (interface{}, error) {
	var out *SecurityGroupStatus
	var ok bool
	if into == nil {
		out = &SecurityGroupStatus{}
	} else {
		out, ok = into.(*SecurityGroupStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SecurityGroupStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SecurityGroupStatus) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *SecurityGroup) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	tenant = m.Tenant

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "meta.tenant"
		uref, ok := resp[tag]
		if !ok {
			uref = apiintf.ReferenceObj{
				RefType: apiintf.ReferenceType("NamedRef"),
				RefKind: "Tenant",
			}
		}

		if m.Tenant != "" {
			uref.Refs = append(uref.Refs, globals.ConfigRootPrefix+"/cluster/"+"tenants/"+m.Tenant)
		}

		if len(uref.Refs) > 0 {
			resp[tag] = uref
		}
	}
}

func (m *SecurityGroup) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for SecurityGroup"))
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "ObjectMeta"
		if errs := m.ObjectMeta.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}

	if !ignoreSpec {

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *SecurityGroup) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *SecurityGroupSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SecurityGroupSpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.WorkloadSelector != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "WorkloadSelector"
			if errs := m.WorkloadSelector.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *SecurityGroupSpec) Normalize() {

	if m.WorkloadSelector != nil {
		m.WorkloadSelector.Normalize()
	}

}

func (m *SecurityGroupStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SecurityGroupStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *SecurityGroupStatus) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&SecurityGroup{},
	)

}
