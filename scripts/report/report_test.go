package main

//
import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pensando/test-infra-tracker/types"

	. "github.com/pensando/sw/venice/utils/testutils"
)

// test vectors
const (
	test           = "github.com/pensando/sw/nic/agent/netagent/state"
	testFailStdOut = `ts=2018-02-06T19:00:16.125059Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.12523Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
proto: no coders for api.TypeMeta
proto: no encoder for TypeMeta api.TypeMeta [GetProperties]
proto: no coders for api.ObjectMeta
proto: no encoder for ObjectMeta api.ObjectMeta [GetProperties]
proto: no coders for netproto.NetworkSpec
proto: no encoder for Spec netproto.NetworkSpec [GetProperties]
proto: no coders for netproto.NetworkStatus
proto: no encoder for Status netproto.NetworkStatus [GetProperties]
proto: no coders for api.Timestamp
proto: no encoder for CreationTime api.Timestamp [GetProperties]
proto: no coders for api.Timestamp
proto: no encoder for ModTime api.Timestamp [GetProperties]
proto: no coders for types.Timestamp
proto: no encoder for Timestamp types.Timestamp [GetProperties]
ts=2018-02-06T19:00:16.125663Z module=Default pid=94604 caller=network.go:29 level=info msg="Received duplicate network create for ep {TypeMeta:<Kind:\"Network\" > ObjectMeta:<Name:\"default\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<IPv4Subnet:\"10.1.1.0/24\" IPv4Gateway:\"10.1.1.254\" > Status:<NetworkID:1 > }"
ts=2018-02-06T19:00:16.125727Z module=Default pid=94604 caller=network.go:24 level=error msg="Network TypeMeta:<Kind:\"Network\" > ObjectMeta:<Name:\"default\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<IPv4Subnet:\"10.1.1.0/24\" IPv4Gateway:\"10.1.1.254\" > Status:<NetworkID:1 >  already exists"
ts=2018-02-06T19:00:16.125779Z module=Default pid=94604 caller=network.go:137 level=error msg="Network {Name:default Tenant:default Namespace: ResourceVersion: UUID: Labels:map[] CreationTime:{Timestamp:{Seconds:0 Nanos:0}} ModTime:{Timestamp:{Seconds:0 Nanos:0}}} not found"
ts=2018-02-06T19:00:16.125839Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.125887Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
ts=2018-02-06T19:00:16.125917Z module=Default pid=94604 caller=network.go:117 level=info msg="Nothing to update."
ts=2018-02-06T19:00:16.125983Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.126025Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
proto: no coders for api.TypeMeta
proto: no encoder for TypeMeta api.TypeMeta [GetProperties]
proto: no coders for api.ObjectMeta
proto: no encoder for ObjectMeta api.ObjectMeta [GetProperties]
proto: no coders for netproto.EndpointSpec
proto: no encoder for Spec netproto.EndpointSpec [GetProperties]
proto: no coders for netproto.EndpointStatus
proto: no encoder for Status netproto.EndpointStatus [GetProperties]
ts=2018-02-06T19:00:16.126308Z module=Default pid=94604 caller=endpoint.go:65 level=info msg="Received duplicate endpoint create for ep {TypeMeta:<Kind:\"Endpoint\" > ObjectMeta:<Name:\"testEndpoint\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<EndpointUUID:\"testEndpointUUID\" WorkloadUUID:\"testWorkloadUUID\" NetworkName:\"default\" > Status:<NodeUUID:\"some-unique-id\" > }"
ts=2018-02-06T19:00:16.126339Z module=Default pid=94604 caller=endpoint.go:73 level=error msg="Error finding the network invalid. Err: Network not found"
ts=2018-02-06T19:00:16.126405Z module=Default pid=94604 caller=endpoint.go:60 level=error msg="Endpoint TypeMeta:<Kind:\"Endpoint\" > ObjectMeta:<Name:\"testEndpoint\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<EndpointUUID:\"testEndpointUUID\" WorkloadUUID:\"testWorkloadUUID\" NetworkName:\"default\" > Status:<NodeUUID:\"some-unique-id\" >  already exists. New ep {TypeMeta:<Kind:\"Endpoint\" > ObjectMeta:<Name:\"testEndpoint\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<EndpointUUID:\"testEndpointUUID2\" WorkloadUUID:\"testWorkloadUUID2\" NetworkName:\"default\" > Status:<NodeUUID:\"some-unique-id\" > }"
ts=2018-02-06T19:00:16.126443Z module=Default pid=94604 caller=endpoint.go:206 level=error msg="Endpoint \"default|testEndpoint\" was not found"
ts=2018-02-06T19:00:16.126515Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.126558Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
ts=2018-02-06T19:00:16.126655Z module=Default pid=94604 caller=endpoint.go:65 level=info msg="Received duplicate endpoint create for ep {TypeMeta:<Kind:\"Endpoint\" > ObjectMeta:<Name:\"testEndpoint\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<EndpointUUID:\"testEndpointUUID\" WorkloadUUID:\"testWorkloadUUID\" NetworkName:\"default\" > Status:<> }"
ts=2018-02-06T19:00:16.126708Z module=Default pid=94604 caller=endpoint.go:60 level=error msg="Endpoint TypeMeta:<Kind:\"Endpoint\" > ObjectMeta:<Name:\"testEndpoint\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<EndpointUUID:\"testEndpointUUID\" WorkloadUUID:\"testWorkloadUUID\" NetworkName:\"default\" > Status:<>  already exists. New ep {TypeMeta:<Kind:\"Endpoint\" > ObjectMeta:<Name:\"testEndpoint\" Tenant:\"default\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<EndpointUUID:\"testEndpointUUID2\" WorkloadUUID:\"testWorkloadUUID2\" NetworkName:\"default\" > Status:<> }"
ts=2018-02-06T19:00:16.126725Z module=Default pid=94604 caller=endpoint.go:206 level=error msg="Endpoint \"default|testEndpoint\" was not found"
ts=2018-02-06T19:00:16.126768Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
--- FAIL: TestSecurityGroupCreateDelete (0.00s)
	testutils.go:35: netagent_state_test.go:498: Failed to create agent &state.NetAgent{Mutex:sync.Mutex{state:0, sema:0x0}, store:(*emstore.MemStore)(0xc420157da0), nodeUUID:"some-unique-id", datapath:(*state.mockDatapath)(0xc4201bc210), ctrlerif:(*state.mockCtrler)(0xc42000e140), networkDB:map[string]*netproto.Network{}, endpointDB:map[string]*netproto.Endpoint{}, secgroupDB:map[string]*netproto.SecurityGroup{}, tenantDB:map[string]*netproto.Tenant{"default|default":(*netproto.Tenant)(0xc420170a50)}}

ts=2018-02-06T19:00:16.12689Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.126914Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
ts=2018-02-06T19:00:16.126945Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type sgID"
ts=2018-02-06T19:00:16.127035Z module=Default pid=94604 caller=endpoint.go:142 level=error msg="Can not change network after endpoint is created. old default, new unknown"
ts=2018-02-06T19:00:16.127065Z module=Default pid=94604 caller=endpoint.go:161 level=error msg="Error finding security group unknown. Err: Security group not found"
ts=2018-02-06T19:00:16.127108Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.127143Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type sgID"
ts=2018-02-06T19:00:16.12717Z module=Default pid=94604 caller=security.go:25 level=error msg="Error finding peer group unknown. Err: Security group not found"
ts=2018-02-06T19:00:16.127182Z module=Default pid=94604 caller=security.go:182 level=error msg="Error adding sg rules. Err: Security group not found"
ts=2018-02-06T19:00:16.127231Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.127278Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
ts=2018-02-06T19:00:16.128205Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
proto: no coders for api.TypeMeta
proto: no encoder for TypeMeta api.TypeMeta [GetProperties]
proto: no coders for api.ObjectMeta
proto: no encoder for ObjectMeta api.ObjectMeta [GetProperties]
proto: no coders for netproto.TenantSpec
proto: no encoder for Spec netproto.TenantSpec [GetProperties]
proto: no coders for netproto.TenantStatus
proto: no encoder for Status netproto.TenantStatus [GetProperties]
ts=2018-02-06T19:00:16.128401Z module=Default pid=94604 caller=tenant.go:30 level=info msg="Received duplicate tenant create {TypeMeta:<Kind:\"Tenant\" > ObjectMeta:<Name:\"testTenant\" Tenant:\"testTenant\" CreationTime:<time:<> > ModTime:<time:<> > > Spec:<> Status:<TenantID:2 > }"
ts=2018-02-06T19:00:16.128452Z module=Default pid=94604 caller=tenant.go:131 level=error msg="Tenant {Name:testTenant Tenant:testTenant Namespace: ResourceVersion: UUID: Labels:map[] CreationTime:{Timestamp:{Seconds:0 Nanos:0}} ModTime:{Timestamp:{Seconds:0 Nanos:0}}} not found"
ts=2018-02-06T19:00:16.128508Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.128572Z module=Default pid=94604 caller=tenant.go:106 level=info msg="Nothing to update."
ts=2018-02-06T19:00:16.128616Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type tenantID"
ts=2018-02-06T19:00:16.128649Z module=Default pid=94604 caller=resid.go:31 level=info msg="New ID generated for resource type networkID"
ts=2018-02-06T19:00:16.128667Z module=Default pid=94604 caller=network.go:50 level=error msg="Could not find the tenant: {tenant not found <nil>}"
FAIL
coverage: 76.8% of statements
FAIL	github.com/pensando/sw/nic/agent/netagent/state	0.021s

2018/02/06 11:00:16 Test Failure: github.com/pensando/sw/nic/agent/netagent/state
2018/02/06 11:00:16 could not get failed tests: test execution failed
2018/02/06 11:00:20 Insufficient code coverage for the following packages:
2018/02/06 11:00:20 github.com/pensando/sw/nic/agent/cmd/halctl`

	testCovFailedStdout      = `ok  	github.com/pensando/sw/nic/agent/netagent/state	0.020s	coverage: 41.7% of statements`
	testInvalidPkgName       = "github.com/pensando/foo"
	testCovIgnoreNoTestFiles = `?   	github.com/pensando/sw/api	[no test files]`
	testCovIgnoreSkipped     = `ok  	github.com/pensando/sw/test/e2e	0.028s`
	testCovIgnoreZeroCov     = `ok  	github.com/pensando/sw/api/integration	0.201s	coverage: 0.0% of statements`
)

