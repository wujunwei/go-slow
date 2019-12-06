package rate

import (
	"sync"
	"time"
)

type Unit float64

const (
	Nanosecond  Unit = 1
	Microsecond      = 1000 * Nanosecond
	Millisecond      = 1000 * Microsecond
	Second           = 1000 * Millisecond
	Minute           = 60 * Second
	Hour             = 60 * Minute
)

type WindowLimiter struct {
	sync.Mutex
	timeUnit Unit
	floor    int64
}

func (limiter *WindowLimiter) acquire() (waitedTime time.Duration) {
	return
}
func (limiter *WindowLimiter) tryAcquire() (ok bool) {
	return
}
func (limiter *WindowLimiter) tryAcquireSome(num int) (ok bool) {
	return
}
func (limiter *WindowLimiter) acquireSome(num int) (waitedTime time.Duration) {
	return
}
func (limiter *WindowLimiter) timeoutAcquire(timeout time.Duration) (ok bool) {
	return
}
func (limiter *WindowLimiter) timeoutAcquireSome(num int, timeout time.Duration) (ok bool) {
	return
}
