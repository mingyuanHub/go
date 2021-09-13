package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{}

	copy(b, a)

	fmt.Println(a, b)
}