package main

import "fmt"

func main() {
	var map1 = map[int]int{}
	add(map1)
	fmt.Println(111111, map1)
}

func add(m map[int]int)  {
	m[1]= 2
}
