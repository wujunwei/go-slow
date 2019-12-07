package rate

import "time"

type Watch struct {
	running      bool
	startTick    int64
	elapsedNanos int64
}

func (w *Watch) start() {
	if w.running {
		return
	}
	w.running = true
	w.startTick = time.Now().UnixNano()
}

func (w *Watch) stop() {
	if !w.running {
		return
	}
	w.running = false
	w.elapsedNanos += time.Now().UnixNano() - w.startTick
}

func (w *Watch) reset() {
	if !w.running {
		return
	}
	w.running = false
	w.elapsedNanos += time.Now().UnixNano() - w.startTick
	w.startTick = 0
}
func (w Watch) elapse() int64 {
	if w.running {
		return w.elapsedNanos + time.Now().UnixNano() - w.startTick
	} else {
		return w.elapsedNanos
	}
}

func (w Watch) isRunning() bool {
	return w.running
}
