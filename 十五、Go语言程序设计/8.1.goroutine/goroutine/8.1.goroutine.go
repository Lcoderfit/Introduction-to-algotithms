package goroutine

import (
	"fmt"
	"time"
)

func Spinner() {
	go spinner(100*time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//package goroutine
//
//import (
//	"fmt"
//	"time"
//)
//
//func PrintDownLine() {
//	str := `—\|/`
//	for {
//		for _, s := range str {
//			fmt.Printf("%c", s)
//			time.Sleep(300000000)
//			fmt.Printf("\b \b")
//		}
//	}
//}
