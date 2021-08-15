package main

func main()  {
	c := make(chan []int)
	for i := 0; i < 3; i++{
		go func(a int, b chan) {
			//请求第三方接口
		}(i, c)
	}

}
