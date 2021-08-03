package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10000,
	})
	defer listen.Close()
	if err != nil {
		fmt.Println("连接服务端失败, err:", err)
		return
	}
	for {
		var data [1024]byte
		// 无需建立连接，直接收发数据
		// 返回字节数，地址。。。
		// 传入一个slice，但是底层数组还是用的[1024]byte
		// addr是客户端的IP+端口 (服务端端口固定，客户端端口随机)
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("读取数据失败: ", err)
			continue
		}
		fmt.Printf("data: %v, addr:%v, err:%v\n", string(data[:n]), addr, err)
		// 与WriteTo类型，但是需要传入一个UDP地址
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}

}
