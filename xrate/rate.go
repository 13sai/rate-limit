package xrate

import (
	"errors"
	"time"

	"github.com/muesli/cache2go"
	"golang.org/x/time/rate"
)

var ErrMinDuration = errors.New("d must larger than 1s")

type RateLimiter struct {
	limiter *cache2go.CacheTable
}

func NewRateLimit(project string, seconds, limit, burst int) (*RateLimiter, error) {
	table := cache2go.Cache(project)
	table.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		val := rate.NewLimiter(rate.Every(time.Duration(1e9/float64(limit*seconds))), burst)
		return cache2go.NewCacheItem(key, time.Duration(seconds*1e9), val)
	})
	return &RateLimiter{
		limiter: table,
	}, nil
}

func (l *RateLimiter) Allow(item string) (btn bool, err error) {
	limit, err := l.limiter.Value(item)
	if err != nil {
		return
	}
	btn = limit.Data().(*rate.Limiter).Allow()
	return
}
