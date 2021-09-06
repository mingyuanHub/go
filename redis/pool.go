package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 1000,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			redis.DialPassword("root")
			fmt.Println(1111)
			return redis.Dial(
					"tcp",
					fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),

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
	fmt.Println(11122, conn)
	fmt.Println(2222, conn.Err())
	a, err := redis.Bool(conn.Do("EXISTS", "a"))
	fmt.Println(3333, a, err)
	conn.Do("set", "test", "d")

	fmt.Println(redis.String(conn.Do("SETEX", "b", 30, 1)))


	var a1 interface{}
	fmt.Println(reflect.TypeOf(a1))
	a1 = a1.(string)


}