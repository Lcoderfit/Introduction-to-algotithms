// 版本1
// package main

// import "fmt"

// const boilingF = 212.0

// func main() {
// 	var f = boilingF
// 	var c = (f - 32) * 5 / 9
// 	fmt.Printf("%gF or %gc\n", f, c)
// }

//版本二
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
	x := 10
	x++
	fmt.Println(x)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
