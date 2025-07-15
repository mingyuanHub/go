package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取本机所有的网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching network interfaces:", err)
		return
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 忽略down状态的接口
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// 遍历接口上的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error fetching addresses for interface:", iface.Name, err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 过滤IPv4和IPv6地址
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4() // 只保留IPv4地址，如果需要IPv6则注释掉这一行
			if ip == nil {
				// 这里处理IPv6地址，如果需要的话
				// 你可以打印出IPv6地址或者进一步处理
				fmt.Printf("IPv6 address for %s: %s\n", iface.Name, addr)
			} else {
				// 打印IPv4地址
				fmt.Printf("IPv4 address for %s: %s\n", iface.Name, ip)
			}
		}
	}
}