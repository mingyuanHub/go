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
	First string `json:"-"`
	Second string
	Full string
}



func main() {

	var nameMy name

	body := "{\"First\":\"1\"}"


	if err := json.Unmarshal([]byte(body), &nameMy); err != nil {
		fmt.Println(err.Error())
	}



	//nameMy = name{
	//	first: "222",
	//	second: "333",
	//	full: nameMy.first + nameMy.second,
	//}

	pr(nameMy)
}

func pr(ademo interface{}) {
	ajson, _ := json.Marshal(ademo)

	var jsonStr = []byte(ajson)
	reader := bytes.NewBuffer(jsonStr)

	fmt.Println(ademo, reader)
}