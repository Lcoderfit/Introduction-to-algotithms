package main

import (
	"flag"
	"fmt"
	"time"
)

/*
1.打印帮助信息
go run main.go --help
  -age int
        请输入真实年龄 (default 9000)
  -ct duration
        结婚多久了 (default 1s)
  -married
        结婚了吗
  -name string
        请输入名字 (default "coder")

2.


*/

// os.Args 获取命令行参数
func main() {
	//fmt.Println(os.Args)
	name := flag.String("name", "coder", "请输入名字")
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了吗")
	// 输入时候例如：1000h
	cTime := flag.Duration("ct", time.Second, "结婚多久了")

	//var name string
	//// 从命令行参数中获取name参数的值，解析后赋给name变量
	//// TypeVar
	//flag.StringVar(&name, "name", "coder", "请输入姓名")
	//// 使用flag
	//flag.Parse()
	//fmt.Println(name)

	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)

	// 返回除命令行参数外的其他参数（一个字符串列表）
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数的总个数
	fmt.Println(flag.NArg())
	// 返回flag参数的个数(-name=lcoder 或 -name lcoder 这种形式才算flag参数)
	// go run main.go -name lcoder -age=23 a b c
	// [a b c]
	// 3
	// 2
	fmt.Println(flag.NFlag())

}
