package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	//递归获取所有目录及子目录下所有文件
	//Walk()

	//只获取单个目录
	Dir()

	//源码
	ReadDir("./assets")
}


func Walk() {
	var files []string

	//filepath.Walk 接受一个string指向根目录(root)，和一个带有签名(signature)的函数类型WalkFunc
	//这个方法将会文件夹扫描的每一次遍历中被调用
	//可以看到另一个变量info的类型是os.FileInfo，这个变量非常重要因为我们从当前的文件（文件夹或者是文件）获取很多有用的信息：文件名，文件大小，模式，更改时间
	err := filepath.Walk("./assets", func(path string, info os.FileInfo, err error) error {

		//判断是否文件夹
		if info.IsDir() {
			return err
		}

		//判断包含拓展名
		if filepath.Ext(path) == ".img" {
			return err
		}

		//获取文件名称
		fmt.Println(info.Name())


		files = append(files, path)
		return err
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(files)
}

func Dir() {
	files, err := ioutil.ReadDir("./assets")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir(), file.Mode())
	}

}

// ioutil.ReadDir 【源码】
// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)  //读取的数量
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}