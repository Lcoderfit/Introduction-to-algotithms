package main

import (
	"GoInterviewQuestion/qi_anxin/Random"
	"GoInterviewQuestion/qi_anxin/dp"
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	ret := Random.LukyRandom(n)
	intPart := int(ret*math.Pow(10.0, 6))
	ret = float64(intPart)/float64(math.Pow(10.0, 6))
	fmt.Printf("%0.6f", ret)
}

func countRabbitNum() {
	var n int
	fmt.Scan(&n)
	ret := dp.CountRabbitNum(n)
	fmt.Println(ret)
}