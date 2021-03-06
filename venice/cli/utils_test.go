package vcli

import (
	"bytes"
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/pensando/sw/api/generated/workload"
	"github.com/pensando/sw/venice/utils/telemetryclient"
)

func TestSliceToMap(t *testing.T) {
	m, err := sliceToMap([]string{"key1" + kvSplitter + "val1", "key2" + kvSplitter + "val2"})
	if err != nil {
		t.Fatalf("erorr convering slice to map: %s", err)
	}

	if m["key1"] != "val1" || m["key2"] != "val2" {
		t.Fatalf("error finding correct keys in the map")
	}
}

func TestMatchRe(t *testing.T) {
	re := regexp.MustCompile(".*buk.*")
	if !matchRe(re, "timbuktoo") {
		t.Fatalf("couldn't match a regex")
	}
	if matchRe(re, "paris") {
		t.Fatalf("could match invalid regex")
	}
}

func TestMatchLabel(t *testing.T) {
	matchLabels := map[string]string{"key1": "val1", "key2": "val2"}
	if !matchLabel(matchLabels, map[string]string{"key2": "val2"}) {
		t.Fatalf("unable to match keys")
	}
	if matchLabel(matchLabels, map[string]string{"key3": "val3"}) {
		t.Fatalf("able to match non existing keys")
	}
	if !matchLabel(map[string]string{}, map[string]string{"key1": "val1"}) {
		t.Fatalf("unable to match empty match criteria")
	}
}

func TestMatchString(t *testing.T) {
	matchStrings := []string{"one", "two", "three"}
	if !matchString(matchStrings, "three") {
		t.Fatalf("unable to match string")
	}
	if matchString(matchStrings, "four") {
		t.Fatalf("able to match non existing string")
	}
	if !matchString([]string{}, "two") {
		t.Fatalf("unable to match empty match criteria")
	}

	matchSubStrings := []string{"usage", "memory"}
	if !matchSubString(matchSubStrings, "CPUUsage") {
		t.Fatalf("unable to find contained substring from the list")
	}
	if !matchSubString(matchSubStrings, "TotalMemory") {
		t.Fatalf("unable to find contained substring from the list")
	}
	if matchSubString(matchSubStrings, "RxBytes") {
		t.Fatalf("able to find contained substring from the list")
	}
}

func TestLookupForJson(t *testing.T) {
	jsonMixedData := `rest get for obj T:<> O:<Name:"vm23" Tenant:"default" CreationTime:<Time:<> > ModTime:<Time:<> > > Spec:<> Status:<> 
ts=2018-08-28T18:16:33.264618214Z module=Default pid=3725 caller=new_logger.go:171 module=cli service=WorkloadV1 method=AutoGetWorkload result=Success duration=2.925191ms level=audit
		{
		  "kind": "Workload",
		  "api-version": "v1",
		  "meta": {
		    "name": "vm23",
		    "tenant": "default",
		    "resource-version": "5",
		    "uuid": "4b589680-817f-43fd-9866-1d092f58c907",
		    "creation-time": "2018-08-28T18:16:33.259122478Z",
		    "mod-time": "2018-08-28T18:16:33.259123321Z",
		    "self-link": "/configs/workload/v1/tenant/default/workloads/vm23"
		  },
		  "spec": {
		    "host-name": "esx-node12",
		    "interfaces": [
			  {
		        "mac-address": "0011.2233.4455",
       		    "micro-seg-vlan": 1001,
		        "external-vlan": 100
		      },
		      {
				"mac-address": "0022.3344.5566",
		        "micro-seg-vlan": 2001,
		        "external-vlan": 200
		      }
		    ]
		  },
		  "status": {}
		}`

	obj := &workload.Workload{}
	if err := lookForJSON(jsonMixedData, obj); err != nil {
		t.Fatalf("error parsing json data: %s", err)
	}

	if obj.Kind != "Workload" || obj.Name != "vm23" {
		t.Fatalf("invalid object returned %+v", obj)
	}
	if obj.Spec.HostName != "esx-node12" {
		t.Fatalf("invalid hostname in obj: %+v", obj)
	}
	if len(obj.Spec.Interfaces) != 2 || !reflect.DeepEqual(obj.Spec.Interfaces,
		[]workload.WorkloadIntfSpec{
			{MACAddress: "0011.2233.4455", ExternalVlan: 100, MicroSegVlan: 1001},
			{MACAddress: "0022.3344.5566", ExternalVlan: 200, MicroSegVlan: 2001}}) {
		t.Fatalf("invalid Interfaces: %+v", obj.Spec.Interfaces)
	}
}

