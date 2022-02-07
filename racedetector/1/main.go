package main

import "sync"

func main() {
	count := 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j < 1000; j ++ {
				count ++
			}
		}()
	}

	wg.Wait()
}
