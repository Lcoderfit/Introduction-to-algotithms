package main

import (
	"fmt"
)

func main() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	a := "Lcoderfit"
	b := "Lcoder"
	c := "fit"
	d := "derf"
	fmt.Println(HasPrefix(a, b))
	fmt.Println(HasSuffix(a, c))
	fmt.Println(Contains(a, d))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
