package main

import "fmt"

func main() {
	a := []float64{1111,1,1}
	add(a)
	fmt.Println(1111, a)
}

func add(list []float64) {
	list[2] = 2222
}