func TestFindJsonRecord(t *testing.T) {
	jsonRecs := `
        {
                "kind": "Workload",
                "meta": {
                    "name": "TestFindJsonRecordVm1"
                },
                "spec": {
                    "host-name": "esx-node12",
                    "interfaces": [
                      {
                        "mac-address": "0011.2233.4455",
                        "micro-seg-vlan": 1001,
                        "external-vlan": 101
                      },
                      {
                        "mac-address": "0022.3344.5566",
                        "micro-seg-vlan": 2001,
                        "external-vlan": 201
                      }
                    ]
                  }
        }
        {
                "kind": "Workload",
                "meta": {
                    "name": "TestFindJsonRecordVm2"
                },
                "spec": {
                    "host-name": "esx-node12",
                    "interfaces": [
                      {
                        "mac-address": "0011.2233.0055",
                        "micro-seg-vlan": 1002,
                        "external-vlan": 102
                      },
                      {
                        "mac-address": "0022.3344.0066",
                        "micro-seg-vlan": 2002,
                        "external-vlan": 202
                      }
                    ]
                }
        }
`
	jsonRec, nextIdx := findJSONRecord(jsonRecs, 0)
	if jsonRec == "" || nextIdx < 0 {
		t.Fatalf("unable to local json record: %s", jsonRecs)
	}
	obj := &workload.Workload{}
	if err := json.Unmarshal([]byte(jsonRec), obj); err != nil {
		t.Fatalf("error '%s' unmarshing record: '%s'\n", err, jsonRec)
	}
	if obj.Kind != "Workload" || obj.Name != "TestFindJsonRecordVm1" {
		t.Fatalf("invalid object returned %+v", obj)
	}
	if obj.Spec.HostName != "esx-node12" {
		t.Fatalf("invalid hostname in obj: %+v", obj)
	}
	if len(obj.Spec.Interfaces) != 2 || !reflect.DeepEqual(obj.Spec.Interfaces,
		[]workload.WorkloadIntfSpec{
			{MACAddress: "0011.2233.4455", ExternalVlan: 101, MicroSegVlan: 1001},
			{MACAddress: "0022.3344.5566", ExternalVlan: 201, MicroSegVlan: 2001}}) {
		t.Fatalf("invalid Interfaces: %+v", obj.Spec.Interfaces)
	}

	jsonRec, nextIdx = findJSONRecord(jsonRecs, nextIdx)
	if jsonRec == "" || nextIdx < 0 {
		t.Fatalf("unable to local json record: %s", jsonRecs)
	}
	obj = &workload.Workload{}
	if err := json.Unmarshal([]byte(jsonRec), obj); err != nil {
		t.Fatalf("error '%s' unmarshing record: '%s'\n", err, jsonRec)
	}
	if obj.Kind != "Workload" || obj.Name != "TestFindJsonRecordVm2" {
		t.Fatalf("invalid object returned json-rec %s %+v", jsonRec, obj)
	}
	if obj.Spec.HostName != "esx-node12" {
		t.Fatalf("invalid hostname in obj: %+v", obj)
	}
	if len(obj.Spec.Interfaces) != 2 || !reflect.DeepEqual(obj.Spec.Interfaces,
		[]workload.WorkloadIntfSpec{
			{MACAddress: "0011.2233.0055", ExternalVlan: 102, MicroSegVlan: 1002},
			{MACAddress: "0022.3344.0066", ExternalVlan: 202, MicroSegVlan: 2002}}) {
		t.Fatalf("invalid Interfaces: %+v", obj.Spec.Interfaces)
	}

	jsonRec, nextIdx = findJSONRecord(jsonRecs, nextIdx)
	if jsonRec != "" || nextIdx >= 0 {
		t.Fatalf("able to locate record after nextIdx %d: %s", nextIdx, jsonRecs)
	}
}

