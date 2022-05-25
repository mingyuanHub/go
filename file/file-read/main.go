package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//直接读取
	ReadFile()

	//分行读取
	ReadLine()

	//分字节读取
	ReadBytes()
}

//直接将数据直接读取入内存，是效率最高的一种方式，但此种方式，仅适用于小文件，对于大文件，则不适合，因为比较浪费内存。
func ReadFile() {
	//使用os.readFile
	content, err :=os.ReadFile("./content.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(content))


	//使用ioutil.ReadFile； 和上面等价
	content, err = ioutil.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))


	//先创建句柄再读取
	//读取使用高级函数 os.Open
	file, err := os.Open("a.txt") // 或者 file, err := os.OpenFile("a.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err = ioutil.ReadAll(file)
	fmt.Println(string(content))
}


//一次性读取所有的数据，太耗费内存，因此可以指定每次只读取一行数据； 缺点：不适用于 不换行的大文件
//bufio.ReadLine()  【低级库，不太适合普通用户使用】
//bufio.ReadBytes('\n')  【】
//bufio.ReadString('\n')
func ReadLine() {
	file, err := os.Open("./content.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)
	for {

		//第一种：ReadBytes
		//lineBytes, err := r.ReadBytes('\n')
		//line := strings.TrimSpace(string(lineBytes))


		//第二种：ReadString
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(string(line))


		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println("读取行数据：", line)
	}
}

//每次只读取固定字节数
//bufio.NewReader 创建一个 Reader
//在 for 循环里调用  Reader 的 Read 函数，每次仅读取固定字节数量的数据。
func ReadBytes() {
	file, err := os.Open("./content.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)

	// 每次读取 5 个字节
	buf := make([]byte, 5)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		fmt.Println(string(buf[:n]))
	}

}
