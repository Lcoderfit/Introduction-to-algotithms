package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	x, y = 3, 4
	a := a(x, y)
	fmt.Println(a)
}

func a(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
