package vchub

import (
	"fmt"
	"sync"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/events/generated/eventtypes"
	"github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/defs"
	"github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/vcprobe"
	"github.com/pensando/sw/venice/utils/events/recorder"
)

// PenPG represents an instance of a port group on a DVS
type PenPG struct {
	*defs.State
	DcName      string
	probe       vcprobe.ProbeInf
	PgName      string
	PgRef       types.ManagedObjectReference
	NetworkMeta api.ObjectMeta
	PgMutex     sync.Mutex
}

// AddPenPG creates a PG
func (d *PenDVS) AddPenPG(pgName string, networkMeta api.ObjectMeta) error {
	primaryVlan, secondaryVlan, err := d.UsegMgr.GetVlansForPG(pgName)
	if err != nil {
		{
			primaryVlan, secondaryVlan, err = d.UsegMgr.AssignVlansForPG(pgName)
			if err != nil {
				d.Log.Errorf("Failed to assign vlans for PG")
				if err.Error() == "PG limit reached" {
					// Generate error
					evtMsg := fmt.Sprintf("Failed to create Network %s in Datacenter %s. Max network limit reached.", networkMeta.Name, d.DcName)
					d.Log.Errorf(evtMsg)

					recorder.Event(eventtypes.ORCH_CONFIG_PUSH_FAILURE, evtMsg, d.State.OrchConfig)
				}
				return err
			}
		}
	}
	d.Log.Debugf("Adding PG %s with pvlan of %d and %d", pgName, primaryVlan, secondaryVlan)
	err = d.AddPenPGWithVlan(pgName, networkMeta, primaryVlan, secondaryVlan)
	if err != nil {
		evtMsg := fmt.Sprintf("Failed to set configuration for network %s in Datacenter %s", networkMeta.Name, d.DcName)
		d.Log.Errorf("%s, err %s", evtMsg, err)
		recorder.Event(eventtypes.ORCH_CONFIG_PUSH_FAILURE, evtMsg, d.State.OrchConfig)
	}
	return err
}

func (d *PenDVS) createPGConfigCheck(pvlan int) vcprobe.IsPGConfigEqual {
	return func(spec *types.DVPortgroupConfigSpec, config *mo.DistributedVirtualPortgroup) bool {

		policy, ok := config.Config.Policy.(*types.VMwareDVSPortgroupPolicy)
		if !ok {
			d.Log.Infof("ConfigCheck: dvs.Config was of type %T", config.Config.Policy)
			return false
		}
		if !policy.VlanOverrideAllowed ||
			!policy.PortConfigResetAtDisconnect {
			d.Log.Infof("ConfigCheck: dvs policy settings were incorrect")
			return false
		}
		portConfig, ok := config.Config.DefaultPortConfig.(*types.VMwareDVSPortSetting)
		if !ok {
			d.Log.Infof("ConfigCheck: portConfig was of type %T", config.Config.DefaultPortConfig)
			return false
		}
		pvlanConfig, ok := portConfig.Vlan.(*types.VmwareDistributedVirtualSwitchPvlanSpec)
		if !ok {
			d.Log.Infof("ConfigCheck: pvlanConfig was of type %T", portConfig.Vlan)
			return false
		}
		if pvlanConfig.PvlanId != int32(pvlan) {
			d.Log.Infof("ConfigCheck: pvlan did not match")
			return false
		}
		return true
	}
}

