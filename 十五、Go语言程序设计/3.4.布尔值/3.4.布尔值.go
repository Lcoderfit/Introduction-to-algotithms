package main

import (
	"fmt"
)

func main() {
	i := 0
	b := true
	if b {
		i = 1
	}
	fmt.Println(i)

	fmt.Println(btoi(b), itob(10))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
