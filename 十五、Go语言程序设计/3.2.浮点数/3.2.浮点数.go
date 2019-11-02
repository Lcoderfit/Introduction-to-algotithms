package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216
	var p float32 = 99999999
	fmt.Println(f == f+1)
	fmt.Println(p == p+1)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	fmt.Printf("%f, %f\n", math.MaxFloat32, math.MaxFloat64)
	fmt.Printf("%d\n", 1<<32)

	k()
	g()
}

func k() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func g() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	fmt.Println(math.Inf, math.NaN())

	nan := math.NaN()
	fmt.Println(nan != nan, nan == nan, nan > nan)
}
