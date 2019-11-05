package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	a()

	a := map[string]int{
		"a": 1,
		"b": 2,
	}
	b(a)

	fmt.Println("C:::::")
	c(a)

	age, ok := a["d"]
	if !ok {
		fmt.Println("d 不是字典中的键")
	} else {
		fmt.Println(age)
	}

	x := map[string]int{
		"a": 1,
		"b": 2,
	}
	y := map[string]int{
		"c": 3,
		"d": 4,
	}
	fmt.Println(equal(x, y))

	// d()
	from, to := "a", "b"
	addEdge(from, to)
	fmt.Println(hasEdge(from, to))
}

func a() {
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2
	b := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(a, b)

	fmt.Println(a["c"])

	c := map[string]int{}
	var d map[string]int
	e := make(map[string]int)
	fmt.Println(c == nil, d == nil, c["a"], e == nil)

	delete(a, "a") // 移除元素
	fmt.Println(a)

	a["b"] += 1
	a["c"]++
	fmt.Println(a)
}

func b(a map[string]int) {
	for name, age := range a {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func c(a map[string]int) {
	var names []string
	for name := range a {
		fmt.Println(name)
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, a[name])
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func d() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
