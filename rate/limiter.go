package rate

import "time"

type Limiter interface {
	tryAcquire() bool
	tryAcquireSome(num int) bool
	acquire() time.Duration
	acquireSome(num int) time.Duration
	timeoutAcquire(timeout time.Duration) bool
	timeoutAcquireSome(num int, timeout time.Duration) bool
}
