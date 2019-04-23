package authz

import (
	"fmt"
	"reflect"

	k8serrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/api/generated/staging"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"
)

// AuthorizedOperations returns authorized operations for an user
func AuthorizedOperations(user *auth.User, authorizer Authorizer) []*auth.OperationStatus {
	var actions []auth.Permission_ActionType
	for _, action := range auth.Permission_ActionType_value {
		actions = append(actions, auth.Permission_ActionType(action))
	}
	operations := authorizer.AuthorizedOperations(user, user.Tenant, globals.DefaultNamespace, actions...)
	// check cluster scoped authorization for user belonging to default tenant
	if user.Tenant == globals.DefaultTenant {
		clusterOperations := authorizer.AuthorizedOperations(user, "", "", actions...)
		operations = append(operations, clusterOperations...)
	}
	var accessReview []*auth.OperationStatus
	for _, authzOp := range operations {
		authzRes := authzOp.GetResource()
		op := &auth.Operation{
			Resource: &auth.Resource{Tenant: authzRes.GetTenant(), Group: authzRes.GetGroup(), Kind: authzRes.GetKind(), Namespace: authzRes.GetNamespace()},
			Action:   authzOp.GetAction(),
		}
		opStatus := &auth.OperationStatus{
			Operation: op,
			Allowed:   true,
		}
		accessReview = append(accessReview, opStatus)
	}
	return accessReview
}

// GetOperationsFromPermissions constructs authz.Operation from auth.Permission
func GetOperationsFromPermissions(permissions []auth.Permission) []Operation {
	var operations []Operation
	for _, permission := range permissions {
		// if actions are defined at a kind, group or namespace level
		if len(permission.ResourceNames) == 0 {
			for _, action := range permission.Actions {
				resource := NewResource(permission.ResourceTenant, permission.ResourceGroup, permission.ResourceKind, permission.ResourceNamespace, "")
				operations = append(operations, NewOperation(resource, action))
			}
			return operations
		}
		for _, resourceName := range permission.ResourceNames {
			for _, action := range permission.Actions {
				resource := NewResource(permission.ResourceTenant, permission.ResourceGroup, permission.ResourceKind, permission.ResourceNamespace, resourceName)
				operations = append(operations, NewOperation(resource, action))
			}
		}
	}
	return operations
}

// IsValidOperationValue validates operation interface value as it is an input coming from authz hooks in API Gateway
func IsValidOperationValue(operation Operation) bool {
	// make sure interface type and value are not nil
	if operation == nil || reflect.ValueOf(operation).IsNil() {
		return false
	}
	resource := operation.GetResource()
	if resource == nil || reflect.ValueOf(resource).IsNil() {
		return false
	}
	return true
}

// ValidatePerm validates that resource kind and group is valid in permission
func ValidatePerm(permission auth.Permission) error {
	s := runtime.GetDefaultScheme()
	switch permission.ResourceKind {
	case "", ResourceKindAll:
		if permission.ResourceGroup != ResourceGroupAll {
			if _, ok := s.Kinds()[permission.ResourceGroup]; !ok {
				return fmt.Errorf("invalid API group [%q]", permission.ResourceGroup)
			}
		}
	case auth.Permission_APIEndpoint.String():
		if permission.ResourceGroup != "" {
			return fmt.Errorf("invalid API group, should be empty instead of [%q]", permission.ResourceGroup)
		}
		if len(permission.ResourceNames) == 0 {
			return fmt.Errorf("missing API endpoint resource name")
		}
		var errs []error
		for _, resourceName := range permission.ResourceNames {
			err := ValidateResource(permission.ResourceTenant, permission.ResourceGroup, permission.ResourceKind, resourceName)
			if err != nil {
				errs = append(errs, err)
			}
		}
		if err := k8serrors.NewAggregate(errs); err != nil {
			return err
		}
	default:
		if err := ValidateResource(permission.ResourceTenant, permission.ResourceGroup, permission.ResourceKind, ""); err != nil {
			return err
		}
	}
	var errs []error
	for _, action := range permission.Actions {
		if err := ValidateAction(permission.ResourceGroup, permission.ResourceKind, action); err != nil {
			errs = append(errs, err)
		}
	}
	return k8serrors.NewAggregate(errs)
}

