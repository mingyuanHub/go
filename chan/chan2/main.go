
package main

import (
	"context"
"fmt"
"net/http"
"time"
)

func main() {
	http.HandleFunc("/test", HelloHandler)
	http.ListenAndServe("0.0.0.0:8002", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	intChan := make(chan int)
	defer close(intChan)

	for i := 3; i <= 8; i++ {
		go Request(ctx, i, intChan)
	}

	for i := 0; i <= 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(time.Now().Unix(), i, "main ctx", ctx.Err())
		case c := <-intChan:
			fmt.Println(time.Now().Unix(), i, "main chan", c)
		}
	}

	fmt.Println(w, "Hello World")
}

// 发起请求
func Request(ctx context.Context, index int, ch chan int) {
	// 第三方请求阻塞时间

	for  {
		time.Sleep(time.Duration(index) * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(time.Now().Unix(), index, "request ctx", ctx.Err())
			return
		default:

			fmt.Println(time.Now().Unix(), index, "request chan ", index)
			ch <- index
		}
	}

	//ch <- index

}