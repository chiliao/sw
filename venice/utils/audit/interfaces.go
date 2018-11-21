package audit

import "github.com/pensando/sw/api/generated/audit"

// Auditor represents an audit event processor that processes, stores and manages audit events. An Auditor implementation
// could store audit events in ElasticSearch, log file etc.
type Auditor interface {
	// ProcessEvents handles an audit event
	ProcessEvents(events ...*audit.Event) error

	// Run will initialize the backend. It must not block, but may run go routines in the background. If
	// stopCh is closed, it is supposed to stop them. Run will be called before the first call to ProcessEvents.
	Run(stopCh <-chan struct{}) error

	// Shutdown will synchronously shut down the backend while making sure that all pending
	// events are delivered. It can be assumed that this method is called after
	// the stopCh channel passed to the Run method has been closed.
	Shutdown()
}
