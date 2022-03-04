package main

import (
	"strings"
	"unicode"
)
import "fmt"
import "strconv"

type Demo struct {
	demo *demo
}

type demo struct {
	Name string
}

func main() {
	a := "4.222rrrrr"

	for _, r := range a {
		fmt.Println(unicode.IsNumber(r))
	}

	_, err := strconv.ParseFloat(a, 64)

	fmt.Println(err)

	b := "aaabbbcccaaa"
	b = strings.Replace(b, "a", "1", -1)
	fmt.Println(b)

	var nihaoa = &demo{Name: "nihao a"}
	var nihaoc = &demo{Name: "nihao c"}

	nihaob := nihaoa
	nihaob = nihaoc
	//nihaob.Name = nihaoc.Name

	fmt.Println(*nihaoa, &nihaob, &nihaoc)
	fmt.Println(nihaoa.Name, nihaob.Name, nihaoc.Name)



	//var nihaod = &demo{Name: "nihao a"}
	//
	//var nihaoD = &Demo{
	//	demo: nihaod,
	//}
	//
	//var nihaoe = &demo{Name: "nihao c"}
	//
	//nihaoD.demo = nihaoe
	//nihaoD.demo.Name = nihaoe.Name
	//
	//fmt.Println(nihaod.Name, nihaoD.demo.Name, nihaoe.Name)
}