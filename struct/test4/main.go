package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)


type Car struct {
	Weight int
	Name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	lunzi int
}

type Train struct {
	Car
	Lun int
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.Name, p.Weight)
	return str
}

func main() {
	var a Bike
	a.Weight = 100
	a.Name = "bike"
	a.lunzi = 2
	fmt.Println(a)
	a.Run()

	var b Train
	b.Weight = 100
	b.Name = "train"
	b.Lun = 4
	b.Run()
	fmt.Printf("%s", &b)

	pr(b)
}


func pr(ademo interface{}) {
	ajson, _ := json.Marshal(ademo)

	var jsonStr = []byte(ajson)
	reader := bytes.NewBuffer(jsonStr)

	fmt.Println(ademo, reader)
}