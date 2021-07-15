package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a()

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	d()
}

func a() {
	ret := strings.Map(func(r rune) rune { return r + 1 }, "ABCD")
	fmt.Println(ret)
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func d() {
	var prereqs = map[string][]string{
		"a": {"b", "d"},
		"b": {"c"},
		"d": {"c"},
	}
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
