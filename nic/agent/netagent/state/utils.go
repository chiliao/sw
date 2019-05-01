package state

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/globals"
)

// ValidateVrf finds the correspnding vrf
func (na *Nagent) ValidateVrf(tenant, namespace, name string) (*netproto.Vrf, error) {
	// Pick default vrf is unspecified or specified default
	if len(name) == 0 || name == globals.DefaultVrf {
		tenant = globals.DefaultTenant
		namespace = globals.DefaultNamespace
		name = globals.DefaultVrf
	}
	meta := api.ObjectMeta{
		Tenant:    tenant,
		Namespace: namespace,
		Name:      name,
	}
	return na.FindVrf(meta)
}
