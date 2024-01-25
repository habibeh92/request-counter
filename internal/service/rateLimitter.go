package service

type RateLimiter struct {
	allowance chan struct{}
}

func NewRateLimiter(maxRequests int) *RateLimiter {
	return &RateLimiter{
		allowance: make(chan struct{}, maxRequests),
	}
}

func (rl *RateLimiter) Allow() bool {
	select {
	case rl.allowance <- struct{}{}:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) Release() {
	<-rl.allowance
}
