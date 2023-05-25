package main

import (
	"fmt"
	"os/exec"

)

func main() {
	out, err := exec.Command("bash", "-c", "git clone").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("The date is", string(out))
}

