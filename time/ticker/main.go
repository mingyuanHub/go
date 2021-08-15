package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)

	for  {
		select {
			case a := <- ticker.C :
				fmt.Println(a)
		}
	}
}
