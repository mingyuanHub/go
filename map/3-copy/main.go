package main

import "fmt"

type test struct {
	p float64
	m map[string]float64
}

func (t *test) copy() *test  {
	c := *t
	return &c
}

//(浅)拷贝对于值类型的话是完全拷贝一份相同的值；而对于引用类型是拷贝其地址，也就是拷贝的对象修改引用类型的变量同样会影响到源对象。
func main() {
	t0 := &test{
		p : 20,
		m: map[string]float64{
			"11": 21,
			"12": 5,
		},
	}

	t1 := t0.copy()

	fmt.Println(100, &*&t0.m)

	fmt.Println(111, &*&t1.m)

	t1.p = 22
	t1.m = map[string]float64{
		"11": 25,
	}

	fmt.Println(100000000, t0)

	fmt.Println(111111111, t1)
}
