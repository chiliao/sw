// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: techsupport.proto
*/
package monitoring

import (
	"errors"
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

var _ validators.DummyVar
var validatorMapTechsupport = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *TechSupportRequest) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "techsupport/", m.Name)
}

func (m *TechSupportRequest) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/techsupport/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *TechSupportNodeResult) Clone(into interface{}) (interface{}, error) {
	var out *TechSupportNodeResult
	var ok bool
	if into == nil {
		out = &TechSupportNodeResult{}
	} else {
		out, ok = into.(*TechSupportNodeResult)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *TechSupportNodeResult) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Status = "Scheduled"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *TechSupportRequest) Clone(into interface{}) (interface{}, error) {
	var out *TechSupportRequest
	var ok bool
	if into == nil {
		out = &TechSupportRequest{}
	} else {
		out, ok = into.(*TechSupportRequest)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *TechSupportRequest) Defaults(ver string) bool {
	m.Kind = "TechSupportRequest"
	m.Tenant, m.Namespace = "", ""
	var ret bool
	ret = m.Status.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *TechSupportRequestSpec) Clone(into interface{}) (interface{}, error) {
	var out *TechSupportRequestSpec
	var ok bool
	if into == nil {
		out = &TechSupportRequestSpec{}
	} else {
		out, ok = into.(*TechSupportRequestSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *TechSupportRequestSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TechSupportRequestSpec_NodeSelectorSpec) Clone(into interface{}) (interface{}, error) {
	var out *TechSupportRequestSpec_NodeSelectorSpec
	var ok bool
	if into == nil {
		out = &TechSupportRequestSpec_NodeSelectorSpec{}
	} else {
		out, ok = into.(*TechSupportRequestSpec_NodeSelectorSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *TechSupportRequestSpec_NodeSelectorSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TechSupportRequestStatus) Clone(into interface{}) (interface{}, error) {
	var out *TechSupportRequestStatus
	var ok bool
	if into == nil {
		out = &TechSupportRequestStatus{}
	} else {
		out, ok = into.(*TechSupportRequestStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *TechSupportRequestStatus) Defaults(ver string) bool {
	var ret bool
	for k := range m.ControllerNodeResults {
		if m.ControllerNodeResults[k] != nil {
			i := m.ControllerNodeResults[k]
			ret = i.Defaults(ver) || ret
		}
	}
	for k := range m.SmartNICNodeResults {
		if m.SmartNICNodeResults[k] != nil {
			i := m.SmartNICNodeResults[k]
			ret = i.Defaults(ver) || ret
		}
	}
	ret = true
	switch ver {
	default:
		m.Status = "Scheduled"
	}
	return ret
}

// Validators

func (m *TechSupportNodeResult) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapTechsupport["TechSupportNodeResult"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTechsupport["TechSupportNodeResult"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *TechSupportRequest) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		ret = m.ObjectMeta.Validate(ver, path+dlmtr+"ObjectMeta", ignoreStatus)
	}
	if m.Tenant != "" {
		ret = append(ret, errors.New("Tenant not allowed for TechSupportRequest"))
	}

	dlmtr := "."
	if path == "" {
		dlmtr = ""
	}
	npath := path + dlmtr + "Spec"
	if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
		ret = append(ret, errs...)
	}
	if !ignoreStatus {

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Status"
		if errs := m.Status.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *TechSupportRequestSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.CollectionSelector != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "CollectionSelector"
		if errs := m.CollectionSelector.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if m.NodeSelector != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "NodeSelector"
		if errs := m.NodeSelector.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *TechSupportRequestSpec_NodeSelectorSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Labels != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Labels"
		if errs := m.Labels.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *TechSupportRequestStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.ControllerNodeResults {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sControllerNodeResults[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	for k, v := range m.SmartNICNodeResults {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sSmartNICNodeResults[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapTechsupport["TechSupportRequestStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTechsupport["TechSupportRequestStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&TechSupportRequest{},
	)

	validatorMapTechsupport = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapTechsupport["TechSupportNodeResult"] = make(map[string][]func(string, interface{}) error)
	validatorMapTechsupport["TechSupportNodeResult"]["all"] = append(validatorMapTechsupport["TechSupportNodeResult"]["all"], func(path string, i interface{}) error {
		m := i.(*TechSupportNodeResult)

		if _, ok := TechSupportJobStatus_value[m.Status]; !ok {
			return fmt.Errorf("%v did not match allowed strings", path+"."+"Status")
		}
		return nil
	})

	validatorMapTechsupport["TechSupportRequestStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapTechsupport["TechSupportRequestStatus"]["all"] = append(validatorMapTechsupport["TechSupportRequestStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*TechSupportRequestStatus)

		if _, ok := TechSupportJobStatus_value[m.Status]; !ok {
			return fmt.Errorf("%v did not match allowed strings", path+"."+"Status")
		}
		return nil
	})

}
