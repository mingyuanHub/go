package main

import (
	"fmt"
	"strings"
)

func main() {

	//使用Builder
	var sts = []string{"123", "456"}

	var b strings.Builder
	for _, st := range sts {
		b.WriteString(st)
	}

	b.WriteByte('7')
	b.Write([]byte("你"))
	b.WriteRune('好')

	fmt.Fprintf(&b, "%d", 8)
	fmt.Fprintf(&b, "%T", 9)

	fmt.Println(b.String(), b.Len(), b.Cap())


	//复制Builder
	c := b
	c.Reset() //必写，不然报错： illegal use of non-zero Builder copied by value
	c.WriteString("abc")
	c.Grow(25)
	fmt.Println(c.String(), c.Len(), c.Cap())

}
