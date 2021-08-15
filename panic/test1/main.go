package main

import (
	"fmt"
	)

type v1 struct {
	a *v2
}

type v2 struct {
	b string
}

func main() {
	defer fmt.Println("22222")

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("3333", msg)
		}
	}()


	defer fmt.Println("44444")


	a := &v1{}

	fmt.Println(a.a.b)



	defer fmt.Println("55555")
}
