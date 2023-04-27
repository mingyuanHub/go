package main

import "fmt"

func main() {
	a := map[string]float64{
		"123":3,
		"1234":4,
	}

	var b map[string]float64

	b = a

	a["123"] = 33

	fmt.Println(b["123"])
}
