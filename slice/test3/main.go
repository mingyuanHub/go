package main

import "fmt"

func main()  {
	var test = []int{1}

	for key, _ := range test {
		fmt.Println(111111, test, key)

		if key == 0 || key == 1 {
			test = append(test[:key], test[key+1:]...)
		}
	}

	a := getA()
	
	switch a.(type) {
	case string:
		fmt.Println(111111)
	case []byte:
		fmt.Println(222222)
	}
}

func getA() interface{} {
	var a = []byte{}

	a = nil
	return a
}
