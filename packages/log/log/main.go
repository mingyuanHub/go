package main

import (
	"log"
	"fmt"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile|log.Ldate|log.Ltime)
	//log.SetFlags(0)
	fmt.Println(log.Flags())

	log.SetPrefix("【my】")

	log.Println(123, 213)

	err := initLogFile()
	log.Println(err)
	logInfo("aaa")

	file, err := os.Create("test.log")
	logger := log.New(file, "", log.Llongfile|log.Ldate|log.Ltime)
	logger.Println("113123")

	log.Fatal(333)
	//log.Fatalln(333)
	//log.Fatalf("%d", 333)
}

