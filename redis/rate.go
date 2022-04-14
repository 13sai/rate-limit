package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type RateLimiter struct {
	redis    *redis.Client
	duration int
	limit    int
	burst    int
	prefix   string
}

func NewRateLimit(project string, seconds, limit, burst int, cfg RedisCfg) (*RateLimiter, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})

	return &RateLimiter{
		redis:    r,
		limit:    limit,
		duration: seconds,
		prefix:   project,
		burst:    burst,
	}, nil
}

func (l *RateLimiter) Allow(item string) (btn bool, err error) {
	now := time.Now().Unix()
	slot := now / int64(l.duration)
	name := fmt.Sprintf("%s-%d", l.prefix, slot)
	incr, err := incrBy.Run(l.redis, []string{name}, l.duration+10).Int()
	if err != nil {
		return
	}
	btn = (incr <= l.limit+l.burst)
	return
}
