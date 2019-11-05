package main

import (
	"fmt"
)

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	fmt.Println(p, p.X, p.Y)

	a := &Point{1, 2}
	b := new(Point)
	c := Point{1, 2}
	fmt.Println(a, b, c)
	a.X = 4
	b.Y = 4
	fmt.Println(a, b)

	d := Point{1, 2}
	e := Point{3, 4}
	fmt.Println(d.X == e.X && d.Y == e.Y)
	fmt.Println(d == e)

	k()
}

func k() {
	type address struct {
		hostname string
		port     int
	}

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
