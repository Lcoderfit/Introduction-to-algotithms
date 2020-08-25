package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "lcoder fit"
	s := strings.ToTitle(a)
	k := strings.ToUpper(a)
	t := strings.Title(a)
	fmt.Println(s, k, t)
}