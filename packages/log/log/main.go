package main

import "log"

func main() {
	log.Println(123)

	log.Fatalln(333)
	log.Println(123)
	log.Fatalf("aadfsafsfasdfsadf")

}