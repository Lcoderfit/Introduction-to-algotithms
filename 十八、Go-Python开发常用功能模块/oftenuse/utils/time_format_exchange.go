package utils

import (
	"fmt"
	"time"
)

func TimeFormatExchange() {
	//datetime := "2019-03-11 21:07:00"  	    //待转化为时间戳的字符串
	////日期转化为时间戳
	//timeLayout := "2006-01-02 15:04:05"     //转化所需模板
	//loc, _ := time.LoadLocation("Local")    //获取时区
	//
	////调用转化方法，传入上面准备好的的三个参数
	//tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	//timestamp := tmp.Unix()                 //转化为时间戳（秒级） 类型是int64
	////timestamp = timestamp * 1000    	    //转化为毫秒级
	//log.Println(timestamp)
	//
	////时间戳转化为日期
	//datetime = time.Unix(timestamp, 0).Format(timeLayout)
	//fmt.Println(datetime)

	// 获取当前时间戳
	timeUnix := time.Now().Unix()         //单位s,打印结果:1491888244
	timeUnixNano := time.Now().UnixNano() //单位纳秒,打印结果：1491888244752784461
	fmt.Println(timeUnixNano)

	// 当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr) //打印结果：2017-04-11 13:24:04

	// 时间戳转时间字符串(int64->string)
	// SwitchTimeStampToDate
	timeUnix = time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39

	// 时间字符串转时间（string->time）
	formatTimeStr = "2017-04-11 13:33:37"
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err == nil {
		fmt.Println(formatTime) //打印结果：2017-04-11 13:33:37 +0000 UTC
	}

	// 时间字符串转时间戳 (string—>int64)
	fmt.Println(formatTime.Unix())
}

// 将时间戳转换为日期
func SwitchTimeStampToData(unixTime int64) string {
	timeStr := time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
	return timeStr
}
