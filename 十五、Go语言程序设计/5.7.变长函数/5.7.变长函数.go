package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) //类似于Python的拆包
	fmt.Printf("%T\n", sum)

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)

	a()
	fmt.Println("\nThe next is B")
	b()

	c()
}

// 变长函数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func a() {
	p := []int{1, 2, 3, 4}
	for _, v := range p {
		fmt.Println(&v)
		fmt.Println(v)
	}
}

func b() {
	p := []int{1, 2, 3, 4}
	for _, v := range p {
		v := v
		fmt.Println(&v)
		fmt.Println(v)
	}
}

func c() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
}
