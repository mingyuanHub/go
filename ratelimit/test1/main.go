package main

import (
	"fmt"
	"time"
)

func main() {

	var a = 10
	var b = 11
	fmt.Println(111111, a/b)

	rule := NewRule()

	rule.AddSingleRule(1 * time.Second, 2)

	rule.AddSingleRule(10 * time.Second, 5)
	//
	rule.AddSingleRule(30 * time.Second, 12)

	for i := 0; i < 40; i ++ {
		res := rule.AllowVisit(0)
		res = rule.AllowVisit(1)
		fmt.Println(res)
		time.Sleep(1 * time.Second)
	}

}



