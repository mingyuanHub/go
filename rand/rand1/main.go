package main

import (
	//"math/rand"
	"fmt"

	"crypto/rand"
	"math/big"
)

func main() {
	//var i = 0
	//for i < 2 {
	//	i ++
	//	fmt.Println(rand.Intn(100))
	//}

	var i = 0
	for i < 2 {
		i ++
		n, _ := rand.Int(rand.Reader, big.NewInt(5))
		fmt.Println(n.Int64())
	}

}
