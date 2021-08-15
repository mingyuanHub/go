package main

import "fmt"

func main() {
	a := make([]int, 0, 3)

	a = append(a, 1, 2, 3, 4)

	i := 0
	for {
		for  {
			if i > 2 {
				break
			}
			i ++
		}

		if i == 8 {
			break
		}
		i ++
		fmt.Println(i)


	}

	//a = append(a[0: 1], a[2:]...)

	//b := make([]int, 0)
	for i, _ := range a {

		if i == 1 {
			a = append(a[0: 1], append([]int{6}, a[1:]...)...)
		}

		if i == 3 {
			a = append(a[0: 3], a[4:]...)
		}

		//fmt.Println(a)

		for i, _ := range a {

			if i == 1 {
				a = append(a[0: 1], append([]int{6}, a[1:]...)...)
			}

			if i == 3 {
				a = append(a[0: 3], a[4:]...)
			}

			break
		}

		fmt.Println(a)
	}



	fmt.Println(a, cap(a))
	//fmt.Println(b, cap(b))
}

