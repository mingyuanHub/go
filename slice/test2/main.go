package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	length := len(arr)

	for k, v := range arr {
		if v%2 == 0 {
			if k == length-1 {
				arr = arr[:k-1]
			} else {
				arr = append(arr[:k], arr[k+1:]...)
			}
		}
	}

	fmt.Println(arr)
}
