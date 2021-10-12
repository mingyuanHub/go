package main

import (
	"fmt"
	"time"
)

func main() {
	rule := NewRule()

	rule.AddSingleRule(3 * time.Second, 2)

	rule.AddSingleRule(10 * time.Second, 5)
	//
	rule.AddSingleRule(30 * time.Second, 12)

	for i := 0; i < 40; i ++ {
		res := rule.AllowVisit(222)
		fmt.Println(res)
		//time.Sleep(1 * time.Second)
	}

}