// Happy path tests
func TestCoverage(t *testing.T) {
	tr := TestReport{
		Results: []*Target{
			{
				Name: test,
			},
		},
	}

	tr.runCoverage()

	_, err := tr.reportToJSON()
	AssertOk(t, err, "Could not convert report to json")

	tr.testCoveragePass()
	AssertEquals(t, false, tr.RunFailed, "Expected the run to pass and it failed")
}

func TestCoverageFail(t *testing.T) {
	tr := TestReport{
		Results: []*Target{
			{
				Name:     test,
				Coverage: 45.0,
				Error:    ErrTestCovFailed.Error(),
			},
		},
	}
	tr.testCoveragePass()
	AssertEquals(t, true, tr.RunFailed, "expected the run to fail due to low coverage")
}

func TestCoverageFailTracker(t *testing.T) {
	if !isJobdCI() {
		t.Skip("Skip because not in job-ci environment")
	}

	tr := TestReport{
		Results: []*Target{
			{
				Name:     test,
				Coverage: 45.0,
				Error:    ErrTestCovFailed.Error(),
			},
		},
	}

	// testtracker should receive corresponding data
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var rep types.Reports
		AssertOk(t, json.NewDecoder(r.Body).Decode(&rep), "Could not parse tracker report")
		AssertEquals(t, int32(-1), rep.Testcases[0].Result, "Receive wrong test case result")
		AssertEquals(t, int32(45), rep.Testcases[0].Coverage, "Receive wrong test case coverage")
		AssertEquals(t, ErrTestCovFailed.Error(), rep.Testcases[0].Detail, "Receive wrong test case detail")
	}))
	defer ts.Close()

	jobRepoBak := os.Getenv("JOB_FORK_REPOSITORY")
	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", swBaseRepo), "Could not set environment variable")
	AssertOk(t, tr.sendToTestTracker(strings.TrimPrefix(ts.URL, "http://")), "Could not send report to test tracker")
	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", jobRepoBak), "Could not set environment variable")
}

