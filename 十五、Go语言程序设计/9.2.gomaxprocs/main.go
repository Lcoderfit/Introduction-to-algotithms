package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 默认为物理线程数（CPU核数,即跑满整个CPU）
	// 一般不用设置，用默认设置即可
	runtime.GOMAXPROCS(8)
	// 返回计算机CPU核数
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	// 返回当前Goroutine数量，main为主协程（即至少有一个协程）
	//
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
