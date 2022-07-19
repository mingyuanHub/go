package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"fmt"
)

//go:embed config/config.json
var f embed.FS

func main() {
	r := gin.New()
	// 初始化默认静态资源
	r.StaticFS("assets", http.FS(Static))

	// 设置模板资源
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(Templates, "templates/*")))

	bs, err := f.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

	r.Run(":8000")
}
