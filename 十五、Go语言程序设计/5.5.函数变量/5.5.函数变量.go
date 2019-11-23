package main

import (
	"fmt"
	"strings"
)

func main() {
	f := a
	fmt.Println("f(3):", f(3))

	f = b
	fmt.Println(f(3))
	fmt.Printf("%T", f)
	// 函数的零值为nil，调用空函数会宕机
	// var k func(int) int
	// fmt.Println(k(3))

	var p func(int) int
	if p != nil {
		p(3)
	}
	fmt.Println()
	d()
}

func a(n int) int {
	return n * n
}

func b(n int) int {
	return -n
}

func c(m, n int) int {
	return m * n
}

func add1(r rune) rune {
	return r + 1
}

func d() {
	ret := strings.Map(add1, "ABCDE")
	fmt.Println(ret)
}
