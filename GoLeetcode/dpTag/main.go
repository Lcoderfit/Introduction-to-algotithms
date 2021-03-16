package main

import (
	"GoLeetcode/dpTag/dp"
	"GoLeetcode/dpTag/medium"
	"fmt"
)

func main() {
	//121. 买卖股票的最佳时机.go
	//maxProfit()

	//221. 最大正方形.go
	MaximalSquare()

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

//221. 最大正方形.go
func MaximalSquare() {
	var m, n int
	for {
		_, err := fmt.Scanln(&m, &n)
		if err != nil {
			fmt.Println(err)
			continue
		}
		matrix := make([][]byte, 0)
		for i := 0; i < m; i++ {
			row := make([]byte, n)
			for j := 0; j < n; j++ {
				_, err := fmt.Scanf("%d", &row[j])
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			matrix = append(matrix, row)
		}
		//ans := medium.MaximalSquare1(matrix)
		ans := medium.MaximalSquare2(matrix)
		fmt.Println(ans)
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
