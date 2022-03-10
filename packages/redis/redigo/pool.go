package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func newRedisPool() *redis.Pool {

	dialOption := []redis.DialOption{
		redis.DialPassword("tradplus123"),
		redis.DialReadTimeout(100 * time.Millisecond),
		redis.DialWriteTimeout(100 * time.Millisecond),
	}

	return &redis.Pool{
		MaxIdle:     1000,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", "172.16.0.210", 6379),
				dialOption...
			)
		},
	}
}

var (
	RedisPool *redis.Pool
)

func init() {
	RedisPool = newRedisPool()
}

func main() {
	conn := RedisPool.Get()
	defer conn.Close()

	conn = RedisPool.Get()
	conn = RedisPool.Get()
	defer conn.Close()

	a, err := redis.Bool(conn.Do("EXISTS", "a"))
	fmt.Println(111, a, err)
	conn.Do("set", "test", "d")

	fmt.Println(redis.String(conn.Do("SETEX", "b", 30, 1)))
}