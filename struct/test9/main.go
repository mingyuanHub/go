package main

import "fmt"

type car struct {
	A map[string]string
	B *B
}

type B struct {
	Name string
	NameList []string
}

func main() {
	c1 := &car{
		B:&B{
			"hhhh",
			[]string{"a"},
		},
	}

	c2 := c1

	c2.B = &B{
		"11111111",
		[]string{"c"},
	}

	//c2.B = nil

	fmt.Println(c1,c2)
}
