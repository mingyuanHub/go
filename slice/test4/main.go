package main

import "fmt"

func main() {
	a := []float64{1111,1,1}
	add(a)
	fmt.Println(1111, a)

	b := []int{1,2,3,4,5}
	c := DeleteSlice(b, 7)
	fmt.Println(b, c)
}

func add(list []float64) {
	list[2] = 2222
}


func DeleteSlice(items []int, item int) []int {
	i := 0
	for _, v := range items {
		if v != item {
			items[i] = v
			i++
		}
	}
	return items[:i]
}