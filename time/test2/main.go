package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(60))

	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC().Unix())
}
