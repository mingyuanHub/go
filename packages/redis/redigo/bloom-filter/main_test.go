package main

import (
	"fmt"
	redisbloom "github.com/RedisBloom/redisbloom-go"
	"math/rand"
	"testing"
)

func BenchmarkXXX(b *testing.B) {
	key := "bench-mark-cars"
	for i := 0; i < b.N; i++ {
		car := fmt.Sprintf("%d", rand.Intn(1000000))
		var client = redisbloom.NewClientFromPool(RedisPool, "bloom-1")
		if ok, err := client.Exists(key, car); err != nil {

		} else {
			if !ok {
				client.Add(key, car)
			}
		}
	}
}