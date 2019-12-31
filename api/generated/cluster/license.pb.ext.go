// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package cluster is a auto generated package.
Input file: license.proto
*/
package cluster

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
func (m *License) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "licenses", "/Singleton")
}

func (m *License) MakeURI(cat, ver, prefix string) string {
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/licenses")
}

// Clone clones the object into into or creates one of into is nil
func (m *Feature) Clone(into interface{}) (interface{}, error) {
	var out *Feature
	var ok bool
	if into == nil {
		out = &Feature{}
	} else {
		out, ok = into.(*Feature)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*Feature))
	return out, nil
}

// Default sets up the defaults for the object
func (m *Feature) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *FeatureStatus) Clone(into interface{}) (interface{}, error) {
	var out *FeatureStatus
	var ok bool
	if into == nil {
		out = &FeatureStatus{}
	} else {
		out, ok = into.(*FeatureStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FeatureStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FeatureStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *License) Clone(into interface{}) (interface{}, error) {
	var out *License
	var ok bool
	if into == nil {
		out = &License{}
	} else {
		out, ok = into.(*License)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*License))
	return out, nil
}

// Default sets up the defaults for the object
func (m *License) Defaults(ver string) bool {
	var ret bool
	m.Kind = "License"
	ret = m.Tenant != "" || m.Namespace != ""
	if ret {
		m.Tenant, m.Namespace = "", ""
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *LicenseSpec) Clone(into interface{}) (interface{}, error) {
	var out *LicenseSpec
	var ok bool
	if into == nil {
		out = &LicenseSpec{}
	} else {
		out, ok = into.(*LicenseSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*LicenseSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *LicenseSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LicenseStatus) Clone(into interface{}) (interface{}, error) {
	var out *LicenseStatus
	var ok bool
	if into == nil {
		out = &LicenseStatus{}
	} else {
		out, ok = into.(*LicenseStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*LicenseStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *LicenseStatus) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *Feature) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *Feature) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *Feature) Normalize() {

}

func (m *FeatureStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FeatureStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *FeatureStatus) Normalize() {

}

func (m *License) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *License) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Tenant != "" {
		ret = append(ret, errors.New("Tenant not allowed for License"))
	}
	if m.Namespace != "" {
		ret = append(ret, errors.New("Namespace not allowed for License"))
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

func (m *License) Normalize() {

	m.ObjectMeta.Normalize()

}

func (m *LicenseSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *LicenseSpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *LicenseSpec) Normalize() {

}

func (m *LicenseStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *LicenseStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *LicenseStatus) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&License{},
	)

}