package main

import (
	"GoLeetcode/dpTab/dp"
	"fmt"
)

func main() {
	//121. 买卖股票的最佳时机.go
	//maxProfit()

	//300. 最长上升子序列.go
	//lengthOfLIS()
}

//121. 买卖股票的最佳时机.go
func maxProfit() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		prices := make([]int, n)
		for i := 0; i < n; i++ {
			var t int
			fmt.Scanf("%d", &t)
			prices[i] = t
		}
		ret := dp.MaxProfit(prices)
		fmt.Println("maxProfit: ", ret)
	}
}

//300. 最长上升子序列.go
func lengthOfLIS() {
	var n int
	var t int
	for {
		fmt.Scanf("%d", &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &t)
			arr[i] = t
		}
		//ret := dp.LengthOfLIS(arr)
		ret := dp.LengthOfLIS1(arr)
		fmt.Println("lengthOfLIS: ", ret)
	}
}