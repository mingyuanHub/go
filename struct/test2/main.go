package main

import (
	"fmt"
	"encoding/json"
	"bytes"
)

type ADemo struct {
	A string
	B string
	C int `json:"c"`
	D string `json:"d"`
}

type ADemo2 struct {
	ADemo
	D string `json:"d"`
}


func main()  {
	ademo := ADemo{
		A : "aaa",
		B : "bbb",
		C : 0,
	}

	cdemo := &ademo

	bdemo := ademo

	bdemo.D = "4444"

	pr(cdemo)
	pr(bdemo)



	ademo2 := ADemo2{
		ADemo:ademo,
		D:"123",
	}

	pr(ademo2)

}


func pr(ademo interface{}) {
	ajson, _ := json.Marshal(ademo)

	var jsonStr = []byte(ajson)
	reader := bytes.NewBuffer(jsonStr)

	fmt.Println(ademo, reader)
}
