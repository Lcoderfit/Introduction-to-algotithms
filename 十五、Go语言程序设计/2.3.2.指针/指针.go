// package main

// import "fmt"

// func main() {
// 	z := 1
// 	p := &z
// 	fmt.Println(*p)
// 	*p = 2
// 	fmt.Println(z)

// 	var x, y int
// 	fmt.Println(&x == &x, &x == &y, &x == nil)

// 	var k = f()
// 	fmt.Println(f() == f(), k)

// 	v := 1
// 	incr(&v)
// 	fmt.Println(incr(&v))
// }

// func f() *int {
// 	v := 1
// 	return &v
// }

// func incr(p *int) int {
// 	*p++
// 	return *p
// }

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")

// func main() {
// 	flag.Parse()
// 	fmt.Print(strings.Join(flag.Args(), *sep))
// 	if !*n {
// 		fmt.Println()
// 	}
// }
