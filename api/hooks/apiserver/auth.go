package impl

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/api/login"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/authn"
	"github.com/pensando/sw/venice/utils/authn/password"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"
)

var (
	// errInvalidInputType is returned when incorrect object type is passed to the hook
	errInvalidInputType = errors.New("invalid input type")
	// errAuthenticatorConfig is returned when authenticator config is incorrect in AuthenticationPolicy.
	errAuthenticatorConfig = errors.New("mis-configured authentication policy, error in authenticator config")
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
)

type authHooks struct {
	logger log.Logger
}

// hashPassword is pre-commit hook to save hashed password in User object when object is created or updated
func (s *authHooks) hashPassword(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiserver.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
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
	case apiserver.CreateOper:
		// password is a required field when local user is created
		if r.Spec.GetPassword() == "" {
			return r, false, errEmptyPassword
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
	case apiserver.UpdateOper:
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
		s.logger.Infof("set the comparator version for [%s] as [%s]", key, cur.ResourceVersion)
		txn.AddComparator(kvstore.Compare(kvstore.WithVersion(key), "=", cur.ResourceVersion))
	default:
	}
	return r, true, nil
}

// changePassword is pre-commit hook registered with PasswordChange method for User service to change password for local user
func (s *authHooks) changePassword(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiserver.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
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
			return oldObj, errors.New("invalid input type")
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
		genID, err := strconv.ParseInt(userObj.GenerationID, 10, 64)
		if err != nil {
			s.logger.Errorf("error parsing generation ID: %v", err)
			return userObj, err
		}
		userObj.GenerationID = fmt.Sprintf("%d", genID+1)
		return userObj, nil
	}); err != nil {
		s.logger.Errorf("error changing password for user key [%s]: %v", key, err)
		return nil, false, err
	}
	// empty out password before returning. Create a copy as cur is pointing to an object in API server cache. Without copy causes intermittent failures in password change integ test
	// where password is empty in user object returned from API server on subsequent GET
	ret, err := cur.Clone(nil)
	if err != nil {
		s.logger.Errorf("error creating a copy of user obj: %v", err)
		return ret, false, err
	}
	user := ret.(*auth.User)
	user.Spec.Password = ""
	s.logger.Debugf("password successfully changed for user key [%s]", key)
	return *user, false, nil
}

// resetPassword is pre-commit hook registered with PasswordReset method for User service to reset password for local user
func (s *authHooks) resetPassword(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiserver.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
	s.logger.DebugLog("msg", "AuthHook called to reset password")
	_, ok := i.(auth.PasswordResetRequest)
	if !ok {
		return nil, false, errInvalidInputType
	}
	b, err := authn.CreateSecret(12)
	if err != nil {
		s.logger.Errorf("Error generating password for user key [%s]: %v", err)
		return nil, false, err
	}
	genPasswd := base64.RawStdEncoding.EncodeToString(b)
	cur := &auth.User{}
	if err := kv.ConsistentUpdate(ctx, key, cur, func(oldObj runtime.Object) (runtime.Object, error) {
		userObj, ok := oldObj.(*auth.User)
		if !ok {
			return oldObj, errors.New("invalid input type")
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
		genID, err := strconv.ParseInt(userObj.GenerationID, 10, 64)
		if err != nil {
			s.logger.Errorf("error parsing generation ID: %v", err)
			return userObj, err
		}
		userObj.GenerationID = fmt.Sprintf("%d", genID+1)
		return userObj, nil
	}); err != nil {
		s.logger.Errorf("error resetting password for user key [%s]: %v", key, err)
		return nil, false, err
	}
	// Create a copy as cur is pointing to an object in API server cache. Without copy causes intermittent failures in password reset integ test
	// where password is corrupted in user object returned from API server on subsequent GET
	ret, err := cur.Clone(nil)
	if err != nil {
		s.logger.Errorf("error creating a copy of user obj: %v", err)
		return ret, false, err
	}
	user := ret.(*auth.User)
	// return non-hashed password to user
	user.Spec.Password = genPasswd
	s.logger.Debugf("password successfully reset for user key [%s]", key)
	return *user, false, nil
}

// validateAuthenticatorConfig hook is to validate that authenticators specified in AuthenticatorOrder are defined
func (s *authHooks) validateAuthenticatorConfig(i interface{}, ver string, ignStatus bool) []error {
	s.logger.DebugLog("msg", "AuthHook called to validate authenticator config")
	var ret = []error{errAuthenticatorConfig}
	r, ok := i.(auth.AuthenticationPolicy)
	if !ok {
		return []error{errInvalidInputType}
	}

	// check if authenticators specified in AuthenticatorOrder are defined
	authenticatorOrder := r.Spec.Authenticators.GetAuthenticatorOrder()
	if authenticatorOrder == nil {
		s.logger.ErrorLog("msg", "Authenticator order config not defined")
		return ret
	}
	for _, authenticatorType := range authenticatorOrder {
		switch authenticatorType {
		case auth.Authenticators_LOCAL.String():
			if r.Spec.Authenticators.GetLocal() == nil {
				s.logger.ErrorLog("msg", "Local authenticator config not defined")
				return ret
			}
		case auth.Authenticators_LDAP.String():
			if r.Spec.Authenticators.GetLdap() == nil {
				s.logger.ErrorLog("msg", "Ldap authenticator config not defined")
				return ret
			}
		case auth.Authenticators_RADIUS.String():
			if r.Spec.Authenticators.GetRadius() == nil {
				s.logger.ErrorLog("msg", "Radius authenticator config not defined")
				return ret
			}
		default:
			s.logger.ErrorLog("msg", "Unknown authenticator type", "authenticator", authenticatorType)
			return ret
		}
	}
	return nil
}

// generateSecret is a pre-commmit hook to generate secret when authentication policy is created or updated
func (s *authHooks) generateSecret(ctx context.Context, kv kvstore.Interface, txn kvstore.Txn, key string, oper apiserver.APIOperType, dryRun bool, i interface{}) (interface{}, bool, error) {
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

func registerAuthHooks(svc apiserver.Service, logger log.Logger) {
	r := authHooks{}
	r.logger = logger.WithContext("Service", "AuthHooks")
	logger.Log("msg", "registering Hooks")
	svc.GetCrudService("User", apiserver.CreateOper).WithPreCommitHook(r.hashPassword)
	svc.GetCrudService("User", apiserver.UpdateOper).WithPreCommitHook(r.hashPassword)
	svc.GetCrudService("AuthenticationPolicy", apiserver.CreateOper).WithPreCommitHook(r.generateSecret).GetRequestType().WithValidate(r.validateAuthenticatorConfig)
	svc.GetCrudService("AuthenticationPolicy", apiserver.UpdateOper).WithPreCommitHook(r.generateSecret).GetRequestType().WithValidate(r.validateAuthenticatorConfig)
	svc.GetCrudService("Role", apiserver.CreateOper).GetRequestType().WithValidate(r.validateRolePerms)
	svc.GetCrudService("Role", apiserver.UpdateOper).GetRequestType().WithValidate(r.validateRolePerms)
	// hook to change password
	svc.GetMethod("PasswordChange").WithPreCommitHook(r.changePassword)
	// hook to reset password
	svc.GetMethod("PasswordReset").WithPreCommitHook(r.resetPassword)
}

func init() {
	apisrv := apisrvpkg.MustGetAPIServer()
	apisrv.RegisterHooksCb("auth.AuthV1", registerAuthHooks)
}
