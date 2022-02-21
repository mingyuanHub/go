package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	st := "hello world. 你好 中国"
	r := strings.NewReader(st)

	//当前长度和大小
	fmt.Println(11111, r.Len(), r.Size())

	//读取3个字节
	a := make([]byte, 3)
	r.Read(a)
	fmt.Println(22222, string(a), r.Len(), r.Size())

	//io.SeekStart：  idx偏移量重新从0开始，偏移5个字节
	//io.SeekCurrent：idx偏移量重新从当前开始，偏移5个字节
	//io.SeekEnd：    idx偏移量重新从结尾开始，偏移5个字节
	r.Seek(5, io.SeekStart)
	fmt.Println(66666, io.SeekStart, r.Len(), r.Size())

	//读取3个字节
	b := make([]byte, 3)
	r.Read(b)
	fmt.Println(33333, string(b), r.Len(), r.Size())

	//读取33个字节
	all := make([]byte, 33)
	r.Read(all)
	fmt.Println(string(all))

	//读取一个utf8字段
	ch, s , _ := r.ReadRune()
	fmt.Println(44444, string(ch), s, r.Len(), r.Size())

	//撤回一个utf8字段
	r.UnreadRune()
	fmt.Println(55555, r.Len(), r.Size())


	//重置
	r.Reset(st)

	//设置偏移量，然后读取3个字节
	c := make([]byte, 3)
	r.ReadAt(c, 2)
	fmt.Println(string(c), r.Len(), r.Size())


	r.Reset("haha")
	fmt.Println(r, r.Len(), r.Size())
}
