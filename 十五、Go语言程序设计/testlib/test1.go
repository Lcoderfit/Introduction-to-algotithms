package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "&I&love&coding"
	res := strings.Split(str, "&")
	fmt.Println("result:", res)
	fmt.Println("len:", len(res))
	fmt.Println("cap:", cap(res))
	fmt.Printf("%T", res)
	for _, s := range res {
		fmt.Println(s)
	}
}
