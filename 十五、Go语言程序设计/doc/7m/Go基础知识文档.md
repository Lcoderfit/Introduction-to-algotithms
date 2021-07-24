# Go基础知识文档

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

## 1.1 Go如何删除和插入切片
```text
1.删除
a := []int{1,2,3}
a = append(a[:index], a[index+1:])

2.插入
a := []int{1,2,3}
a = append(a[:index], 0, a[index+1:]...)

```

## 1.2 Go扩容时的特点
```text
1.如果定义一个slice时，不定义大小而定义容量，则append效率会很快
a := make([]int, 0, 10)
a = append(a, 2)

2.如果append时超过了slice的容量，则会进行扩容，Go底层会重新创建一个新的slice，所以slice对应的变量名的地址会放生变化;
例如：
a := make([]int, 0, 1)
a = append(a, 1,2,3) // 此时超出容量，则a的地址会发生变化

但是slice的地址不会发生变化

pa := &a
*pa = append(*pa, 1, 2, 3, 4) // pa的值是不会发生变化的


newCap := oldCap
doubleCap := 2 * newCap
if cap > doubleCap {
    newCap = Cap
} else {
    if old.len < 1024 {
        newCap = doubleCap
    } else {
        for 0 < newCap && newCap < cap {
            newCap += newCap/4     
        }
        
        if newCap <= 0 {
            newCap = cap
        }
    }
}



```


## 1.3
```text
%p打印地址

a := []{1,2,3}
b := a
copy(c, a) // 把内存拷贝一份
a[0] = 100  // a: [100, 2, 3], c: [100, 2, 3]; c: [1,2,3]


// 将a中索引为1的数删除
a = append(a[:1], s[2:]..) // 假设底层数组为[1,2,3], 则经过这一步操作后
                    // a为[1,3], 但底层数组为[1,3,3], 就是将索引为2的数给赋值到索引为1的位置



a := make([]int, 5, 10)
for i := 0; i < 10; i++ {
    a = append(a, i)
}
// 输出0 0 0 0 0 0 1 2 3 4 5 6 7 8 9
// 因为一开始a初始化之为[0 0 0 0 0], 而后面又append了10个元素，所以最后长度
// 为15，由于当长度小于1024时每次扩容为原来的2倍，所以最后容量为20
fmt.Println(a) 
                

对切片进行排序
sort.Ints(a)

```


```text
GO语言中不需要对地址进行操
作，只需要记住两个符号:
1. &: 取地址
2. *: 根据地址取值

n := 10
p := &n // *int类型
fmt.Printf("%p %T", p) 


new和make均可用于申请内存，区别在于new返回地址，make返回类型本身
new一般用于基本数据类型的内存申请
make只用于slice map chan的内存创建

// 以下代码会panic
var a *int
*a = 100


// 正确
var a = new(int)
*a = 100



if v, ok := m["key"]; !ok {
    fmt.Println("不存在")
}

for k, v := range m {
    
}

// 删除key
// 没有返回值, 若m中不包含指定的key或m为nil，则delete不操作
delete(m, "key")

%v %+v %#v的区别

rand.Seed(time.Now().UnixNano())
rand.Intn(100) // 0~99

// 对string slice进行排序
sort.Strings()


// GO语言中没有默认参数
// y为不定参数，可以传递 >=1个参数
func f(x int, y ...int) int {
    
}

// 延迟调用,
// reture x者条语句其实是一个非原子操作，第一步是将x赋给返回值，第二是执行return指令
// derfer执行时机位于两个操作之间,先将x赋值给返回值，然后执行defer，最后执行return
defer会把当前语句放入栈中，按照先进后出的顺序执行



// 先将x赋值给返回值t，然后x自增1，组后返回t的值，结果为5
func f() {
    x := 5
    defer func() {
        x++
    }
    
    return x 
}

// 先将x赋值给返回值x，然后x自增1，组后返回x的值，结果为6
func f() (x int) {
    defer func() {
        x++
    }
    
    return x
}

// 返回值为5
func f() (x int) {
    // 传入参数x，是值传递，会将x的值拷贝一份，
    defer func(x int) {
        // 此时自增的是拷贝后的x
        x++
    }(x)
    return 5
}

// 返回值为5, 
func f() (x int) {
    defer func(x int) int {
        x++
        // 这个返回没有变量接收，注意，这里不是f的返回值，是defer func的返回值  
        return x
    }(x)
    
    return 5
}

// 结果是6，首先将5赋值给x，然后x自增1变为6，最后return
func f() (x int) {
    defer func(x *int) int {
        (*x)++  
    }(&x)
    
    return 5
}

// 注意，f(2)会先执行，假设其返回值为k，之后defer延迟调用的是f(1,k)
defer f(1, f(2))




1.全局作用域
2.函数内作用域
3.语句块作用域


// 接受 func() int类型的函作为参数，注意func(int) int 类型的函数不能作为参数
func f(x func() int) {
    x()
}


// 匿名函数
func main() {
    x := func(x, y int) int {}
}







```


