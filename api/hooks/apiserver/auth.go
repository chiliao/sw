package impl

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8serrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/api/login"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/authn"
	"github.com/pensando/sw/venice/utils/authn/password"
	authzgrpc "github.com/pensando/sw/venice/utils/authz/grpc"
	authzgrpcctx "github.com/pensando/sw/venice/utils/authz/grpc/context"
	"github.com/pensando/sw/venice/utils/ctxutils"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"
)

const (
	minTokenExpiry = 2 * time.Minute // token expiry cannot be less than 2 minutes
)

var (
	// errInvalidInputType is returned when incorrect object type is passed to the hook
	errInvalidInputType = errors.New("invalid input type")
	// errInvalidRolePermissions is returned when tenant in permission's resource does not match tenant of the Role
	errInvalidRolePermissions = errors.New("invalid tenant in role permission")
	// errExtUserPasswordChange is returned when change or reset password hook is called for external user
	errExtUserPasswordChange = errors.New("cannot change or reset password of external user")
	// errInvalidOldPassword is returned when invalid old password is provided for changing password
	errInvalidOldPassword = errors.New("invalid old password")
	// errEmptyPassword is returned when password is specifying as empty while creating user or changing password for user
	errEmptyPassword = errors.New("password is empty")
	// errCreatingPasswordHash is returned when there is error generating password hash
	errCreatingPasswordHash = errors.New("error creating password hash")
	// errAdminRoleUpdateNotAllowed is returned when AdminRole for a tenant is being updated
	errAdminRoleUpdateNotAllowed = fmt.Errorf("%s create, update or delete is not allowed", globals.AdminRole)
	// errAdminRoleBindingDeleteNotAllowed is returned when AdminRoleBinding for a tenant is being deleted
	errAdminRoleBindingDeleteNotAllowed = fmt.Errorf("%s delete is not allowed", globals.AdminRoleBinding)
	// errAdminRoleBindingRoleUpdateNotAllowed is returned when AdminRoleBinding role is updated to something other than AdminRole
	errAdminRoleBindingRoleUpdateNotAllowed = fmt.Errorf("%s can bind to only %s", globals.AdminRoleBinding, globals.AdminRole)
)

type authHooks struct {
	logger log.Logger
}

// hashPassword is pre-commit hook to save hashed password in User object when object is created or updated
func (s *authHooks) hashPassword(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to hash password")
	r, ok := i.(auth.User)
	if !ok {
		return i, false, errInvalidInputType
	}

	// don't save password for external user
	if r.Spec.GetType() == auth.UserSpec_External.String() {
		r.Spec.Password = ""
		return r, true, nil
	}

	// if user type is not external then assume local by default
	r.Spec.Type = auth.UserSpec_Local.String()

	switch oper {
	case apiintf.CreateOper:
		// password is a required field when local user is created
		if r.Spec.GetPassword() == "" {
			return r, false, errEmptyPassword
		}
		// validate password
		pc := password.NewPolicyChecker()
		if err := k8serrors.NewAggregate(pc.Validate(r.Spec.GetPassword())); err != nil {
			s.logger.ErrorLog("method", "hashPassword", "msg", "password validation failed", "error", err)
			return r, false, err
		}
		// get password hasher
		hasher := password.GetPasswordHasher()
		// generate hash
		passwdhash, err := hasher.GetPasswordHash(r.Spec.GetPassword())
		if err != nil {
			s.logger.Errorf("error creating password hash: %v", err)
			return r, false, errCreatingPasswordHash
		}
		r.Spec.Password = passwdhash
		s.logger.InfoLog("msg", "Created password hash", "user", r.GetName())
	case apiintf.UpdateOper:
		cur := &auth.User{}
		if err := kv.Get(ctx, key, cur); err != nil {
			s.logger.Errorf("error getting user with key [%s] in API server hashPassword pre-commit hook for update cluster", key)
			return r, false, err
		}
		// decrypt password as it is stored as secret. Cannot use passed in context because peer id in it is APIGw and transform returns empty password in that case
		if err := cur.ApplyStorageTransformer(context.Background(), false); err != nil {
			s.logger.Errorf("error decrypting password field: %v", err)
			return r, false, err
		}
		r.Spec.Password = cur.Spec.Password
		// add a comparator for CAS
		if !dryRun {
			s.logger.Infof("set the comparator version for [%s] as [%s]", key, cur.ResourceVersion)
			txn.AddComparator(kvstore.Compare(kvstore.WithVersion(key), "=", cur.ResourceVersion))
		}
	default:
	}
	return r, true, nil
}

