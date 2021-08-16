package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
1.为什么需要context
如何优雅的控制goroutine

2.在select中使用break默认是跳出select语句，但是不会跳出包含select的for循环
可以使用掉标签的break，go中的break标签指的是跳出标签所指定的部分，执行之后的操作

跳出select语句的三种办法
1.使用带标签的break
2.return
3.使用flag
*/

var wg sync.WaitGroup
var notify bool
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
	for {
		fmt.Println("周玲")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func f1() {
	defer wg.Done()
	// 使用带标签的break跳出forLoop所指定的部分，执行后续操作
forLoop:
	for {
		fmt.Println("周林")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			fmt.Println("run case<-exitChan, break")
			break forLoop
		default:
			fmt.Println("run default")
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
	// ctx一级一级往下传，在goroutine中再启一个子goroutine，则只需要把当前
	// goroutine中的context传到子goroutine启动的函数中，调用ctx.Done()即可
	go f3(ctx)
forLoop:
	for {
		fmt.Println("周林2")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break forLoop
		default:
		}
	}
}

func f3(ctx context.Context) {
	defer wg.Done()
forLoop:
	for {
		fmt.Println("周林3")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break forLoop
		default:

		}
	}
}

func main() {
	//使用Go语言官方推荐的退出goroutine的方法：context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	//go f()
	//go f1()
	go f2(ctx)
	//time.Sleep(time.Second * 5)
	//notify = true
	//exitChan <- true
	//fmt.Println("exitChan is true")

	// 调用cancel通知所有子goroutine退出
	cancel()
	wg.Wait()
	// 如何通知子goroutine退出
}
