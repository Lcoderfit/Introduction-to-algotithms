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
}
