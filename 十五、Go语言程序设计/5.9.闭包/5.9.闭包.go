package main

import "fmt"

/*
Go语言中，函数是一等公民（可以像普通类型一样被赋值给变量，作为参数传递，或作为函数的返回值）

一、什么是闭包
引用外部局部变量的函数


二、闭包的作用是什么：
1.延长局部变量的生命周期（当没有变量引用闭包时，闭包才会被销毁；闭包内引用的局部变量的生命周期，与外部引用该闭包的变量的生命周期一致）
2.让函数外部可以操作函数内部的数据

三、闭包的缺点
闭包会使得函数中的变量都被保存在内存中，生命周期延长，导致内存消耗较大；所以不能滥用

四、一般什么情况下使用闭包
例如变量自增，我想操作一个变量能像全局变量一样，每一次操作都能使该变量自增1，但又不想定义全局变量（全局变量容易被其他程序修改，获取变量太大不好管理）

五、闭包的原理
1.函数可以作为返回值
2.函数内部查找变量的顺序，先在自己内部找，找不到再到外层找


一、defer的作用
1.资源清理、文件关闭、解锁及记录时间, 利用延迟调用的特性，可以用于处理资源释放
*/

func main() {
	k := f()
	k()
	k()

	r := add(100)
	fmt.Println(r(100))
	fmt.Println(r(200))

	a, s := calc(100)
	fmt.Println(a(1), s(1))
	fmt.Println(a(5), s(1))
	fmt.Println(a(4), s(10))
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func f() func() {
	var a = 1
	return func() {
		a++
		fmt.Println(a)
	}
}

func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
