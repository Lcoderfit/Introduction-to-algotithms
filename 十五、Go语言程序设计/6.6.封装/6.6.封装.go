package main

import (
	"fmt"
)

func main() {
	a := Counter{1}
	a.Increment()
	a.Increment()
	res := a.N()
	fmt.Println(res)
	a.Reset()
	res = a.N()
	fmt.Println(res)

	p := person{age: 10}
	p.addAge1()
	fmt.Println(p.age)
}

// type Buffer struct {
// 	buf     []byte
// 	initial [64]byte
// }

// func (b *Buffer) Grow(n int) {
// 	if b.buf == nil {
// 		b.buf = b.initial[:0] // 最初使用预分配的空间
// 	}
// 	if len(b.buf)+n > cap(b.buf) {
// 		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
// 		copy(buf, b.buf)
// 		b.buf = buf
// 	}
// }

type person struct {
	age int
}

func (p person) addAge() {
	p.age++
}

func (p *person) addAge1() {
	p.age++
}

func newPerson(age int) *person {
	return &person{
		age: age,
	}
}

type Counter struct {
	n int
}

func (c *Counter) N() int {
	return c.n
}

func (c *Counter) Increment() {
	c.n++
}

func (c *Counter) Reset() {
	c.n = 0
}
