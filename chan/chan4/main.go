package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	<-c
	<-c

	fmt.Println(111)

	<-c

	for val := range c {
		fmt.Println(val)
	}
}
