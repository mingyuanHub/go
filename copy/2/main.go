package main

import (
	"fmt"
	"reflect"
)

type name struct {
	m string
}

func main()  {
	nameArr := []*name{
		{"maaa"},
	}

	var name2Arr = make([]*name, len(nameArr))
	copy(name2Arr, nameArr)

	fmt.Println(name2Arr)
	fmt.Println(nameArr)

	name3Arr := copyDspAccountList(nameArr)

	fmt.Println(name3Arr)
	fmt.Println(nameArr)
}

func copyDspAccountList(dspAccountList []*name) []*name {
	var newDspAccountList []*name
	for _, dspAccount := range dspAccountList {
		newDspAccount := copyPoint(dspAccount)
		newDspAccountList = append(newDspAccountList, newDspAccount)
	}
	return newDspAccountList
}

func copyPoint(dspAccount *name) *name{
	vt := reflect.TypeOf(dspAccount).Elem()
	newOby := reflect.New(vt)
	newOby.Elem().Set(reflect.ValueOf(dspAccount).Elem())
	return newOby.Interface().(*name)
}