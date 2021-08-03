package main

import (
	"fmt"
	"strings"
)

/*
互斥锁，读写锁 都是结构体，所以当作函数参数传入时候必须传入
指针，否则传入的就不是原来那个锁

函数参数是值传递，但是如果参数是一个slice，则参数传递是地址拷贝
即将底层数组的地址复制给参数slice，
(slice本质是一个结构体，该结构体中有一个字段存储底层数组的指针
所以对slice进行拷贝的时候，拷贝的是这个结构体，但是结构体内指向
底层数组的指针的值是不变的，因此，在函数内部修改slice会影响原来的slice
)

和其他语言不同的是，如果函数参数是一个数组，则在参数传递时是对数组的拷贝
即在函数内对数组修改，不会影响到原来的数组
*/

// 不传入指针的话，不会修改v的值
func f(v interface{}) {
	v = 10
}

//函数参数是值传递，但是如果参数是一个slice，则参数传递是地址拷贝
//即将底层数组的地址复制给参数slice，
// 所以下面的函数k是会对slice b造成影响的
func k(v []int) {
	v[0] = 10
}

func main() {
	//a := 5
	//b := []int{1, 2}
	//f(a)
	//k(b)
	//fmt.Println(a)
	//fmt.Println(b)
	ret := Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	fmt.Printf("%+v\n", ret)
	fmt.Printf("%v\n", ret)
}

// 实现字符串分割函数
func Split(str string, sep string) []string {
	var ret []string
	index := strings.Index(str, sep)
	for index > -1 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	return ret
}
