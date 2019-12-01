package main

import (
	"fmt"
	"net/url"
)

func main() {
	a()
	b()
	c()

	m := url.Values{}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
}

type Point struct {
	X, Y float64
}

type P struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p P) S(fac float64) {
	p.X *= fac
	p.Y *= fac
}

func a() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	r.ScaleBy(2)
	fmt.Println(*r)

	k := P{1, 2}
	k.S(2)
	fmt.Println(k)
}

func b() {
	r := Point{1, 2}
	pptr := &r
	pptr.ScaleBy(2)
	fmt.Println(r)

	p := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)
}

func c() {
	p := Point{1, 2}
	(&p).ScaleBy(10)
	// 错误写法
	// Point{1, 2}.ScaleBy(2)
	fmt.Println(p)

	q := &P{1, 2}
	// 解引用
	q.S(12)
	fmt.Println(*q)
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
