package main

import (
	"GoLeetcode/binaryFindTag/binaryFind"
	"fmt"
)

func main() {
	//35. 搜索插入位置.go
	//searchInsert()

	//392. 判断子序列.go
	//isSubsequence()

	//875. 爱吃香蕉的珂珂.go
	//minEatingSpeed()

	//1064. 不动点.go
	fixedPoint()

	//1351. 统计有序矩阵中的负数.go
	countNegatives()
}

//35. 搜索插入位置.go
func searchInsert() {
	var target, n int
	var nums [] int
	for {
		fmt.Scanln(&target, &n)
		var t int
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &t)
			nums = append(nums, t)
		}
		//fmt.Scanln(&line)
		//for _, s := range strings.Split(line, " ") {
		//	if len(strings.Trim(s, " ")) == 0 {
		//		continue
		//	}
		//	i, _ := strconv.Atoi(s)
		//	nums = append(nums, i)
		//}

		ret := binaryFind.SearchInsert(nums, target)
		fmt.Println(ret)
	}
}

//392. 判断子序列.go
func isSubsequence() {
	var s, t string
	for {
		fmt.Scan(&s, &t)
		ret := binaryFind.IsSubsequence(s, t)
		fmt.Println("isSubsequence: ", ret)
	}
}

//875. 爱吃香蕉的珂珂.go
func minEatingSpeed() {
	var n int
	var h int
	for {
		fmt.Scanln(&n, &h)
		piles := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &piles[i])
		}
		fmt.Println(piles)

		ret := binaryFind.MinEatingSpeed(piles, h)
		fmt.Println("minEatingSpeed: ", ret)
	}
}

//1064. 不动点.go
func fixedPoint() {
	var n int
	for {
		fmt.Scanln(&n)
		piles := make([]int, 0)
		var t int
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &t)
			piles = append(piles, t)
		}
		res := binaryFind.FixedPoint(piles)
		fmt.Println("fixedPoint: ", res)
	}
}

//1351. 统计有序矩阵中的负数.go
func countNegatives() {
	var m, n int
	for {
		fmt.Scanln(&m, &n)
		grid := make([][]int, 0)
		for i := 0; i < m; i++ {
			piles := make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scanf("%d", &piles[j])
			}
			grid = append(grid, piles)
		}
		res := binaryFind.CountNegatives(grid)
		fmt.Println("countNegatives: ", res)
	}
}