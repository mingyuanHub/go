package main

import (
	"flag"
	_ "net/http/pprof"
	"time"
	"fmt"
)

func demo (){
	//ch := make(chan int)	//1
	ch := make(chan int, 2)	//2
	go func() {  //写chan
		time.Sleep(2 * time.Second)
		ch <- 0	//执行完成
		ch <- 0	//执行完成
	}()

	select {
	case <-ch:	//读chan
		fmt.Printf("exec success\n")
		return
	case <- time.After(1 *time.Second):
		fmt.Printf("exec timeout\n")
		return
	}
}

func main() {
	flag.Parse()

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8080", nil))
	//}()

	for i := 0; i < 400; i++ {
		go demo()
	}

	fmt.Printf("sleep 1hour")
	time.Sleep(10 * time.Second)
}