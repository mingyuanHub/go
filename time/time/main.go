package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTimer(time.Second * 2)

	for {
		select {
		case <- timer.C:
			fmt.Println("time done")
			timer.Reset(time.Second * 2)
		}
	}
}

//-------------------------------------------------------

//package main
//
//import (
//	"time"
//	"fmt"
//)
//
//func main() {
//	timer := time.NewTimer(time.Second * 2)
//
//	select {
//	case <- timer.C:
//		fmt.Println("time done")
//	}
//}