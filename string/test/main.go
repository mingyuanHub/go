package main

import "fmt"

func main() {
	a := "abc"

	//error：strings are immutable
	//a[1] = "a"

	for _,v := range a {
		fmt.Println(v)
	}

	for _,v := range []byte(a) {
		fmt.Println(v)
	}

	c := make([]int, 1, 3)

	c[6] = 7
}
