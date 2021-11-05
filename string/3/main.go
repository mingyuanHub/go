package main

import (
	"fmt"
	"go/types"
	"reflect"
)

func main() {
	var a string= ""

	fmt.Println(a)

	var price interface{}

	fmt.Println(reflect.TypeOf(price))

	fmt.Println(reflect.TypeOf(nil))



	switch price.(type) {
	case float64:
		fmt.Println(111)
	case string:
		fmt.Println(222)
	case types.Nil:
		fmt.Println(333)
	default:

	}
}
