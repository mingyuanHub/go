package main

import (
	"fmt"
	"reflect"
)

type car struct {
	name string
}

func main() {
	var a = "abcdddddd"
	getType(a)
	
	var b = car{
		name: "dddddd",
	}
	getType(b)
}

func getType(itf interface{})  {

	kind := reflect.TypeOf(itf).Kind()
	fmt.Println(1111111, kind == reflect.String, kind == reflect.Struct)

	//switch itf.(type) {
	//case string:
	//	print("string")
	//case reflect.Struct:
	//}
}
