package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

//funcStr: 函数路径字符串，如：oftenuse/utils.PrintCurFuncName
//callerLevel: 调用者层次，0表示当前函数
//seps: 分割符，如'/', 表示去掉"/"之前的部分
func GetFuncName(funcStr string, seps string) (ret string) {
	//获取函数名称
	//pc, _, _, _ := runtime.Caller(callerLevel)
	//fn := runtime.FuncForPC(pc)
	retList := strings.Split(funcStr, seps)
	ret = retList[len(retList) - 1]
	return ret
}

func GetLevelCallerFuncName() {
	//PrintCurFuncName
	a := PrintCurFuncName()
	//oftenuse/utils.init.0
	b := PrintUpOneLevelFuncName()
	//main.getFuncInfo，
	c := PrintUpTwoLevelFuncName()
	//runtime.main
	d := PrintUpThreeLevelFuncName()

	fmt.Println("当前函数：", a)
	fmt.Println("上一层调用者：", b)
	fmt.Println("上两层调用者：", c)
	fmt.Println("上三层调用者：", d)
}

//打印当前函数：PrintCurFuncName
func PrintCurFuncName() string {
	//skip == 0，表示当前函数
	pc, _, _, _ := runtime.Caller(0)
	return runtime.FuncForPC(pc).Name()
}

//打印上一层调用者
func PrintUpOneLevelFuncName() string {
	//skip == 1, 表示当前函数的上一层，即init
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

//打印上两层调用者
func PrintUpTwoLevelFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

//打印上三层调用者
func PrintUpThreeLevelFuncName() string {
	pc, _, _, _ := runtime.Caller(3)
	return runtime.FuncForPC(pc).Name()
}

func GetFuncInfo() {
	// This will print "name: main.foo"
	fmt.Println("name:", GetFunctionNameSpit(Foo, '.'))

	// runtime/debug.FreeOSMemory
	fmt.Println(GetFunctionNameSpit(debug.FreeOSMemory))
	// FreeOSMemory
	fmt.Println(GetFunctionNameSpit(debug.FreeOSMemory, '.'))
	// FreeOSMemory
	fmt.Println(GetFunctionNameSpit(debug.FreeOSMemory, '/', '.'))
}

//函数的路径
func GetFunctionNameSpit(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func Foo() {

}