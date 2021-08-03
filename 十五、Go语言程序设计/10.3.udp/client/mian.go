package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10000,
	})
	if err != nil {
		fmt.Println("连接服务器失败, err:", err)
		return
	}
	defer socket.Close()

	// 从终端输入读取
	reader := bufio.NewReader(os.Stdin)
	var reply [1024]byte
	for {
		fmt.Print("请输入内容: ")
		// 每次读取一行
		msg, _ := reader.ReadString('\n')
		_, err := socket.Write([]byte(msg))
		if err != nil {
			fmt.Println("write data error:", err)
			return
		}

		// 收到回复的数据
		n, addr, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read data error:", err)
			return
		}
		fmt.Printf("count:%v, addr:%v, data:%v", n, addr, string(reply[:n]))

	}

	//sendData := []byte("hello world!")
	//_, err = socket.Write(sendData)
	//if err != nil {
	//	fmt.Println("send data error:", err)
	//	return
	//}
	//data := make([]byte, 4096)
	//// 返回字符长度，服务端IP+端口，
	//n, addr, err := socket.ReadFromUDP(data)
	//if err != nil {
	//	fmt.Println("接收数据失败, err:", err)
	//	return
	//}
	//fmt.Printf("recv:%v, addr:%v, count:%v\n", string(data[:n]), addr, n)

}
