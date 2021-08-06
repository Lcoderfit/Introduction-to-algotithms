package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 程序启动之后会创建一个主goroutine
func main() {
	//goroutine.LockTest()

	//9.1.goroutine.go
	//goroutine.Spinner()

	//8.4.2.管道.go
	//goroutine.Pipeline()

	// 比操作系统的线程更轻量，用户态线程；顺序通信进程(CSP模型)
	// 开启一个单独的goroutine去执行hello函数
	go hello()
	fmt.Println("main") // 如果main先结束了，由main函数启动的goroutine也都结束了

	// 并发打印出来会有重复值，因为假设当i的值位k时，启动了一个goroutine，如果循环运行的够快，在该goroutine还没开始执行到
	// 打印语句时，i就自增了，所以本应该打印k结果却打印了k+1(而本应该打印k+1的那个协程可能刚好在i自增前就执行了，所以也打印出k+1);
	// 因此会打印出重复值
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//
	//}

	//for i := 0; i < 1000; i++ {
	//	// 通过传参的方式将i的值传给匿名函数，这样每次fmt.Println操作的就是当前循环的i的一个副本
	//	// goroutine对应的函数执行结束了，goroutine就结束了
	//	// main函数执行结束了，由main创建的其他goroutine也就结束了
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}

	for i := 0; i < 10; i++ {
		// wg计数器+1， 调用wg.Done()会使wg计数器-1
		wg.Add(1)
		go f1(i)
	}
	// 等待wg的计数器减为0
	wg.Wait()
}

// 注意，不能copyWaitGroup是一个值类型，如果copy了那就不是原来那个wg了
var wg sync.WaitGroup

func f() {
	// 设置随机数种子，如果不设置，则每次跑出来的结果都是一样的
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		// 生成一个int64的随机数
		r1 := rand.Int()
		// 生成一个int类型的随机数
		r2 := rand.Intn(10)
		fmt.Println(0-r1, 0-r2)
	}
}

func f1(i int) {
	// wg计数器减去1
	defer wg.Done()
	fmt.Println(i)
}

func hello() {
	fmt.Println("hello")
}
