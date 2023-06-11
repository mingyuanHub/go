package main

import (
	"fmt"
	"time"
)

func GetZeroTime2() int64 {
	nowTime := time.Now().Unix()

	fmt.Println(1111, nowTime, nowTime%86400)

	return nowTime - nowTime%86400
}

func main() {
	fmt.Println(22222, GetZeroTime2())
}

//func GetZeroTime() int64 {
//
//	//timeZone := "Asia/Shanghai"
//
//	var l = time.FixedZone("CST", 8*3600)
//
//	nowTime := time.Now()
//	now := nowTime.In(l).Unix()
//	_, offsetSeconds := nowTime.In(l).Zone()          //相对于utc时区偏移秒数
//
//	fmt.Println(1111111, nowTime.Unix(), now, int64(offsetSeconds))
//	seconds := (now + int64(offsetSeconds)) % 86400 //相对于当前时区0点偏移秒数
//	return now - seconds
//}