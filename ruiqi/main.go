package main

import "fmt"

func main()  {
	fmt.Println("300 * 20 = ", 300 * 20)

	fmt.Println("1 * 2 * 3 * 4 * 5= ", 1 * 2 * 3 * 4 * 5)

	//sum := 1
	//for i := 1; i <= 20; i ++ {
	//	sum = sum * i
	//}
	//
	//fmt.Println("sum=", sum)

	sum := 1
	for i := 1; i <= 20; i ++ {
		sum = sum + i
	}

	fmt.Println("sum=", sum)
}
