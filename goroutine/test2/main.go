package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Car struct {
	Name string
}

func (c *Car) Log()  {
	String(*c)
}

//todo：高并发下 for循环中的defer 会导致 c.log 会挂掉：c = nil；  问题还待复现
func main() {
	for i := 0; i < 100000; i ++ {
		c := &Car{}
		defer c.Log()

		go setName(c)
		time.Sleep(100*time.Millisecond)
	}
	time.Sleep(5*time.Second)
}

func setName(c *Car) {
	c.Name = "hahaha"
	c = nil
	time.Sleep(1*time.Second)
}

func String(ad interface{}) string {
	b, err := json.Marshal(ad)
	if err != nil {
		return fmt.Sprintf("%v", ad)
	}
	return string(b)
}