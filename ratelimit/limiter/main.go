package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main()  {
	var rateLimitR = 2
	var rateLimitB = 5
	var rateLimit = rate.NewLimiter(rate.Limit(rateLimitR), rateLimitB)

	for  {
		fmt.Println(111)
		rateLimit.AllowN(time.Now(), 4)
		fmt.Println(222)
		if rateLimit.AllowN(time.Now().Add(time.Second), 4) {
			fmt.Println("allow to request")
		} else {
			fmt.Println("too manny request")
			break
		}
	}

	time.Sleep(3 * time.Second)

	for  {
		if rateLimit.Allow() {
			fmt.Println("allow to request")
		} else {
			fmt.Println("too manny request")
			break
		}
	}

}
