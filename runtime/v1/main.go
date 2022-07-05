package main

import (
	"fmt"
	"runtime"
)

func say() {
	for i := 0; i < 10 ; i ++{

		fmt.Println(i, 222)

		//runtime.Gosched()
	}
}

//runtime.Gosched()
//这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。同时，当前的goroutine也会在未来的某个时间点继续运行。

func main() {
	runtime.GOMAXPROCS(2)

	go say()
	runtime.Gosched()
	fmt.Println(1111)  //(每次结果不一定,但"1111" 一定输出且在最后)
	//runtime.Gosched()
	//fmt.Println(2222)
}
