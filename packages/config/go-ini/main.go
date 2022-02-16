package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main(){
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println(cfg.Section("").Key("name").String())
	fmt.Println(cfg.Section("").Key("age").Int())
	fmt.Println(cfg.Section("server").Key("host").String())
	fmt.Println(cfg.Section("server").Key("port").Value())
	fmt.Println(cfg.Section("server").Key("port").In("222", []string{"123"}))
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("config.ini")
}
