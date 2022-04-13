package xrate

import (
	"time"

	"github.com/muesli/cache2go"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	limiter *cache2go.CacheTable
	burst   int
}

func New(d time.Duration, limit, burst int, project string) *IPRateLimiter {
	cache := cache2go.Cache(project)
	cache.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		b := time.Duration((float64(1000000000/limit) * d.Seconds()))
		val := rate.NewLimiter(rate.Every(b), burst)
		return cache2go.NewCacheItem(key, d, val)
	})
	return &IPRateLimiter{
		limiter: cache,
		burst:   burst,
	}
}

func (i *IPRateLimiter) GetLimiter(item string) (*rate.Limiter, error) {
	v, err := i.limiter.Value(item)
	if err != nil {
		return nil, err
	}
	return v.Data().(*rate.Limiter), nil
}

func (i *IPRateLimiter) Allow(item string) bool {
	l, err := i.GetLimiter(item)
	return err != nil || l.Allow()
}
