package main

import "fmt"

type Common interface {
	GetName() string

	GetSize() int
}

type Goods struct {
	Name   string
	Size   int
	Weight int
}

func (g Goods) GetName() string {
	return g.Name
}

func (g Goods) GetSize() int {
	return g.Size
}

func main() {
	var good = Goods{}

	good.Name = "cut"

	fmt.Println(good.GetName())

	var com Common

	com = good

	fmt.Println(com.GetName())

	a := getCommonName()

	fmt.Println(a.GetSize())
}


func getCommonName() Common {
	return Goods{}
}
