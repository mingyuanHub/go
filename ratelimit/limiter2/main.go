package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main()  {
	limit := rate.Every(5 * 1000 * time.Millisecond);
	limiter := rate.NewLimiter(limit, 10);

	for  {
		fmt.Println(limiter.Allow())

		time.Sleep(100 * time.Millisecond)
	}
}
