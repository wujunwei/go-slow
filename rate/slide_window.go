package rate

import (
	"github.com/wujunwei/go-slow/prop"
	"sync"
	"time"
)

type windowLimiter struct {
	stopWatch      prop.Watch
	lock           sync.Mutex
	timeUnit       time.Duration
	maxPermits     int
	nextMaxPermits int
	storedPermits  int
}

func (limiter *windowLimiter) Acquire() time.Duration {
	return limiter.AcquireSome(1)
}
func (limiter *windowLimiter) TryAcquire() bool {
	return limiter.TryAcquireSome(1)

}
func (limiter *windowLimiter) TimeoutAcquire(timeout time.Duration) bool {
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
func (limiter *windowLimiter) AcquireSome(num int) time.Duration {
	start := time.Now()
	limiter.TimeoutAcquireSome(num, limiter.timeUnit)
	return time.Since(start)
}

//todo fix condition > num and foreach the condition util enough permits are produced
func (limiter *windowLimiter) TimeoutAcquireSome(num int, timeout time.Duration) bool {
	limiter.lock.Lock()
	limiter.lazyProduce()
	if num <= limiter.storedPermits {
		limiter.storedPermits -= num
		limiter.lock.Unlock()
		return true
	}
	duration := limiter.timeToProduce()
	if duration > timeout {
		return false
	}
	num = num - limiter.storedPermits
	limiter.storedPermits = 0
	// in order to lock before sleep ,we give num permits of the next-time-unit permits
	limiter.preProduce(num)
	limiter.lock.Unlock()
	time.Sleep(duration)
	return true
}

func (limiter windowLimiter) timeToProduce() time.Duration {
	return limiter.timeUnit - limiter.stopWatch.Elapse()
}

func (limiter *windowLimiter) preProduce(num int) {
	if limiter.nextMaxPermits < num {
		return
	}
	limiter.nextMaxPermits -= num
}

func (limiter *windowLimiter) SetRate(perUnit int, timeUnit time.Duration) {
	limiter.maxPermits = perUnit
	limiter.nextMaxPermits = perUnit
	limiter.storedPermits = perUnit
	limiter.timeUnit = timeUnit
}

func (limiter *windowLimiter) lazyProduce() {
	if limiter.stopWatch.Elapse() >= limiter.timeUnit {
		limiter.produce()
	}
}

func (limiter *windowLimiter) produce() {
	limiter.storedPermits = limiter.nextMaxPermits
	limiter.nextMaxPermits = limiter.maxPermits
	limiter.stopWatch.Reset()
	limiter.stopWatch.Start()
}

func Create(perUnit int, timeUnit time.Duration) *windowLimiter {
	watch := prop.Watch{}
	watch.Start()
	l := &windowLimiter{
		stopWatch: watch,
		lock:      sync.Mutex{},
	}
	l.SetRate(perUnit, timeUnit)
	return l
}
