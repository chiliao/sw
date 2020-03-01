// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: telemetry.proto
*/
package monitoring

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

// FlowExportPolicySpec_Formats_normal is a map of normalized values for the enum
var FlowExportPolicySpec_Formats_normal = map[string]string{
	"ipfix": "ipfix",
}

var FlowExportPolicySpec_Formats_vname = map[int32]string{
	0: "ipfix",
}

var FlowExportPolicySpec_Formats_vvalue = map[string]int32{
	"ipfix": 0,
}

func (x FlowExportPolicySpec_Formats) String() string {
	return FlowExportPolicySpec_Formats_vname[int32(x)]
}

// FwlogFilter_normal is a map of normalized values for the enum
var FwlogFilter_normal = map[string]string{
	"all":             "all",
	"allow":           "allow",
	"deny":            "deny",
	"implicit-reject": "implicit-reject",
	"none":            "none",
	"reject":          "reject",
}

var FwlogFilter_vname = map[int32]string{
	0: "none",
	1: "allow",
	2: "deny",
	3: "reject",
	4: "implicit-reject",
	5: "all",
}

var FwlogFilter_vvalue = map[string]int32{
	"none":            0,
	"allow":           1,
	"deny":            2,
	"reject":          3,
	"implicit-reject": 4,
	"all":             5,
}

func (x FwlogFilter) String() string {
	return FwlogFilter_vname[int32(x)]
}

var _ validators.DummyVar
var validatorMapTelemetry = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *FlowExportPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "flowExportPolicy/", m.Tenant, "/", m.Name)
}

func (m *FlowExportPolicy) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/flowExportPolicy/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *FwlogPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "fwlogPolicy/", m.Tenant, "/", m.Name)
}

