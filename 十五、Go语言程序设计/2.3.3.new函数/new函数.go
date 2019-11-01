package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 4
	fmt.Println(*p)

	a := newInt()
	b := newInt1()
	fmt.Println(a == b)

	c := ""
	d := ""

	e := new(string)
	f := new(string)
	*e = c
	*f = d
	fmt.Println(c == d)
}

func newInt() *int {
	return new(int)
}

func newInt1() *int {
	var dummy int
	return &dummy
}
