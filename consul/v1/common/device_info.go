package common

import (
	"net"
	"net/url"
	"strings"
	"strconv"
)

func GetPort(addr string) (int, error) {
	urlInfo, err := url.Parse(addr)
	if err != nil {
		return 80, err
	}

	return strconv.Atoi(urlInfo.Port())
}

func GetLANHost() string {
	// 思路来自于Python版本的内网IP获取，其他版本不准确
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	// udp 面向无连接，所以这些东西只在你本地捣鼓
	res := conn.LocalAddr().String()
	res = strings.Split(res, ":")[0]
	return res
}