// changePassword is pre-commit hook registered with PasswordChange method for User service to change password for local user
func (s *authHooks) changePassword(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to change password")
	r, ok := i.(auth.PasswordChangeRequest)
	if !ok {
		return nil, false, errInvalidInputType
	}
	if r.NewPassword == "" || r.OldPassword == "" {
		return nil, false, errEmptyPassword
	}
	cur := &auth.User{}
	if err := kv.ConsistentUpdate(ctx, key, cur, func(oldObj runtime.Object) (runtime.Object, error) {
		userObj, ok := oldObj.(*auth.User)
		if !ok {
			return oldObj, errInvalidInputType
		}
		// error if external user
		if userObj.Spec.GetType() == auth.UserSpec_External.String() {
			return userObj, errExtUserPasswordChange
		}
		// decrypt password as it is stored as secret. Cannot use passed in context because peer id in it is APIGw and transform returns empty password in that case
		if err := userObj.ApplyStorageTransformer(context.Background(), false); err != nil {
			s.logger.Errorf("error decrypting password field: %v", err)
			return userObj, err
		}
		hasher := password.GetPasswordHasher()
		ok, err := hasher.CompareHashAndPassword(userObj.Spec.Password, r.OldPassword)
		if err != nil {
			s.logger.Errorf("error comparing password hash: %v", err)
			return userObj, errInvalidOldPassword
		}
		if !ok {
			return userObj, errInvalidOldPassword
		}
		passwdhash, err := hasher.GetPasswordHash(r.NewPassword)
		if err != nil {
			s.logger.Errorf("error creating password hash: %v", err)
			return userObj, errCreatingPasswordHash
		}
		userObj.Spec.Password = passwdhash
		// encrypt password as it is stored as secret
		if err := userObj.ApplyStorageTransformer(ctx, true); err != nil {
			s.logger.Errorf("error encrypting password field: %v", err)
			return userObj, err
		}
		// update last password change time
		m, err := types.TimestampProto(time.Now())
		if err != nil {
			return userObj, err
		}
		userObj.Status.LastPasswordChange = &api.Timestamp{
			Timestamp: *m,
		}
		genID, err := strconv.ParseInt(userObj.GenerationID, 10, 64)
		if err != nil {
			s.logger.Errorf("error parsing generation ID: %v", err)
			genID = 2
		}
		userObj.GenerationID = fmt.Sprintf("%d", genID+1)
		return userObj, nil
	}); err != nil {
		s.logger.Errorf("error changing password for user key [%s]: %v", key, err)
		return nil, false, err
	}
	// The ConsistentUpdate op will happen at a later time due to overlay. Retrieve the current object and return so the ResponseWriter has a object to work on.
	ret := auth.User{}
	ret.ObjectMeta = r.ObjectMeta
	ret.Spec.Password = ""
	s.logger.Debugf("password successfully changed for user key [%s]", key)
	return ret, false, nil
}

