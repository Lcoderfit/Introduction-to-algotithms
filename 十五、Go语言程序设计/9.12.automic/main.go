package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
原子级别操作：atomic
https://www.liwenzhou.com/posts/Go/14_concurrence/#autoid-1-7-3

通过(*int32)(unsafe.Pointer(&y))将*int转换为*int32
*/

var (
	x  int32
	y  int
	wg sync.WaitGroup
)

func f() {
	x = x + 1
	defer wg.Done()
}

func f2() {
	defer wg.Done()
	// 通过(*int32)(unsafe.Pointer(&y))将*int转换为*int32
	atomic.AddInt32((*int32)(unsafe.Pointer(&y)), 1)
}

func main() {
	for i := 0; i < 2000; i++ {
		wg.Add(2)
		go f()
		go f()
	}

	for i := 0; i < 2000; i++ {
		wg.Add(2)
		go f2()
		go f2()
	}

	wg.Wait()
	fmt.Println(x)
	fmt.Println(y)
}
