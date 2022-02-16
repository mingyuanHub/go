package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Name string
	Age  int
}

var Configuration = &Config{}

func main(){
	var err error
	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(){
		if err = file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&Configuration); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Configuration)
}
