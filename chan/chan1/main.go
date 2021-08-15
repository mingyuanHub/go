package main

import (
	"fmt"
	"time"
	"golang.org/x/net/context"
)

func main() {
	test()
	fmt.Println("main退出")
}


func test() {
	ctx, cancel:= context.WithTimeout(context.Background(), time.Millisecond * time.Duration(5))
	defer cancel()

	chan1 := make(chan bool)
	defer close(chan1)

	go func(ctx context.Context, chan1 chan bool) {
		select {
		case  <-ctx.Done():
			fmt.Println("退出携程")
			return
		default:
			fmt.Println("监控中...")
			time.Sleep(2 * time.Second)
			chan1 <- true
		}
	}(ctx, chan1)

	select {
	case <-ctx.Done():
		fmt.Println("ctx退出")
	case <-chan1:
		fmt.Println("chan1")
	}

	time.Sleep(10 * time.Second)
	fmt.Println("程序退出")
}
