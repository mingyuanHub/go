package main

import (
	"fmt"
	"reflect"
	"time"
)

type Car struct {
	Name string
}


type CarCar struct{
	Car *Car
}

func (adRequest *CarCar) Copy() *CarCar{
	//cp := *adRequest
	//return &cp

	vt := reflect.TypeOf(adRequest).Elem()
	newOby := reflect.New(vt)
	newOby.Elem().Set(reflect.ValueOf(adRequest).Elem())
	return newOby.Interface().(*CarCar)
}

func main() {
	var a = &CarCar{
		Car: &Car{
			"a1111111111",
		},
	}


	go func(a *CarCar) {
		b := a.Copy()
		b.Car.Name = "bbbb1111111111"

		time.Sleep(2 * time.Second)
		fmt.Println("b1:", b.Car.Name)
	}(a)

	go func(a *CarCar) {
		c := a.Copy()
		c.Car.Name = "cccc1111111111"
		fmt.Println("c1:",c.Car.Name)
	}(a)

	time.Sleep(5 * time.Second)

	fmt.Println("d1:", a.Car.Name)
}
