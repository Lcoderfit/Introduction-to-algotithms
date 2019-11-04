package main

import (
	"fmt"
	"math"
	"time"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
)

func main() {
	const pi = 3.12
	fmt.Println(pi, math.Pi)

	const (
		a = 1.2
		b = 3.4
	)

	const (
		k = 1
		g
		y = 3
		c
	)
	fmt.Println(k, g, y, c)

	f()

	type weekday int
	const (
		sunday = iota
		m
		t
		w
		th
		fri
		sat
	)
	fmt.Println(sunday, m, t, w, th, fri, sat)
	fmt.Println(math.Pi, float64(math.Pi), float32(math.Pi))
}

func f() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

}
