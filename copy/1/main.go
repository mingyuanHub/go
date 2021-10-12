package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 2)
	r1 := copy( b, a)
	fmt.Println(a, b, r1)

	c := "1234"
	d := "222222222"
	e := []byte(d)
	r2 := copy(e , c)
	fmt.Println(string(e), d, r2)


	h := []*Apple{{M:11}}

	i := make([]*Apple, 2)
	r3 := copy(i , h)
	fmt.Println(h, i, r3)

	fmt.Println(i[0].M)
	h[0].M = 8
	fmt.Println(i[0].M)
}

type Apple struct {
	M int
}