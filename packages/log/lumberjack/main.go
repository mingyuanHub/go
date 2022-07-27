package main

import (
	"github.com/natefinch/lumberjack"
	"log"
)

// pkg地址：github.com/natefinch/lumberjack
// 源码分析：https://www.jianshu.com/p/a53963036f69

func main()  {

	log.SetFlags(0)

	log.Println(111111)

	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/ev.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   true, // disabled by default
	})

	log.Println(1111)
}
