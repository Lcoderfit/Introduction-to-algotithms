package main

import (
	"fmt"
	"math"
)

func main() {
	a()
}

type Point struct {
	X, Y float64
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func a() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))
}

func b() {
	p := Point{1, 2}
	q := Point{3, 4}
	distance := (*Point).Distance
	fmt.Println(distance(&p, q))
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (q Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p q Point) Point
	if add {
		op = Point.Add 
	} else {
		op = Point.Sub 
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}