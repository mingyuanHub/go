package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	M string
}

func main()  {

	m := &Name{"haha"}

	js(m)

}

func js(name interface{}) {
	byte, _ := json.Marshal(name)

	fmt.Println(string(byte))
}
