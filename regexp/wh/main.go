package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := "abcd1234efgh567ab8j"
	reg := regexp.MustCompile(`ab(.*?)(ef|j)`)
	results := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(results)
	for _, result := range results {
		fmt.Println(result[1]) // 输出: 1234 和 5678
	}
}
