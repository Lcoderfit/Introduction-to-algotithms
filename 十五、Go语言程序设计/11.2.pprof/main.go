package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/*
性能优化
CPU profile：报告CPU使用情况
Memory Profile: 报告程序的内存使用情况
Block Profiling: 报告goroutine不在运行状态的情况(用来分析查找死锁等性能瓶颈)
Goroutine Profiling: 报告goroutine使用情况，有哪些goroutine，他们调用关系怎样

runtime/pprof: 采集工具型应用运行数据进行分析(例如一些实用工具)
net/http/pprof: 采集服务型应用运行时的数据进行分许(例如http服务)

pprof开启后，每隔10ms会收集当下的堆栈信息，至少运行20s，获取各个函数
占用的CPU及内存资源等，对这些采样数据进行分析，形成分析报告
只在性能测试时实用pprof，因为很消耗性能


go run main.go -cpu=true // 采集cpu使用信息
go run main.go -mem=true // 采集内存使用信息

// 通过tool工具查看信息采集文件
go tool pprof cpu.pprof

// 查看占用cpu前三的函数, ctrl+c 或 quit退出
// flag 当前函数占用CPU耗时
// flat% 当前函数占用CPU耗时百分比
// sum% 函数占用cpu耗时累计百分比
// cum 当前函数加上调用当前函数的函数占用cpu的总耗时
// 最后一列：函数名称

(pprof) top3
Showing nodes accounting for 119.41s, 99.90% of 119.53s total
Dropped 21 nodes (cum <= 0.60s)
      flat  flat%   sum%        cum   cum%
    55.95s 46.81% 46.81%     96.58s 80.80%  runtime.selectnbrecv
    40.61s 33.97% 80.78%     40.63s 33.99%  runtime.chanrecv
    22.85s 19.12% 99.90%    119.45s 99.93%  main.logicCode

// 在select的default中加上 time.Sleep() 将cpu时间让出
Showing top 3 nodes out of 19
      flat  flat%   sum%        cum   cum%
      20ms   100%   100%       20ms   100%  runtime.stdcall4
         0     0%   100%       10ms 50.00%  runtime.(*mcache).nextFr
ee
         0     0%   100%       10ms 50.00%  runtime.(*mcache).refill


// (pprof)lis logicCode 可以单独查看指定函数的分析
// 通过分析大部分CPU资源被64行占用
//
// Total: 2mins
// ROUTINE ======================== main.logicCode in E:\SocialProject\
// Algorithms-Tags\Introduction-to-algotithms\十五、Go语言程序设计\11.2
// .pprof\main.go
//     23.22s      2mins (flat, cum)   100% of Total
//          .          .     55:
//          .          .     56:// (pprof)lis logicCode 可以单独查看指
// 定函数的分析
//          .          .     57:
//          .          .     58:
//          .          .     59:
//     23.22s      2mins     60:func logicCode() {
//          .          .     61:   var c chan int
//          .          .     62:   for{
//          .          .     63:           select {
//          .          .     64:           case v := <-c:
//          .          .     65:                   fmt.Printf("recv fro
// m chan:%v\n", v)


6.图形化界面 graphviz
https://graphviz.org/download/

7.

*/
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan:%v\n", v)
		default:
			// 休息一段时间，将CPU时间让出，避免程序一直占用
			//time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUProf bool
	var isMemProf bool // 是否开启内存profile的标志位

	flag.BoolVar(&isCPUProf, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemProf, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUProf {
		// 创建文件，并往文件中写入CPU信息
		f1, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMemProf {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		// 写入堆栈信息
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}
