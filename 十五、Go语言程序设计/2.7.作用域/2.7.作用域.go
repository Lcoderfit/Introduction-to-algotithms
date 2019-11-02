package main

import "fmt"

//版本一
// func main() {
// 	x := "hello!"
// 	for i := 0; i < len(x); i++ {
// 		x := x[i]
// 		if x != '!' {
// 			x := x + 'A' - 'a'
// 			fmt.Printf("%c", x)
// 		}
// 	}
// 	fmt.Println()
// }

//版本二
// func main() {
// 	x := "hello"
// 	for _, x := range x {
// 		x := x + 'A' - 'a'
// 		fmt.Printf("%c", x)
// 	}
// 	fmt.Println()
// }
