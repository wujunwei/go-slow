package rate

import "time"

type Limiter interface {
	tryAcquire() Token
	tryAcquireSome(num int) Token
	acquire() Token
	acquireSome(num int) Token
	timeoutAcquire(timeout time.Duration) Token
	timeoutAcquireSome(num int, timeout time.Duration) Token
}