## 闭包
```text

```

## 内置函数
```text
close 主要用来关闭channel

panic recover  用来做错误处理

// defer会在panic之前执行
defer f()...
panic("aksdjf")


// 1.recover()必须搭配defer使用(因为如果panic了，而没有defer的话，是不会执行到恢复的语句的, 而在panic之前执行recover恢复又没有意义)
// 2.defer一定要在可能引发panic之前的语句定义(因为在panic之后定义的话就执行不到了,panic之后直接就退出了)

// 下面程序的执行顺序：
// 1.defer延迟
    2.先开始触发panic，然后执行defer，recover会将程序恢复，然后打印err和"恢复数据库连接"
func f() {
    defer func() {
        // 拿到panic的错误，尝试恢复
        err := recover()
        fmt.Println(err)
        fmt.Println("恢复数据库连接")
    }
    
    panic("程序崩溃了")
    
}


%e 科学记数法 1e10
%E 科学计数法 1E10
%f 有小数部分但无指数部分  1.11
%F 等价于%f
%g 根据实际情况采用%e %f格式
%G 根据实际情况采用%E或%F格式

%q 对应对双引号括起来的字符串字面值
%x 每个字节用两字符16进制数表示(a-z)
%X 每个字节用两字符16进制数表示(A-Z)
%9.2f  宽带9位，小数保留2位

占位符：

%m5s  用m进行填充

fmt.Scan()
fmt.Scanf()

%% 百分号
%+v 输出结构体会添加字段名
%#v 值的Go语法表示

```

## 自定义类型和类型别名
```text
// 自定义类型，作用：为自定义类型添加自己的方法
// 类型为main.myInt
type myInt int

// 类型别名, 类型为int
// 类型别名的作用: 为了编写代码时更加直观
// 例如rune类型其实是int32类型，但是如果用int32会让人以为是整型
// 其实是存储unicode
type yourInt = int

// 结构体字段名小写，在当前包可以赋值，但是不可导出到其他包
type person struct {
    name string
    b string
}

var p person
p.name = "string" // 这个是可以的


匿名结构体：

// 多用于临时场景，用一次就不再使用了
var s struct{
    name string
}

```


