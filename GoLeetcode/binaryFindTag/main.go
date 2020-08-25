package main

import (
	"GoLeetcode/binaryFindTag/binaryFind"
	"fmt"
)

func main() {
	//875. 爱吃香蕉的珂珂.go
	minEatingSpeed()
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