func (m *FwlogPolicy) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/fwlogPolicy/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *FlowExportPolicy) Clone(into interface{}) (interface{}, error) {
	var out *FlowExportPolicy
	var ok bool
	if into == nil {
		out = &FlowExportPolicy{}
	} else {
		out, ok = into.(*FlowExportPolicy)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FlowExportPolicy))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FlowExportPolicy) Defaults(ver string) bool {
	var ret bool
	m.Kind = "FlowExportPolicy"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FlowExportPolicySpec) Clone(into interface{}) (interface{}, error) {
	var out *FlowExportPolicySpec
	var ok bool
	if into == nil {
		out = &FlowExportPolicySpec{}
	} else {
		out, ok = into.(*FlowExportPolicySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FlowExportPolicySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FlowExportPolicySpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Format = "ipfix"
		m.Interval = "10s"
		m.TemplateInterval = "5m"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FlowExportPolicyStatus) Clone(into interface{}) (interface{}, error) {
	var out *FlowExportPolicyStatus
	var ok bool
	if into == nil {
		out = &FlowExportPolicyStatus{}
	} else {
		out, ok = into.(*FlowExportPolicyStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FlowExportPolicyStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FlowExportPolicyStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogPolicy) Clone(into interface{}) (interface{}, error) {
	var out *FwlogPolicy
	var ok bool
	if into == nil {
		out = &FwlogPolicy{}
	} else {
		out, ok = into.(*FwlogPolicy)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogPolicy))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogPolicy) Defaults(ver string) bool {
	var ret bool
	m.Kind = "FwlogPolicy"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogPolicySpec) Clone(into interface{}) (interface{}, error) {
	var out *FwlogPolicySpec
	var ok bool
	if into == nil {
		out = &FwlogPolicySpec{}
	} else {
		out, ok = into.(*FwlogPolicySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogPolicySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogPolicySpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		for k := range m.Filter {
			m.Filter[k] = "all"
		}
		m.Format = "syslog-bsd"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogPolicyStatus) Clone(into interface{}) (interface{}, error) {
	var out *FwlogPolicyStatus
	var ok bool
	if into == nil {
		out = &FwlogPolicyStatus{}
	} else {
		out, ok = into.(*FwlogPolicyStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogPolicyStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogPolicyStatus) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *FlowExportPolicy) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *FlowExportPolicy) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for FlowExportPolicy"))
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

func (m *FlowExportPolicy) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *FlowExportPolicySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FlowExportPolicySpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Exports {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sExports[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	for k, v := range m.MatchRules {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sMatchRules[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapTelemetry["FlowExportPolicySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTelemetry["FlowExportPolicySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *FlowExportPolicySpec) Normalize() {

	for k, v := range m.Exports {
		v.Normalize()
		m.Exports[k] = v

	}

	m.Format = FlowExportPolicySpec_Formats_normal[strings.ToLower(m.Format)]

	for k, v := range m.MatchRules {
		if v != nil {
			v.Normalize()
			m.MatchRules[k] = v
		}
	}

}

func (m *FlowExportPolicyStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FlowExportPolicyStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *FlowExportPolicyStatus) Normalize() {

}

func (m *FwlogPolicy) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *FwlogPolicy) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for FwlogPolicy"))
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

func (m *FwlogPolicy) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *FwlogPolicySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FwlogPolicySpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Config != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Config"
			if errs := m.Config.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	for k, v := range m.Targets {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sTargets[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapTelemetry["FwlogPolicySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTelemetry["FwlogPolicySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *FwlogPolicySpec) Normalize() {

	if m.Config != nil {
		m.Config.Normalize()
	}

	for k, v := range m.Filter {
		m.Filter[k] = FwlogFilter_normal[strings.ToLower(v)]
	}

	m.Format = MonitoringExportFormat_normal[strings.ToLower(m.Format)]

	for k, v := range m.Targets {
		v.Normalize()
		m.Targets[k] = v

	}

}

func (m *FwlogPolicyStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FwlogPolicyStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *FwlogPolicyStatus) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&FlowExportPolicy{},
		&FwlogPolicy{},
	)

	validatorMapTelemetry = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapTelemetry["FlowExportPolicySpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapTelemetry["FlowExportPolicySpec"]["all"] = append(validatorMapTelemetry["FlowExportPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FlowExportPolicySpec)

		if _, ok := FlowExportPolicySpec_Formats_vvalue[m.Format]; !ok {
			vals := []string{}
			for k1, _ := range FlowExportPolicySpec_Formats_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Format", vals)
		}
		return nil
	})

	validatorMapTelemetry["FlowExportPolicySpec"]["all"] = append(validatorMapTelemetry["FlowExportPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FlowExportPolicySpec)
		args := make([]string, 0)
		args = append(args, "1s")
		args = append(args, "24h")

		if err := validators.Duration(m.Interval, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Interval", err.Error())
		}
		return nil
	})

	validatorMapTelemetry["FlowExportPolicySpec"]["all"] = append(validatorMapTelemetry["FlowExportPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FlowExportPolicySpec)
		args := make([]string, 0)
		args = append(args, "1m")
		args = append(args, "30m")

		if err := validators.Duration(m.TemplateInterval, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"TemplateInterval", err.Error())
		}
		return nil
	})

	validatorMapTelemetry["FwlogPolicySpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapTelemetry["FwlogPolicySpec"]["all"] = append(validatorMapTelemetry["FwlogPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogPolicySpec)

		for k, v := range m.Filter {
			if _, ok := FwlogFilter_vvalue[v]; !ok {
				vals := []string{}
				for k1, _ := range FwlogFilter_vvalue {
					vals = append(vals, k1)
				}
				return fmt.Errorf("%v[%v] did not match allowed strings %v", path+"."+"Filter", k, vals)
			}
		}
		return nil
	})

	validatorMapTelemetry["FwlogPolicySpec"]["all"] = append(validatorMapTelemetry["FwlogPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogPolicySpec)

		if _, ok := MonitoringExportFormat_vvalue[m.Format]; !ok {
			vals := []string{}
			for k1, _ := range MonitoringExportFormat_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Format", vals)
		}
		return nil
	})

}
