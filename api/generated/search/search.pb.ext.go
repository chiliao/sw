// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package search is a auto generated package.
Input file: search.proto
*/
package search

import (
	fmt "fmt"
	"strings"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// PolicySearchResponse_MatchStatus_normal is a map of normalized values for the enum
var PolicySearchResponse_MatchStatus_normal = map[string]string{
	"match": "match",
	"miss":  "miss",
}

var PolicySearchResponse_MatchStatus_vname = map[int32]string{
	0: "match",
	1: "miss",
}

var PolicySearchResponse_MatchStatus_vvalue = map[string]int32{
	"match": 0,
	"miss":  1,
}

func (x PolicySearchResponse_MatchStatus) String() string {
	return PolicySearchResponse_MatchStatus_vname[int32(x)]
}

// SearchRequest_RequestMode_normal is a map of normalized values for the enum
var SearchRequest_RequestMode_normal = map[string]string{
	"full":    "full",
	"preview": "preview",
}

var SearchRequest_RequestMode_vname = map[int32]string{
	0: "full",
	1: "preview",
}

var SearchRequest_RequestMode_vvalue = map[string]int32{
	"full":    0,
	"preview": 1,
}

func (x SearchRequest_RequestMode) String() string {
	return SearchRequest_RequestMode_vname[int32(x)]
}

// SearchRequest_SortOrderEnum_normal is a map of normalized values for the enum
var SearchRequest_SortOrderEnum_normal = map[string]string{
	"ascending":  "ascending",
	"descending": "descending",
}

var SearchRequest_SortOrderEnum_vname = map[int32]string{
	0: "ascending",
	1: "descending",
}

var SearchRequest_SortOrderEnum_vvalue = map[string]int32{
	"ascending":  0,
	"descending": 1,
}

func (x SearchRequest_SortOrderEnum) String() string {
	return SearchRequest_SortOrderEnum_vname[int32(x)]
}

var _ validators.DummyVar
var validatorMapSearch = make(map[string]map[string][]func(string, interface{}) error)

// Clone clones the object into into or creates one of into is nil
func (m *CategoryAggregation) Clone(into interface{}) (interface{}, error) {
	var out *CategoryAggregation
	var ok bool
	if into == nil {
		out = &CategoryAggregation{}
	} else {
		out, ok = into.(*CategoryAggregation)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*CategoryAggregation))
	return out, nil
}

// Default sets up the defaults for the object
func (m *CategoryAggregation) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *CategoryPreview) Clone(into interface{}) (interface{}, error) {
	var out *CategoryPreview
	var ok bool
	if into == nil {
		out = &CategoryPreview{}
	} else {
		out, ok = into.(*CategoryPreview)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*CategoryPreview))
	return out, nil
}

// Default sets up the defaults for the object
func (m *CategoryPreview) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *ConfigEntry) Clone(into interface{}) (interface{}, error) {
	var out *ConfigEntry
	var ok bool
	if into == nil {
		out = &ConfigEntry{}
	} else {
		out, ok = into.(*ConfigEntry)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*ConfigEntry))
	return out, nil
}

// Default sets up the defaults for the object
func (m *ConfigEntry) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *Entry) Clone(into interface{}) (interface{}, error) {
	var out *Entry
	var ok bool
	if into == nil {
		out = &Entry{}
	} else {
		out, ok = into.(*Entry)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*Entry))
	return out, nil
}

// Default sets up the defaults for the object
func (m *Entry) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *EntryList) Clone(into interface{}) (interface{}, error) {
	var out *EntryList
	var ok bool
	if into == nil {
		out = &EntryList{}
	} else {
		out, ok = into.(*EntryList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EntryList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EntryList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *Error) Clone(into interface{}) (interface{}, error) {
	var out *Error
	var ok bool
	if into == nil {
		out = &Error{}
	} else {
		out, ok = into.(*Error)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*Error))
	return out, nil
}

// Default sets up the defaults for the object
func (m *Error) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *KindAggregation) Clone(into interface{}) (interface{}, error) {
	var out *KindAggregation
	var ok bool
	if into == nil {
		out = &KindAggregation{}
	} else {
		out, ok = into.(*KindAggregation)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*KindAggregation))
	return out, nil
}

