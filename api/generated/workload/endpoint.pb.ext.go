// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package workload is a auto generated package.
Input file: endpoint.proto
*/
package workload

import (
	"errors"
	fmt "fmt"
	"strings"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// EndpointMigrationStatus_MigrationState_normal is a map of normalized values for the enum
var EndpointMigrationStatus_MigrationState_normal = map[string]string{
	"aborted": "aborted",
	"done":    "done",
	"failed":  "failed",
	"none":    "none",
	"start":   "start",
}

var EndpointMigrationStatus_MigrationState_vname = map[int32]string{
	0: "none",
	1: "start",
	2: "done",
	3: "failed",
	4: "aborted",
}

var EndpointMigrationStatus_MigrationState_vvalue = map[string]int32{
	"none":    0,
	"start":   1,
	"done":    2,
	"failed":  3,
	"aborted": 4,
}

func (x EndpointMigrationStatus_MigrationState) String() string {
	return EndpointMigrationStatus_MigrationState_vname[int32(x)]
}

var _ validators.DummyVar
var validatorMapEndpoint = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *Endpoint) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "endpoints/", m.Tenant, "/", m.Name)
}

func (m *Endpoint) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/endpoints/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *Endpoint) Clone(into interface{}) (interface{}, error) {
	var out *Endpoint
	var ok bool
	if into == nil {
		out = &Endpoint{}
	} else {
		out, ok = into.(*Endpoint)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*Endpoint))
	return out, nil
}

// Default sets up the defaults for the object
func (m *Endpoint) Defaults(ver string) bool {
	var ret bool
	m.Kind = "Endpoint"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Status.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *EndpointMigrationStatus) Clone(into interface{}) (interface{}, error) {
	var out *EndpointMigrationStatus
	var ok bool
	if into == nil {
		out = &EndpointMigrationStatus{}
	} else {
		out, ok = into.(*EndpointMigrationStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EndpointMigrationStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EndpointMigrationStatus) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.MigrationTimeout = "3m"
		m.Status = "none"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *EndpointSpec) Clone(into interface{}) (interface{}, error) {
	var out *EndpointSpec
	var ok bool
	if into == nil {
		out = &EndpointSpec{}
	} else {
		out, ok = into.(*EndpointSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EndpointSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EndpointSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *EndpointStatus) Clone(into interface{}) (interface{}, error) {
	var out *EndpointStatus
	var ok bool
	if into == nil {
		out = &EndpointStatus{}
	} else {
		out, ok = into.(*EndpointStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EndpointStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EndpointStatus) Defaults(ver string) bool {
	var ret bool
	if m.Migration != nil {
		ret = m.Migration.Defaults(ver) || ret
	}
	return ret
}

// Validators and Requirements

func (m *Endpoint) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	tenant = m.Tenant

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "status"

		m.Status.References(tenant, tag, resp)

	}
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

func (m *Endpoint) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for Endpoint"))
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

	if !ignoreStatus {

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Status"
		if errs := m.Status.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *Endpoint) Normalize() {

	m.ObjectMeta.Normalize()

	m.Status.Normalize()

}

func (m *EndpointMigrationStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EndpointMigrationStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapEndpoint["EndpointMigrationStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapEndpoint["EndpointMigrationStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *EndpointMigrationStatus) Normalize() {

	m.Status = EndpointMigrationStatus_MigrationState_normal[strings.ToLower(m.Status)]

}

func (m *EndpointSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EndpointSpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *EndpointSpec) Normalize() {

}

func (m *EndpointStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "network"
		uref, ok := resp[tag]
		if !ok {
			uref = apiintf.ReferenceObj{
				RefType: apiintf.ReferenceType("WeakRef"),
				RefKind: "Network",
			}
		}

		if m.Network != "" {
			uref.Refs = append(uref.Refs, globals.ConfigRootPrefix+"/network/"+"networks/"+tenant+"/"+m.Network)
		}

		if len(uref.Refs) > 0 {
			resp[tag] = uref
		}
	}
}

func (m *EndpointStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Migration != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Migration"
			if errs := m.Migration.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	if vs, ok := validatorMapEndpoint["EndpointStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapEndpoint["EndpointStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *EndpointStatus) Normalize() {

	if m.Migration != nil {
		m.Migration.Normalize()
	}

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&Endpoint{},
	)

	validatorMapEndpoint = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapEndpoint["EndpointMigrationStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapEndpoint["EndpointMigrationStatus"]["all"] = append(validatorMapEndpoint["EndpointMigrationStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*EndpointMigrationStatus)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "0")

		if err := validators.EmptyOr(validators.Duration, m.MigrationTimeout, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"MigrationTimeout", err.Error())
		}
		return nil
	})

	validatorMapEndpoint["EndpointMigrationStatus"]["all"] = append(validatorMapEndpoint["EndpointMigrationStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*EndpointMigrationStatus)

		if _, ok := EndpointMigrationStatus_MigrationState_vvalue[m.Status]; !ok {
			vals := []string{}
			for k1, _ := range EndpointMigrationStatus_MigrationState_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Status", vals)
		}
		return nil
	})

	validatorMapEndpoint["EndpointStatus"] = make(map[string][]func(string, interface{}) error)

	validatorMapEndpoint["EndpointStatus"]["all"] = append(validatorMapEndpoint["EndpointStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*EndpointStatus)
		if err := validators.EmptyOr(validators.MacAddr, m.MacAddress, nil); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"MacAddress", err.Error())
		}
		return nil
	})

}
