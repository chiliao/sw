// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package restapi

import (
	"expvar"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"

	"github.com/pensando/sw/venice/utils/log"
)

// this package contains the REST API provided by the agent

// RestServer is the REST api server
type RestServer struct {
	listenURL  string       // URL where http server is listening
	listener   net.Listener // socket listener
	httpServer *http.Server // HTTP server
}

// Response captures the HTTP Response sent by Agent REST Server
type Response struct {
	StatusCode int      `json:"status-code,omitempty"`
	Error      string   `json:"error,omitempty"`
	References []string `json:"references,omitempty"`
}

type routeAddFunc func(*mux.Router, *RestServer)

var prefixRoutes map[string]routeAddFunc

func init() {
	if prefixRoutes == nil {
		prefixRoutes = make(map[string]routeAddFunc)
	}
}

// NewRestServer creates a new HTTP server servicg REST api
func NewRestServer(listenURL string) (*RestServer, error) {
	// create server instance
	srv := RestServer{
		listenURL: listenURL,
	}

	// if no URL was specified, just return (used during unit/integ tests)
	if listenURL == "" {
		return &srv, nil
	}

	// setup the top level routes
	router := mux.NewRouter()

	for prefix, subRouter := range prefixRoutes {
		sub := router.PathPrefix(prefix).Subrouter().StrictSlash(true)
		subRouter(sub, &srv)
	}
	router.Methods("GET").Subrouter().Handle("/debug/vars", expvar.Handler())
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/", pprof.Index)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/block", pprof.Handler("block").ServeHTTP)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	router.Methods("GET").Subrouter().HandleFunc("/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)

	log.Infof("Starting server at %s", listenURL)

	// listener
	listener, err := net.Listen("tcp", listenURL)
	if err != nil {
		log.Errorf("Error starting listener. Err: %v", err)
		return nil, err
	}
	srv.listener = listener

	// create a http server
	srv.httpServer = &http.Server{Addr: listenURL, Handler: router}
	go srv.httpServer.Serve(listener)

	return &srv, nil
}

// GetListenURL returns the listen URL of the http server
func (s *RestServer) GetListenURL() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}

	return ""
}

// Stop stops the http server
func (s *RestServer) Stop() error {
	if s.httpServer != nil {
		s.httpServer.Close()
	}

	return nil
}
