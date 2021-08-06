package main

import (
	"fmt"
	"sync"
)

/*
通过共享内存进行数据交换的确定：共享内存在不同的goroutine中容易发生竞态问题，为了保证数据交换的正确性，
	须通过互斥量进行加锁,这种做法势必造成心性能问题
1.Go语言并发模型：CSP（communicating sequential process），提倡通过通信共享内存而不是通过共享内存实现通信
2.Go语言通过channel（通道）在不同的goroutine之间进行通信, channel类似一个传送带或一个队列，
	在通道中数据是按顺序传输的，先入先出;
	每一个channel都是一个具体类型的导管，声明时需要指定元素类型
3.通道必须初始化才能使用
4.channel操作
	发送：ch <- 10
	接收：x := <- ch
	关闭: close(ch)
5.往一个无缓冲区的channel发送数据(但是不接收)会导致阻塞(fatal error: all goroutines are asleep - deadlock!)
	如果有接收者接收则不会阻塞
  往通道中发送数据，如果超过缓冲区容量，则会导致阻塞（fatal error: all goroutines are asleep - deadlock!），直到有接收者接收

6.	c := make(chan int) // 不带缓冲区的通道的初始化
	d := make(chan int, 16) // 带缓冲区的通道的初始化

7.遍历通道值
for v := range ch {} // 只能用一个值，没有索引(写成for k, v := range ch是错误的)
注意：for v:= range ch 这种写法会一直读取ch，所有数据都读完了之后如果channel未关闭会阻塞（报fatal error）
所以写完之后一定要close(ch),从关闭的channel中读可以读数据，缓冲中的数据读完后，仍然可以读，不过此时读取的是通道的默认值，例如：
x, ok := <- ch 如果ch已关闭，且缓冲中的数据已经读完，这条语句仍然可以正常执行，只不过此时ok未false，x读取出来的是通道类型默认值

for v := range ch 这种写法内部会对ok的值进行判断，如果ch已关闭，循环会一直读取，直到ok为false（即缓存已读完）

另一种写法：
for {
	x, ok := <- ch
	// 如果通道已经读完了就退出循环
	if !ok {
		break
	}
}

8.往已关闭的带缓存的channel中写数据会阻塞，但是读可以正常读???

9.从channel读数据，如果channel缓冲区为空或者没有缓冲区，当前goroutine会被阻塞。
向channel写数据，如果channel缓冲区已满或者没有缓冲区，当前goroutine会被阻塞。

10.确保某个操作只执行一次
sync.Once

sync.Do(func() { close(ch) })

// done用来标志是否执行过，m表示锁
// 会先检测是否执行过该操作，如果执行过则不再执行，否则上锁后执行
type Once struct {
	done uint32
	m    Mutex
}

11.单向通道:
	chan<- int	接收通道(只能接收不能发送)
	ch <-chan in 发送通道(只能发送不能接收)


声明和定义的区别：定义涉及到内存的分配，声明不涉及

例如：var a int //声明了一个int类型变量a
	var a int = 10 //定义了一个int类型变量, 初始化其值为10
*/

var a []int
var b chan int // 声明时需要指定元素类型, 默认值为nil

var wg sync.WaitGroup
var once sync.Once

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	// 如果wg减未负数会报：panic: sync: negative WaitGroup counter
	wg.Add(3)
	go f1(a)
	// 对同一个channel连续关闭两次：panic: close of closed channel
	// 通过sync.Once控制关闭操作只执行一次
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	// 如果b未关闭，则会一直读，读完后会报fatal error
	// 如果b关闭了，写会阻塞，但是读可以读，读完后则退出循环
	for ret := range b {
		fmt.Println(ret)
	}
}

//
func f() {
	b = make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 无变量接收，即将值丢弃
		<-b
		fmt.Println("接收通道值")
	}()
	// 向通道发送10
	b <- 10

	wg.Wait()
}

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		// 如果ch1不关闭，则ok一直为true
		// 只有当读完ch1，并且ch1关闭，才会返回false
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	fmt.Println("ch2运行完了")
	// 只关闭一次ch2
	//
	once.Do(func() {
		close(ch2)
	})
}
