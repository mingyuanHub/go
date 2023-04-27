package main

import "fmt"

func main() {

	a := &AoDi{
		Name: "aoooooo",
	}

	c := &Car{
		AoDi: a,
	}

	fmt.Println(a, c.AoDi)

	c.AoDi = copy(c.AoDi)

	fmt.Println(a, c.AoDi)

	c.AoDi.Name = "dddddddddd"

	fmt.Println(a, c.AoDi)
}

type Car struct {
	AoDi *AoDi
}

type AoDi struct {
	Name string
}

func copy(a *AoDi) *AoDi {
	a1 := *a
	return &a1
}
