package memkv

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

const (
	outCount = 128
)

// watcher implements kvstore.Watcher interface.
type watcher struct {
	f           *MemKv                   // kvstore pointer
	keyOrPrefix string                   // what is being watched
	fromVersion int64                    // version to watch from
	recursive   bool                     // set to true for prefix based watches
	keys        []string                 // all keys (expanded in case of prefix)
	outCh       chan *kvstore.WatchEvent // channel for watch events
	ctx         context.Context
	cancel      context.CancelFunc
}

// newWatcher creates a new watcher interface for key based watches.
// TBD: this function is common among all kv stores
func (f *MemKv) newWatcher(ctx context.Context, key string, fromVersion string) (*watcher, error) {
	if strings.HasSuffix(key, "/") {
		return nil, fmt.Errorf("Watch called on a prefix")
	}
	return f.watch(ctx, key, fromVersion, false)
}

// newPrefixWatcher creates a new watcher interface for prefix based watches.
// TBD: this function is common among all kv stores
func (f *MemKv) newPrefixWatcher(ctx context.Context, prefix string, fromVersion string) (*watcher, error) {
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return f.watch(ctx, prefix, fromVersion, true)
}

// watch sets up the watcher context and starts the watch.
// TBD: this function is common among all kv stores
func (f *MemKv) watch(ctx context.Context, keyOrPrefix string, fromVersion string, recursive bool) (*watcher, error) {
	newCtx, cancel := context.WithCancel(ctx)
	if fromVersion == "" {
		fromVersion = "0"
	}
	version, err := strconv.ParseInt(fromVersion, 10, 64)
	if err != nil {
		cancel()
		return nil, err
	}
	w := &watcher{
		f:           f,
		keyOrPrefix: keyOrPrefix,
		fromVersion: version,
		recursive:   recursive,
		keys:        []string{},
		outCh:       make(chan *kvstore.WatchEvent, outCount),
		ctx:         newCtx,
		cancel:      cancel,
	}
	w.startWatching()
	go w.waitForCancel()
	return w, nil
}

func sendEvent(w *watcher, key string, v *memKvRec, deleted bool) {
	evType := kvstore.Updated
	if deleted {
		evType = kvstore.Deleted
	} else if v.revision == v.createRev {
		evType = kvstore.Created
	}
	w.sendEvent(evType, key, []byte(v.value), v.revision)
}

// setupWatchers looks for all watchers and find the ones that
// have been watching the newly added key. It sends notifications in case
// there was a watcher and send them events
func (f *MemKv) setupWatchers(key string, v *memKvRec) {
	for watchKey, wl := range f.cluster.watchers {
		for _, w := range wl {
			if w.recursive {
				if strings.HasPrefix(key, watchKey) {
					v.watchers = append(v.watchers, w)
					sendEvent(w, key, v, false)
				}
			} else if watchKey == key {
				v.watchers = append(v.watchers, w)
				sendEvent(w, key, v, false)
			}
		}
	}
}

// sendWatchEvents sends watch events on updates or deletes, it doesn't
// scan for all watchers to see which ones matches a given key
func (f *MemKv) sendWatchEvents(key string, v *memKvRec, deleted bool) {
	var deleteRecKey = func(w *watcher, key string) {
		for idx, value := range w.keys {
			if value == key {
				w.keys = append(w.keys[:idx], w.keys[idx+1:]...)
				return
			}
		}
	}

	for _, w := range v.watchers {
		sendEvent(w, key, v, deleted)
		if deleted {
			deleteRecKey(w, key)
		}
	}
}

// startWatching starts the watch.
func (w *watcher) startWatching() {
	f := w.f

	// insert watcher in memKv's global list of watchers
	// insert watcher into key's watcher list
	// insert memKvRec into watcher's list (to handle Stop)
	f.cluster.Lock()
	defer f.cluster.Unlock()

	wl, ok := f.cluster.watchers[w.keyOrPrefix]
	if !ok {
		wl = []*watcher{}
	}
	wl = append(wl, w)
	f.cluster.watchers[w.keyOrPrefix] = wl

	if w.recursive {
		for key, v := range f.cluster.kvs {
			if strings.HasPrefix(key, w.keyOrPrefix) {
				w.keys = append(w.keys, key)
				v.watchers = append(v.watchers, w)
			}
		}
	} else {
		if v, ok := f.cluster.kvs[w.keyOrPrefix]; ok {
			w.keys = []string{w.keyOrPrefix}
			v.watchers = append(v.watchers, w)
		}
	}

	// If starting from a lower version that current object's version
	// send current object(s) on the channel
	for _, key := range w.keys {
		v := f.cluster.kvs[key]
		if v.revision >= w.fromVersion {
			sendEvent(w, key, v, false)
		}
	}
}

// sendEvent sends out the event unless the watch is stopped.
func (w *watcher) sendEvent(evType kvstore.WatchEventType, key string, value []byte, version int64) {
	f := w.f

	obj, err := f.codec.Decode(value, nil)
	if err != nil {
		log.Errorf("Failed to decode %v with error %v", string(value), err)
		w.sendError(err)
		return
	}

	err = f.objVersioner.SetVersion(obj, uint64(version))
	if err != nil {
		log.Errorf("Failed to set version %v with error: %v", version, err)
		w.sendError(err)
		return
	}

	e := &kvstore.WatchEvent{
		Type:   evType,
		Object: obj,
		Key:    key,
	}

	if len(w.outCh) == outCount {
		log.Warnf("Number of buffered watch events hit max count of %v", outCount)
	}

	select {
	case w.outCh <- e:
	case <-w.ctx.Done():
	}
}

// sendError sends out the status object for the given error.
func (w *watcher) sendError(err error) {
	status := &api.Status{
		Result:  api.StatusResultInternalError,
		Message: []string{err.Error()},
		Code:    http.StatusInternalServerError,
	}

	event := &kvstore.WatchEvent{
		Type:   kvstore.WatcherError,
		Object: status,
	}
	select {
	case w.outCh <- event:
	case <-w.ctx.Done():
	}
}

func (f *MemKv) deleteWatchers(w *watcher) {
	var deleteWatcher = func(v *memKvRec, w *watcher) {
		if v == nil {
			return
		}
		for idx, value := range v.watchers {
			if value == w {
				v.watchers = append(v.watchers[:idx], v.watchers[idx+1:]...)
				return
			}
		}
	}

	f.cluster.Lock()
	defer f.cluster.Unlock()

	for _, key := range w.keys {
		v := f.cluster.kvs[key]
		deleteWatcher(v, w)
	}

	wl, ok := f.cluster.watchers[w.keyOrPrefix]
	if ok {
		for idx, value := range wl {
			if value == w {
				wl = append(wl[:idx], wl[idx+1:]...)
			}
		}
		if len(wl) == 0 {
			delete(f.cluster.watchers, w.keyOrPrefix)
		} else {
			f.cluster.watchers[w.keyOrPrefix] = wl
		}
	}
}

// EventChan returns the channel for watch events.
func (w *watcher) EventChan() <-chan *kvstore.WatchEvent {
	return w.outCh
}

// Stop stops the watcher.
func (w *watcher) Stop() {
	w.f.deleteWatchers(w)
	w.cancel()
}

// waitForCancel waits for cancel and stops the watcher.
func (w *watcher) waitForCancel() {
	select {
	case <-w.ctx.Done():
		w.Stop()
	}
}
