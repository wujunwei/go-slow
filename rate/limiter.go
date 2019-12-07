package rate

import "time"

type Limiter interface {
	TryAcquire() bool
	TryAcquireSome(num int) bool
	Acquire() time.Duration
	AcquireSome(num int) time.Duration
	TimeoutAcquire(timeout time.Duration) bool
	TimeoutAcquireSome(num int, timeout time.Duration) bool
}
