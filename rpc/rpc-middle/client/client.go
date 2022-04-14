package main

import (
	"fmt"
	"log"
)

func main() {
	client, err := DiaHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err.Error())
	}
	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(reply)
}



