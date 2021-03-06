package apigwpkg

import (
	"fmt"

	auditapi "github.com/pensando/sw/api/generated/audit"
	apiintf "github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/apigw"
)

type svcProfile struct {
	kind       string
	group      string
	auditLevel *string
	oper       apiintf.APIOperType
	defProf    apigw.ServiceProfile
	preauthn   []apigw.PreAuthNHook
	preauthz   []apigw.PreAuthZHook
	precall    []apigw.PreCallHook
	postcall   []apigw.PostCallHook
}

// GetKind gets the kind on which this Service profile operates on, "" if it is none or more than one kind
func (s *svcProfile) GetKind() string {
	return s.kind
}

// GetAPIGroup returns the API group to which this profile belongs.
func (s *svcProfile) GetAPIGoup() string {
	return s.group
}

// GetOper returns the operation involved, Unknown oper if none or more than one oper.
func (s *svcProfile) GetOper() apiintf.APIOperType {
	return s.oper
}

// GetAuditLevel returns the audit level if it is set. if not isSet is returned as false
func (s *svcProfile) GetAuditLevel() (level string, isSet bool) {
	if s.auditLevel != nil {
		return *s.auditLevel, true
	}
	return auditapi.Level_Basic.String(), false
}

// SetAuditLevel sets the audit level for the service profile
func (s *svcProfile) SetAuditLevel(level string) error {
	switch level {
	case auditapi.Level_Basic.String(), auditapi.Level_Request.String(), auditapi.Level_Response.String(), auditapi.Level_RequestResponse.String():
		s.auditLevel = &level
	default:
		return fmt.Errorf("Unknown level [%v]", level)
	}
	return nil
}

// preauthNHooks returns all registered pre authn hooks
func (s *svcProfile) PreAuthNHooks() []apigw.PreAuthNHook {
	if len(s.preauthn) == 0 && s.defProf != nil {
		return s.defProf.PreAuthNHooks()
	}
	return s.preauthn
}

// preauthzHooks returns all registered pre authn hooks
func (s *svcProfile) PreAuthZHooks() []apigw.PreAuthZHook {
	if len(s.preauthz) == 0 && s.defProf != nil {
		return s.defProf.PreAuthZHooks()
	}
	return s.preauthz
}

// PreCallHooks returns all registered pre call hooks
func (s *svcProfile) PreCallHooks() []apigw.PreCallHook {
	if len(s.precall) == 0 && s.defProf != nil {
		return s.defProf.PreCallHooks()
	}
	return s.precall
}

// PostCallHooks returns all registered post call hooks
func (s *svcProfile) PostCallHooks() []apigw.PostCallHook {
	if len(s.postcall) == 0 && s.defProf != nil {
		return s.defProf.PostCallHooks()
	}
	return s.postcall
}

// AddPreAuthNHook registers a pre authn hook
func (s *svcProfile) AddPreAuthNHook(hook apigw.PreAuthNHook) error {
	s.preauthn = append(s.preauthn, hook)
	return nil
}

// ClearPreAuthNHooks clears any hooks registered
func (s *svcProfile) ClearPreAuthNHooks() {
	s.preauthn = nil
}

// AddPreAuthZHook registers a pre authn hook
func (s *svcProfile) AddPreAuthZHook(hook apigw.PreAuthZHook) error {
	s.preauthz = append(s.preauthz, hook)
	return nil
}

// ClearPreAuthZHooks clears any hooks registered
func (s *svcProfile) ClearPreAuthZHooks() {
	s.preauthz = nil
}

// AddPreCallHook registers a pre call hook
func (s *svcProfile) AddPreCallHook(hook apigw.PreCallHook) error {
	s.precall = append(s.precall, hook)
	return nil
}

// ClearPreCallHooks clears any hooks registered
func (s *svcProfile) ClearPreCallHooks() {
	s.precall = nil
}

// AddPostCallHook registers a post call hook
func (s *svcProfile) AddPostCallHook(hook apigw.PostCallHook) error {
	s.postcall = append(s.postcall, hook)
	return nil
}

// ClearPostCallHooks clears any hooks registered
func (s *svcProfile) ClearPostCallHooks() {
	s.postcall = nil
}

// SetDefaults sets any system wide defaults to the service profile. This
//  is usually called during init and overridden if needed while registering
//  hooks.
func (s *svcProfile) SetDefaults() error {
	// All defaults for the service profile go here.
	return nil
}

// NewServiceProfile creates a new service profile object
func NewServiceProfile(fallback apigw.ServiceProfile, kind, group string, oper apiintf.APIOperType) apigw.ServiceProfile {
	return &svcProfile{
		defProf: fallback,
		kind:    kind,
		group:   group,
		oper:    oper,
	}
}
