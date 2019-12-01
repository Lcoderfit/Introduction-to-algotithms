package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	a()
	b()
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func a() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
}

func b() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
}

type ColoredPoint1 struct {
	*Point
	Color color.RGBA
}

func c() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint1{&Point{1, 2}, red}
	q := ColoredPoint1{&Point{1, 2}, blue}
	fmt.Println(p.Distance(*q.Point)) // 相当于*(q.Point)
	q.Point = p.Point                 // p和q共享一个Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}

type ColoredPoint1 struct {
	Point
	color.RGBA
}

var (
	mu      sync.Mutex
	mapping = make(map[string][]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

var cache = struct {
	sync.Mutex
	mapping map[string][]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
