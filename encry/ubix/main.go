package main

import (
	"fmt"
)

func main() {
	token := "a0332873aa7caea849e87e516bf3cdd2"

	p, _ := AesCBCDecrypte("24d31cb12ac949f1fb340d736a3b2325", token)
	fmt.Println(1111111, p)
}
