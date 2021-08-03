package main

import (
	"fmt"
	"net"
)

/*
1.监听端口
2.接收客户端请求建立连接
3.与客户端通信
*/

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, err:", err)
		return
	}

	//// 接收客户端请求并建立连接
	//conn, err := listener.Accept()
	//if err != nil {
	//	fmt.Println("accept failed, err: ", err)
	//	return
	//}

	// 3.与客户端通信
	for {
		// 可以连接多个客户端
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}
		go processConn(conn)
	}
}

func processConn(conn net.Conn) {
	var tmp [128]byte
	// for使得服务端可以多次接收同一个客户端的消息
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read data error: ", err)
			return
		}
		// tmp[:n] => 数组转切片
		fmt.Println(string(tmp[:n]))
	}
}
