package main

import "unicode"
import "fmt"
import "strconv"

func main() {
	a := "4.222rrrrr"

	for _, r := range a {
		fmt.Println(unicode.IsNumber(r))
	}

	_, err := strconv.ParseFloat(a, 64)

	fmt.Println(err)
}