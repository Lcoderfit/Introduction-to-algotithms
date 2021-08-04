package test

import (
	"reflect"
	"testing"
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
