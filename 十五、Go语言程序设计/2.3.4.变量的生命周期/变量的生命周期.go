package main

import "fmt"

var global *int

func main() {
	f()
	g()
	fmt.Println(*global)

	for i := 1; i < 10; i++ {
		x := i
		y := 2 * i
		fmt.Println(x, y)
	}
	// fmt.Println(x, y)
}

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
