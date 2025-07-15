package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	M string `json:"m"`
	N *int `json:"n,omitempty"`
}

// 只验证最基础的解析
// 不做其他处理
func main()  {

	var str1 = "{\"m\":\"aaa\"}"
	var name1 *Name
	json.Unmarshal([]byte(str1), &name1)
	fmt.Println(name1)
	js(name1)

	var str2 = "{\"m\":\"aaa\", \"n\":0}"
	var name2 *Name
	json.Unmarshal([]byte(str2), &name2)
	fmt.Println(name2)
	js(name2)

	var str3 = "{\"m\":\"aaa\", \"n\":1}"
	var name3 *Name
	json.Unmarshal([]byte(str3), &name3)
	fmt.Println(name3)
	js(name3)

	//var a int = 9
	//name := &Name{M:"haha",N: &a }
	//js(name)

}

func js(name interface{}) {
	byte1, _ := json.Marshal(name)
	fmt.Println(string(byte1))
}
