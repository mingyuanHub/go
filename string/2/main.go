package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.HasPrefix("a-bc", "b"))


	a := "abc"
	add(&a)

	fmt.Println(a)
}

func add(v *string) {
	*v += "ddd"
}