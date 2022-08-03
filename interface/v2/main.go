package main

import "fmt"

type PCommon interface {
	GetName() string
}

type CCommon interface {
	PCommon
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
	var common CCommon

	var goods = Goods{
		Name: "cup",
		Size: 20,
	}

	common = goods

	fmt.Println(common.GetSize())
}
