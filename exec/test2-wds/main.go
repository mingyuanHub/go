package main

import (
	"fmt"
	"syscall"
	"os/exec"
	"runtime"
)



func main() {
	cmd := exec.Command("cmd", "/C", "D:/GoPath/bin/gojson.exe")
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	//Start执行不会等待命令完成，Run会阻塞等待命令完成。
	//err := cmd.Start()
	//err := cmd.Run()
	//cmd.Output()函数的功能是运行命令并返回其标准输出。
	buf, err := cmd.Output()

	fmt.Println(string(buf), err)
}
