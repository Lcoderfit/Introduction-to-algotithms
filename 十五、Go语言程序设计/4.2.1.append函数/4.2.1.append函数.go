package main

import (
	"fmt"
)

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Println(runes)

	a := [...]int{1, 2, 3}
	b := appendInt(a[:], 4)
	fmt.Println(len(a), cap(a))
	fmt.Println(a)
	fmt.Println(b, len(b), cap(b))

	g()

	k()
	n := [...]int{1, 2, 3}
	m := appendInt1(n[:], 4, 5, 6, 7)
	fmt.Println(m)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func g() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func k() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)
}

func appendInt1(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
