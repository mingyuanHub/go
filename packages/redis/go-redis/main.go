package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"fmt"
)

var ctx = context.Background()

func main() {
	single()

	//cluster()
}

func single() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "172.16.0.210", 6379),
		Password: "tradplus123", // no password set
		DB:       0,             // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}

func cluster() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{fmt.Sprintf("%s:%d", "172.16.0.210", 6379)},
		Password: "tradplus123", // no password set
	})

	err := rdb.Set(ctx, "key1", "value2", 0).Err()
	if err != nil {
		panic(err)
	}
}