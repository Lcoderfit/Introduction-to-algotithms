package main

import (
	"fmt"
	"sync"
)

/*
sync.Once 让某个操作只执行一次

// done为标志位，m为sync.Mutex 互斥锁
// 如果标志位为1，则不操作，否则先上互斥锁，
// 然后判断标志位是否为0，为0则执行操作，最后将标志位置为1 后释放锁
type Once struct {
	done uint32
	m Mutex
}
*/

var (
	once sync.Once
	x    = 0
)

func f() {
	x = x + 1
}

func f1() {
	once.Do(func() {
		f()
	})
}

func main() {
	for i := 1; i <= 10; i++ {
		go f1()
	}
	fmt.Println(x)
}