func TestDumpStruct(t *testing.T) {
	obj := &workload.Workload{}
	obj.Kind = "workload"
	obj.Name = "workload1"
	obj.Labels = map[string]string{"label-key1": "label-value1", "label-key2": "label=value2"}
	obj.Spec = workload.WorkloadSpec{
		HostName: "node021",
		Interfaces: []workload.WorkloadIntfSpec{
			{MACAddress: "1111.1111.1111", ExternalVlan: 11, MicroSegVlan: 1000},
			{MACAddress: "2222.2222.2222", ExternalVlan: 22, MicroSegVlan: 2000},
			{MACAddress: "3333.3333.3333", ExternalVlan: 33, MicroSegVlan: 3000},
		},
	}

	// dump yml output
	out := dumpStruct(true, obj)
	if !strings.Contains(string(out), "3333.3333.3333") {
		t.Fatalf("unable to find interface object: out %s", out)
	}

	// dump json output
	out = dumpStruct(false, obj)
	if !strings.Contains(string(out), "3333.3333.3333") {
		t.Fatalf("unable to find interface object: out %s", out)
	}
}

func TestMatchLineFields(t *testing.T) {
	out := `name         labels     mac-address external-vlan         host-name       micro-seg-vlan
		----         ------     ----------- -------------         ---------       --------------
		TestReadVm1  key1=val1  00de.edde.edd0 55  dc12_rack3_bm4  2222
		             key2=val2  00f0.0df0.0dd0 66                  3333
		TestReadVm3  key1=val1  00de.edde.edd0 55  dc12_rack3_bm4  2224
		             key3=val3                                  `
	if !matchLineFields(out, []string{"TestReadVm1", "dc12_rack3_bm4", "00de.edde.edd0", "55", "2222"}) {
		t.Fatalf("unable to match fields")
	}

	if !matchLines(out, []string{"key3=val3"}) {
		t.Fatalf("unable to match lines")
	}

	if matchLineFields(out, []string{"TestReadVm11", "dc12_rack3_bm4", "00de.edde.edd0", "55", "2224"}) {
		t.Fatalf("able to match unexpected fields")
	}

	if matchLines(out, []string{"key3=val33"}) {
		t.Fatalf("able to match expected lines")
	}

}

func TestPrintMetrics(t *testing.T) {
	mem := make([]byte, 10000, 10000)
	b := bytes.NewBuffer(mem)

	series := &telemetryclient.MetricsResultSeries{
		Columns: []string{"time", "CPUUsedPercent", "DiskFree", "DiskTotal"},
		Values: [][]interface{}{
			[]interface{}{"2019-01-28T08:17:53.162022171Z", "19.7", "33911", "38945", "5034"},
			[]interface{}{"2019-01-28T08:18:03.199000923Z", "0", "33909", "38945", "5035"},
			[]interface{}{"2019-01-28T08:18:13.214700016Z", "0", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:18:23.222706563Z", "1.81", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:18:33.229694889Z", "0", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:18:43.237116035Z", "0", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:18:53.24621946Z", "0", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:19:03.252830807Z", "1.12", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:19:13.261806661Z", "0.79", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:19:23.269494761Z", "0.71", "33912", "38945", "5032"},
			[]interface{}{"2019-01-28T08:19:33.276044028Z", "0", "33912", "38945", "5032"},
		},
	}

	printSeries(b, 0, []string{}, series)

	outString := string(b.Bytes())
	if !strings.Contains(outString, "2019-01-28T08:18:43.237116035Z") {
		t.Fatalf("unable to match output:\n\"\n%s\n\"\n", outString)
	}
}
