package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	1. 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	2. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	3. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	4. 主goroutine从resultChan取出结果并打印到终端输出
*/

// Job为任务结构体，其中包含的字段可以按需添加
type Job struct {
	value int64
}

// 存储原始值和结果
type Result struct {
	job *Job
	sum int
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan Result, 100)
var wg sync.WaitGroup

func main() {
	go sendValue(jobChan)
	for i := 1; i <= 24; i++ {
		go receiveVal(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value: %d, sum: %d\n", result.job.value, result.sum)
	}
	wg.Wait()
}

// 循环发送随机数
func sendValue(ch chan<- *Job) {
	wg.Add(1)
	defer wg.Done()
	for {
		// 返回一个非负的63bit的int64类型整数(最高位64为符号位0，表示正数)
		x := rand.Int63()
		job := &Job{
			value: x,
		}
		jobChan <- job
	}
}

// 接收结果
// 一直循环不间断读
func receiveVal(ch1 <-chan *Job, ch2 chan<- Result) {
	wg.Add(1)
	defer wg.Done()
	for {
		val := <-ch1
		sum := sumVal(val.value)
		ch2 <- Result{
			job: val,
			sum: sum,
		}
	}
}

// 计算所有位之和
func sumVal(n int64) int {
	result := 0
	for n > 0 {
		result += int(n % 10)
		n /= 10
	}
	return result
}
