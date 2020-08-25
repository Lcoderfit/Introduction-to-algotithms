package main

import (
	"GoLeetcode/hashTableTag/hashTable"
	"fmt"
	"runtime"
)

func main() {
	//1365. 有多少小于当前数字的数字.go
	smallNumberThanCurrent()
}

//1365. 有多少小于当前数字的数字.go
func smallNumberThanCurrent() {
	nums := []int{8, 1, 2, 2, 3}
	ret := hashTable.SmallNumberThanCurrent(nums)

	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()
	fmt.Println(fn, " ret: ", ret)
}
