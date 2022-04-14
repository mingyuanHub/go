package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

func main() {
	logf, err := rotatelogs.New(
		"./logs/app.%Y%m%d%H.log",
		//rotatelogs.WithLinkName("./logs/access_log"),
		rotatelogs.WithMaxAge(1 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {

	}

	defer logf.Close()

	n := 1

	for  {
		logf.Write([]byte("haha\r\n"))

		time.Sleep(5 * time.Second)

		if n == 30 {
			return
		}

		n ++
	}

	//log.SetOutput(logf)
	//log.Printf("Hello, World!")
}
