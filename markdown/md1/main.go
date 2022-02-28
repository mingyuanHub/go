package main

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func main() {

	input := []byte(`#### title1`)

	output := blackfriday.MarkdownBasic(input)

	fmt.Println(string(output))
}

