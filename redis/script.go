package redis

import "github.com/go-redis/redis/v7"

var incrBy = redis.NewScript(`
local key = KEYS[1]
local expire = ARGV[1]
local value = redis.call("INCR", key)
if not value then
  value = 0
end
redis.call("EXPIRE", key, expire)
return value
`)
