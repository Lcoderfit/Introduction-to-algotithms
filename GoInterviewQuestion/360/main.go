package main

import (
	"GoInterviewQuestion/360/arr"
	"fmt"
)

func main() {
	//最后赢家
	var n, m int
	fmt.Scan(&n, &m)
	var array []int
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		array = append(array, t)
	}
	ret := arr.LastWinner(array, m)
	fmt.Printf("%d", ret)
}