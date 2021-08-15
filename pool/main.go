package main

import (
	"fmt"
	"sync"
)

type Demo struct {
	Name string
	Num int
}

var strPool = sync.Pool{
	New: func() interface{} {
		return Demo{
			Name: "",
		}
	},
}

func main() {
	demo := strPool.Get().(Demo)
	demo.Name = "hello111"
	fmt.Println(demo.Name)
	strPool.Put(demo)

	demo = strPool.Get().(Demo)
	fmt.Println(demo.Name)
	strPool.Put(demo)
	fmt.Println(demo.Name)
}

