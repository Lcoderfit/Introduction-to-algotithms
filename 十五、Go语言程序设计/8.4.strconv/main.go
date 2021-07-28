package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	// base和bitSize参数表示转换为64位的10进制数
	// 如果指定的bitSiz而不够存储转换后的值，会报错：strconv.ParseInt: parsing "10000": value out of range
	// 如果bitSize设置位0，则是根据base的前缀来设置转换后的值所占的位数
	// 例如0b表示8进制，则会设置位8位
	// 0x表示16进制，则会设置位16位
	ret1, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", ret1)

	// 从字符串中解析出bool
	str = "true"
	ret2, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", ret2)

	// 解析float
	str = "1.234"
	// 将字符串解析为float类型，bitSize表示解析后的浮点型所占的位数
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(floatValue)

	// strconv.Atoi  strconv.ItoA()
	t := 10
	ret3 := strconv.Itoa(t)
	fmt.Println(ret3)

	str = "10"
	ret4, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret4)
}
