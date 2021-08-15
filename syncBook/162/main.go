package main

import (
	"fmt"
	"syscall"
	"os"
	"os/signal"
)

func main()  {
	fmt.Println(1)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("quit, %v", <-sig)
}