func TestFilterFailTests(t *testing.T) {
	tr := TestReport{
		RunFailed: true,
		Results: []*Target{
			{
				Name:     "test",
				Coverage: 45.0,
				Error:    ErrTestCovFailed.Error(),
			},
			{
				Name:     "test2",
				Coverage: 88.0,
				Error:    "",
			},
		},
	}
	tr2 := tr.filterFailedTests()
	AssertEquals(t, tr.RunFailed, tr2.RunFailed, "RunFailed should be same after filterFailedTests")
	AssertEquals(t, len(tr2.Results), 1, "Only failed test should be present")
	AssertEquals(t, tr.Results[0], tr2.Results[0], "Only failed test should be present")
}

func TestFilterFailTestsTracker(t *testing.T) {
	if !isJobdCI() {
		t.Skip("Skip because not in job-ci environment")
	}

	tr := TestReport{
		RunFailed: true,
		Results: []*Target{
			{
				Name:     "test",
				Coverage: 45.0,
				Error:    ErrTestCovFailed.Error(),
			},
			{
				Name:     "test2",
				Coverage: 88.0,
				Error:    "",
			},
		},
	}
	// testtracker should receive corresponding data
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rep types.Reports
		AssertOk(t, json.NewDecoder(r.Body).Decode(&rep), "Could not parse tracker report")
		AssertEquals(t, 2, len(rep.Testcases), "Receive wrong number of test case ")

		foundtest := false
		foundtest2 := false
		for _, tc := range rep.Testcases {
			if tc.Name == "test" {
				foundtest = true
				AssertEquals(t, int32(-1), tc.Result, "Receive wrong test case result")
				AssertEquals(t, int32(45), tc.Coverage, "Receive wrong test case coverage")
				AssertEquals(t, ErrTestCovFailed.Error(), tc.Detail, "Receive wrong test case detail")
			}
			if tc.Name == "test2" {
				foundtest2 = true
				AssertEquals(t, int32(1), tc.Result, "Receive wrong test case result")
				AssertEquals(t, int32(88), tc.Coverage, "Receive wrong test case coverage")
				AssertEquals(t, "", tc.Detail, "Receive wrong test case detail")
			}
		}
		AssertEquals(t, true, foundtest, "test not found")
		AssertEquals(t, true, foundtest2, "test2 not found")
	}))

	jobRepoBak := os.Getenv("JOB_FORK_REPOSITORY")
	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", swBaseRepo), "Could not set environment variable")
	AssertOk(t, tr.sendToTestTracker(strings.TrimPrefix(ts.URL, "http://")), "Could not send report to test tracker")
	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", jobRepoBak), "Could not set environment variable")
}

