package main

import (
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	logs.Debug("test debug %v", "hhahahah")
	logs.Notice("test notice %d", 456)
	logs.Warn("test warn %s", "abc")
	logs.Error("test error %d", 123)

	logger := logs.NewLogger(65535)
	logger.SetLogFuncCallDepth(3)
	logger.EnableFuncCallDepth(true)
	logger.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log","level":7,"maxlines":0,"maxsize":0,"hourly":true,"maxhours":12,"color":true,"formatter": ""}`)

	logger.Debug("123")
}