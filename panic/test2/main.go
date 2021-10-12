package main

import "fmt"

type Exp struct {
	Code int
	Msg  string
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			if se, ok := err.(*Exp); ok {
				fmt.Println(111111, se.Code)
			} else {
				panic(err)
			}
		}
	}()


	panic(&Exp{4000, "w44444adfasdfasdfasfasf"})
}
