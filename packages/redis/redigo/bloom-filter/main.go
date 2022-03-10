package main

import (
	"fmt"
	redisbloom "github.com/RedisBloom/redisbloom-go"
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
	var client = redisbloom.NewClientFromPool(RedisPool, "aaa")

	a, err := client.Add("userssss", "aaa")
	fmt.Println(a, err)

	a, err = client.Add("userssss", "aaa")
	fmt.Println(a, err)

	a, err = client.Exists("userssss", "aaa")
	fmt.Println(a, err)

	var res []int64
	res, err = client.BfExistsMulti("rooms", []string{"aa", "bb", "kk"})
	fmt.Println(res, err)
}