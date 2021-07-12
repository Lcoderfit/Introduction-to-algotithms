# Go环境配置

```text
1.GOROOT: 安装GO的路径
2.GOPATH: GO的项目路径
    bin/    GO项目执行文件（exe）
    pkg/    GO相关的库
    src/    GO项目存放路径

    GO1.12之后支持go modules, 项目可以随便放，不用放到src下
    
3.go env: 查看当前GO的环境变量
4.GOPROXY: go下载第三方库所使用的代理


go build 命令大全
    go build -o 指定名称
    
go run命令大全


go install
    先编译成exe可执行文件，然后移动到GOBIN目录下,
   
类型推到：var a = 10 
短变量声明：s := "sadfj" // 只能用于函数内
匿名变量：_, y := f() // 忽略某个值

函数外必须以某个关键字开始(var func const)


// 批量声明常量时，如果某一行没写，则默认与上一行一样
const (
    n = 100
    n1
    n2
)

// 常量计数器, 
const (
    a1 = iota // 0
    a2          // 1
    a3          // 2
)

const (
    a1 = iota // 0
    a2          // 1
    _           // 2
    a3          // 3
)

const (
    a1 = iota // 0, 没新增一行常量声明，iota计数加一次
    a2 = iota // 1
    a3 = iota // 2

)


const (
    c1 = iota // 0
    c2 = 100 // 100, iota还是会加1
    c3 = iota // 2
)

const (
    d1, d2 = iota+1, iota+2 // 1, 2
    
    d3, d4 = iota+1, iota+2 // iota计数加一次，变为1, d3 d4分别为 2, 3
    
)

定义数量级
const (
    _ = iota
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
    TB = 1 << (10 * iota)
    PB = 1 << (10 * iota)
)

整型分为两类，
有符号类型和无符号类型
int8 int16 int32 int64

无符号类型：
uint8 uint16 uint32 uint64

特殊整型
int 在32位操作系统上为int32, 64位操作系统上为int64
uint 在32位操作系统上为int32, 64位操作系统上为int64
uintptr  存放一个指针，无符号整型

GO语言种无法直接定义二进制
var a int = 10
fmt.Printf("%b", a) // 输出二进制表示

var c int = 0xff
fmt.Printf("%x", c)

var d int = 077
fmt.Printf("%o", d)

fmt.Printf("%d", d) 表示十进制数

各种占位符 %v %T

math.MaxFloat32 32位浮点型最大值
math.Max....

float32不能直接赋值给float64
int....

复数类型：complex64 complex128
var i complex64 = 1+2i

布尔值默认为false
无法与其他类型进行转换

fmt.Printf("%#v")

GO语言中字符串必须用双引号进行包裹
单引号包裹的是字符(单独的字母，汉字，符号)

一个字符占一个字节
一个utf8编码汉字一般占三个字节


反引号：原生字符串

fmt多个内置函数

strings.Split(s, "")

strings.Contains()
strings.HasSuffix

strings.Index(s, "c") 查找位置
strings.LastIndex(s, "c") 最后出现的位置

strings.Join(s, "" )

for _, c := range s {
    fmt.Printf("%c", c)
}

处理非ascii码的类型，使用rune类型

英文类型使用byte，中文字符使用rune

字符串不可修改,需要转换为rune列表再修改

s2 := "阿道夫"
s3 := []rune(s2)
s3[0]='红'

s4 := 'h' //int32

[]rune类型是int32的一个别名

统计"hello沙河小王子" 中中文的数量

// age用完即销毁
if age:=1; age > 0 {
    fmt.Printf("%d", age)
}


go中只有for循环，没有while循环

i := 1
for ; i < 10; i++ {

}

for i < 10 {
    i++
}


for {

}


for range 键值循环，专门用于遍历数组、切片、字符串、map及通道

for k, v := range s {

}

```