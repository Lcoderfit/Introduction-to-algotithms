package main

import (
	"GoOffer/algorithm"
	"fmt"
)

func main() {
	// 10-2.青蛙跳台阶
	jumpFloor()
}

// 10-2.青蛙跳台阶
func jumpFloor() {
	var n int64
	var res int
	fmt.Println("/********begin**********/")
	for {
		fmt.Scanf("%d\n", &n)
		res = algorithm.JumpFloor(n)
		fmt.Printf("%d级台阶跳法有：%d种\n", n, res)
	}
}