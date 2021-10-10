//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	s, sep := "", ""
//	for _, arg := range os.Args[1:] {
//		s += sep + arg
//		sep = " "
//	}
//	fmt.Println(s)
//	s := ""
//	var s1 string
//	var s = ""
//	var s string
//}

//package main
//
//import (
//	"fmt"
//	"os"
//	"strings"
//)
//
//func main() {
//	fmt.Println(strings.Join(os.Args[1:], " "))
//	fmt.Println(os.Args[1:])
//}

// 练习题1.2
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
