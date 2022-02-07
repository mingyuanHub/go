package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Car struct {
	Age int
	Data map[string]interface{}
	DataStr []string
}

func main() {
	car := &Car{
		Age: 10,
		Data: map[string]interface{}{
			"name":"lb",
		},
		DataStr: []string{"a1,a2"},
	}

	fmt.Println(car.Data)

	b := car.Data

	b["name"] = "l1111"

	fmt.Println(b)

	fmt.Println(car.Data)


	a := cloneJsonMap(car.Data)

	fmt.Println(a)

	a["name"] = "l222222"

	fmt.Println(a)

	fmt.Println(car.Data)



	fmt.Println("----------------------------------")

	c := car.Age
	fmt.Println(c)

	c = 111
	fmt.Println(car.Age)


	fmt.Println("----------------------------------")

	d := car.DataStr
	d = []string{"b1", "b2"}

	fmt.Println(d)
	fmt.Println(car.DataStr)



	test1()
	test2()
}

func test1() {

	fmt.Println("----------------------------------")


	m1 := map[string]interface{}{
		"name":"aa",
	}

	m2 := m1

	fmt.Println(m1)

	m2["name"] = "bb"

	fmt.Println(m1)
}

func test2() {

	fmt.Println("----------------------------------")


	m1 := map[string]string{
		"name":"aa",
	}

	m2 := m1

	fmt.Println(&m1, &m2)

	fmt.Println(m1)

	m2["name"] = "bb"

	fmt.Println(m1)
}


func cloneJsonMap(data map[string]interface{}) map[string]interface{} {
	var other = make(map[string]interface{})
	if len(data) > 0 {
	var jsonStr, _ = json.Marshal(data) // 因data原本就是由json解析而来，这里不应报error
	_ = json.NewDecoder(strings.NewReader(string(jsonStr))).Decode(&other)
	}
	return other
}
