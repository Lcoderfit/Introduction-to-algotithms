// package io

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Closer interface {
// 	Close() error
// }

// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// // 嵌入式接口
// type ReadWriteCloser interface {
// 	Reader
// 	Writer
// 	Closer
// }

// type ReadWriter interface {
// 	Read(p []byte) (n int, err error)
// 	Write(p []byte) (n int, err error)
// }

// // 混用多种形式
// type ReadWriter interface {
// 	Read(p []byte) (n int, err error)
// 	Writer
// }

// var any interface{}
// any = true
// any = 12.34
// any = "hello"
// any = map[string]int{"one", 1}
// any = new(bytes.Buffer
// )

package main

import (
	"fmt"
	"time"
)

func main() {
	var a = `123` + `321`
	fmt.Println(a)
	var b = '1' + '2'
	fmt.Println(b)

	fmt.Println(time.Now().Unix())

	c()
}

func c() {
	ch := make(chan int, 1)
	fmt.Println("begin:", ch)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("i:", i, "    x:", x)
		case ch <- i:
			fmt.Println(ch, "Lcoderfit")
		}
	}
}
