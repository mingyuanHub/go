package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个新的 Replacer。
	// "博客" 将被替换为 "所思所想"，"精彩" 将被替换为 "值得学习"。
	stringArr := []string{
		"博客", "所思所想",
		"精彩", "值得学习",
		"路", "1",
	}
	r := strings.NewReplacer(stringArr...)

	// 使用 Replace 方法替换字符串中的子串。
	s := "路多辛的博客非常精彩"
	s = r.Replace(s)

	fmt.Println(s)
	// Output: 路多辛的所思所想非常值得学习
}