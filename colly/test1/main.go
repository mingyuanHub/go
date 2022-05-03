package main

import (
	"fmt"
	"regexp"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.baidu.com"),
	)

	detailRegex, _ := regexp.Compile(`/go/go\?p=\d+$`)

	//listRegex, _ := regexp.Compile(`/t/\d+#\w+`)

	fmt.Println(c, detailRegex)


}


