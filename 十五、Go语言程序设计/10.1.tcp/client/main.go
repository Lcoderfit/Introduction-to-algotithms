package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connect to server error:", err)
		return
	}

	// 只能发送一次
	//_, err = conn.Write([]byte("hello golang http server"))
	//if err != nil {
	//	fmt.Println("send data to server error:", err)
	//	return
	//}

	// 发送多次
	// 从标准输入读取数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请说话: ")
		// 读到换行
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		// 退出
		if msg == "exit" {
			break
		}
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write data error:", err)
			return
		}

	}
}
