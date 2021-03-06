package manager

import (
	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/venice/utils/authz"
	"github.com/pensando/sw/venice/utils/authz/orb"
	"github.com/pensando/sw/venice/utils/resolver"
)

// AuthorizationManager authorizes user and returns authorization information
type authorizationManager struct {
	authz.AbstractAuthorizer
	authorizers []authz.Authorizer
}

// NewAuthorizationManager returns an instance of AuthorizationManager
func NewAuthorizationManager(name, apiServer string, rslver resolver.Interface) authz.Authorizer {
	authorizers := make([]authz.Authorizer, 1)
	authorizers[0] = orb.NewORBAuthorizer(name, apiServer, rslver)
	authzMgr := &authorizationManager{
		authorizers: authorizers,
	}
	authzMgr.AbstractAuthorizer.Authorizer = authzMgr
	return authzMgr
}

// IsAuthorized checks if user is authorized for the given operations. If multiple authorizers are configured and enabled, it will execute them in the order specified.
// If any authorizer fails, access is denied and remaining authorizers are not tried.
// Returns
//   true if all authorizers succeed
func (authzmgr *authorizationManager) IsAuthorized(user *auth.User, operations ...authz.Operation) (bool, error) {
	for _, authorizer := range authzmgr.authorizers {
		ok, err := authorizer.IsAuthorized(user, operations...)
		if !ok {
			// if any authorizer fails to authorize return false
			return false, err
		}
	}
	return true, nil
}

// Stop stops the underlying authorizers
func (authzmgr *authorizationManager) Stop() {
	for _, authorizer := range authzmgr.authorizers {
		authorizer.Stop()
	}
}
