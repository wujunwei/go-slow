package rate

import "time"

type Limiter interface {
	setRate() bool
	getRate() float64
	tryAcquire() (Token, error)
	tryAcquireSome(num int) (Token, error)
	acquire() (Token, error)
	acquireSome(num int) (Token, error)
	timeoutAcquire(timeout time.Duration) (Token, error)
	timeoutAcquireSome(num int, timeout time.Duration) (Token, error)
}
