package main

import "fmt"

func main() {
	a := "abc"
	add(&a)

	fmt.Println(a)
}

func add(v *string) {
	*v += "ddd"
}