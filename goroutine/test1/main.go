package main

import (
	"time"
	"fmt"
)

func main() {

	//qqqq
	//line path/to/file:4
	go func() {
		fmt.Println("1")
	}()

	time.Sleep(time.Second)
}