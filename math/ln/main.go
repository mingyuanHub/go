package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Log10(0.01)
	fmt.Println(a)

	b := math.Ln2
	fmt.Println(4.60517/(b*b))

	m := 4.60517/(b*b) * 100000000 / (1024 * 1024)
	fmt.Println(m)
}
