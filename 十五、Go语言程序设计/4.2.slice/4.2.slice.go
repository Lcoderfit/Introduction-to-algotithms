package main

import (
	"fmt"
)

func main() {
	a := [...]string{1: "a", 2: "b", 3: "c"}
	for i, s := range a {
		fmt.Println(i, s)
	}

	b := [...]string{"b", "c", "d"}

	for _, p := range a {
		for _, q := range b {
			if p == q {
				fmt.Printf("%s appears in both\n", p)
			}
		}
	}

	c := [...]int{0, 1, 2, 3, 4, 5}
	reverse(c[:])
	fmt.Println(c)

	revern(3, c[:])
	fmt.Println(c)

	x := [...]string{"a", "b", "c"}
	y := [...]string{"a", "b", "c"}
	z := [...]string{"a", "b", "a"}
	fmt.Println(equal(x[:], y[:]), equal(y[:], z[:]))

	NL()

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func revern(n int, a []int) {
	length := len(a)
	n = n % length
	reverse(a[:n])
	reverse(a[n:])
	// 由于传入的是c[:], 已经是一个slice了，所以reverse(a)就相当于reverse(c[:])
	reverse(a)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func NL() {
	var a []int
	fmt.Println(len(a), cap(a), a == nil)
	a = nil
	c := []int(nil)
	d := []int{}
	fmt.Println(len(a), cap(a), a == nil)
	fmt.Println(len(c), cap(c), c == nil)
	fmt.Println(len(d), cap(d), d == nil)

	e := make([]int, 3)[3:]
	f := make([]int, 3)[:2
	]
	fmt.Println(len(e), cap(e), e == nil)
	fmt.Println(len(f), cap(f), f == nil)
}
