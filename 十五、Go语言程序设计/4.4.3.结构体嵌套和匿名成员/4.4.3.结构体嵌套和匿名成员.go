package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	a()
	b()

	c()

	d()
}

func a() {

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w)
	fmt.Printf("%#v\n", w)
}

func b() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	a := Circle{Point{1, 2}, 5}
	fmt.Println(w, a)
}

func c() {
	w := Wheel{
		Circle: Circle{
			Center: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}

type Circle1 struct {
	Point
	Radius int
}

type Wheel1 struct {
	Circle1
	Spokes int
}

func d() {
	w := Wheel1{
		Circle1: Circle1{
			Point:  Point{X: 1, Y: 2},
			Radius: 5,
		},
		Spokes: 10,
	}
	fmt.Println(w)
}
