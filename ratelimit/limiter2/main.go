package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main()  {
	//v1()
	//v1_2()
	//v1_3()
	//v1_4()
	//v2()
	v2_1()
}

func v1() {
	limit := rate.Every(1 * time.Second);
	limiter := rate.NewLimiter(limit, 10);

	for  {
		fmt.Println(limiter.Allow())
		time.Sleep(50 * time.Millisecond)
	}
}

func v1_2() {
	limit := rate.Every(200 * time.Millisecond);
	limiter := rate.NewLimiter(limit, 1);

	for  {
		fmt.Println(limiter.Allow(),1)
		time.Sleep(50 * time.Millisecond)
	}
}

func v1_3() {
	limiter := rate.NewLimiter(10, 1);  //qps = 10

	t := 0
	f := 0

	max := 1000/50 * 5
	n := 0

	for  {
		if limiter.Allow() {
			t ++
		} else {
			f ++
		}
		time.Sleep(50 * time.Millisecond)   //每秒20次请求 ， 10次通过，10次失败，通过占比 10/20 = 50%

		n++

		if n > max {
			break
		}
	}

	fmt.Println(float64(t)/float64(f+t))
}

func v1_4() {
	limit := rate.Every(100 * time.Millisecond)
	limiter := rate.NewLimiter(limit, 1);  //qps = 10

	t := 0
	f := 0

	max := 5
	n := 0

	for  {
		for i :=0; i < 20; i ++ {
			time.Sleep(50 * time.Millisecond)
			go func() {
				if limiter.Allow() {
					t ++
				} else {
					f ++
				}
			}()

		}

		time.Sleep(1 * time.Second)   //每秒20次请求 ， 10次通过，10次失败，通过占比 10/20 = 50%
		n++
		if n > max {
			break
		}
	}

	fmt.Println(t,f,float64(t)/float64(f+t))
}

func v2() {
	limiter := rate.NewLimiter(10, 10)
	for {
		fmt.Println(limiter.Allow())
		time.Sleep(50 * time.Millisecond)
	}
	//
	//time.Sleep(3000 * time.Millisecond)
	//fmt.Println("----------------------------")
	//
	//for i:=0;i<50;i++ {
	//	fmt.Println(limiter.Allow())
	//	time.Sleep(100 * time.Millisecond)
	//}
}

func v2_1() {
	limiter := rate.NewLimiter(2, 2)
	n := 0
	for {
		fmt.Println(n, limiter.Allow())
		time.Sleep(100 * time.Millisecond)
		n ++
	}
	//
	//time.Sleep(3000 * time.Millisecond)
	//fmt.Println("----------------------------")
	//
	//for i:=0;i<50;i++ {
	//	fmt.Println(limiter.Allow())
	//	time.Sleep(100 * time.Millisecond)
	//}
}

func v3() {
	limiter := rate.NewLimiter(1, 1)
	for {
		if err := limiter.Wait(context.Background()); err != nil {
			fmt.Println("false")
		} else {
			fmt.Println("true")
		}
		time.Sleep(50 * time.Millisecond)
	}
}