```text
1.Go语言中函数传参永远是拷贝
2.Go语言不支持对指针进行操作，所以(*x).name与x.name所表示的含义是一样的，
一定是对指针所指向的变量进行操作

3.内存对齐
4.Go语言是面向接口编程;通过结构体构造对象
5.定义结构体时，什么时候返回结构体，什么是否返回结构体指针
a := struct{}
a := &struct{} // 当结构体字段较大，该结构体比较大时候，返回指针，这样可以节省内存

构造函数:

type person struct {
    name string
    age int
}

func newPerson(name string, age int) {
    return &person{
        name: name,
        age: age,
    }
}

// 方法是作用于特定类型的函数
type dog struct {
    fmt.Printf("%f", d.name)
}

// d是接收者, 命名一般用类型名首字母小写
func (d dog) wang() {
    fmt.Printf()
}
p := newPerson
p.wang()

标识符：变量名，函数名，类型名，方法名

// Dog是一个结构体类型，var d Dog // d是变量名，Dog是类型
type Dog struct {

}


// 什么时候使用指针类型接收者
1.需要修改接收者中的值
2.接收者是拷贝代价比较大的对象
3.如果有某个方法使用了指针接收者，则需要保持一致性，即其他方法也使用指针类型接收者


// 不能给GO内置类型添加方法，只能给自定义类型添加方法
// 不能给别的包里的类型添加方法，只能给自己包中的类型添加方法
// 如果要给别的包里的类型添加方法，则可以根据别的包中的类型添加自定义类型

type myInt int

func (i myInt) hello() {
    fmt.Println("我是一个int")
}


//结构体匿名字段
// 适用于字段比较少且简单的场景，不能有两种相同类型的匿名字段
type person struct {
    string
    int
}

//使用匿名字段

p = person{}
p.string // 直接加类型名


// 嵌套结构体

type AS struct {
    name string
}

type BS struct {
    age int
    a AS
}

bs := BS{}
bs.a.name // 访问

// 匿名嵌套结构体

type AS struct {
    name string
}

type BS struct {
    age int
    AS // 匿名嵌套访问时可以省去中间结构体，但是如果嵌套的两个结构体中都包含相同字段，则不能省略(匿名嵌套冲突)
}

bs := BS{}
bs.name



// 结构体模拟实现继承

type animal struct {
    name string
}

func (a animal) move() {
    fmt.Printf("%s", a.name)
}

type dog struct {
    feet uint8
    animal
}

func (d dog) wang() {
    fmt.Printf("%s在叫：汪汪汪", d.name)
}

d := dog{}
d.wang()
d.move() // animal有的方法，通过匿名嵌套调用


```


```text
1.序列化：把Go语言中结构体转换为json格式字符串
2.反序列化：json格式的字符串 -》 Go语言中的结构体变量



p := person{}

// 定义结构体时字段应该大写，否则打印出来字段将不可见
// 因为json格式化是在json这个包中操作的，所以字段必须可导
b,err := json.Marshal(p)


type person {
    // 解析成不同格式时有设置不同名称
    Name string `json:"name" db:"name" ini:"name"`
    Age int `json:"age"`
}


str := `{"name": "lcoder"}`
var p person

// 传递指针是为了能在json.Unmarshal内部能修改p2的值
json.Unmarshal([]byte(str), &p2)


var b = tmp{
    name string
    age int
}{"lcoder", "23"}


Go语言中把错误当成值返回，通常作为第二值返回


类型别名只在编译过程中有效，编译完后就不存在了，byte和rune都是类型别名

```


```text
1.接口是一种类型, 规定了变量有哪些方法(不关心变量是什么类型，只关心是什么方法)

type speaker {
    speak() // 只要实现了speak方法的变量都能当成一个speaker接口
}

func (c cat) speak() {
 
}

func (d dog) speak() {

}

func da(x speaker) speak() {

}


// 接口定义, 一个变量如果实现了接口中的所有方法，则该变量就实现了这个接口（可以作为这个接口类型）

type reader interface {
    f1(int) int
    f2(string) string
}


type animal interface {
    move()
    eat(string)
}

type cat struct {
    name string
    feet int8
}

func (c cat) move() {

}

// 如果实现的eat方法不带参数，则就不能算实现了animal接口
func (c cat) eat(s string) {
    
}


//接口具有动态类型和动态值, 这样设计才能使得接口可以接收任意类型
初始化时为动态类型和动态值为nil, nil
接口接收什么类型，就具有该变量的类型和值

// 接口是引用类型


// 使用值接收者与使用指针接收者实现接口的区别
1.使用值接收者实现接口，结构体类型和结构体指针类型变量都能存
2.指针接收者实现接口，只能存结构体指针类型的变量

// 多个类型可以实现同一个接口，一个类型可以实现多个接口

// 空接口类型

var m map[string]interface{}
m = make(map[string]interface{}, 10)
m["name"] "lu"

// 类型断言, 如果为true，则将a转换为string类型，如果为false，则为该类型对应的空值
str, ok := a.(string)
if !ok {
    fmt.Println("猜错了")
} else {
    fmt.Println("传进来的是一个字符串")
}


switch t := a.(type) {
    case string:
    case xxx
}


// 该什么时候使用接口
当两个及以上的类型需要以相同方式运行时，才需要定义接口

```


