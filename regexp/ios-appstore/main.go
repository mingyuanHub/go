package main

import (
	"fmt"
	"regexp"
)

func main() {
	url := "https://apps.apple.com/id/app/id479516143?123123"

	var key = "/id"
	reg := regexp.MustCompile(fmt.Sprintf(`%s\d+`, key))

	id := reg.FindString(url)

	fmt.Println(id)
}
