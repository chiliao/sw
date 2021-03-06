package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"path"

	"github.com/go-martini/martini"

	cmd "github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/venice/cmd/env"
	"github.com/pensando/sw/venice/cmd/ops"
	"github.com/pensando/sw/venice/globals"
	// Import utils/debug pkg to publish runtime stats as part of its pkg init
	_ "github.com/pensando/sw/venice/utils/debug/stats"
	"github.com/pensando/sw/venice/utils/errors"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

// constants used by REST interface
const (
	uRLPrefix   = "/api/v1"
	clusterURL  = "/cluster"
	servicesURL = "/services"
	expvarURL   = "/vars"
)

// NewClusterRESTHandler creates a handler for the for cluster create/get API.
func NewClusterRESTHandler() *martini.ClassicMartini {
	m := martini.Classic()

	m.Post(uRLPrefix+clusterURL, ClusterCreateHandler)
	m.Get(uRLPrefix+clusterURL+"/:id", ClusterGetHandler)

	return m
}

// ClusterCreateHandler handles the REST call for cluster creation.
func ClusterCreateHandler(w http.ResponseWriter, req *http.Request) {
	env.Mutex.Lock()
	defer env.Mutex.Unlock()

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	cluster := cmd.Cluster{}
	if err := decoder.Decode(&cluster); err != nil {
		errors.SendBadRequest(w, err.Error())
		return
	}

	log.Infof("Cluster create args: %+v", cluster)

	ops.RunHTTP(w, ops.NewClusterCreateOp(&cluster))
}

// ClusterGetHandler returns the cluster information.
func ClusterGetHandler(w http.ResponseWriter, params martini.Params) {
	cluster := cmd.Cluster{}
	id := params["id"]

	if id == "" {
		errors.SendNotFound(w, "Cluster", "")
		return
	}

	if env.KVStore == nil {
		errors.SendNotFound(w, "Cluster", "")
		return
	}

	if err := env.KVStore.Get(context.Background(), path.Join(globals.ClusterKey, id), &cluster); err != nil {
		if kvstore.IsKeyNotFoundError(err) {
			errors.SendNotFound(w, "Cluster", id)
			return
		}
		errors.SendInternalError(w, err)
		return
	}

	cluster.Status.Leader = env.LeaderService.Leader()

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(&cluster); err != nil {
		log.Errorf("Failed to encode with error: %v", err)
	}
}
