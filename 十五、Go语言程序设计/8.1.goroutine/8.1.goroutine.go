// package main

// import (
// 	"fmt"
// 	// "time"
// )

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	const n = 45
// 	fibN := fib(n)
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("%c\b ", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	str := `—\|/`
	for {
		for _, s := range str {
			fmt.Printf("%c", s)
			time.Sleep(300000000)
			fmt.Printf("\b \b")
		}
	}
}
