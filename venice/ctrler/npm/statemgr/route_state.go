// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package statemgr

import (
	"errors"
	"fmt"
	"sync"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/ctkit"
	"github.com/pensando/sw/api/generated/network"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"
	"github.com/pensando/sw/venice/utils/runtime"
)

var smgrRoute *SmRoute

// SmRoute is object statemgr for Route object
type SmRoute struct {
	featureMgrBase
	sm *Statemgr
}

// CompleteRegistration is the callback function statemgr calls after init is done
func (sma *SmRoute) CompleteRegistration() {
	// if featureflags.IsOVerlayRoutingEnabled() == false {
	// 	return
	// }

	sma.sm.SetRoutingConfigReactor(smgrRoute)
}

func init() {
	mgr := MustGetStatemgr()
	smgrRoute = &SmRoute{
		sm: mgr,
	}

	mgr.Register("statemgrroute", smgrRoute)
}

// RoutingConfigState is a wrapper for RoutingConfig object
type RoutingConfigState struct {
	sync.Mutex
	RoutingConfig *ctkit.RoutingConfig `json:"-"` // RoutingConfig object
}

// RoutingConfigStateFromObj converts from memdb object to RoutingConfig state
func RoutingConfigStateFromObj(obj runtime.Object) (*RoutingConfigState, error) {
	switch obj.(type) {
	case *ctkit.RoutingConfig:
		rtcfg := obj.(*ctkit.RoutingConfig)
		switch rtcfg.HandlerCtx.(type) {
		case *RoutingConfigState:
			state := rtcfg.HandlerCtx.(*RoutingConfigState)
			return state, nil
		default:
			return nil, errors.New("incorrect object type")
		}
	default:
		return nil, errors.New("incorrect object type")
	}
}

func convertRoutingConfig(rtcfg *RoutingConfigState) *netproto.RoutingConfig {
	meta := api.ObjectMeta{
		Tenant:          globals.DefaultTenant,
		Namespace:       globals.DefaultNamespace,
		Name:            rtcfg.RoutingConfig.Name,
		GenerationID:    rtcfg.RoutingConfig.GenerationID,
		ResourceVersion: rtcfg.RoutingConfig.ResourceVersion,
		UUID:            rtcfg.RoutingConfig.UUID,
	}

	obj := &netproto.RoutingConfig{
		TypeMeta:   rtcfg.RoutingConfig.TypeMeta,
		ObjectMeta: meta,
	}

	obj.Spec = netproto.RoutingConfigSpec{}

	if rtcfg.RoutingConfig.Spec.BGPConfig != nil {
		obj.Spec.BGPConfig = &netproto.BGPConfig{
			DSCAutoConfig:     rtcfg.RoutingConfig.Spec.BGPConfig.DSCAutoConfig,
			ASNumber:          rtcfg.RoutingConfig.Spec.BGPConfig.ASNumber.ASNumber,
			KeepaliveInterval: rtcfg.RoutingConfig.Spec.BGPConfig.KeepaliveInterval,
			Holdtime:          rtcfg.RoutingConfig.Spec.BGPConfig.Holdtime,
		}

		for _, nbr := range rtcfg.RoutingConfig.Spec.BGPConfig.Neighbors {
			neighbor := new(netproto.BGPNeighbor)
			neighbor.Shutdown = nbr.Shutdown
			neighbor.IPAddress = nbr.IPAddress
			neighbor.RemoteAS = nbr.RemoteAS.ASNumber
			neighbor.MultiHop = nbr.MultiHop
			neighbor.Password = nbr.Password
			neighbor.DSCAutoConfig = nbr.DSCAutoConfig
			neighbor.KeepaliveInterval = nbr.KeepaliveInterval
			neighbor.Holdtime = nbr.Holdtime
			for _, addr := range nbr.EnableAddressFamilies {
				neighbor.EnableAddressFamilies = append(neighbor.EnableAddressFamilies, addr)
			}

			obj.Spec.BGPConfig.Neighbors = append(obj.Spec.BGPConfig.Neighbors, neighbor)
		}
	}
	log.Infof("Converted Routing Config [%+v]", obj)
	return obj
}

// FindRoutingConfig finds routingconfig by name
func (sma *SmRoute) FindRoutingConfig(tenant, ns, name string) (*RoutingConfigState, error) {
	// find it in db
	obj, err := sma.sm.FindObject("RoutingConfig", tenant, ns, name)
	if err != nil {
		return nil, err
	}

	return RoutingConfigStateFromObj(obj)
}

