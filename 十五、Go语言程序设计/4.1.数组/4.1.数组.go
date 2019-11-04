package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d", i, v)
	}

	d := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(d)

	b := d[:4]
	c := b[:6]
	fmt.Println(b, c)

	f()
	g()
	k()

	e := []byte("x")
	f := []byte("X")
	sh(e, f)

	h := [...]int{1, 3, 5, 7, 9}
	j := ch(h)
	fmt.Println(j)
	fmt.Println(h)

	chvalue(&h)
	fmt.Println(h)

	zero(&h)
	fmt.Println(h)
}

func f() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], q[2])

	a := [...]int{1, 2, 3}
	fmt.Printf("%T\n", a)
}

func g() {
	a := [...]string{"a", "b", "c"}
	b := [...]int{99: -1}
	fmt.Println(a)
	fmt.Println(b)
}

func k() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	// d := [3]int{1, 2}
	//fmt.Println(a == d) //编译错误
}

func sh(a []byte, b []byte) {
	c1 := sha256.Sum256(a)
	c2 := sha256.Sum256(b)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func ch(a [5]int) [5]int {
	for i, _ := range a {
		a[i] = 1
	}
	return a
}

func chvalue(a *[5]int) {
	for i := range a {
		a[i] = 1
	}
}

func zero(ptr *[5]int) {
	*ptr = [5]int{}
}
