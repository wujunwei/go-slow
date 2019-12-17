package rate

import (
	"errors"
	"github.com/wujunwei/go-slow/prop"
	"sync"
	"time"
)

type windowLimiter struct {
	stopWatch      prop.Watch
	lock           sync.Mutex
	timeUnit       time.Duration
	maxLoopLevel   int
	maxPermits     int
	nextMaxPermits []int
	storedPermits  int
}

func (limiter *windowLimiter) Acquire() (time.Duration, error) {
	return limiter.AcquireSome(1)
}
func (limiter *windowLimiter) TryAcquire() bool {
	return limiter.TryAcquireSome(1)
}
func (limiter *windowLimiter) TimeoutAcquire(timeout time.Duration) error {
	return limiter.TimeoutAcquireSome(1, timeout)
}
func (limiter *windowLimiter) TryAcquireSome(num int) bool {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()
	limiter.lazyProduce()
	if num <= limiter.storedPermits {
		limiter.storedPermits -= num
		return true
	}
	return false
}
func (limiter *windowLimiter) AcquireSome(num int) (duration time.Duration, err error) {
	start := time.Now()
	err = limiter.TimeoutAcquireSome(num, time.Duration(limiter.maxLoopLevel)*limiter.timeUnit)
	if err != nil {
		return 0, err
	}
	return time.Since(start), err
}

func (limiter *windowLimiter) TimeoutAcquireSome(num int, timeout time.Duration) (err error) {
	limiter.lock.Lock()
	limiter.lazyProduce()
	if num <= limiter.storedPermits {
		limiter.storedPermits -= num
		limiter.lock.Unlock()
		return nil
	}
	// in order to lock before sleep ,we give num permits of the next-time-unit permits
	duration, ok := limiter.timeoutPreProduce(num, timeout)
	limiter.lock.Unlock()
	if !ok {
		return errors.New("up to the max level, or request timeout")
	}
	time.Sleep(duration)
	return nil
}

func (limiter windowLimiter) timeToNextProduce() time.Duration {
	return limiter.timeUnit - limiter.stopWatch.Elapse()
}

func (limiter *windowLimiter) SetRate(perUnit int, timeUnit time.Duration) {
	limiter.maxPermits = perUnit
	for i := 0; i < limiter.maxLoopLevel; i++ {
		limiter.nextMaxPermits = append(limiter.nextMaxPermits, perUnit)
	}
	limiter.storedPermits = perUnit
	limiter.timeUnit = timeUnit
}

func (limiter *windowLimiter) lazyProduce() {
	if limiter.stopWatch.Elapse() >= limiter.timeUnit {
		limiter.produce()
	}
}

func (limiter *windowLimiter) produce() {
	limiter.storedPermits = limiter.nextMaxPermits[0]
	limiter.nextMaxPermits = append(limiter.nextMaxPermits[1:], limiter.maxPermits)
	limiter.stopWatch.Reset()
	limiter.stopWatch.Start()
}

func (limiter *windowLimiter) timeoutPreProduce(num int, timeout time.Duration) (time.Duration, bool) {
	waitTime := limiter.timeToNextProduce()
	if waitTime > timeout {
		return 0, false
	}
	j := 0
	gains := limiter.storedPermits
	for i, permits := range limiter.nextMaxPermits {
		if i != 0 {
			waitTime += limiter.timeUnit
		}
		gains += permits
		if num <= gains {
			if timeout < waitTime {
				return 0, false
			}
			j = i
			break
		}
	} //judge if can we borrow time from future or not

	if num > gains {
		return 0, false
	}
	limiter.storedPermits = 0
	for i := 0; i <= j; i++ {
		if num >= limiter.nextMaxPermits[i] {
			limiter.nextMaxPermits[i] = 0
		} else {
			limiter.nextMaxPermits[i] -= num
		}
	}
	return waitTime, true
}

func Create(perUnit int, maxLevel int, timeUnit time.Duration) *windowLimiter {
	watch := prop.Watch{}
	watch.Start()
	l := &windowLimiter{
		stopWatch:    watch,
		maxLoopLevel: maxLevel,
		lock:         sync.Mutex{},
	}
	l.SetRate(perUnit, timeUnit)
	return l
}
