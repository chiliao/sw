package debug

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	. "github.com/pensando/sw/venice/utils/testutils"
	"github.com/pensando/sw/venice/utils/tsdb"
)

var socketInfo = map[string]string{
	"k1": "v1",
	"k2": "v2",
}

func socketInfoFunction() interface{} {
	return socketInfo
}

func TestDebugInfo(t *testing.T) {
	dbgSock := "test.sock"
	debugSocket := New(socketInfoFunction)
	err := debugSocket.StartServer(dbgSock)
	AssertOk(t, err, "Failed to start debug socket")

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", dbgSock)
			},
		},
	}

	resp, err := client.Get("http://localhost/debug")
	AssertOk(t, err, "Http client received an error")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	AssertOk(t, err, "failed to read response")

	expJSON, err := json.Marshal(socketInfo)
	AssertOk(t, err, "failed to marshall expected response")
	exp := string(expJSON) + "\n"

	AssertEquals(t, exp, string(body),
		fmt.Sprintf("expected returned object %v, got %v", exp, string(body)))

	err = debugSocket.Destroy()
	AssertOk(t, err, "failed to close debug socket")
}

func createSocketClient(dbgSock string) http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", dbgSock)
			},
		},
	}
}

func TestDebugMetrics(t *testing.T) {
	dbgSock := "test.sock"
	debugSocket := New(socketInfoFunction)
	err := debugSocket.BuildMetricObj("debugTable", nil)
	Assert(t, err != nil, "building table before tsdb init should have failed")

	err = debugSocket.StartServer(dbgSock)
	AssertOk(t, err, "Failed to start debug socket")

	tsdb.Init(context.Background(), &tsdb.Opts{})

	err = debugSocket.BuildMetricObj("debugTable", nil)
	AssertOk(t, err, "Failed to build metrics table")
	table := debugSocket.MetricObj

	table.Counter("rx_ep_create_msg").Inc()
	table.Counter("rx_ep_create_msg").Inc()
	table.Counter("rx_ep_create_msg").Inc()
	table.Counter("peer_disconnects").Inc()
	table.Counter("peer_rpc_failure").Inc()
	table.Gauge("cpu_in_use").Set(34.4)
	table.Gauge("mem_in_use").Set(102)
	table.String("version").Set("v0.1", time.Time{})

	lms := []tsdb.LocalMetric{}
	client := createSocketClient(dbgSock)

	AssertEventually(t, func() (bool, interface{}) {
		resp, err := client.Get("http://localhost/debugMetrics")
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		err = json.Unmarshal(body, &lms)
		if err != nil {
			return false, err
		}

		return len(lms) > 0, nil

	}, "failed to get response from localhost")

	Assert(t, len(lms) == 1, fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["rx_ep_create_msg"] == "3", fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["peer_disconnects"] == "1", fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["peer_rpc_failure"] == "1", fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["cpu_in_use"] == "34.4", fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["mem_in_use"] == "102", fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["version"] == "v0.1", fmt.Sprintf("invalid lms attributes %+v", lms))

	// Server is running by now, so we don't need an assert eventually
	lms = []tsdb.LocalMetric{}
	resp, err := client.Get("http://localhost/debugMetrics/debugTable/rx_ep_create_msg")
	AssertOk(t, err, "GET request failed")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	AssertOk(t, err, "Failed to read response")

	err = json.Unmarshal(body, &lms)
	AssertOk(t, err, "Failed to unmarshall response")

	Assert(t, len(lms) == 1, fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, len(lms[0].Attributes) == 1, fmt.Sprintf("invalid lms attributes %+v", lms))
	Assert(t, lms[0].Attributes["rx_ep_create_msg"] == "3", fmt.Sprintf("invalid lms attributes %+v", lms))

	err = debugSocket.Destroy()
	AssertOk(t, err, "Failed to destory ")
	Assert(t, debugSocket.MetricObj == nil, "Failed to set Metric Obj to nil")

}
