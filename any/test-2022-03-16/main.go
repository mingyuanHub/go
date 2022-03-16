package main

import (
	"fmt"
)

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

	var b interface{}
	b = 123
	b = "123"
	b = &car{
		"123",
	}
	fmt.Println(b)

	var int any = 1
	fmt.Println(int)
}
