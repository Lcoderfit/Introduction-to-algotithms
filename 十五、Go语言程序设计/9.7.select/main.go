package main

import "fmt"

/*
select可以同时响应多个通道操作，每个case对应一个通道接收或发送操作
select会一直等待，直到某一个通道操作完成，则执行对应的case语句，
如果有多个case满足条件，则随机执行一个,
如果所有条件都不满足则执行default

select {
case <-ch1:
	...
case data := <-ch2:
	...
default:
	默认操作
}
*/

func main() {
	ch := make(chan int, 1)
	// i为0时，读阻塞，写不阻塞，则执行写操作（ch<-i）
	// 当i为1时，写阻塞（缓冲区满），读不阻塞，执行读操作(x:=<-ch)
	// 所以只会打印偶数
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println(i, "发送成功")
		}
	}
}
