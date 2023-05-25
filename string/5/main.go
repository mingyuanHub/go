package main

import "fmt"

type name struct {
	a string
}

func main() {

	var name1 = &name{}
	name1.a = "a"
	defer name1.log()


	name1.a += "1c"


}

func (c *name) log()  {
	fmt.Println(c.a)
}
