package main

import (
	"fmt"
	"time"
)

type pp struct {
	Name string
	P1 *p1
}

type p1 struct {
	Age int
}

func main() {
	p := &pp{
		Name: "aaa",
		P1: &p1{
			Age: 22,
		},
	}

	q := *p
	//q.Name = "bbb"
	c := &q

	fmt.Println(&p, &c)
	fmt.Println(&p.P1, &c.P1)

	for i := 0; i < 10; i ++ {
		go func(){
			a := *p
			fmt.Println("1:", p.Name)
			a.Name = "11111"
		}()

		go func(){
			b := *p
			fmt.Println("2:", p.Name)
			b.Name = "22222"
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("3:", p.Name)
}
