package main

import (
	"fmt"
	"os"
)

func main() {

	//判断目录是否存在
	if ok, _ := PathExists("./test"); !ok {
		//创建目录
		//如果目录存在，则报错
		err := os.Mkdir("./test", 0666)
		if err != nil {
			fmt.Println(err)
		}
	}

	//不需要判断是否存在
	os.MkdirAll("./test", 0766)

	//递归创建目录
	err1 := os.MkdirAll("./test1/test2/test3/test4", 0766)

	if err1 != nil {
		fmt.Println(err1)
	}
}

//golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断：
//
//如果返回的错误为nil，说明文件或文件夹存在
//
//如果返回的错误类型使用os.IsNotExist()判断为true，说明文件或文件夹不存在
//
//如果返回的错误为其它类型，则不确定是否在存在

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}
