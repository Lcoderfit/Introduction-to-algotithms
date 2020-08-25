package main

import (
	"fmt"
)

func main() {
	//base-蓄水池采样算法.go
	a := make([]int, 0)
	fmt.Printf("%p\n", &a)
	test(a)
	fmt.Println(a)
}

func test(ret []int) {
	fmt.Printf("%p\n", &ret)
	for i := 0; i < 10; i++ {
		ret = append(ret, i)
	}
}
