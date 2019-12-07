package rate

import "time"

type Watch struct {
	running      bool
	startTick    int64
	elapsedNanos int64
}

func (w *Watch) start() {
	w.running = true
	w.startTick = time.Now().UnixNano()
}

func (w *Watch) stop() {
	w.running = false
	w.elapsedNanos += time.Now().UnixNano() - w.startTick
}

func (w *Watch) reset() {

}

func (w Watch) isRunning() bool {
	return w.running
}
