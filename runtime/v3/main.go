package main

import (
	"fmt"
)

type Bar struct {
	name map[string]int
}

func main() {
	var nameMap = map[int]map[string]int{
		2: {"aaa222":22},
	}

	var bar = &Bar{}
	if name, ok := nameMap[2]; ok {
		fmt.Println(11111111,  &name)
		bar.name = name
	}

	nameMap = map[int]map[string]int{
		3: {"aaa222":22},
	}

	fmt.Println(222222,  bar.name)
	fmt.Println(333333,  &nameMap)
}