func TestStdOutParsing(t *testing.T) {
	tgt := Target{
		Name: test,
	}

	err := tgt.parseCmdOutput([]byte(testFailStdOut))
	AssertEquals(t, ErrTestFailed, err, "parsing stdout failed")

	err = tgt.getCoveragePercent([]byte(testCovFailedStdout))
	AssertOk(t, err, "failed to parse coverage output")
	AssertEquals(t, tgt.Coverage, 41.7, "did not get expected coverage results")
}

func TestInvalidPackageName(t *testing.T) {
	tgt := Target{
		Name: testInvalidPkgName,
	}
	var ignoredPackages []string
	tgt.test(ignoredPackages)
	AssertEquals(t, ErrTestFailed.Error(), tgt.Error, "expected the test to fail, it passed instead")
}

func TestCoverageIgnore(t *testing.T) {
	tgt := Target{
		Name: test,
	}
	err := tgt.getCoveragePercent([]byte(testCovIgnoreNoTestFiles))
	AssertOk(t, err, "coverage parsing expected to pass")
	AssertEquals(t, 100.0, tgt.Coverage, "Expected coverage 100%% for missing test files")

	tgt.Coverage = 0.0
	err = tgt.getCoveragePercent([]byte(testCovIgnoreSkipped))
	AssertOk(t, err, "coverage parsing expected to pass")
	AssertEquals(t, 100.0, tgt.Coverage, "Expected coverage 100%% for missing test files")

	tgt.Coverage = 0.0
	err = tgt.getCoveragePercent([]byte(testCovIgnoreZeroCov))
	AssertOk(t, err, "coverage parsing expected to pass")
	AssertEquals(t, 100.0, tgt.Coverage, "Expected coverage 100%% for missing test files")
}

