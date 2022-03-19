package main

import (
	"fmt"
	"sync"
	"math/rand"
)

type singleton struct {
	name int
}


var (
	instance *singleton
	once sync.Once
)

func Instance() int {
	once.Do(func() {
		instance = &singleton{
			name : rand.Int(),
		}
	})

	//instance = &singleton{
	//	name : rand.Int(),
	//}

	fmt.Println(instance.name)

	return  instance.name
}

func main()  {
	go Instance()
	go Instance()
}
