package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = map[int]interface{

	}{}

	//a[4] = map[int]string{2:"123"}

	fmt.Println(reflect.TypeOf(a[4]).String())

	if reflect.TypeOf(a[4]).String() != "string" {
		a[4] = "123123"
	}

	fmt.Println(a[4].(string))
}
