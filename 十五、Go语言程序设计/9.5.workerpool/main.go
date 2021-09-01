package main

import (
	"fmt"
	"time"
)

/*
3个协程分6个任务
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	// 因为管道中的元素是读一个少一个，所以多个goroutine并发执行时
	// 不会出现两个goroutine读取到同一个元素的情况
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 6; i++ {
		jobs <- i
	}
	close(jobs)

	// 输出结果(没有通道接收者，所以不会输出到控制台)
	for i := 1; i <= 6; i++ {
		<-results
	}
}