// resetPassword is pre-commit hook registered with PasswordReset method for User service to reset password for local user
func (s *authHooks) resetPassword(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to reset password")
	in, ok := i.(auth.PasswordResetRequest)
	if !ok {
		return nil, false, errInvalidInputType
	}
	genPasswd, err := password.CreatePassword()
	if err != nil {
		s.logger.ErrorLog("method", "resetPassword", "msg", "error generating password", "user", key, "error", err)
		return nil, false, err
	}
	cur := &auth.User{}
	if err := kv.ConsistentUpdate(ctx, key, cur, func(oldObj runtime.Object) (runtime.Object, error) {
		userObj, ok := oldObj.(*auth.User)
		if !ok {
			return oldObj, errInvalidInputType
		}
		// error if external user
		if userObj.Spec.GetType() == auth.UserSpec_External.String() {
			return userObj, errExtUserPasswordChange
		}
		hasher := password.GetPasswordHasher()

		passwdhash, err := hasher.GetPasswordHash(genPasswd)
		if err != nil {
			s.logger.Errorf("error creating password hash: %v", err)
			return userObj, errCreatingPasswordHash
		}
		userObj.Spec.Password = passwdhash
		// encrypt password as it is stored as secret
		if err := userObj.ApplyStorageTransformer(ctx, true); err != nil {
			s.logger.Errorf("error encrypting password field: %v", err)
			return userObj, err
		}
		// update last password change time
		m, err := types.TimestampProto(time.Now())
		if err != nil {
			return userObj, err
		}
		userObj.Status.LastPasswordChange = &api.Timestamp{
			Timestamp: *m,
		}
		genID, err := strconv.ParseInt(userObj.GenerationID, 10, 64)
		if err != nil {
			s.logger.Errorf("error parsing generation ID: %v", err)
			genID = 2
		}
		userObj.GenerationID = fmt.Sprintf("%d", genID+1)
		return userObj, nil
	}); err != nil {
		s.logger.Errorf("error resetting password for user key [%s]: %v", key, err)
		return nil, false, err
	}
	// The ConsistentUpdate op will happen at a later time due to overlay. Retrieve the current object and return so the ResponseWriter has a object to work on.
	ret := auth.User{}
	ret.ObjectMeta = in.ObjectMeta
	ret.Spec.Password = genPasswd
	s.logger.Debugf("password successfully reset for user key [%s]", key)
	return ret, false, nil
}

// validateAuthenticatorConfig hook is to validate that authenticators specified in AuthenticatorOrder are defined
func (s *authHooks) validateAuthenticatorConfig(i interface{}, ver string, ignStatus bool) []error {
	s.logger.DebugLog("msg", "AuthHook called to validate authenticator config")
	var ret []error
	r, ok := i.(auth.AuthenticationPolicy)
	if !ok {
		return []error{errInvalidInputType}
	}
	// TokenExpiry has already been validated by Duration() validator
	exp, _ := time.ParseDuration(r.Spec.TokenExpiry)
	// token expiry cannot be less than 2 minutes
	if exp < minTokenExpiry {
		ret = append(ret, fmt.Errorf("token expiry (%s) should be atleast 2 minutes", r.Spec.TokenExpiry))
	}
	// check if authenticators specified in AuthenticatorOrder are defined
	authenticatorOrder := r.Spec.Authenticators.GetAuthenticatorOrder()
	if authenticatorOrder == nil || len(authenticatorOrder) == 0 {
		ret = append(ret, errors.New("authenticator order config not defined"))
		return ret
	}
	for _, authenticatorType := range authenticatorOrder {
		switch authenticatorType {
		case auth.Authenticators_LOCAL.String():
			if r.Spec.Authenticators.GetLocal() == nil {
				ret = append(ret, errors.New("local authenticator config not defined"))
			}
		case auth.Authenticators_LDAP.String():
			if r.Spec.Authenticators.GetLdap() == nil {
				ret = append(ret, errors.New("ldap authenticator config not defined"))
			}
		case auth.Authenticators_RADIUS.String():
			if r.Spec.Authenticators.GetRadius() == nil {
				ret = append(ret, errors.New("radius authenticator config not defined"))
			}
		default:
			ret = append(ret, fmt.Errorf("unknown authenticator type: %v", authenticatorType))
		}
	}
	return ret
}

// generateSecret is a pre-commmit hook to generate secret when authentication policy is created or updated
func (s *authHooks) generateSecret(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to generate JWT secret")
	r, ok := i.(auth.AuthenticationPolicy)
	if !ok {
		return i, false, errInvalidInputType
	}
	secret, err := authn.CreateSecret(128)
	if err != nil {
		s.logger.Errorf("Error generating secret, Err: %v", err)
		return r, false, err
	}
	r.Spec.Secret = secret
	s.logger.InfoLog("msg", "Generated JWT Secret")
	return r, true, nil
}

