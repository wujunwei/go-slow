package rate

import (
	"goslow/prop"
	"sync"
	"time"
)

type Limiter interface {
	TryAcquire() bool
	TryAcquireSome(num int) bool
	Acquire() (time.Duration, error)
	AcquireSome(num int) (time.Duration, error)
	TimeoutAcquire(timeout time.Duration) error
	TimeoutAcquireSome(num int, timeout time.Duration) error
	SetRate(perUnit int, timeUnit time.Duration)
}

//Create
func Create(perUnit, maxLevel int, timeUnit time.Duration) (l Limiter) {
	watch := prop.Watch{}
	watch.Start()
	if maxLevel == 0 {
		l = &windowLimiter{
			stopWatch:    watch,
			maxLoopLevel: maxLevel,
			lock:         sync.Mutex{},
		}
	} else {
		l = &smoothWindow{
			stopWatch: watch,
			lock:      sync.Mutex{},
		}
	}

	l.SetRate(perUnit, timeUnit)
	return l
}
