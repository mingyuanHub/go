package main

import (
	"reflect"
	"fmt"
)

type Demo struct {
	Name string `json:"name"`
}

func main()  {
	demo := &Demo{
		Name: "mingyuan",
	}

	appConfValue := reflect.ValueOf(demo).Elem()
	appConfType := reflect.TypeOf(demo).Elem()
	for i := 0; i < appConfValue.NumField(); i ++ {
		v := fmt.Sprintf("%v", appConfValue.Field(i))
		fmt.Println("invalid appConf.%v : %s", appConfType.Field(i).Name, v, appConfType.Field(i).Tag.Get("json"))
	}
}
