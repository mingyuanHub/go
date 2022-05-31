package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Car struct {
	Name string
}


type CarCar struct{
	Car *Car
}

//func (adRequest *CarCar) Copy() *CarCar{
//	cp := *adRequest
//	return &cp
//}

//func (adRequest *CarCar) Copy() *CarCar{
//	vt := reflect.TypeOf(adRequest).Elem()
//	newOby := reflect.New(vt)
//	newOby.Elem().Set(reflect.ValueOf(adRequest).Elem())
//	return newOby.Interface().(*CarCar)
//}

func (adRequest *CarCar) Copy() (*CarCar, error){
	newAdRequest := new(CarCar)

	b, err := json.Marshal(adRequest)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, newAdRequest)
	if err != nil {
		return nil, err
	}

	return newAdRequest, nil
}

//当struct1中有map类型及其他struct2时，必须深层拷贝
//即使copy struct1 后不是指针， 因为map类型的问题，也不能直接 a := b, 否则导致 a 和 b 的指针不一致， 但是 a.map  和 b.map 指针是一致的
func main() {
	var a = &CarCar{
		Car: &Car{
			"aaaa1111111111",
		},
	}

	for i:=0; i < 6; i ++ {
		acopy2, _ := a.Copy()
		go func(acopy *CarCar, i int) {
			fmt.Println(&acopy, &(acopy.Car))
			acopy.Car.Name = fmt.Sprintf("acopy%d", i)
			time.Sleep(2 * time.Second)

			fmt.Println("a:", acopy.Car.Name)
			fmt.Println("acopy:", acopy.Car.Name)
		}(acopy2, i)
	}


	//go func(a *CarCar) {
	//	b, _ := a.Copy()
	//	b.Car.Name = "bbbb1111111111"
	//
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("b1:", b.Car.Name)
	//}(a)
	//
	//go func(a *CarCar) {
	//	c, _ := a.Copy()
	//	c.Car.Name = "cccc1111111111"
	//	fmt.Println("c1:",c.Car.Name)
	//}(a)

	time.Sleep(5 * time.Second)

	//fmt.Println("d1:", a.Car.Name)
}
