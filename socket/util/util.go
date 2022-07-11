package util

import (
	"socket/config"
	"bytes"
	"fmt"
	"net"
)

//Write 向连接中写入数据 用于客户端传递数据给服务端或服务端返回消息给客户端
func Write(conn net.Conn, content string) (int, error) {
	//fmt.Printf("send %v: %v\n", conn.RemoteAddr(), content)
	var bytebuf bytes.Buffer
	bytebuf.WriteString(content)
	bytebuf.WriteByte(config.MessageDelimiter)

	bytearr := bytebuf.Bytes()
	return conn.Write(bytearr)
}

//Read 从连接中读取字节流 以结束符位标记
func Read(conn net.Conn) (string, error) {
	var str string
	var bytebuf bytes.Buffer
	bytearr := make([]byte, 1)
	for {
		if _, err := conn.Read(bytearr); err != nil {
			return str, err
		}
		item := bytearr[0]
		if item == config.MessageDelimiter {
			break
		}
		bytebuf.WriteByte(item)
	}
	str = bytebuf.String()
	fmt.Printf("recv %v: %v\n", conn.RemoteAddr(), str)
	return str, nil
}