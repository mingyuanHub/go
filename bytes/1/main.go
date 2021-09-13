package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var buffer bytes.Buffer

	buffer.WriteString("aaaaaaaaaaaaaaaaaaaa")

	buffer.WriteString("\n")

	buffer.WriteString("bbbbbbbbbbbbbbbbbbbbbb")

	buffer.WriteString("ccccccccccccccccccccc")

	buffer.WriteString("dddddddddddddddddddd")

	buffer.WriteString("eeeeeeeeeeeeeeeeeeeeeee")

	fmt.Println(buffer.String(), reflect.TypeOf(buffer.String()))
}