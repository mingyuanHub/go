package main

import (
	"fmt"
)

//func main()  {
//	client, err := rpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//
//	}
//
//	var reply string
//	err = client.Call(HelloServiceName+".Hello", "haha", &reply)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(reply)
//}


func main() {
	client, err := DiaHelloService("tcp", "localhost:1234")
	if err != nil {

	}
	var reply string
	err = client.Hello("haha", &reply)
	if err != nil {

	}

	fmt.Println(reply)
}



