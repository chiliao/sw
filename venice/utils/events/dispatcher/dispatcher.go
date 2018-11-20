// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

package dispatcher

import (
	"encoding/json"
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	evtsapi "github.com/pensando/sw/api/generated/events"
	"github.com/pensando/sw/venice/utils"
	memcache "github.com/pensando/sw/venice/utils/cache"
	"github.com/pensando/sw/venice/utils/events"
	"github.com/pensando/sw/venice/utils/log"
)

const (
	// defaultSendInterval to be used if the user given value is invalid <=0
	defaultSendInterval = 10 * time.Second

	// defaultDedupInterval to be used if the user given value is invalid <=0
	defaultDedupInterval = 24 * (60 * time.Minute) // 24 hrs
)

// dispatcherImpl implements the `Dispatcher` interface. It is responsible for
// dispatching events to all the registered writers.
type dispatcherImpl struct {
	sync.Mutex                  // for protecting the dispatcher object
	sendInterval  time.Duration // i.e, batch interval; events are sent to the writers in this interval
	dedupInterval time.Duration // events are de-duped for the given interval
	logger        log.Logger    // logger

	// store, cache and batch needs to be in sync always; so, they should be handled through a common lock
	eventsStore   events.PersistentStore // persistent store for events
	dedupCache    *cache                 // in-memory cache to de-dup events
	eventsBatcher *eventsBatcher         // batcher batches the list of events to be sent out in the next send (batch) interval

	// any operation on writers (events distribution, registration, un-registration) should not stall the events pipeline
	writers *eventWriters // event writers

	start sync.Once // used for starting the dispatcher

	stop     sync.Once      // used for shutting down the dispatcher
	shutdown chan struct{}  // to send shutdown signal to the daemon go routines (i.e. event distribution)
	wg       sync.WaitGroup // used to wait for the graceful shutdown of daemon go routines
	stopped  bool           // indicates whether the dispatcher is running or not
}

// upon registering the writer, each writers get a events channel to watch for
// events. `eventWriters` maintains the list of registered writers.
type eventWriters struct {
	sync.Mutex                    // to protect the writers map without having to stall the events pipeline
	list       map[string]*writer // map of writers with their name; writers are given a name during creation NewWriter(...).
}

// writer ties the writer with it's associated channel receiving events and the offset tracker.
type writer struct {
	eventsCh      events.Chan
	offsetTracker events.OffsetTracker
	wr            events.Writer
}

// NewDispatcher creates a new dispatcher instance with the given send interval.
func NewDispatcher(dedupInterval, sendInterval time.Duration, eventsStorePath string, logger log.Logger) (events.Dispatcher, error) {
	if dedupInterval <= 0 {
		dedupInterval = defaultDedupInterval
	}

	if sendInterval <= 0 {
		sendInterval = defaultSendInterval
	}

	if utils.IsEmpty(eventsStorePath) {
		return nil, fmt.Errorf("empty events store path")
	}

	// create persistent event store
	eventsStore, err := newPersistentStore(eventsStorePath)
	if err != nil {
		logger.Errorf("failed to create dispatcher; could not create events store, err: %v", err)
		return nil, fmt.Errorf("failed to create dispatcher, err: %v", err)
	}

	dispatcher := &dispatcherImpl{
		dedupInterval: dedupInterval,
		sendInterval:  sendInterval,
		logger:        logger.WithContext("submodule", "events_dispatcher"),
		dedupCache:    newDedupCache(dedupInterval),
		eventsBatcher: newEventsBatcher(),
		eventsStore:   eventsStore,
		writers:       &eventWriters{list: map[string]*writer{}},
		shutdown:      make(chan struct{}),
	}

	return dispatcher, nil
}

// start notifying writers of the events every send interval
func (d *dispatcherImpl) Start() {
	d.start.Do(func() {
		// start processing failed events
		d.ProcessFailedEvents()

		// start sending events from the cache
		d.wg.Add(1)
		go d.notifyWriters()

		d.logger.Info("started events dispatcher")
	})
}

// Action implements the action to be taken when the event reaches the dispatcher.
// 1. Writes the events to persistent store.
// 2. Add event to the de-dup cache.
// 3. Add the de-duped event to the batch which will be sent to the writers.
func (d *dispatcherImpl) Action(event evtsapi.Event) error {
	return d.addEvent(&event)
}

// helper function to write event to the persistent store and add events to the dedup cache & batch.
func (d *dispatcherImpl) addEvent(event *evtsapi.Event) error {
	d.Lock()
	defer d.Unlock()

	if d.stopped {
		d.logger.Errorf("dispatcher stopped, cannot process event: {%s}", event.GetSelfLink())
		return fmt.Errorf("dispatcher stopped, cannot process events")
	}

	if err := events.ValidateEvent(event); err != nil {
		d.logger.Errorf("event {%s} validation failed, err: %v", event.GetUUID(), err)
		return err
	}

	// write event to the persistent store (i.e. file)
	if err := d.writeToEventsStore(event); err != nil {
		d.logger.Errorf("failed to write event {%s} to persistent store (file), err: %v", event.GetUUID(), err)
		return err
	}

	// de-dup and add the event to batch
	if err := d.dedupAndBatch(events.GetEventKey(event), event); err != nil {
		d.logger.Errorf("failed to dedup and batch event {%s}, err: %v", event.GetUUID(), err)
		return err
	}

	return nil
}

