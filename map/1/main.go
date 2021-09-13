package main

import (
	"fmt"
	"sync"
	"time"
)

var haha = map[int]int{}

var lock sync.RWMutex

func main() {
	go func() {
		for {
			lock.Lock()
			haha[2] = 11
			lock.Unlock()
			//for i := 0; i <10; i ++ {
			//	if i % 6 == 0 {
			//		haha[2] = 11
			//	}
			//}
		}
	}()

	for i := 0; i < 10; i ++ {
		go func() {
			lock.RLock()
			if index, ok := haha[i]; ok {
				fmt.Println(index)
			}
			lock.RUnlock()
		}()
	}

	time.Sleep(5 * time.Second)
}
