package main

import (
	"fmt"
	"net"
	"stickypacket/proto"
)

func main() {
	// Dial拨号，即对ip：port发起请求
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// 需要在err之后defer，因为如果报错，则返回的conn可能为一个空值，该空置
	// 可能没有Close方法，所以需要先判断err，在defer
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

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
