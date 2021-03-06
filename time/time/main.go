package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)

	for {
		select {
		case <- timer.C:
			fmt.Println("time done")
			timer.Reset(time.Second * 2)
		}
	}
}