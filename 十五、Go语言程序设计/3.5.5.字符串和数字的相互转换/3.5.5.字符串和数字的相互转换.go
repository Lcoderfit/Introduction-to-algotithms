package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)

	p, _ := strconv.Atoi("123")
	q := strconv.Itoa(123)
	z, _ := strconv.ParseInt("110", 2, 64)
	fmt.Println(p, q, z)
}
