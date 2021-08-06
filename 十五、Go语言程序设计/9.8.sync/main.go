package main

import (
	"fmt"
	"sync"
)

/*
互斥锁：用来控制访问共享资源的方法
sync.Mutex: 如果需要当函数参数传必须传指针，因为Mutex是一个结构体

1.互斥锁：完全互斥，读写都会加锁
2.读写互斥锁：sync.RWMutex  goroutine如果获得读锁，其他goroutine可继续获得读锁（即其他协程任可以读）
	如果获取到写锁，则其他协程无论获取读锁还是写锁均会等待
*/

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add1() {
	for i := 0; i < 50000; i++ {
		// lock.Lock()对后面的程序加锁，遇到lock.Unlock()释放锁
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		// x=x+1其实包含3步操作，第一步取到x的值，第二步将值+1,第三步将结果赋到x
		// 在main中启动两个协程运行add函数，则可能出现两个协程同时取出x的值后，都进行+1操作
		// 然后先后又赋给x，导致本来x应该自增两次 变成了 自增1次
		x = x + 1
	}
}

func main() {
	// wg.Add()不能放到go协程里，因为可能启动协程后，还没运行到wg.Add()这一步主协程就退出了
	wg.Add(2)
	//go add()
	//go add()
	go add1()
	go add1()
	wg.Wait()
	fmt.Println(x)
}
