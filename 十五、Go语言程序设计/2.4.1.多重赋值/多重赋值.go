package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, y := 3, 4
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		array := strings.Split(input.Text(), " ")
		p, q := array[0], array[1]
		a := make(strconv.Atoi(p))[0:1]
		b, _ := strconv.Atoi(q)

		res1 := gcd1(a, b)
		res2 := gcd2(a, b)
		fmt.Println(res1, res2)
	}
}

//计算两个数的最大公约数
func gcd1(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func gcd2(x, y int) int {
	for y != 0 {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return x
}

//斐波那契堆
// func fib()
