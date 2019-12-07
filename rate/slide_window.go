package rate

import (
	"sync"
	"time"
)

type WindowLimiter struct {
	stopWatch        Watch
	lock             sync.Mutex
	maxPermits       int
	availablePermits int
}

func (limiter *WindowLimiter) Acquire() (waitedTime time.Duration) {

	return
}
func (limiter *WindowLimiter) TryAcquire() (ok bool) {
	ok = limiter.TryAcquireSome(1)
	return
}
func (limiter *WindowLimiter) TryAcquireSome(num int) (ok bool) {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()
	if num <= limiter.availablePermits {
		limiter.maxPermits -= num
		ok = true
	}
	return
}
func (limiter *WindowLimiter) AcquireSome(num int) (waitedTime time.Duration) {
	return
}
func (limiter *WindowLimiter) TimeoutAcquire(timeout time.Duration) (ok bool) {
	return
}
func (limiter *WindowLimiter) TimeoutAcquireSome(num int, timeout time.Duration) (ok bool) {
	return
}
func (limiter *WindowLimiter) setRate(perSecond int) {
	limiter.maxPermits = perSecond
}
