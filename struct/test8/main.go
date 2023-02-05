package main

import "fmt"

type test8 struct {
	a *adSize
	b *adSize
}

type adSize struct {
	w int
	h int
}

func main() {
	var test = &test8{
		a: &adSize{
			w: 10,
			h: 20,
		},
	}

	a := test.a

	test.b = a

	a.w = 30
	a.h = 40

	test.b.w = 50
	test.b.h = 60

	fmt.Println(test.a, test.b)
}
