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

switch n {
case 1:

case 2:
}

switch n := 4; n {

}

switch n := 4; n {
case 1,2,3,4:
}

switch {
case n > 1:
case n  < 10:
}


fallthrough // 执行满足条件的case的下一个case

switch {
case s == "a":
    fmt.Println("a")
    fallthrough // 兼容C语言, 为了向前兼容，但是不鼓励用这个
case s == "b":

default:
}


算术运算符 + - * / %
关系运算符 > < == >= ...
逻辑运算符 && || !
位运算符 & | ^  // 针对二进制数进行操作
赋值运算符 *= /= <<= >>= &=

++和--在GO语言中是单独的语句，并不是运算符

a=a++ // 错误的写法，因为a++是一个单独语句，不是表达式

GO语言是强类型，相同类型才能比较

运算符优先级

```


```text
复合数据类型

1.数组，必须指定存放的元素类型和长度

var a [3]bool


bool类型默认值为false, 整型浮点型默认值为0, 字符串为空字符

a1 := [3]bool{true, true, false}

// 根据初始值自动推断长度
a := [...]int{1,2,3}


a := [5]{1,2} // 

a := [5]{0:1, 4:3} // 根据索引初始化

a := [2][3]int{
    [3]int{1,2,3},
    [3]int{1,4,3},
}

// 数组是值类型，赋值操作是将数组拷贝一份，两个占地址不同
a := [2]{1,2}
b := a
b[0] = 3 // 修改b对a无影响


a := [...]{1,2}
b := a[:2] // 基于一个数组得到切片

a:= [...]{1,2,3,4,5,6,7}

// 长度为4，容量也为4
b := a[3:] // 切片指向一个底层数组；切片的长度指的是切片元素的个数
            //切片的容量指的是底层数组从切片第一个元素到数组最后一个元素的长度
 
切片是一个引用类型，底层是一个数组，如果修改了底层数组，切片的值也会发生改变

a[6] = 100 // 则b变为 [4,5,6,100]


切片的本质(指向底层数组的指针、长度、容量)


切片属于引用类型，真正的数据保存在底层数组中
nil值的切片没有底层数组


a = []{1,2}
b := a
b[0] = 100 // 会改变底层数组的值，从而导致a的值也改变


// 如果append操作导致底层数组扩容(而数组实际上定义之后
// 长度是不可变的，所以事实上扩容是重新定义了一个新的数组)，
// 则会导致a的地址发生变化
a := append(a, 10)

a := append(a, k...) // 将切片k拆开
```