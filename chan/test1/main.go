package main

import "fmt"

func main() {
	done := make(chan int)

	go func() {
		fmt.Println("hello")
		<- done

		<- done
	}()


	done <- 1

	fmt.Println("hello002222")




	done <- 1

	fmt.Println("hello0000")


}
