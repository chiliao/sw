package vchub

import (
	"github.com/pensando/sw/api/generated/network"
	"github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/defs"
	"github.com/pensando/sw/venice/utils/kvstore"
)

func (v *VCHub) startEventsListener() {
	v.Log.Infof("Running store")
	v.Log.Infof("Starting probe channel watch for %s", v.OrchConfig.Name)
	defer v.Wg.Done()
	apiServerCh, err := v.StateMgr.GetProbeChannel(v.OrchConfig.GetName())
	if err != nil {
		v.Log.Errorf("Could not get probe channel for %s. Err : %v", v.OrchConfig.GetKey(), err)
		return
	}

	for {
		select {
		case <-v.Ctx.Done():
			return

		case m, active := <-v.vcReadCh:
			if !active {
				return
			}

			switch m.MsgType {
			case defs.VCEvent:
				v.handleVCEvent(m.Val.(defs.VCEventMsg))
			default:
				v.Log.Errorf("Unknown event %s", m.MsgType)
			}
		case evt, ok := <-apiServerCh:
			if ok {
				nw := evt.Object.(*network.Network)
				v.handleNetworkEvent(evt.Type, nw)
			}
		}
	}
}

func (v *VCHub) handleVCEvent(m defs.VCEventMsg) {
	v.Log.Infof("Msg from %v, key: %s prop: %s", m.Originator, m.Key, m.VcObject)
	switch m.VcObject {
	case defs.VirtualMachine:
		v.handleWorkload(m)
	case defs.HostSystem:
		v.handleHost(m)
	case defs.Datacenter:
		v.handleDC(m)
	default:
		v.Log.Errorf("Unknown object %s", m.VcObject)
	}
}

func (v *VCHub) handleDC(m defs.VCEventMsg) {
	// Check if we have a DC object
	v.Log.Infof("Handle DC called")
	for _, prop := range m.Changes {
		name := prop.Val.(string)
		v.Log.Infof("Handle DC %s", name)
		v.DcMapLock.Lock()
		if penDc, ok := v.DcMap[m.Key]; ok {
			penDc.Lock()
			if _, ok := penDc.DvsMap[defs.DefaultDVSName]; ok {
				v.DcMapLock.Unlock()
				penDc.Unlock()
				return
			}
			penDc.Unlock()
		}
		v.DcMapLock.Unlock()
		// We create DVS and check networks
		if v.forceDCname != "" && name != v.forceDCname {
			v.Log.Infof("Skipping DC event for DC %s", name)
			continue
		}
		v.Log.Infof("new DC %s", name)
		v.NewPenDC(name)
	}
}

func (v *VCHub) handleNetworkEvent(evtType kvstore.WatchEventType, nw *network.Network) {
	v.Log.Infof("Handling network event. %v", nw)

	switch evtType {
	case kvstore.Created:
		if len(nw.Spec.Orchestrators) == 0 {
			return
		}
		dcs := []string{}
		for _, orch := range nw.Spec.Orchestrators {
			if orch.Name == v.OrchConfig.GetName() {
				dcs = append(dcs, orch.Namespace)
			}
		}
		v.Log.Infof("Create network %s event for dcs %v", nw.Name, dcs)
		for _, dc := range dcs {
			v.DcMapLock.Lock()
			penDC, ok := v.DcMap[dc]
			v.DcMapLock.Unlock()
			if !ok {
				continue
			}
			pgName := createPGName(nw.Name)
			penDC.AddPG(pgName, nw.ObjectMeta, "")
		}
	case kvstore.Updated:
		// If wire vlan changes, workloads should be modified
		v.Log.Info("Update network event")
	case kvstore.Deleted:
		if len(nw.Spec.Orchestrators) == 0 {
			return
		}
		dcs := []string{}
		for _, orch := range nw.Spec.Orchestrators {
			if orch.Name == v.OrchConfig.GetName() {
				dcs = append(dcs, orch.Namespace)
			}
		}
		v.Log.Info("Delete network %s event for dcs %v", nw.Name, dcs)
		for _, dc := range dcs {
			v.DcMapLock.Lock()
			penDC, ok := v.DcMap[dc]
			v.DcMapLock.Unlock()
			if !ok {
				continue
			}
			pgName := createPGName(nw.Name)
			penDC.RemovePG(pgName, "")
		}
	}
}