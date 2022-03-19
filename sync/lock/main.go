package main

import (
	"fmt"
	"sync"
	"time"
)

var a sync.Mutex

func main() {
	a.Lock()

	go func(){
		time.Sleep(5*time.Second)
		a.Unlock()
		fmt.Println(33)
	}()

	fmt.Println(11)
	a.Lock()
	fmt.Println(22)
}