// Default sets up the defaults for the object
func (m *KindAggregation) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *KindPreview) Clone(into interface{}) (interface{}, error) {
	var out *KindPreview
	var ok bool
	if into == nil {
		out = &KindPreview{}
	} else {
		out, ok = into.(*KindPreview)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*KindPreview))
	return out, nil
}

// Default sets up the defaults for the object
func (m *KindPreview) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PolicyMatchEntries) Clone(into interface{}) (interface{}, error) {
	var out *PolicyMatchEntries
	var ok bool
	if into == nil {
		out = &PolicyMatchEntries{}
	} else {
		out, ok = into.(*PolicyMatchEntries)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PolicyMatchEntries))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PolicyMatchEntries) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PolicyMatchEntry) Clone(into interface{}) (interface{}, error) {
	var out *PolicyMatchEntry
	var ok bool
	if into == nil {
		out = &PolicyMatchEntry{}
	} else {
		out, ok = into.(*PolicyMatchEntry)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PolicyMatchEntry))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PolicyMatchEntry) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PolicySearchRequest) Clone(into interface{}) (interface{}, error) {
	var out *PolicySearchRequest
	var ok bool
	if into == nil {
		out = &PolicySearchRequest{}
	} else {
		out, ok = into.(*PolicySearchRequest)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PolicySearchRequest))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PolicySearchRequest) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Namespace = "default"
		m.Tenant = "default"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *PolicySearchResponse) Clone(into interface{}) (interface{}, error) {
	var out *PolicySearchResponse
	var ok bool
	if into == nil {
		out = &PolicySearchResponse{}
	} else {
		out, ok = into.(*PolicySearchResponse)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PolicySearchResponse))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PolicySearchResponse) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Status = "match"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *SearchQuery) Clone(into interface{}) (interface{}, error) {
	var out *SearchQuery
	var ok bool
	if into == nil {
		out = &SearchQuery{}
	} else {
		out, ok = into.(*SearchQuery)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SearchQuery))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SearchQuery) Defaults(ver string) bool {
	var ret bool
	for k := range m.Texts {
		if m.Texts[k] != nil {
			i := m.Texts[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *SearchRequest) Clone(into interface{}) (interface{}, error) {
	var out *SearchRequest
	var ok bool
	if into == nil {
		out = &SearchRequest{}
	} else {
		out, ok = into.(*SearchRequest)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SearchRequest))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SearchRequest) Defaults(ver string) bool {
	var ret bool
	if m.Query != nil {
		ret = m.Query.Defaults(ver) || ret
	}
	ret = true
	switch ver {
	default:
		m.Aggregate = true
		m.MaxResults = 50
		m.Mode = "full"
		m.SortOrder = "ascending"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *SearchResponse) Clone(into interface{}) (interface{}, error) {
	var out *SearchResponse
	var ok bool
	if into == nil {
		out = &SearchResponse{}
	} else {
		out, ok = into.(*SearchResponse)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SearchResponse))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SearchResponse) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TenantAggregation) Clone(into interface{}) (interface{}, error) {
	var out *TenantAggregation
	var ok bool
	if into == nil {
		out = &TenantAggregation{}
	} else {
		out, ok = into.(*TenantAggregation)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*TenantAggregation))
	return out, nil
}

// Default sets up the defaults for the object
func (m *TenantAggregation) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TenantPreview) Clone(into interface{}) (interface{}, error) {
	var out *TenantPreview
	var ok bool
	if into == nil {
		out = &TenantPreview{}
	} else {
		out, ok = into.(*TenantPreview)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*TenantPreview))
	return out, nil
}

// Default sets up the defaults for the object
func (m *TenantPreview) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TextRequirement) Clone(into interface{}) (interface{}, error) {
	var out *TextRequirement
	var ok bool
	if into == nil {
		out = &TextRequirement{}
	} else {
		out, ok = into.(*TextRequirement)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*TextRequirement))
	return out, nil
}