// GetRoutingConfigWatchOptions gets options
func (sma *SmRoute) GetRoutingConfigWatchOptions() *api.ListWatchOptions {
	opts := api.ListWatchOptions{}
	return &opts
}

// NewRoutingConfigState creates a new RoutingConfigState
func NewRoutingConfigState(routecfg *ctkit.RoutingConfig, sma *SmRoute) (*RoutingConfigState, error) {
	rtcfg := &RoutingConfigState{
		RoutingConfig: routecfg,
	}
	routecfg.HandlerCtx = rtcfg
	return rtcfg, nil
}

// OnRoutingConfigCreate creates local routingcfg state based on watch event
func (sma *SmRoute) OnRoutingConfigCreate(obj *ctkit.RoutingConfig) error {
	log.Info("OnRoutingConfigCreate: received: ", obj.Spec)

	// create new routingconfig state
	rtcfg, err := NewRoutingConfigState(obj, sma)
	if err != nil {
		log.Errorf("Error creating routing config state. Err: %v", err)
		return err
	}

	log.Info("OnRoutingConfigCreate created: ", rtcfg.RoutingConfig)

	// store it in local DB
	err = sma.sm.mbus.AddObjectWithReferences(rtcfg.RoutingConfig.MakeKey(string(apiclient.GroupNetwork)), convertRoutingConfig(rtcfg), references(obj))
	if err != nil {
		log.Errorf("could not add RoutingConfig to nimbus (%s)", err)
		return err
	}
	return nil
}

// OnRoutingConfigUpdate handles RoutingConfig update
func (sma *SmRoute) OnRoutingConfigUpdate(oldcfg *ctkit.RoutingConfig, newcfg *network.RoutingConfig) error {
	log.Info("OnRoutingConfigUpdate: received: ", oldcfg.Spec, newcfg.Spec)

	// see if anything changed
	_, ok := ref.ObjDiff(oldcfg.Spec, newcfg.Spec)
	if (oldcfg.GenerationID == newcfg.GenerationID) && !ok {
		oldcfg.ObjectMeta = newcfg.ObjectMeta
		return nil
	}

	// update old state
	oldcfg.ObjectMeta = newcfg.ObjectMeta
	oldcfg.Spec = newcfg.Spec

	// find the routingconfig state
	rtcfg, err := RoutingConfigStateFromObj(oldcfg)
	if err != nil {
		log.Errorf("Can't find an routingconfig state for updating {%+v}. Err: {%v}", oldcfg.ObjectMeta, err)
		return fmt.Errorf("Can not find routingconfig state")
	}

	err = sma.sm.mbus.UpdateObjectWithReferences(newcfg.MakeKey(string(apiclient.GroupNetwork)), convertRoutingConfig(rtcfg), references(newcfg))
	if err != nil {
		log.Errorf("could not update RoutingConfig to nimbus (%s)", err)
		return err
	}
	log.Info("OnRoutingConfigUpdate, found: ", rtcfg.RoutingConfig)

	return nil
}

// OnRoutingConfigDelete deletes the routingcfg
func (sma *SmRoute) OnRoutingConfigDelete(obj *ctkit.RoutingConfig) error {
	log.Info("OnRoutingConfigDelete: received: ", obj.Spec)

	rtcfg, err := sma.FindRoutingConfig(obj.Tenant, obj.Namespace, obj.Name)

	if err != nil {
		log.Error("FindRoutingConfig returned an error: ", err, "for: ", obj.Tenant, obj.Namespace, obj.Name)
		return errors.New("Object doesn't exist")
	}

	log.Info("OnRoutingConfigDelete, found: ", rtcfg.RoutingConfig)
	err = sma.sm.mbus.DeleteObjectWithReferences(rtcfg.RoutingConfig.MakeKey(string(apiclient.GroupNetwork)), convertRoutingConfig(rtcfg), references(obj))
	if err != nil {
		log.Errorf("could not delete RoutingConfig to nimbus (%s)", err)
		return err
	}
	return nil
}

// OnRoutingConfigReconnect is called when ctkit reconnects to apiserver
func (sma *SmRoute) OnRoutingConfigReconnect() {
	return
}
