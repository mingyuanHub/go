package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

//打印堆栈错误报错
func main()  {

	//主协程
	defer func() {
		if panicMsg := recover(); panicMsg != nil {
			fmt.Println(11111111, panicMsg, GetStack(), 2222)
			fmt.Println(33333333333, GetDebugStack(), 444444444)
		}
	}()

	var a []string

	//次协程
	go func() {

		defer func() {
			if panicMsg := recover(); panicMsg != nil {
				fmt.Println(555555555, GetDebugStack(), 666666666666666)
				fmt.Println(777777777777, GetAllStack(), 8888888888888)
			}
		}()

		fmt.Println(a[2])
	}()

	time.Sleep(2 * time.Second)
}

//PrintStack 打印调用堆栈错误
func GetStack()string{
	var buf [20]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

//当前堆栈错误
func GetDebugStack() string {
	return string(debug.Stack())
}

//全局堆栈错误
func GetAllStack() string {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	return  string(buf)
}