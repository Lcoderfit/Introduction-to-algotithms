package goroutine

import (
	"fmt"
	"sync"
	"time"
)

//var set = make(map[int]bool, 0)
//
//func printOnce(num int) {
//	if _, exist := set[num]; !exist {
//		fmt.Println(num)
//	}
//	set[num] = true
//}
//
//func main() {
//	for i := 0; i < 10; i++ {
//		go printOnce(100)
//	}
//	time.Sleep(time.Second)
//}

var set = make(map[int]bool, 0)
var m sync.Mutex

//func printOnce(num int) {
//	m.Lock()
//	if _, exist := set[num]; !exist {
//		fmt.Println(num)
//	}
//	set[num] = true
//	m.Unlock()
//}

//另外一种写法
func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

func LockTest() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}