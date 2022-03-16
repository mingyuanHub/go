package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	s := "abcdefgh"
	s = Reverse(s)
	fmt.Println(s)

	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	a:=genMd5("123")
	fmt.Println(a)
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func genMd5(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str) //将str写入到w中
	hexstr := hex.EncodeToString(w.Sum(nil))
	return hexstr
}

