package rate

import "time"

type BucketLimiter struct {
}

func (limiter *BucketLimiter) acquire() (waitedTime time.Duration) {
	return
}
func (limiter *BucketLimiter) tryAcquire() (ok bool) {
	return
}
func (limiter *BucketLimiter) tryAcquireSome(num int) (ok bool) {
	return
}
func (limiter *BucketLimiter) acquireSome(num int) (waitedTime time.Duration) {
	return
}
func (limiter *BucketLimiter) timeoutAcquire(timeout time.Duration) (ok bool) {
	return
}
func (limiter *BucketLimiter) timeoutAcquireSome(num int, timeout time.Duration) (ok bool) {
	return
}