// Default sets up the defaults for the object
func (m *TextRequirement) Defaults(ver string) bool {
	var ret bool
	return ret
}

// Validators and Requirements

func (m *CategoryAggregation) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *CategoryAggregation) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *CategoryAggregation) Normalize() {

}

func (m *CategoryPreview) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *CategoryPreview) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *CategoryPreview) Normalize() {

}

func (m *ConfigEntry) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *ConfigEntry) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	return ret
}

func (m *ConfigEntry) Normalize() {

	m.ObjectMeta.Normalize()

}

func (m *Entry) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *Entry) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *Entry) Normalize() {

}

func (m *EntryList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EntryList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *EntryList) Normalize() {

}

func (m *Error) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *Error) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *Error) Normalize() {

}

func (m *KindAggregation) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *KindAggregation) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *KindAggregation) Normalize() {

}

func (m *KindPreview) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *KindPreview) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *KindPreview) Normalize() {

}

func (m *PolicyMatchEntries) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "entries"

		for _, v := range m.Entries {
			if v != nil {
				v.References(tenant, tag, resp)
			}
		}
	}
}

func (m *PolicyMatchEntries) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Entries {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEntries[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *PolicyMatchEntries) Normalize() {

	for k, v := range m.Entries {
		if v != nil {
			v.Normalize()
			m.Entries[k] = v
		}
	}

}

func (m *PolicyMatchEntry) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "rule"

		if m.Rule != nil {
			m.Rule.References(tenant, tag, resp)
		}

	}
}

func (m *PolicyMatchEntry) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Rule != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Rule"
			if errs := m.Rule.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *PolicyMatchEntry) Normalize() {

	if m.Rule != nil {
		m.Rule.Normalize()
	}

}

func (m *PolicySearchRequest) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *PolicySearchRequest) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *PolicySearchRequest) Normalize() {

}

func (m *PolicySearchResponse) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "results"

		for _, v := range m.Results {
			if v != nil {
				v.References(tenant, tag, resp)
			}
		}
	}
}

func (m *PolicySearchResponse) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Results {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sResults[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapSearch["PolicySearchResponse"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapSearch["PolicySearchResponse"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *PolicySearchResponse) Normalize() {

	for k, v := range m.Results {
		if v != nil {
			v.Normalize()
			m.Results[k] = v
		}
	}

	m.Status = PolicySearchResponse_MatchStatus_normal[strings.ToLower(m.Status)]

}

func (m *SearchQuery) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SearchQuery) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Fields != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Fields"
			if errs := m.Fields.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}

	if m.Labels != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Labels"
			if errs := m.Labels.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	for k, v := range m.Texts {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sTexts[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapSearch["SearchQuery"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapSearch["SearchQuery"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *SearchQuery) Normalize() {

	if m.Fields != nil {
		m.Fields.Normalize()
	}

	if m.Labels != nil {
		m.Labels.Normalize()
	}

	for k, v := range m.Texts {
		if v != nil {
			v.Normalize()
			m.Texts[k] = v
		}
	}

}

func (m *SearchRequest) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SearchRequest) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Query != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Query"
			if errs := m.Query.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	if vs, ok := validatorMapSearch["SearchRequest"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapSearch["SearchRequest"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *SearchRequest) Normalize() {

	m.Mode = SearchRequest_RequestMode_normal[strings.ToLower(m.Mode)]

	if m.Query != nil {
		m.Query.Normalize()
	}

	m.SortOrder = SearchRequest_SortOrderEnum_normal[strings.ToLower(m.SortOrder)]

}

func (m *SearchResponse) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SearchResponse) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *SearchResponse) Normalize() {

}

func (m *TenantAggregation) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *TenantAggregation) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *TenantAggregation) Normalize() {

}

func (m *TenantPreview) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *TenantPreview) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *TenantPreview) Normalize() {

}

