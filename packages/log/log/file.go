package main

import (
	"os"
	"strconv"
	"time"
	"fmt"
)

var (
	logFIle *os.File
	err error
)

func initLogFile() error{
	logFIle, err = OpenLogFile()
	return err
}

func OpenLogFile() (*os.File, error) {
	timeLocation := time.FixedZone("UTC+8", +8*3600)

	//当前时间戳
	nowTime := time.Now().In(timeLocation)

	//当前日期
	day := nowTime.Format("2006-01-02")
	//当前时间小时
	hour := strconv.Itoa(nowTime.Hour())

	//按照所在时区 日期零点 转换为 时间戳
	timea, _ := time.ParseInLocation("2006-01-02", day, timeLocation)
	timeStampA := timea.Unix() //当天日期时间戳
	timeStampB := timeStampA + 3600*24 - 1 //当天日期最大事件时间戳
	fmt.Println(timeStampA, timeStampB)

	file, err := os.OpenFile("./logs/"+day+"-"+hour+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	fmt.Println("logfileName:", day+"-"+hour+".log")

	if err != nil {
		return nil, err
	}

	return file, nil
}

func logInfo(ct string)  {
	logFIle.WriteString(ct + "\r\n")
}
