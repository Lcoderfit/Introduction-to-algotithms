
package main

import (
	"fmt"
	"time"
)

func main() {
	// 方案一：在同一个位置实现进度打转用英文的"-"
	// 方案二：一般进度打转，一边进度条向右延申，用中文的"—"
	//str := `-\|/`
	// 用于方案二
	str := `—\|/`
	for {
		for _, s := range str {
			fmt.Printf("%c", s)
			time.Sleep(3e8)
			fmt.Printf("\b \b")
		}
	}
}