package main

import (
	"fmt"
	"encoding/json"
	"bytes"
)

type ADemo struct {
	A string
	B string `json:"b"`
	C int `json:"c"`
	D string `json:"d"`
	Ext *ADemoExt
}

type ADemoExt struct {
	F int
}

type BDemo struct {
	ADemo ADemo
	B string `json:"-"`
	Ext *ADemoExt
}

type BDemoExt struct {
	ADemoExt
	G int
}

func main()  {
	ademo := &ADemo{
		A : "aaa",
		B : "bbb",
		C : 0,
		D : "",
		Ext: &ADemoExt{
			F:1,
		},
	}

	var bDemoExt = &BDemoExt{
		ADemoExt: *ademo.Ext,
	}

	b2Demo := &BDemo{ADemo: *ademo, Ext: bDemoExt}
	b3Demo := &BDemo{ADemo: *ademo}


	b2Demo.A = "22123123"

	pr(b2Demo)
	pr(b3Demo)



	pr(ademo)


	bdemo := &BDemo{
	}

	bdemo.A = "222"
	bdemo.B = "333"

	pr(bdemo)
}

func pr(ademo interface{}) {
	ajson, _ := json.Marshal(ademo)

	var jsonStr = []byte(ajson)
	reader := bytes.NewBuffer(jsonStr)

	fmt.Println(ademo, reader)
}
