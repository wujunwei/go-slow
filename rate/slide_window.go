package rate

import (
	"sync"
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

type WindowLimiter struct {
	sync.Mutex
	timeUnit time.Duration
	floor    int64
}

func (limiter *WindowLimiter) acquire() (token Token) {
	return
}
func (limiter *WindowLimiter) tryAcquire() (token Token) {
	return
}
func (limiter *WindowLimiter) tryAcquireSome(num int) (token Token) {
	return
}
func (limiter *WindowLimiter) acquireSome(num int) (token Token) {
	return
}
func (limiter *WindowLimiter) timeoutAcquire(timeout time.Duration) (token Token) {
	return
}
func (limiter *WindowLimiter) timeoutAcquireSome(num int, timeout time.Duration) (token Token) {
	return
}
