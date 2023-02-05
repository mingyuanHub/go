package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	M string
}

// 只验证最基础的解析
// 不做其他处理
func main()  {

	m := &Name{"haha"}

	js(m)

}

func js(name interface{}) {
	byte, _ := json.Marshal(name)

	fmt.Println(string(byte))
}
