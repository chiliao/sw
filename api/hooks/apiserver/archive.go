package impl

import (
	"context"
	"errors"

	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/api/generated/monitoring"
	apiintf "github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

type archiveHooks struct {
	logger log.Logger
}

func (a *archiveHooks) archiveRequestDeleteCheck(ctx context.Context, kvs kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryrun bool, i interface{}) (interface{}, bool, error) {
	into := &monitoring.ArchiveRequest{}
	if err := kvs.Get(ctx, key, into); err != nil {
		return i, true, err
	}
	switch into.Status.Status {
	// archive request can only be deleted after it is processed
	case "", monitoring.ArchiveRequestStatus_Scheduled.String(), monitoring.ArchiveRequestStatus_Running.String():
		return i, true, errors.New("archive request not processed yet")
	default:
		txn.AddComparator(kvstore.Compare(kvstore.WithVersion(key), "=", into.ResourceVersion))
		return i, true, nil
	}
}

func (a *archiveHooks) archiveRequestUpgradeCheck(ctx context.Context, kvs kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryrun bool, i interface{}) (interface{}, bool, error) {
	// Check if Rollout is in progress. Reject if so.
	verObj := cluster.Version{}
	vkey := verObj.MakeKey(string(apiclient.GroupCluster))
	err := kvs.Get(ctx, vkey, &verObj)
	if err == nil {
		txn.AddComparator(kvstore.Compare(kvstore.WithVersion(vkey), "=", verObj.ResourceVersion))
		if verObj.Status.RolloutBuildVersion != "" {
			return i, true, errors.New("rollout in progress, restore operation not allowed")
		}
	}
	return i, true, nil
}

func registerArchiveHooks(svc apiserver.Service, logger log.Logger) {
	r := archiveHooks{}
	r.logger = logger.WithContext("Service", "ArchiveHooks")
	logger.Log("msg", "registering Hooks")
	svc.GetCrudService("ArchiveRequest", apiintf.DeleteOper).WithPreCommitHook(r.archiveRequestDeleteCheck)
	svc.GetCrudService("ArchiveRequest", apiintf.CreateOper).WithPreCommitHook(r.archiveRequestUpgradeCheck)
}
