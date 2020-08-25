package main

import (
	"GoLeetcode/bfsAndDfsTag/bfsAndDfs"
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//329. 矩阵中的最长递增路径.go
	//longestIncreasingPath()

	//994. 腐烂的橘子.go
	//orangesRotting()
	a := "3.1"
	b := "32.     "
	fmt.Println(strings.Split(a, "."))
	fmt.Println(strings.Split(b, "."))
	fmt.Println(strings.Trim(b, " "), "asdjf", "nasdkf")
	c := "2 "

	//字符串可以相加减
	d := byte('1')
	d = d + 1
	fmt.Println(d)
	//对slice排序
	e := []int{4, 2, 3, 1}
	sort.Ints(e)
	fmt.Println(e)

	//d := []byte("123")
	//d = append(d, []byte("123")...)
	fmt.Println(strconv.Atoi(c))
	fmt.Println(strconv.ParseFloat(c, 32))
}

//329. 矩阵中的最长递增路径.go
func longestIncreasingPath() {
	matrix := [][]int{
		{9,9,4},
		{6,6,8},
		{2,1,1},
	}
	ret := bfsAndDfs.LongestIncreasingPath(matrix)
	fmt.Println("拓扑排序法：", ret)

	ret1 := bfsAndDfs.LongestIncreasingPath1(matrix)
	fmt.Println("记忆化+DFS: ", ret1)
}

//994. 腐烂的橘子.go
func orangesRotting() {
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()
	fmt.Println("Starting: ", fn)
	gird := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	ret := bfsAndDfs.OrangesRotting(gird)
	fmt.Println(ret)
	fmt.Println("End: ", fn)
}
