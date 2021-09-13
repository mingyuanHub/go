package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkString(b *testing.B)  {
	for i := 0; i <= b.N; i ++ {
		str()
	}
}

func str() {
	var buffer string

	buffer += "aaaaaaaaaaaaaaaaaaaa" + "\n" +"bbbbbbbbbbbbbbbbbbbbbb" +"ccccccccccccccccccccc"+"dddddddddddddddddddd"+"eeeeeeeeeeeeeeeeeeeeeee"

	//buffer += "\n"
	////
	//buffer += "bbbbbbbbbbbbbbbbbbbbbb"
	//
	//buffer += "ccccccccccccccccccccc"
	//
	//buffer += "dddddddddddddddddddd"
	//
	//buffer += "eeeeeeeeeeeeeeeeeeeeeee"

	fmt.Sprintf("%s", buffer)
}


func BenchmarkBuf(b *testing.B)  {
	for i := 0; i <= b.N; i ++ {
		buf()
	}
}

func buf() {
	//var buffer bytes.Buffer

	buffer := bytes.NewBufferString("123")

	buffer.WriteString("aaaaaaaaaaaaaaaaaaaa")

	buffer.WriteString("\n")
	//
	buffer.WriteString("bbbbbbbbbbbbbbbbbbbbbb")
	//
	buffer.WriteString("ccccccccccccccccccccc")

	buffer.WriteString("dddddddddddddddddddd")

	buffer.WriteString("eeeeeeeeeeeeeeeeeeeeeee")

	fmt.Sprintf("%s", buffer.String())
}


func BenchmarkSpr(b *testing.B)  {
	for i := 0; i <= b.N; i ++ {
		spr()
	}
}

func spr()  {
	buffer := fmt.Sprintf("%s%s%s%s%s%s", "aaaaaaaaaaaaaaaaaaaa", "\n", "bbbbbbbbbbbbbbbbbbbbbb", "ccccccccccccccccccccc", "dddddddddddddddddddd", "eeeeeeeeeeeeeeeeeeeeeee")
	fmt.Sprintf("%s", buffer)
}