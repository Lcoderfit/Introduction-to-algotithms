package sortTag

import (
	"fmt"
	"math/rand"
	"time"
)

//产生长度为length，数据范围为n的slice
func RandomArray(length, n int) []int {
	rand.Seed(time.Now().Unix())
	arr := make([]int, 0)
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(n))
	}
	return arr
}

func Print2DArrayForLine(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}