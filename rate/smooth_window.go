package rate

import (
	"errors"
	"sync"
	"time"
)

type smoothWindow struct {
	stopWatch     Watch
	lock          sync.Mutex
	produceRate   time.Duration
	maxPermits    int
	storedPermits int
}

func (limiter *smoothWindow) Acquire() (time.Duration, error) {
	return limiter.AcquireSome(1)
}
func (limiter *smoothWindow) TryAcquire() bool {
	return limiter.TryAcquireSome(1)
}
func (limiter *smoothWindow) TimeoutAcquire(timeout time.Duration) error {
	return limiter.TimeoutAcquireSome(1, timeout)
}
func (limiter *smoothWindow) TryAcquireSome(num int) bool {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()
	limiter.lazyProduce()
	if num <= limiter.storedPermits {
		limiter.storedPermits -= num
		return true
	}
	return false
}

//AcquireSome this function will never return error
func (limiter *smoothWindow) AcquireSome(num int) (duration time.Duration, err error) {
	start := time.Now()
	balance := limiter.storedPermits - num
	balance = (balance >> 31) ^ balance - (balance >> 31) //calculate the abs of balance
	err = limiter.TimeoutAcquireSome(num, time.Duration(balance)*limiter.produceRate)
	duration = time.Since(start)
	return
}

func (limiter *smoothWindow) TimeoutAcquireSome(num int, timeout time.Duration) (err error) {
	limiter.lock.Lock()
	limiter.lazyProduce()
	var duration time.Duration = 0
	if limiter.storedPermits-num < 0 {
		duration = limiter.produceRate * time.Duration(num-limiter.storedPermits)
		if duration > timeout {
			limiter.lock.Unlock()
			return errors.New("up to the max level, or request timeout")
		}
	}
	limiter.storedPermits -= num
	limiter.lock.Unlock()
	time.Sleep(duration)
	return
}

func (limiter *smoothWindow) SetRate(perUnit int, timeUnit time.Duration) {
	limiter.maxPermits = perUnit
	limiter.storedPermits = perUnit
	limiter.produceRate = timeUnit / time.Duration(perUnit)
}

func (limiter *smoothWindow) lazyProduce() {
	num := int(limiter.stopWatch.Elapse() / limiter.produceRate)
	limiter.produce(num)
}

func (limiter *smoothWindow) produce(num int) {
	if limiter.storedPermits+num > limiter.maxPermits {
		limiter.storedPermits = limiter.maxPermits
	} else {
		limiter.storedPermits += num
	}
	limiter.stopWatch.Reset()
	limiter.stopWatch.Start()
}
