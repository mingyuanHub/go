package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseFloat(fmt.Sprintf("%.0f", 3.550), 64))

	var a float64 = 0.0000

	fmt.Println(a == 0)
}
