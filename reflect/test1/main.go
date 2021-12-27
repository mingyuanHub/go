package main

import (
	"fmt"
	"reflect"
)

type good struct {
	name int
}

func (g *good) GetName() int {
	fmt.Println(g.name)
	return g.name
}

func main() {
	mingyuan := &good{
		name : 4,
	}

	my := reflect.ValueOf(mingyuan)

	method := my.MethodByName("GetName")

	fmt.Println(method.Kind())

	if method.Kind() == reflect.Func {
		args := []reflect.Value{}
		method.Call(args)
	}
}
