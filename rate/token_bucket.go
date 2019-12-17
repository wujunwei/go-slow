package rate

import "time"

type BucketLimiter struct {
}

func (limiter *BucketLimiter) Acquire() (waitedTime time.Duration, err error) {
	return
}
func (limiter *BucketLimiter) TryAcquire() (ok bool) {
	return
}
func (limiter *BucketLimiter) TryAcquireSome(num int) (ok bool) {
	return
}
func (limiter *BucketLimiter) AcquireSome(num int) (waitedTime time.Duration, err error) {
	return
}
func (limiter *BucketLimiter) TimeoutAcquire(timeout time.Duration) (ok error) {
	return
}
func (limiter *BucketLimiter) TimeoutAcquireSome(num int, timeout time.Duration) (ok error) {
	return
}
