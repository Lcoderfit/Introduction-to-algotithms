package main

import (
	"GoLeetcode/iQTag/iQ"
	"fmt"
	"runtime"
)

func main() {
	//1227. 飞机座位分配概率.go
	//nthPersonGetsNthSeat()
}

//1227. 飞机座位分配概率.go
func nthPersonGetsNthSeat() {
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()
	fmt.Println(fn, ": starting")

	var n int
	k, _ := fmt.Scanf("%d\n", &n)
	for k != 0{
		ret := iQ.NthPersonGetsNthSeat(n)
		fmt.Println(ret)
		k, _ = fmt.Scanf("%d\n", &n)
	}

	fmt.Println(fn, ": End")
}
