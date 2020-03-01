// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package network is a auto generated package.
Input file: lb.proto
*/
package network

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
func (m *LbPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "lb-policy/", m.Tenant, "/", m.Name)
}

func (m *LbPolicy) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/lb-policy/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *HealthCheckSpec) Clone(into interface{}) (interface{}, error) {
	var out *HealthCheckSpec
	var ok bool
	if into == nil {
		out = &HealthCheckSpec{}
	} else {
		out, ok = into.(*HealthCheckSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*HealthCheckSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *HealthCheckSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LbPolicy) Clone(into interface{}) (interface{}, error) {
	var out *LbPolicy
	var ok bool
	if into == nil {
		out = &LbPolicy{}
	} else {
		out, ok = into.(*LbPolicy)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*LbPolicy))
	return out, nil
}

// Default sets up the defaults for the object
func (m *LbPolicy) Defaults(ver string) bool {
	var ret bool
	m.Kind = "LbPolicy"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *LbPolicySpec) Clone(into interface{}) (interface{}, error) {
	var out *LbPolicySpec
	var ok bool
	if into == nil {
		out = &LbPolicySpec{}
	} else {
		out, ok = into.(*LbPolicySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*LbPolicySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *LbPolicySpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LbPolicyStatus) Clone(into interface{}) (interface{}, error) {
	var out *LbPolicyStatus
	var ok bool
	if into == nil {
		out = &LbPolicyStatus{}
	} else {
		out, ok = into.(*LbPolicyStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*LbPolicyStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *LbPolicyStatus) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *HealthCheckSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *HealthCheckSpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *HealthCheckSpec) Normalize() {

}

func (m *LbPolicy) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *LbPolicy) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for LbPolicy"))
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
	return ret
}

func (m *LbPolicy) Normalize() {

	m.ObjectMeta.Normalize()

}

func (m *LbPolicySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *LbPolicySpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *LbPolicySpec) Normalize() {

}

func (m *LbPolicyStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *LbPolicyStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *LbPolicyStatus) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&LbPolicy{},
	)

}
