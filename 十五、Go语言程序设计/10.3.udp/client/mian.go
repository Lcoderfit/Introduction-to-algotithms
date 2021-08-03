package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10000,
	})
	defer socket.Close()
	if err != nil {
		fmt.Println("连接服务器失败, err:", err)
		return
	}
	sendData := []byte("hello world!")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("send data error:", err)
		return
	}
	data := make([]byte, 4096)
	// 返回字符长度，服务端IP+端口，
	n, addr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收数据失败, err:", err)
		return
	}
	fmt.Printf("recv:%v, addr:%v, count:%v\n", string(data[:n]), addr, n)

}
