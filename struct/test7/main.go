package main

import "fmt"

type meta struct {
	name string
	age  int
}

func (m *meta) copy(dest *meta) {
	*dest = *m
}

func main() {
	a := &meta{name: "aaa", age: 11}
	var b = &meta{}
	a.copy(b)
	b.name = "bbb"
	b.age = 22
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func main1() {
	a := meta{name: "aaa", age: 11}
	var b meta
	b = a
	b.name = "bbb"
	b.age = 22
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