func (m *TextRequirement) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *TextRequirement) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapSearch["TextRequirement"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapSearch["TextRequirement"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *TextRequirement) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

	validatorMapSearch = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapSearch["PolicySearchResponse"] = make(map[string][]func(string, interface{}) error)
	validatorMapSearch["PolicySearchResponse"]["all"] = append(validatorMapSearch["PolicySearchResponse"]["all"], func(path string, i interface{}) error {
		m := i.(*PolicySearchResponse)

		if _, ok := PolicySearchResponse_MatchStatus_vvalue[m.Status]; !ok {
			vals := []string{}
			for k1, _ := range PolicySearchResponse_MatchStatus_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Status", vals)
		}
		return nil
	})

	validatorMapSearch["SearchQuery"] = make(map[string][]func(string, interface{}) error)
	validatorMapSearch["SearchQuery"]["all"] = append(validatorMapSearch["SearchQuery"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchQuery)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "64")

		for _, v := range m.Categories {
			if err := validators.EmptyOr(validators.StrLen, v, args); err != nil {
				return fmt.Errorf("%v failed validation: %s", path+"."+"Categories", err.Error())
			}
		}
		return nil
	})

	validatorMapSearch["SearchQuery"]["all"] = append(validatorMapSearch["SearchQuery"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchQuery)
		for k, v := range m.Kinds {
			if err := validators.EmptyOr(validators.ValidKind, v, nil); err != nil {
				return fmt.Errorf("%v[%v] failed validation: %s", path+"."+"Kinds", k, err.Error())
			}
		}

		return nil
	})

	validatorMapSearch["SearchRequest"] = make(map[string][]func(string, interface{}) error)
	validatorMapSearch["SearchRequest"]["all"] = append(validatorMapSearch["SearchRequest"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchRequest)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "1023")

		if err := validators.IntRange(m.From, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"From", err.Error())
		}
		return nil
	})

	validatorMapSearch["SearchRequest"]["all"] = append(validatorMapSearch["SearchRequest"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchRequest)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "8192")

		if err := validators.IntRange(m.MaxResults, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"MaxResults", err.Error())
		}
		return nil
	})

	validatorMapSearch["SearchRequest"]["all"] = append(validatorMapSearch["SearchRequest"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchRequest)

		if _, ok := SearchRequest_RequestMode_vvalue[m.Mode]; !ok {
			vals := []string{}
			for k1, _ := range SearchRequest_RequestMode_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Mode", vals)
		}
		return nil
	})

	validatorMapSearch["SearchRequest"]["all"] = append(validatorMapSearch["SearchRequest"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchRequest)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "256")

		if err := validators.EmptyOr(validators.StrLen, m.QueryString, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"QueryString", err.Error())
		}
		return nil
	})

	validatorMapSearch["SearchRequest"]["all"] = append(validatorMapSearch["SearchRequest"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchRequest)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "256")

		if err := validators.EmptyOr(validators.StrLen, m.SortBy, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"SortBy", err.Error())
		}
		return nil
	})

	validatorMapSearch["SearchRequest"]["all"] = append(validatorMapSearch["SearchRequest"]["all"], func(path string, i interface{}) error {
		m := i.(*SearchRequest)

		if _, ok := SearchRequest_SortOrderEnum_vvalue[m.SortOrder]; !ok {
			vals := []string{}
			for k1, _ := range SearchRequest_SortOrderEnum_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"SortOrder", vals)
		}
		return nil
	})

	validatorMapSearch["TextRequirement"] = make(map[string][]func(string, interface{}) error)
	validatorMapSearch["TextRequirement"]["all"] = append(validatorMapSearch["TextRequirement"]["all"], func(path string, i interface{}) error {
		m := i.(*TextRequirement)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "256")

		for _, v := range m.Text {
			if err := validators.EmptyOr(validators.StrLen, v, args); err != nil {
				return fmt.Errorf("%v failed validation: %s", path+"."+"Text", err.Error())
			}
		}
		return nil
	})

}
