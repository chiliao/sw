// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package security is a auto generated package.
Input file: svc_security.proto
*/
package security

import (
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
func (m *AppList) MakeKey(prefix string) string {
	obj := App{}
	return obj.MakeKey(prefix)
}

func (m *AppList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *CertificateList) MakeKey(prefix string) string {
	obj := Certificate{}
	return obj.MakeKey(prefix)
}

func (m *CertificateList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *FirewallProfileList) MakeKey(prefix string) string {
	obj := FirewallProfile{}
	return obj.MakeKey(prefix)
}

func (m *FirewallProfileList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *NetworkSecurityPolicyList) MakeKey(prefix string) string {
	obj := NetworkSecurityPolicy{}
	return obj.MakeKey(prefix)
}

func (m *NetworkSecurityPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *SecurityGroupList) MakeKey(prefix string) string {
	obj := SecurityGroup{}
	return obj.MakeKey(prefix)
}

func (m *SecurityGroupList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *TrafficEncryptionPolicyList) MakeKey(prefix string) string {
	obj := TrafficEncryptionPolicy{}
	return obj.MakeKey(prefix)
}

func (m *TrafficEncryptionPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgAppWatchHelper) MakeKey(prefix string) string {
	obj := App{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgCertificateWatchHelper) MakeKey(prefix string) string {
	obj := Certificate{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgFirewallProfileWatchHelper) MakeKey(prefix string) string {
	obj := FirewallProfile{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgNetworkSecurityPolicyWatchHelper) MakeKey(prefix string) string {
	obj := NetworkSecurityPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgSecurityGroupWatchHelper) MakeKey(prefix string) string {
	obj := SecurityGroup{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgTrafficEncryptionPolicyWatchHelper) MakeKey(prefix string) string {
	obj := TrafficEncryptionPolicy{}
	return obj.MakeKey(prefix)
}

// Clone clones the object into into or creates one of into is nil
func (m *AppList) Clone(into interface{}) (interface{}, error) {
	var out *AppList
	var ok bool
	if into == nil {
		out = &AppList{}
	} else {
		out, ok = into.(*AppList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AppList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AppList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAppWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAppWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgAppWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgAppWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgAppWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAppWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAppWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAppWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgAppWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgAppWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgAppWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAppWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgCertificateWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgCertificateWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgCertificateWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgCertificateWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgCertificateWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgCertificateWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgCertificateWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgCertificateWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgCertificateWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgCertificateWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgCertificateWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgCertificateWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgFirewallProfileWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgFirewallProfileWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgFirewallProfileWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgFirewallProfileWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgFirewallProfileWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgFirewallProfileWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgFirewallProfileWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgFirewallProfileWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgFirewallProfileWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgFirewallProfileWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgFirewallProfileWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgFirewallProfileWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNetworkSecurityPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNetworkSecurityPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgNetworkSecurityPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgNetworkSecurityPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNetworkSecurityPolicyWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNetworkSecurityPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgSecurityGroupWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgSecurityGroupWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgSecurityGroupWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgSecurityGroupWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgSecurityGroupWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgSecurityGroupWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgSecurityGroupWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgSecurityGroupWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgSecurityGroupWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgSecurityGroupWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgSecurityGroupWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgSecurityGroupWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgTrafficEncryptionPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgTrafficEncryptionPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgTrafficEncryptionPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgTrafficEncryptionPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgTrafficEncryptionPolicyWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgTrafficEncryptionPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *CertificateList) Clone(into interface{}) (interface{}, error) {
	var out *CertificateList
	var ok bool
	if into == nil {
		out = &CertificateList{}
	} else {
		out, ok = into.(*CertificateList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*CertificateList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *CertificateList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *FirewallProfileList) Clone(into interface{}) (interface{}, error) {
	var out *FirewallProfileList
	var ok bool
	if into == nil {
		out = &FirewallProfileList{}
	} else {
		out, ok = into.(*FirewallProfileList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FirewallProfileList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FirewallProfileList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *NetworkSecurityPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *NetworkSecurityPolicyList
	var ok bool
	if into == nil {
		out = &NetworkSecurityPolicyList{}
	} else {
		out, ok = into.(*NetworkSecurityPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*NetworkSecurityPolicyList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *NetworkSecurityPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *SecurityGroupList) Clone(into interface{}) (interface{}, error) {
	var out *SecurityGroupList
	var ok bool
	if into == nil {
		out = &SecurityGroupList{}
	} else {
		out, ok = into.(*SecurityGroupList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SecurityGroupList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SecurityGroupList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TrafficEncryptionPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *TrafficEncryptionPolicyList
	var ok bool
	if into == nil {
		out = &TrafficEncryptionPolicyList{}
	} else {
		out, ok = into.(*TrafficEncryptionPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*TrafficEncryptionPolicyList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *TrafficEncryptionPolicyList) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *AppList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AppList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AppList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *AutoMsgAppWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgAppWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAppWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgAppWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgAppWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgAppWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgCertificateWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgCertificateWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgCertificateWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgCertificateWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgCertificateWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgCertificateWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgFirewallProfileWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgFirewallProfileWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgFirewallProfileWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgFirewallProfileWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgFirewallProfileWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgFirewallProfileWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgNetworkSecurityPolicyWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "events"

		for _, v := range m.Events {
			if v != nil {
				v.References(tenant, tag, resp)
			}
		}
	}
}

func (m *AutoMsgNetworkSecurityPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgNetworkSecurityPolicyWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "object"

		if m.Object != nil {
			m.Object.References(tenant, tag, resp)
		}

	}
}

func (m *AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgNetworkSecurityPolicyWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgSecurityGroupWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgSecurityGroupWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgSecurityGroupWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgSecurityGroupWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgSecurityGroupWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgSecurityGroupWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgTrafficEncryptionPolicyWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgTrafficEncryptionPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgTrafficEncryptionPolicyWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgTrafficEncryptionPolicyWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *CertificateList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *CertificateList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *CertificateList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *FirewallProfileList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FirewallProfileList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *FirewallProfileList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *NetworkSecurityPolicyList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "items"

		for _, v := range m.Items {
			if v != nil {
				v.References(tenant, tag, resp)
			}
		}
	}
}

func (m *NetworkSecurityPolicyList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *NetworkSecurityPolicyList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *SecurityGroupList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SecurityGroupList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *SecurityGroupList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *TrafficEncryptionPolicyList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *TrafficEncryptionPolicyList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *TrafficEncryptionPolicyList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

}
