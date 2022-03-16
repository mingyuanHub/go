package main

import "fmt"

type car struct{
	Name string
}

func main() {
	var a any
	a = 123
	a = "123"
	a = &car{
		"123",
	}
	fmt.Println(a)
}
