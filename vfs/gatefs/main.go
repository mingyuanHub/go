package main

import (
	"fmt"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/gatefs"
)

func main() {
	fs := gatefs.New(vfs.OS("./gor"), make(chan bool, 8))
	s := fs.String()
	fmt.Println(s)



}
