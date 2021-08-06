package main

import (
	"fmt"
	"sync"
	"time"
)

/*
读锁，其他协程仍可读但不可写
写锁，其他协程不可读也不可写

用互斥锁：2.836473s
用读写锁：1.0063103s
*/

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	//lock.Lock()
	// 读取数据，只需要设置读锁
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	//lock.Lock()
	// 写入数据，需要设置读写锁
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	//lock.Unlock()
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