func TestTrackerBasic(t *testing.T) {
	if !isJobdCI() {
		t.Skip("Skip because not in job-ci environment")
	}

	tr := TestReport{
		Results: []*Target{
			{
				Name: test,
			},
		},
	}

	tr.runCoverage()

	// testtracker should receive data
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rep types.Reports
		AssertOk(t, json.NewDecoder(r.Body).Decode(&rep), "Could not parse tracker report")
		AssertEquals(t, swBaseRepo, rep.Repository, "Receive wrong repo")
		Assert(t, rep.TargetID != int32(0), "Receive wrong TargetID")
		Assert(t, rep.SHA != "", "Receive wrong SHA")
		Assert(t, rep.SHATitle != "", "Receive wrong SHA title")
		AssertEquals(t, testbed, rep.Testbed, "Receive wrong testbed")
		AssertEquals(t, 1, len(rep.Testcases), "Receive wrong number of test cases")
		AssertEquals(t, "nic/agent/netagent/state", rep.Testcases[0].Name, "Receive wrong test case name")
		AssertEquals(t, "nic/agent/netagent/state", rep.Testcases[0].Description, "Receive wrong test case description")
		Assert(t, time.Now().Format("Mon Jan 2 -0700 MST 2006") == rep.Testcases[0].FinishTime.Format("Mon Jan 2 -0700 MST 2006"), "Receive wrong test case finish time")
		AssertEquals(t, int32(tr.Results[0].Coverage), rep.Testcases[0].Coverage, "Receive wrong test case coverage")
		AssertEquals(t, fmt.Sprintf("http://jobd/logs/%d", rep.TargetID), rep.Testcases[0].LogURL, "Receive wrong test case log url")
		AssertEquals(t, int32(1), rep.Testcases[0].Result, "Receive wrong test case result")
		AssertEquals(t, "nic", rep.Testcases[0].Area, "Receive wrong test case area")
		AssertEquals(t, "agent/netagent/state", rep.Testcases[0].Subarea, "Receive wrong test case subarea")
	}))
	defer ts.Close()

	u := strings.TrimPrefix(ts.URL, "http://")

	jobRepoBak := os.Getenv("JOB_FORK_REPOSITORY")
	targetIDBak := os.Getenv("TARGET_ID")
	AssertOk(t, os.Unsetenv("JOB_FORK_REPOSITORY"), "Could not unset environment variable")
	AssertOk(t, os.Unsetenv("TARGET_ID"), "Could not unset environment variable")
	// empty job repo environment should not sent report
	AssertEquals(t, errNotBaseRepo.Error(), tr.sendToTestTracker(u).Error(), "Could not check empty sw repo")

	// non pensando/sw repo should report failure too
	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", "pensando/swtest"), "Could not set environment variable")

	AssertEquals(t, errNotBaseRepo.Error(), tr.sendToTestTracker(u).Error(), "Could not check non sw repo")

	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", swBaseRepo), "Could not set environment variable")

	// empty TARGET_ID should fail too
	Assert(t, tr.sendToTestTracker(u) != nil, "Could not check empty target id")

	AssertOk(t, os.Setenv("TARGET_ID", targetIDBak), "Could not set environment variable")
	AssertOk(t, tr.sendToTestTracker(u), "Could not send report to test tracker")

	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", jobRepoBak), "Could not set environment variable")
}

func TestTrackerRetry(t *testing.T) {
	if !isJobdCI() {
		t.Skip("Skip because not in job-ci environment")
	}

	tr := TestReport{
		RunFailed: true,
		Results: []*Target{
			{
				Name:     "test",
				Coverage: 45.0,
				Error:    ErrTestCovFailed.Error(),
			},
		},
	}

	// reduce retryTimeout to speed up test
	retryTrackerInterval = time.Second

	// testtracker should receive corresponding data
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	u := "127.0.0.1:7890"

	go func() {
		// start http test server after 3 seconds
		time.Sleep(3 * time.Second)
		l, err := net.Listen("tcp", u)
		AssertOk(t, err, "Could not listen on 127.0.0.1:7890")
		ts.Listener = l
		ts.Start()
	}()

	jobRepoBak := os.Getenv("JOB_FORK_REPOSITORY")
	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", swBaseRepo), "Could not set environment variable")
	AssertOk(t, tr.sendToTestTracker(u), "Could not send report to test tracker")

	// Not retry if server return error
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test failure", 500)
	}))

	err := tr.sendToTestTracker(strings.TrimPrefix(ts2.URL, "http://"))
	Assert(t, err != nil, "Could receive failure without error")
	AssertEquals(t, "test failure", err.Error(), "Could not receive expected failure reply")

	AssertOk(t, os.Setenv("JOB_FORK_REPOSITORY", jobRepoBak), "Could not set environment variable")
}