// ProcessFailedEvents processes failed events; used to replay events during restarts based on the
// bookmarked offset of each writer.
func (d *dispatcherImpl) ProcessFailedEvents() {
	d.Lock()
	defer d.Unlock()

	d.writers.Lock()
	defer d.writers.Unlock()

	d.logger.Info("processing failed/pending events")

	// get the current offset of the persistent store which will be bookmarked by the writers once the failed events are processed
	currentEvtsOffset, err := d.eventsStore.GetCurrentOffset()
	if err != nil {
		d.logger.Errorf("couldn't get the current events file offset, err: %v", err)
		return
	}

	// nothing in the events file to be sent to the writers
	if currentEvtsOffset == 0 {
		d.logger.Debugf("current events file offset is 0; nothing to be sent to the writers")
		return
	}

	for _, w := range d.writers.list {
		writerName := w.wr.Name()
		writerOffset, err := w.wr.GetLastProcessedOffset()
		if err != nil {
			d.logger.Errorf("cannot process failed/pending events; failed to get the bookmarked offset for writer {%s}, err: %v", writerName, err)
			continue
		}

		// get the list of events pending events from persistent store
		evts, err := d.eventsStore.GetEventsFromOffset(writerOffset)
		if err != nil {
			d.logger.Errorf("cannot process failed/pending events; failed to get the events using offset {%v}", writerOffset)
			continue
		}

		if len(evts) > 0 {
			select {
			case <-w.eventsCh.Stopped():
				d.logger.Debugf("event receiver channel for writer {%s} stopped; cannot deliver events", writerName)
			case w.eventsCh.Chan() <- newBatch(evts, currentEvtsOffset):
				d.logger.Infof("sent failed/pending events to the writer {%s}", writerName)
			default: // to avoid blocking
				d.logger.Debugf("could not send failed/pending events to the writer {%s}", writerName)
			}
		}
	}
}

// RegisterWriter creates a watch channel and offset tracker for the caller and returns it.
// the caller can watch the channel for events and once done, can stop the channel.
// each channel maintains a name which is useful for stopping the watch. Offset tracker is used to
// bookmark the offset.
func (d *dispatcherImpl) RegisterWriter(w events.Writer) (events.Chan, events.OffsetTracker, error) {
	d.Lock()
	if d.stopped {
		d.logger.Errorf("dispatcher stopped, cannot register writer: {%s}", w.Name())
		d.Unlock()
		return nil, nil, fmt.Errorf("dispatcher stopped, cannot register writers")
	}
	d.Unlock()

	d.writers.Lock()
	defer d.writers.Unlock()

	writerName := w.Name()
	if _, ok := d.writers.list[writerName]; ok {
		d.logger.Errorf("writer {%s} exists already", writerName)
		return nil, nil, fmt.Errorf("writer with the given name exists already")
	}

	// to record and manage file offset
	offsetTracker, err := newOffsetTracker(path.Join(d.eventsStore.GetStorePath(), "offset"), writerName)
	if err != nil {
		d.logger.Errorf("could not create offset tracker, err: %v", err)
		return nil, nil, errors.Wrap(err, "failed to register writer")
	}

	writerOffset, err := offsetTracker.GetOffset()
	if err != nil {
		d.logger.Errorf("could not read from offset tracker, err: %v", err)
		return nil, nil, errors.Wrap(err, "failed to register writer")
	}

	// during restart, it is possible that the new writer could end up receiving more events than
	// intended (ones that were generated before the writer registration). To avoid such issue, new writer
	// is given the current events store offset. So, that it starts receiving events from now on(from current offset).
	if writerOffset == 0 { // new writer
		esCurrOffset, err := d.eventsStore.GetCurrentOffset()
		if err != nil {
			d.logger.Errorf("could not read current events store offset, err: %v", err)
			return nil, nil, errors.Wrap(err, "failed to register writer")
		}

		if err := offsetTracker.UpdateOffset(esCurrOffset); err != nil {
			d.logger.Errorf("could not update the writer offset, err: %v", err)
			return nil, nil, errors.Wrap(err, "failed to register writer")
		}
	}

	e := newEventsChan(w.ChLen())
	d.writers.list[writerName] = &writer{eventsCh: e, offsetTracker: offsetTracker, wr: w}
	d.logger.Debugf("writer {%s} registered with the dispatcher successfully", writerName)

	return e, offsetTracker, nil
}

