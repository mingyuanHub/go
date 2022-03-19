package main

import "fmt"

//素数筛
func GenerateNatural() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i ++ {
			ch <- i
		}
	}()

	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int  {

	out := make(chan int)

	go func() {
		if i := <- in; i%prime != 0 {
			out <- i
		}
	}()

	return out
}

func main()  {
	ch := GenerateNatural()

	for i:= 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v : %v\n", i + 1, prime)
		ch = PrimeFilter(ch, prime)
	}

}
