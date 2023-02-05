package main

import (
	"encoding/json"
	"fmt"
)

//验证json key 大小写同时存在时的解析情况：{"A":1,"a":true}
func main() {
	js1()
	js2()
}

type Test1 struct {
	A int
}

//【tip:该字段会解析两次，一个成功一次报错】
//json: cannot unmarshal bool into Go struct field Test1.A of type int
//2
//

func js1() {
	json_str := "{\"A\":2,\"a\":true}"

	var test = &Test1{}

	err := json.Unmarshal([]byte(json_str), &test)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(test.A)
}




type Test2 struct {
	A interface{}
}

//【tip:该字段会解析两次，最后一个覆盖前一个】
//true
//

func js2() {
	json_str := "{\"A\":2,\"a\":true}"

	var test = &Test2{}

	err := json.Unmarshal([]byte(json_str), &test)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(test.A)
}