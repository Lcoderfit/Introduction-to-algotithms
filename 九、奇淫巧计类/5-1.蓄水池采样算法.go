package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var n, k int
	// for condition: condition只能返回一个值，如果返回两个值，会报-》use as value错误
	for {
		// fmt.Scanf("%d%d\n", &n, &k)： win下的\r\n会导致输出两次
		stdin := bufio.NewReader(os.Stdin)
		fmt.Fscanf(stdin, "%d%d", &n, &k)
		if k > n {
			fmt.Println("ERROR INPUT!!!")
			continue
		}
		array := make([]int, n)
		for i := 0; i < n; i++ {
			array[i] = i + 1
		}
		ret := array[:k]
		rand.Seed(time.Now().UnixNano())
		for i := k; i < len(array); i++ {
			if rand.Intn(k+1) < k {
				ret[rand.Intn(k)] = array[i]
			}
		}
		fmt.Println(ret)
	}
}
