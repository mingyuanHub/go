package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	params := map[string]string{
		"aa": "a b c",
	}

	querys := url.Values{}
	for idx, item := range params {
		querys.Add(idx, item)

		fmt.Println(1111, url.PathEscape(item))
	}

	fmt.Println(2222, querys.Encode())

	fmt.Println(3333, strings.Replace(querys.Encode(), "+", "%20", -1))
}
