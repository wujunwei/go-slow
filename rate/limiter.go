package rate

import "time"

type Limiter interface {
	TryAcquire() bool
	TryAcquireSome(num int) bool
	Acquire() (time.Duration, error)
	AcquireSome(num int) (time.Duration, error)
	TimeoutAcquire(timeout time.Duration) error
	TimeoutAcquireSome(num int, timeout time.Duration) error
}
