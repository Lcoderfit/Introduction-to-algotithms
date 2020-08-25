package main

import "fmt"
import "GoContest/weekContest-179/ques"

func main() {
	//5352. 生成每种字符都是奇数个的字符串.
	//generateTheString()

	//5353. 灯泡开关 III.go
	numTimesAllBlue()
}

//5353. 灯泡开关 III.go
//2 1 3 5 4
func numTimesAllBlue() {
	var light []int
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scanf("%d", &t)
		light = append(light, t)
	}
	ret := ques.NumTimesAllBlue(light)
	fmt.Println("numTimesAllBlue: ", ret)
}

//5352. 生成每种字符都是奇数个的字符串.
func generateTheString() {
	var n int
	fmt.Scanf("%d\n", &n)
	for {
		if n == -1 {
			break
		}
		ret := ques.GenerateTheString(n)
		fmt.Println("generateTheString: ", ret)
		fmt.Scanf("%d\n", &n)
	}
}