// validateRolePerms is hook to validate that resource kind and group is valid in permission and a role in non default tenant doesn't contain permissions to other tenants
func (s *authHooks) validateRolePerms(i interface{}, ver string, ignStatus bool) []error {
	s.logger.DebugLog("msg", "AuthHook called to validate role")
	r, ok := i.(auth.Role)
	if !ok {
		return []error{errInvalidInputType}
	}

	var errs []error
	for _, perm := range r.Spec.Permissions {
		// "default" tenant role can have permissions for objects in other tenants
		if r.Tenant != globals.DefaultTenant && perm.GetResourceTenant() != r.Tenant {
			errs = append(errs, errInvalidRolePermissions)
		}
		// validate resource kind and group
		if err := login.ValidatePerm(perm); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		s.logger.Errorf("validation failed for role [%#v], %s", r, errs)
	}
	return errs
}

// privilegeEscalationCheck is a pre-commit hook to check if user is trying to escalate his privileges while creating or updating role binding
func (s *authHooks) privilegeEscalationCheck(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to check privilege escalation")
	r, ok := i.(auth.RoleBinding)
	if !ok {
		return i, false, errInvalidInputType
	}
	cluster := &cluster.Cluster{}
	if err := kv.Get(ctx, cluster.MakeKey("cluster"), cluster); err != nil {
		s.logger.ErrorLog("method", "privilegeEscalationCheck",
			"msg", "error getting cluster with key",
			"error", err)
		return i, false, err
	}
	// check authorization only if request is coming from API Gateway and auth has been bootstrapped
	if ctxutils.GetPeerID(ctx) == globals.APIGw && cluster.Status.AuthBootstrapped {
		role := &auth.Role{ObjectMeta: api.ObjectMeta{Name: r.Spec.Role, Tenant: r.Tenant}}
		roleKey := role.MakeKey("auth")
		if err := kv.Get(ctx, roleKey, role); err != nil {
			s.logger.ErrorLog("method", "privilegeEscalationCheck",
				"msg", fmt.Sprintf("error getting role with key [%s]", roleKey),
				"error", err)
			return i, false, err
		}
		userMeta, ok := authzgrpcctx.UserMetaFromIncomingContext(ctx)
		if !ok {
			s.logger.ErrorLog("method", "privilegeEscalationCheck", "msg", "no user in grpc metadata")
			return i, false, status.Errorf(codes.Internal, "no user in context")
		}
		user := &auth.User{ObjectMeta: *userMeta}
		// check if user is authorized to create the role binding
		authorizer, err := authzgrpc.NewAuthorizer(ctx)
		if err != nil {
			s.logger.ErrorLog("method", "privilegeEscalationCheck", "msg", "error creating grpc authorizer", "error", err)
			return i, false, status.Error(codes.Internal, err.Error())
		}
		ops := login.GetOperationsFromPermissions(role.Spec.Permissions)
		ok, _ = authorizer.IsAuthorized(user, ops...)
		if !ok {
			return i, false, status.Error(codes.PermissionDenied, fmt.Sprintf("unauthorized to create role binding (%s|%s)", r.GetTenant(), r.GetName()))
		}
		s.logger.InfoLog("method", "privilegeEscalationCheck", "msg", "success")
	}
	return i, true, nil
}

func (s *authHooks) returnUser(ctx context.Context, kvs kvstore.Interface, prefix string, in, old, resp interface{}, oper apiintf.APIOperType) (interface{}, error) {
	ic := resp.(auth.User)
	s.logger.Infof("Got user [%+v]", ic)
	key := ic.MakeKey("auth")
	cur := auth.User{}
	if err := kvs.Get(ctx, key, &cur); err != nil {
		s.logger.Errorf("Error getting user with key [%s] in API server response writer hook", key)
		return nil, err
	}
	err := cur.ApplyStorageTransformer(ctx, false)
	cur.Spec.Password = ic.Spec.Password
	return cur, err
}