// AddPenPGWithVlan creates a PG with the given pvlan values
func (d *PenDVS) AddPenPGWithVlan(pgName string, networkMeta api.ObjectMeta, primaryVlan, secondaryVlan int) error {
	d.Lock()
	defer d.Unlock()

	spec := types.DVPortgroupConfigSpec{
		Name: pgName,
		Type: string(types.DistributedVirtualPortgroupPortgroupTypeEarlyBinding),
		Policy: &types.VMwareDVSPortgroupPolicy{
			DVPortgroupPolicy: types.DVPortgroupPolicy{
				BlockOverrideAllowed:               true,
				ShapingOverrideAllowed:             false,
				VendorConfigOverrideAllowed:        false,
				LivePortMovingAllowed:              false,
				PortConfigResetAtDisconnect:        true,
				NetworkResourcePoolOverrideAllowed: types.NewBool(false),
				TrafficFilterOverrideAllowed:       types.NewBool(false),
			},
			VlanOverrideAllowed:           true,
			UplinkTeamingOverrideAllowed:  false,
			SecurityPolicyOverrideAllowed: false,
			IpfixOverrideAllowed:          types.NewBool(false),
		},
		DefaultPortConfig: &types.VMwareDVSPortSetting{
			Vlan: &types.VmwareDistributedVirtualSwitchPvlanSpec{
				PvlanId: int32(secondaryVlan),
			},
		},
	}
	moPG, err := d.probe.GetPGConfig(d.DcName, pgName, []string{"config"}, 1)

	if err == nil {
		policySpec, ok := moPG.Config.Policy.(*types.VMwareDVSPortgroupPolicy)
		if ok {
			policySpec.VlanOverrideAllowed = true
			policySpec.PortConfigResetAtDisconnect = true
			spec.Policy = policySpec
		}
	}

	err = d.probe.AddPenPG(d.DcName, d.DvsName, &spec, d.createPGConfigCheck(secondaryVlan), defaultRetryCount)
	if err != nil {
		return err
	}

	pg, err := d.probe.GetPenPG(d.DcName, pgName, defaultRetryCount)
	if err != nil {
		return err
	}

	penPG := d.getPenPG(pgName)
	if penPG == nil {
		penPG = &PenPG{
			State:  d.State,
			probe:  d.probe,
			DcName: d.DcName,
			PgName: pgName,
		}
		d.Pgs[pgName] = penPG
	}
	penPG.PgRef = pg.Reference()
	d.pgIDMap[pg.Reference().Value] = penPG

	penPG.NetworkMeta = networkMeta

	err = d.probe.TagObjAsManaged(penPG.PgRef)
	if err != nil {
		d.Log.Errorf("Failed to tag PG %s as managed, %s", pgName, err)
		// Error isn't worth failing the operation for
	}

	nw, err := d.StateMgr.Controller().Network().Find(&networkMeta)
	if err == nil {
		externalVlan := int(nw.Spec.VlanID)
		err = d.probe.TagObjWithVlan(penPG.PgRef, externalVlan)
		if err != nil {
			d.Log.Errorf("Failed to tag PG %s with vlan tag, %s", pgName, err)
			// Error isn't worth failing the operation for
		}
	} else {
		d.Log.Errorf("Couldn't tag PG %s with vlan tag since we couldn't find the network info: networkMeta %v, err %s", pgName, networkMeta, err)
		// Error isn't worth failing the operation for
	}

	return nil
}

// GetPenPG returns the PenPG with the given name
func (d *PenDVS) GetPenPG(pgName string) *PenPG {
	d.Lock()
	defer d.Unlock()
	return d.getPenPG(pgName)
}

func (d *PenDVS) getPenPG(pgName string) *PenPG {
	pg, ok := d.Pgs[pgName]
	if !ok {
		return nil
	}
	return pg
}

// GetPenPGByID fetches the pen PG object by ID
func (d *PenDVS) GetPenPGByID(pgID string) *PenPG {
	d.Lock()
	defer d.Unlock()
	pg, ok := d.pgIDMap[pgID]
	if !ok {
		return nil
	}
	return pg
}

// RemovePenPG removes the pg from the dvs
func (d *PenDVS) RemovePenPG(pgName string) error {
	d.Lock()
	defer d.Unlock()

	var ref types.ManagedObjectReference
	if penPG, ok := d.Pgs[pgName]; ok {
		ref = penPG.PgRef
		id := penPG.PgRef.Value
		delete(d.Pgs, pgName)

		if _, ok := d.pgIDMap[id]; ok {
			delete(d.pgIDMap, id)
		} else {
			d.Log.Errorf("Removed entry in PG map that wasn't in pgIDMap, pgName %s", pgName)
		}
	}

	err := d.probe.RemovePenPG(d.DcName, pgName, defaultRetryCount)
	if err != nil {
		d.Log.Errorf("Failed to delete PG %s, removing management tag. Err %s", pgName, err)
		tagErrs := d.probe.RemovePensandoTags(ref)
		if len(tagErrs) != 0 {
			d.Log.Errorf("Failed to remove tags, errs %v", tagErrs)
		}
		return err
	}

	return nil

}