// UnregisterWriter removes the writer identified by given name from the list of writers. As a result
// the channel associated with the given name will no more receive the events
// from the dispatcher. And, offset tracker will be stopped as well.
// call does nothing if the writer identified by given name is not found in the dispatcher's writer list.
func (d *dispatcherImpl) UnregisterWriter(name string) {
	d.writers.Lock()
	if w, ok := d.writers.list[name]; ok {
		delete(d.writers.list, name)
		w.eventsCh.Stop()
		d.logger.Debugf("writer {%s} unregistered from the dispatcher successfully", name)
	}
	d.writers.Unlock()
}

// Shutdown sends shutdown signal to the notifier, flushes all the de-duped events to all
// registered writers and closes all the writers.
func (d *dispatcherImpl) Shutdown() {
	d.stop.Do(func() {
		d.Lock()
		d.stopped = true // helps to stop accepting any more events
		close(d.shutdown)
		d.Unlock()

		// wait for the notifier to complete
		d.wg.Wait()

		// process the pending events and send to writers
		d.logger.Debug("flush the batched events to registered writers")
		d.Lock()
		evts := d.eventsBatcher.getEvents()
		offset, err := d.eventsStore.GetCurrentOffset()
		if err != nil {
			d.logger.Errorf("failed to get the offset from events store to flush, err: %v", err)
		} else {
			d.distributeEvents(evts, offset)
		}

		d.eventsStore.Close()
		d.Unlock()

		// close all the writers
		d.closeAllWriters()

		d.logger.Debugf("dispatcher shutdown")
	})
}

// notifyWriters is a daemon which processes the de-duped/cached events every send interval
// and distributes it to all the writers. This daemon stops when it receives shutdown
// signal.
func (d *dispatcherImpl) notifyWriters() {
	ticker := time.NewTicker(d.sendInterval)
	defer d.wg.Done()

	for {
		select {
		case <-ticker.C: // distribute current batch with offset
			d.Lock()
			evts := d.eventsBatcher.getEvents()
			if len(evts) == 0 {
				d.Unlock()
				continue
			}
			d.eventsBatcher.clear()

			offset, err := d.eventsStore.GetCurrentOffset()
			d.Unlock()

			if err != nil {
				d.logger.Errorf("failed to get the offset from events store to distribute events; skipping distribute, err: %v", err)
				continue
			}

			d.distributeEvents(evts, offset)
		case <-d.shutdown:
			return
		}
	}
}

// distributeEvents helper function to distribute given event list and offset to all writers.
func (d *dispatcherImpl) distributeEvents(evts []*evtsapi.Event, offset int64) {
	if len(evts) == 0 {
		return
	}

	d.writers.Lock()
	defer d.writers.Unlock()

	resp := newBatch(evts, offset)
	// notify all the watchers
	for _, w := range d.writers.list {
		select {
		case <-w.eventsCh.Stopped():
			d.logger.Debugf("writer event channel {%s} stopped; cannot deliver events", w.wr.Name())
		case w.eventsCh.Chan() <- resp:
			// slow writers will block this. So, it is highly recommended to set a large enough
			// channel length for them.
		default:
			d.logger.Debugf("writer event channel {%s} failed to receive events", w.wr.Name())
			// for non-blocking send; writer failing to receive the event for a any reason (channel full)
			// will lose the event.
		}
	}
}

// helper function to write the given to persistent event store.
func (d *dispatcherImpl) writeToEventsStore(event *evtsapi.Event) error {
	evt, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshal the given event")
	}

	if err = d.eventsStore.Write(append(evt, '\n')); err != nil {
		return errors.Wrap(err, "failed to write event to file")
	}

	return nil
}

// dedupAndBatch dedups the given event and adds it to the batch to be sent to the writers.
func (d *dispatcherImpl) dedupAndBatch(hashKey string, event *evtsapi.Event) error {
	evt := *event

	// look for potential de-duplication
	srcCache := d.getCacheByEventSource(event.GetSource())
	if existingEvt, ok := srcCache.Get(hashKey); ok { // found, update the count of the existing event and timestamp
		d.logger.Debugf("event {%s} found in cache, updating the counter and timestamp", event.GetSelfLink())
		evt = existingEvt.(evtsapi.Event)

		// update count and timestamp
		timestamp, _ := types.TimestampProto(time.Now())
		evt.EventAttributes.Count++
		evt.ObjectMeta.ModTime.Timestamp = *timestamp
	}

	// add to de-dup cache
	srcCache.Add(hashKey, evt)

	// add event to the batch
	d.eventsBatcher.add(hashKey, &evt)
	return nil
}

// closeAllWriters helper function to close all the writers
func (d *dispatcherImpl) closeAllWriters() {
	d.logger.Debug("closing all the registered writers")
	d.writers.Lock()
	defer d.writers.Unlock()

	for _, w := range d.writers.list {
		delete(d.writers.list, w.wr.Name())
		w.eventsCh.Stop()
	}

	// efficient than deleting the elements one by one
	d.writers.list = map[string]*writer{}
}

// getCacheByEventSource helper function that fetches the underlying cache of the given source.
func (d *dispatcherImpl) getCacheByEventSource(source *evtsapi.EventSource) memcache.Cache {
	return d.dedupCache.getSourceCache(events.GetSourceKey(source))
}
