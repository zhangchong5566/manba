package proxy

import (
	"time"

	"github.com/zhangchong5566/manba/pkg/pb/metapb"
	"github.com/juju/ratelimit"
)

type rateLimiter struct {
	limiter *ratelimit.Bucket
	option  metapb.RateLimitOption
}

func newRateLimiter(max int64, option metapb.RateLimitOption) *rateLimiter {
	return &rateLimiter{
		limiter: ratelimit.NewBucket(time.Second/time.Duration(max), max),
		option:  option,
	}
}

func (l *rateLimiter) do(count int64) bool {
	if l.option == metapb.Wait {
		l.limiter.Wait(count)
		return true
	}

	return l.limiter.TakeAvailable(count) > 0
}