# 导入包
```text
1. 导入内置包和第三方包时空一行

import (
    "fmt"
    
    "github.com/xxxx/xx"
    
    
    alias "github.com/x/x//x" // 取别名
    
    _ xxxxx // 匿名导入，执行包内的init，导入包但是不适用包中的数据
    . xxxx // 点试导入，可以不写包名
)


init执行顺序：全局声明-》init->main


如果main中导入了A，A中导入了B，则先执行的是B的init，然后执行A，最后执行main


```

# 文件操作
```text

fileObj, err := os.Open("./main.go")
if err != nil {
    return
}

defer fileObj.Close()

// 指定读的长度
var t = make([]byte, 100)
n, err := fileObj.Read(t)
if err != nil {

}
// 读取了n个字节
fmt.Println(n)


for {
    n, err := fileObj.Read(t)
    if err == io.EOF {
        return
    }
    if err != nil {
        adsfjkjasf
    }
    if n < 128 {
        return
    }
}



// bufio //缓冲区

reader := bufio.NewReader(file)
for {
    // 一行一行读取  
    line, err := reader.ReadString('\n')
    if err == io.EOF {
        break
    }
    if err != nil {
        return
    }
    
    fmt.Println(line)
}


// ioutils读取整个文件

content, err := ioutil.ReadFile("./main.go")
if err != nil {
    fmt.Println()
    return
}

fmt.Println(content)

// 写入文件

os.OpenFile(name string, flag int, perm FileMode)
name: 要打开的文件名
flag: 打开文件的模式

os.O_WRONLY 只写
os.O_CREATE 创建文件
os.O_RDONLY 只读
os.O_RDWR  读写
os.O_TRUNC 清空
os.O_APPEND 追加


r 04
w 02
x 01


// 写入字节切片和字符串
// O_xxx 通过不同的二进制位表示不同的意义
// O_TRUNC 相当于每次都重新写入

file, err := os.OpenFile("a.txt", O_APPEND|O_CREATE, 0644)
if err != nil {
    
}
file.Write([]byte("lkjsdf"))
file.WriteString("hello world")

// 写入缓冲区
wr := bufio.NewWriter(fileObj)
wr.WriteString("kjasdf")
wr.Flush() // 将缓存中的内容写入文件

// io.utils

fileObj, err := ioutils.WriteFile(filename, []byte(), 0666)
if err != nil {
    fmt.Println("Lcoderfit")
    return
}

// 读取到空白符停止     
fmt.Scanln(&s)


// 从标准输入读取
var s string
reader := bufio.NewReader(os.Stdin)
s, _ := reader.ReadString('\n') // 读取到\n结束



// 在文件中间插入
file, err := os.OpenFile(file_name, os.O_WRRD, 0644)
file.Seek(3, 0) // 第一个参数是偏移量，第二个参数是相对位置，0表示相对文件开头
// 相对文件开头偏移三个字节后写入文件
file.Write([]byte('c'))


```

# time
```text
time.Now() //获取当前时间(2021-01-01 10:10:10.1111 +0800 .....)

now.Year() 

// 


```

```text
日志级别

// Trace和Debug的区别，Debug会在Debug版本下输出，但是在Release下不会输出，Trace会在Release下输出
1.Debug
2.Trace
3.Info
4.Warning
5.Error
6.Fatal


```