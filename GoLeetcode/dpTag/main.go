package main

import (
	"GoLeetcode/dpTag/dp"
	"GoLeetcode/dpTag/easy"
	"GoLeetcode/dpTag/medium"
	"fmt"
)

func main() {
	//121. 买卖股票的最佳时机.go
	//maxProfit()

	//198. 打家劫舍.go

	//221. 最大正方形.go
	//MaximalSquare()

	//300. 最长上升子序列.go
	//lengthOfLIS()

	//1143. 最长公共子序列.go
	longestCommonSubsequence()
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
		ret := easy.MaxProfit(prices)
		fmt.Println("maxProfit: ", ret)
	}
}

//198. 打家劫舍.go
func rob() {
	var n int
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println(err)
			continue
		}
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			_, err := fmt.Scanf("%d", &nums[i])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		//ans := medium.Rob1(nums)
		ans := medium.Rob2(nums)
		fmt.Println(ans)
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

//1143. 最长公共子序列.go
func longestCommonSubsequence() {
	var text1, text2 string
	for {
		_, err := fmt.Scanln(&text1, &text2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//ans := medium.LongestCommonSubsequence1(text1, text2)
		ans := medium.LongestCommonSubsequence(text1, text2)
		fmt.Println(ans)
	}
}
