package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"stickypacket/proto"
)

/*
TCP粘包：指发送方发送的多个包粘在一起形成一个包（TCP是发送数据是流式发送的，不知道一个包的开头和结尾）
所以出现一个包的头部紧挨着另一个包的尾部发送

2.导致粘包的主要原因：tcp数据传递模式是流模式，保持长连接的时候可以进行多次收和发

2.1 发送端粘包：Nagle算法（改善网络传输效率的方法，在未收到确认数据已发送的ACK时，
发送器会将数据放到缓存里，直到收到ACK或达到小于MSS的最大值时才发送）

为什么Nagle能提高效率：因为如果一个数据包包括首部和实体，
例如如果我发送两个实体为10字节的数据包，由于每个数据包首部为40，
所以两个数据包首部耗费了80字节；
如果我把两个包合成一个，则发送同样的实体数据，

2.2 接收端粘包：接收端会把接收到的数据放入缓存，然后通知应用层取数据，如果
由于某些原因不能及时把数据取出来，则会造成TCP缓冲区中存放了几段数据


mspaint: 打开画图功能


*********************************************************************
什么情况下传递一个指针？
1.需要修改值的情况下
2.如果实现接口的方法的接收器是指针类型，则传参给接口类型时必须是指针，如果不是，则传不传指针都无所谓

bufio.NewReader和bytes.NewReader的区别

1.只要实现了Read方法，就实现了io.Reader接口
type Reader interface {
	Read(p []byte) (n int, err error)
}

2.bufio.NewReader接收一个io.Reader接口类型作为参数，然后返回一个默认缓冲区大小为4096字节的读取器(Reader)
bytes.NewReader 接收一个[]byte字节数组作为参数，返回一个读取器，并以该字节数组的内容作为该读取器缓冲区中的内容

总结：bytes包主要涉及一些bytes操作，很多函数参数都是字节数组
bufio主要涉及一个写io读写操作，很多函数参数都是io.Reader类型或者io.Writer

*/


func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}
		// 为什么这里不传一个指针？因为如果传递指针有被修改的风险
		go processConn(conn)
	}
}

// 发送数据
func processConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		// end of file: 即表示读取到文件尾(读完了)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg error:", err)
			return
		}
		fmt.Println("收到来自client的数据：", msg)
	}

}
