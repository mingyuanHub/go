package main

import (
	"fmt"
	"time"
)

func work(cannel chan bool, i int) {
	for {
		select {
		default:
			fmt.Printf("%d", i)
			time.Sleep(  time.Duration(i) * time.Second)
			fmt.Println("mmmmmmmmmm")
			return
		case <- cannel:
			fmt.Println("tui chu")
			return
		}
	}
}

func main()  {
	cancel := make(chan bool)

	for i := 0; i < 10; i ++ {
		go work(cancel, i)
	}

	time.Sleep(time.Second * time.Duration(4))

	close(cancel)

	select {

	}
}
