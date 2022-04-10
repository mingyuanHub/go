package main

import "fmt"

func main() {
	done := make(chan int,5 )

	//go func() {
	//	fmt.Println("hello")
	//	<- done
	//
	//	<- done
	//}()


	done <- 1
	done <- 1

	fmt.Println("hello002222")


	done <- 1

	<-done

	l := len(done)

	fmt.Println(l, "hello0000")


}
