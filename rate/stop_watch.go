package rate

import "time"

type Watch struct {
	running      bool
	startTick    int64
	elapsedNanos int64
}

func (w *Watch) Start() {
	if w.running {
		return
	}
	w.running = true
	w.startTick = time.Now().UnixNano()
	w.elapsedNanos = 0
}

func (w *Watch) Stop() {
	if !w.running {
		return
	}
	w.running = false
	w.elapsedNanos += time.Now().UnixNano() - w.startTick
}

func (w *Watch) Reset() {
	if !w.running {
		return
	}
	w.running = false
	w.elapsedNanos += time.Now().UnixNano() - w.startTick
	w.startTick = 0
}
func (w Watch) Elapse() (elapsed time.Duration) {
	if w.running {
		elapsed = time.Duration(w.elapsedNanos + time.Now().UnixNano() - w.startTick)
	} else {
		elapsed = time.Duration(w.elapsedNanos)
	}
	return
}

func (w Watch) Running() bool {
	return w.running
}
