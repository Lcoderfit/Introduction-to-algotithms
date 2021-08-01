package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
go内置的map不是 并发安全的

当开的协程多了之后：fatal error: concurrent map writes （并发映射写入）

2.并发安全的Map， sync.Map

// 设置key value
m1.Store(key, n)
// Load返回key对应的value，如果不存在，则ok为false
value, _ := m1.Load(key)

m1.Delete() 删除键值对
m1.LoadAndDelete() 删除键值对，并返回值
m1.LoadAndStore
m1.Range() 遍历键值对

3.

*/

var m = make(map[string]int)

// 并发安全的map
var m1 = sync.Map{}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
		}(i)
	}
	wg.Wait()
}

func f() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// 设置key value
			m1.Store(key, n)
			// Load返回key对应的value，如果不存在，则ok为false
			value, _ := m1.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
}
