package main

import (
	"golang.org/x/net/context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1000))

	defer fmt.Println(111)
	defer cancel()

	go func(ctx context.Context) {
		time.Sleep(time.Millisecond * 800)
		//done <- struct{}{}
		return
	}(ctx)

	select {
		case a := <-ctx.Done():
			fmt.Println(a, "call succ")
			return
		case b := <-time.After(time.Duration(time.Millisecond * 1200)):
			fmt.Println(b, "timeout!")
			return
	}
}
