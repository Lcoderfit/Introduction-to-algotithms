package test

import (
	"reflect"
	"testing"
	"time"
)

/*
1.切片不能直接比较，需要使用reflect.DeepEqual进行比较

2.go test

// 显示详细信息
go test -v

// 子测试（运行一个测试函数中的一个case）
go test -run=TestSplit4/case_3

// 运行一个测试函数
go test -run=TestSplit4

// 显示测试覆盖率(Goland 用go xxx.go with Coverage)
go test -cover

// 将覆盖率信息写入文件
go test -cover -coverprofile=c.out
go tool cover -html=c.out // 处理生成的覆盖率信息，
							// 在本地浏览器窗口生成HTML报告
测试函数覆盖率：100%
测试代码覆盖率：60%


二、基准测试: 在一定负载下检测程序性能的方法
1. go test -bench=Split

// 查看内存使用情况
2.go test -bench=Split -benchmem

// 预分配容量，减少内存分配次数
ret := make([]string, 0, strings.Count(str, sep)+1)

// 基准测试，至少会运行1s
// 指定基准测试时间(-bench=后面跟的是基准测试函数名，例如Fib40指的是
// 函数名中包含bench的名字)
go test -bench=Fib40 -benchtime=20s

*/

// 失败的测试用例
func TestSplit(t *testing.T) {
	// 函数结果
	ret := Split("babcbef", "b")
	// 预期结果
	want := []string{"", "a", "c", "ef"}
	// 切片不能直接比较，需要使用reflect.DeepEqual进行比较
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

// 成功的测试用例
func TestSplit2(t *testing.T) {
	ret := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

// 测试组
func TestSplit3(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		{"a:b:c", ",", []string{"a:b:c"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcef", "bc", []string{"a", "ef"}},
		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
		}
	}
}

// 子测试, 会将每一个测试用例单独罗列出来,
// 如果想单独运行一个子case: go test -run=TestSplit4/case_1
//    --- PASS: TestSplit4/case_1 (0.00s)
//    --- PASS: TestSplit4/case_2 (0.00s)
//    --- PASS: TestSplit4/case_3 (0.00s)
//    --- PASS: TestSplit4/case_4 (0.00s)
//PASS
func TestSplit4(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	// 如果是"case 1", 则用: go test -run=TestSplit4/"case 1"
	testGroup := map[string]testCase{
		"case_1": {"a:b:c", ",", []string{"a:b:c"}},
		"case_2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": {"abcef", "bc", []string{"a", "ef"}},
		"case_4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("want: %#v got: %#v\n", tc.want, got)
			}
		})
	}
}

// 基准测试, goos表示操作系统，goarch表示平台
// pkg表示当前所在的包，BenchmarkSplit-8中的8表示GoMaxProcs
// 3346797表示执行的次数， 379 ns/op表示平均每次操作的耗时
// go test -bench=Split -benchmem  查看内存分配情况
//
// 240 B/op 表示平均每次占用240byte
// 4 allocs/op 平均每次操作分配4次内存
//
// 创建slice map时，可以预分配容量，减少内存分配次数

// 优化前：
// goos: windows
// goarch: amd64
// pkg: test
// BenchmarkSplit-8         3346797               379 ns/op
// PASS
// ok      test    5.449s
//
// 优化后:
// goos: windows
// goarch: amd64
// pkg: test
// BenchmarkSplit-8         3580920               361 ns/op
//   80 B/op              1 allocs/op
// PASS
// ok      test    3.090s
func BenchmarkSplit(b *testing.B) {
	// 执行了N遍Split函数
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// 性能比较函数
// 上面的基准测试只能得到给定操作的绝对耗时, 下面的示例用于计算相对耗时
// 例如：同一个函数处理1000个元素的耗时与处理1w甚至100w个元素的耗时差别
// 或者处理同一个任务的不同算法的耗时
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 30)
}

// goos: windows
// goarch: amd64
// pkg: test
// BenchmarkFib1-8         228748588                5.18 ns/op
// BenchmarkFib2-8         85927262                14.2 ns/op
// BenchmarkFib3-8         50043787                24.2 ns/op
// BenchmarkFib20-8              92          13898079 ns/op
// BenchmarkFib40-8               1        1644599400 ns/op
// PASS
//
// Fib40只运行了1次（默认情况下，每个基准测试至少运行1s）
// 如果要多次运行Fib40，可以指定最小基准测试时间
// go test -bench=Fib40 -benchtime=20s
// 跑长一点时间计算出来的平均操作耗时更准确
func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}

// 重置基准测试时间, 去除掉一些无关的耗时操作
func BenchmarkSplit2(b *testing.B) {
	// 假设需要做一些耗时的无关操作
	time.Sleep(5 * time.Second)
	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split("沙河又沙又有河", "沙")
	}
}

// 并行测试
// 会以并行的方式执行给定的基准测试
// go test -bench=Split -cpu=1
func BenchmarkSplit3(b *testing.B) {
	// 假设需要做一些耗时的无关操作
	time.Sleep(5 * time.Second)
	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split("沙河又沙又有河", "沙")
	}
}
