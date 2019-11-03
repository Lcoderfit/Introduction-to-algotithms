package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	f(s)

	n := 0
	for _, _ = range s {
		n++
	}

	n = 0
	for range s {
		n++
	}
	fmt.Println(n)

	f(s)
	frange(s)

	fmt.Println(string(65))
}

func f(s string) {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

func frange(s string) {
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