// validatePassword is a hook to check user password against password policy
func (s *authHooks) validatePassword(i interface{}, ver string, ignStatus bool) []error {
	s.logger.DebugLog("msg", "AuthHook called to validate password")
	pc := password.NewPolicyChecker()
	switch obj := i.(type) {
	case auth.User:
		if obj.Spec.Type == auth.UserSpec_Local.String() {
			return pc.Validate(obj.Spec.Password)
		}
	case auth.PasswordChangeRequest:
		return pc.Validate(obj.NewPassword)
	default:
		return []error{errInvalidInputType}
	}
	return nil
}

// adminRoleCheck is a pre-commit hook to prevent create/update/delete of admin role
func (s *authHooks) adminRoleCheck(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, in interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to prevent admin role create, update, delete")
	obj, ok := in.(auth.Role)
	if !ok {
		return in, false, errInvalidInputType
	}
	if obj.Name == globals.AdminRole {
		return in, false, errAdminRoleUpdateNotAllowed
	}
	return in, true, nil
}

// adminRoleBindingCheck is a pre-commit hook to prevent deletion of admin role binding and update of referred AdminRole
func (s *authHooks) adminRoleBindingCheck(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiintf.APIOperType, dryRun bool, in interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to prevent admin role binding delete")
	obj, ok := in.(auth.RoleBinding)
	if !ok {
		return in, false, errInvalidInputType
	}
	if obj.Name == globals.AdminRoleBinding {
		switch oper {
		case apiintf.UpdateOper:
			if obj.Spec.Role != globals.AdminRole {
				return in, false, errAdminRoleBindingRoleUpdateNotAllowed
			}
		case apiintf.DeleteOper:
			return in, false, errAdminRoleBindingDeleteNotAllowed
		}
	}
	return in, true, nil
}

func registerAuthHooks(svc apiserver.Service, logger log.Logger) {
	r := authHooks{}
	r.logger = logger.WithContext("Service", "AuthHooks")
	logger.Log("msg", "registering Hooks")
	svc.GetCrudService("User", apiintf.CreateOper).WithPreCommitHook(r.hashPassword)
	svc.GetCrudService("User", apiintf.UpdateOper).WithPreCommitHook(r.hashPassword)
	svc.GetCrudService("AuthenticationPolicy", apiintf.CreateOper).WithPreCommitHook(r.generateSecret).GetRequestType().WithValidate(r.validateAuthenticatorConfig)
	svc.GetCrudService("AuthenticationPolicy", apiintf.UpdateOper).WithPreCommitHook(r.generateSecret).GetRequestType().WithValidate(r.validateAuthenticatorConfig)
	svc.GetCrudService("Role", apiintf.CreateOper).WithPreCommitHook(r.adminRoleCheck).GetRequestType().WithValidate(r.validateRolePerms)
	svc.GetCrudService("Role", apiintf.UpdateOper).WithPreCommitHook(r.adminRoleCheck).GetRequestType().WithValidate(r.validateRolePerms)
	svc.GetCrudService("Role", apiintf.DeleteOper).WithPreCommitHook(r.adminRoleCheck)
	svc.GetCrudService("RoleBinding", apiintf.CreateOper).WithPreCommitHook(r.privilegeEscalationCheck)
	svc.GetCrudService("RoleBinding", apiintf.UpdateOper).WithPreCommitHook(r.adminRoleBindingCheck).WithPreCommitHook(r.privilegeEscalationCheck)
	svc.GetCrudService("RoleBinding", apiintf.DeleteOper).WithPreCommitHook(r.adminRoleBindingCheck)
	// hook to change password
	svc.GetMethod("PasswordChange").WithPreCommitHook(r.changePassword).GetRequestType().WithValidate(r.validatePassword)
	svc.GetMethod("PasswordChange").WithResponseWriter(r.returnUser)
	// hook to reset password
	svc.GetMethod("PasswordReset").WithPreCommitHook(r.resetPassword)
	svc.GetMethod("PasswordReset").WithResponseWriter(r.returnUser)
}

func init() {
	apisrv := apisrvpkg.MustGetAPIServer()
	apisrv.RegisterHooksCb("auth.AuthV1", registerAuthHooks)
}