// ValidatePerms validates that resource kind and group is valid in permissions
func ValidatePerms(permissions []auth.Permission) error {
	var errs []error
	for _, perm := range permissions {
		if err := ValidatePerm(perm); err != nil {
			errs = append(errs, err)
		}
	}
	return k8serrors.NewAggregate(errs)
}

// ValidateOperation validates operation
func ValidateOperation(op *auth.Operation) (Operation, error) {
	// make sure interface type and value are not nil
	if op == nil || reflect.ValueOf(op).IsNil() {
		return nil, fmt.Errorf("operation not specified")
	}
	res := op.GetResource()
	if res == nil || reflect.ValueOf(res).IsNil() {
		return nil, fmt.Errorf("resource not specified")
	}
	if err := k8serrors.NewAggregate(op.Validate("all", "", true)); err != nil {
		return nil, err
	}
	if err := ValidateResource(res.Tenant, res.Group, res.Kind, res.Name); err != nil {
		return nil, err
	}
	if err := ValidateAction(res.Group, res.Kind, op.GetAction()); err != nil {
		return nil, err
	}
	return NewOperation(NewResource(res.Tenant, res.Group, res.Kind, res.Namespace, res.Name), op.Action), nil
}

// ValidateResource validates resource information
func ValidateResource(tenant, group, kind, name string) error {
	s := runtime.GetDefaultScheme()
	switch kind {
	case auth.Permission_Event.String(), auth.Permission_Search.String(), auth.Permission_MetricsQuery.String(), auth.Permission_FwlogsQuery.String(), auth.Permission_AuditEvent.String():
		if group != "" {
			return fmt.Errorf("invalid API group, should be empty instead of [%q]", group)
		}
	case auth.Permission_APIEndpoint.String():
		if group != "" {
			return fmt.Errorf("invalid API group, should be empty instead of [%q]", group)
		}
		if name == "" {
			return fmt.Errorf("missing API endpoint resource name")
		}
	default:
		if s.Kind2APIGroup(kind) != group {
			return fmt.Errorf("invalid resource kind [%q] and API group [%q]", kind, group)
		}
		ok, err := s.IsClusterScoped(kind)
		if err != nil {
			log.Infof("unknown resource kind [%q], err: %v", kind, err)
		}
		if ok && tenant != "" {
			return fmt.Errorf("tenant should be empty for cluster scoped resource kind [%q]", kind)
		}
		ok, err = s.IsTenantScoped(kind)
		if err != nil {
			log.Infof("unknown resource kind [%q], err: %v", kind, err)
		}
		if ok && tenant == "" {
			return fmt.Errorf("tenant should not be empty for tenant scoped resource kind [%q]", kind)
		}
	}
	return nil
}

// ValidateAction validates if an action is valid for a kind
func ValidateAction(group, kind, action string) error {
	switch kind {
	case string(staging.KindBuffer):
	case auth.Permission_AuditEvent.String(), auth.Permission_MetricsQuery.String(), auth.Permission_FwlogsQuery.String(), auth.Permission_Event.String(), auth.Permission_Search.String():
		switch action {
		case auth.Permission_Read.String():
		default:
			return fmt.Errorf("invalid resource kind [%q] and action [%q]", kind, action)
		}
	default:
		switch action {
		case auth.Permission_Clear.String(), auth.Permission_Commit.String():
			return fmt.Errorf("invalid resource kind [%q] and action [%q]", kind, action)
		default:
		}
	}
	return nil
}

// PrintOperations creates a string out of operations for logging
func PrintOperations(operations []Operation) string {
	var message string
	for _, oper := range operations {
		if oper != nil {
			res := oper.GetResource()
			if res != nil {
				owner := res.GetOwner()
				var ownerTenant, ownerName string
				if owner != nil {
					ownerTenant = owner.Tenant
					ownerName = owner.Name
				}
				message = message + fmt.Sprintf("resource(tenant: %v, group: %v, kind: %v, namespace: %v, name: %v, owner: %v|%v), action: %v; ",
					res.GetTenant(),
					res.GetGroup(),
					res.GetKind(),
					res.GetNamespace(),
					res.GetName(),
					ownerTenant,
					ownerName,
					oper.GetAction())
			}
		}
	}
	return message
}
