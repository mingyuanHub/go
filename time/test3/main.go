package main

import (
	"fmt"
	"time"
)

func main() {

	//+8 Asia/Shanghai
	//+0 Europe/London
	//-8 America/Los_Angeles

	//获取时区，日期对应的时间戳
	timeZone := "America/Los_Angeles"
	date := "2022-02-18 16:00:00"

	if l, err := time.LoadLocation(timeZone); err != nil {

	} else {
		lt, _ := time.ParseInLocation("2006-01-02 15:04:05", date, l)
		fmt.Println(111111111, lt.Unix())
	}

	//测试
	nowDate := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(nowDate)

	if l, err := time.LoadLocation("America/Los_Angeles"); err != nil {
		fmt.Println(111, time.Time{})
	} else {
		fmt.Println(222, l)

		lt, _ := time.ParseInLocation("2006-01-02 15:04:05", nowDate, l)
		fmt.Println(lt.Unix())

		fmt.Println(lt.Format("2006-01-02 15:04:05"))

		fmt.Println(time.Now().In(l).Unix())
		fmt.Println(time.Now().In(l).Format("2006-01-02 15:04:05"))
	}

	//var cstZone = time.FixedZone("CST",  - 8 * 3600)
	//fmt.Println(cstZone)
	//fmt.Println(time.Now().In(cstZone).Format("2006-01-02 15:04:05"))


	//fmt.Println(time.LoadLocation(""))
}
