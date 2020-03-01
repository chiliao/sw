// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package auth is a auto generated package.
Input file: svc_auth.proto
*/
package auth

import (
	"context"
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
func (m *AuthenticationPolicyList) MakeKey(prefix string) string {
	obj := AuthenticationPolicy{}
	return obj.MakeKey(prefix)
}

func (m *AuthenticationPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *RoleBindingList) MakeKey(prefix string) string {
	obj := RoleBinding{}
	return obj.MakeKey(prefix)
}

func (m *RoleBindingList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *RoleList) MakeKey(prefix string) string {
	obj := Role{}
	return obj.MakeKey(prefix)
}

func (m *RoleList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *UserList) MakeKey(prefix string) string {
	obj := User{}
	return obj.MakeKey(prefix)
}

func (m *UserList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *UserPreferenceList) MakeKey(prefix string) string {
	obj := UserPreference{}
	return obj.MakeKey(prefix)
}

func (m *UserPreferenceList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgAuthenticationPolicyWatchHelper) MakeKey(prefix string) string {
	obj := AuthenticationPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgRoleBindingWatchHelper) MakeKey(prefix string) string {
	obj := RoleBinding{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgRoleWatchHelper) MakeKey(prefix string) string {
	obj := Role{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgUserPreferenceWatchHelper) MakeKey(prefix string) string {
	obj := UserPreference{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgUserWatchHelper) MakeKey(prefix string) string {
	obj := User{}
	return obj.MakeKey(prefix)
}

// Clone clones the object into into or creates one of into is nil
func (m *AuthenticationPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *AuthenticationPolicyList
	var ok bool
	if into == nil {
		out = &AuthenticationPolicyList{}
	} else {
		out, ok = into.(*AuthenticationPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AuthenticationPolicyList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AuthenticationPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAuthenticationPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAuthenticationPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgAuthenticationPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgAuthenticationPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgAuthenticationPolicyWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAuthenticationPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgAuthenticationPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgAuthenticationPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgAuthenticationPolicyWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgRoleBindingWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgRoleBindingWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgRoleBindingWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgRoleBindingWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgRoleBindingWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgRoleBindingWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgRoleBindingWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgRoleBindingWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgRoleBindingWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgRoleBindingWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgRoleBindingWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgRoleBindingWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgRoleWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgRoleWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgRoleWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgRoleWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgRoleWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgRoleWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgRoleWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgRoleWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgRoleWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgRoleWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgRoleWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgRoleWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgUserPreferenceWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgUserPreferenceWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgUserPreferenceWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgUserPreferenceWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgUserPreferenceWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgUserPreferenceWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgUserPreferenceWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgUserPreferenceWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgUserPreferenceWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgUserPreferenceWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgUserPreferenceWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgUserPreferenceWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgUserWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgUserWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgUserWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgUserWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgUserWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgUserWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgUserWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgUserWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgUserWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgUserWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgUserWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgUserWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleBindingList) Clone(into interface{}) (interface{}, error) {
	var out *RoleBindingList
	var ok bool
	if into == nil {
		out = &RoleBindingList{}
	} else {
		out, ok = into.(*RoleBindingList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*RoleBindingList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleBindingList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleList) Clone(into interface{}) (interface{}, error) {
	var out *RoleList
	var ok bool
	if into == nil {
		out = &RoleList{}
	} else {
		out, ok = into.(*RoleList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*RoleList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *UserList) Clone(into interface{}) (interface{}, error) {
	var out *UserList
	var ok bool
	if into == nil {
		out = &UserList{}
	} else {
		out, ok = into.(*UserList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*UserList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *UserList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *UserPreferenceList) Clone(into interface{}) (interface{}, error) {
	var out *UserPreferenceList
	var ok bool
	if into == nil {
		out = &UserPreferenceList{}
	} else {
		out, ok = into.(*UserPreferenceList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*UserPreferenceList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *UserPreferenceList) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *AuthenticationPolicyList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AuthenticationPolicyList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AuthenticationPolicyList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *AutoMsgAuthenticationPolicyWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgAuthenticationPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgAuthenticationPolicyWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgRoleBindingWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *AutoMsgRoleBindingWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgRoleBindingWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgRoleBindingWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *AutoMsgRoleBindingWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgRoleBindingWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgRoleWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgRoleWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgRoleWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgRoleWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgRoleWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgRoleWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgUserPreferenceWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgUserPreferenceWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgUserPreferenceWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgUserPreferenceWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgUserPreferenceWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgUserPreferenceWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgUserWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgUserWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgUserWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgUserWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgUserWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgUserWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *RoleBindingList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *RoleBindingList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *RoleBindingList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *RoleList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *RoleList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *RoleList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *UserList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *UserList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *UserList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *UserPreferenceList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *UserPreferenceList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *UserPreferenceList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

// Transformers

func (m *AuthenticationPolicyList) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	for i, v := range m.Items {
		c := *v
		if err := c.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
		m.Items[i] = &c
	}
	return nil
}

func (m *AuthenticationPolicyList) EraseSecrets() {
	for _, v := range m.Items {
		v.EraseSecrets()
	}
	return
}

func (m *AutoMsgAuthenticationPolicyWatchHelper) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	for i, v := range m.Events {
		c := *v
		if err := c.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
		m.Events[i] = &c
	}
	return nil
}

func (m *AutoMsgAuthenticationPolicyWatchHelper) EraseSecrets() {
	for _, v := range m.Events {
		v.EraseSecrets()
	}
	return
}

func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {

	if m.Object == nil {
		return nil
	}
	if err := m.Object.ApplyStorageTransformer(ctx, toStorage); err != nil {
		return err
	}
	return nil
}

func (m *AutoMsgAuthenticationPolicyWatchHelper_WatchEvent) EraseSecrets() {

	if m.Object == nil {
		return
	}
	m.Object.EraseSecrets()

	return
}

func (m *AutoMsgUserWatchHelper) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	for i, v := range m.Events {
		c := *v
		if err := c.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
		m.Events[i] = &c
	}
	return nil
}

func (m *AutoMsgUserWatchHelper) EraseSecrets() {
	for _, v := range m.Events {
		v.EraseSecrets()
	}
	return
}

func (m *AutoMsgUserWatchHelper_WatchEvent) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {

	if m.Object == nil {
		return nil
	}
	if err := m.Object.ApplyStorageTransformer(ctx, toStorage); err != nil {
		return err
	}
	return nil
}

func (m *AutoMsgUserWatchHelper_WatchEvent) EraseSecrets() {

	if m.Object == nil {
		return
	}
	m.Object.EraseSecrets()

	return
}

func (m *UserList) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	for i, v := range m.Items {
		c := *v
		if err := c.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
		m.Items[i] = &c
	}
	return nil
}

func (m *UserList) EraseSecrets() {
	for _, v := range m.Items {
		v.EraseSecrets()
	}
	return
}

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

}
