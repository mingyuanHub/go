package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type a struct {
	ha  string
}

type name struct {
	a
	first string
	second string
	full string
}

var nameMy = &name{}

func main() {

	body := "{\"ha\":\"lo\"}"


	json.Unmarshal([]byte(body), &nameMy)

	//nameMy = &name{
	//	first: "222",
	//	second: "333",
	//	full: nameMy.first + nameMy.second,
	//	ha: 222
	//}

	pr(nameMy)
}

func pr(ademo interface{}) {
	ajson, _ := json.Marshal(ademo)

	var jsonStr = []byte(ajson)
	reader := bytes.NewBuffer(jsonStr)

	fmt.Println(ademo, reader)
}