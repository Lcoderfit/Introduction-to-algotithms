package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var a Employee
	fmt.Println(a)

	a.Salary -= 500
	position := &a.Position
	*position = "Lcoder" + "fit"
	fmt.Println(a)

	var b *Employee = &a
	b.Position += "This is Fuck"
	(*b).Address = "ECJTU"
	fmt.Println(*b)

	values := [...]int{1, 2, 3, 4, 5}
	Sort(values[:])
	fmt.Println(values)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// func Sort(values []int) {
// 	var root *tree
// 	for _, v := range values {
// 		root = add(root, v)
// 	}
// 	appendValues(values[:0], root)
// }

// func appendValues(values []int, t *tree) []int {
// 	if t != nil {
// 		values = appendValues(values, t.left)
// 		values = appendValues(values, t.value)
// 		values = appendValues(values, t.right)
// 	}
// 	return values
// }

// func add(t *tree, value int) *tree {
// 	if t == nil {
// 		t = new(tree)
// 		t.value = value
// 		return t
// 	}
// 	if value < t.value {
// 		t.left = add(t.left, value)
// 	} else {
// 		t.right = add(t.right, value)
// 	}
// 	return t
// }
