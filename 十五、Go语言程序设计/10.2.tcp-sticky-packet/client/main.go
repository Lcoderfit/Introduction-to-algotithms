package main

import (
	"fmt"
	"net"
	"stickypacket/proto"
)

func main() {
	// Dial拨号，即对ip：port发起请求
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// 需要在err之前defer，否则如果报错了，则不会运行到defer那一步
	defer conn.Close()
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	for i := 0; i < 20; i++ {
		msg := "hello, hello, how are you?"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode data error:", err)
			return
		}
		conn.Write(data)
	}
}


