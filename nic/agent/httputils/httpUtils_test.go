// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package httputils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"

	. "github.com/pensando/sw/venice/utils/testutils"
)

type testResponse struct {
	StatusCode int    `json:"status-code"`
	SelfLink   string `json:"self-link"`
}

func TestHTTPUtils(t *testing.T) {
	listenURL := "/tmp/test.sock"
	reqCount := 0

	os.Remove(listenURL)

	// register handlers for cni
	router := mux.NewRouter()
	sr := router.Headers("Content-Type", "application/json").Methods("POST").Subrouter()
	sr.HandleFunc("/dummy", MakeHTTPHandler(func(r *http.Request) (interface{}, error) {
		reqStr := ""
		err := ReadJSON(r, &reqStr)
		AssertOk(t, err, "reading json")
		if reqStr == "fail" {
			return nil, fmt.Errorf("Error")
		}

		Assert(t, (reqStr == "test"), "invalid req string")
		reqCount++
		return "success", nil
	}))

	sr.Methods("POST").Subrouter().HandleFunc("/dummy_resp", MakeHTTPHandler(dummyRespErrHandler))

	// create a listener
	l, err := net.ListenUnix("unix", &net.UnixAddr{Name: listenURL, Net: "unix"})
	AssertOk(t, err, "Listening to unix socket")
	defer os.Remove(listenURL)

	// start serving HTTP requests
	// http.Serve is a blocking call. so, do this in a separate go routine..
	go func() {
		http.Serve(l, router)
		defer l.Close()
	}()

	// make a call
	// create a HTTP client
	transport := &http.Transport{Dial: func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", listenURL)
	}}
	client := &http.Client{Transport: transport}

	// make a call to cni server
	buf, _ := json.Marshal("test")
	body := bytes.NewBuffer(buf)
	r, err := client.Post("http://localhost/dummy", "application/json", body)
	AssertOk(t, err, "Making http call")
	r.Body.Close()

	// check the HTTP status code
	Assert(t, (r.StatusCode == 200), "Incorrect http error code")
	Assert(t, (reqCount == 1), "incorrect req count")

	// make a call to cni server
	buf, _ = json.Marshal("test")
	body = bytes.NewBuffer(buf)
	r, err = client.Post("http://localhost/dummy_resp", "application/json", body)
	AssertOk(t, err, "Making http call")
	r.Body.Close()

	// verify failures are handled correctly
	buf, _ = json.Marshal("fail")
	body = bytes.NewBuffer(buf)
	r, err = client.Post("http://localhost/dummy", "application/json", body)
	AssertOk(t, err, "Making http call")
	r.Body.Close()
	Assert(t, (r.StatusCode != 200), "http succes while expecting failure")
}

func dummyRespErrHandler(r *http.Request) (interface{}, error) {
	resp := testResponse{
		StatusCode: http.StatusBadRequest,
		SelfLink:   r.RequestURI,
	}
	return &resp, errors.New("bad request test")
}
