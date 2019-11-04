package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "", "b", "c"}
	s := nonempty(a)
	fmt.Println(s)

	b := []string{"a", "b", "", "c"}
	s = nonempty2(b)
	fmt.Println(s)

	m := []int{1, 2, 3, 4}
	c := remove(m, 2)
	fmt.Println(c)

	c = remove(m, 0)
	fmt.Println(c)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// 移除下标为i位置的元素
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove1(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
