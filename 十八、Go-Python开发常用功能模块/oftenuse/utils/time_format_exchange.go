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

//time.Now()
func TimeFunction() {
	now := time.Now() //获取当前时间
	fmt.Println("Now: ", now)
	//==>2019-08-21 11:30:51.2470317 +0800 CST m=+0.004501101
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())        //10位
	fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())    //19位
	fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)        //或者秒*1000也可
	fmt.Printf("时间戳（纳秒-->秒）：%v;\n",time.Now().UnixNano() / 1e9)
}

//获取指定时间前的时间
func GetTimeBefore() {
	// 获取50秒前的时间，方式1
	st,_ := time.ParseDuration("-50s")
	fmt.Println("50秒前的时间：",time.Now().Add(st))

	// 获取1分钟前的时间，n秒前则是time.Second * -n，方式2
	t := time.Now().Add(time.Minute * -1)
	fmt.Println("一分钟前的时间：",t)

	//获取1小时前的时间
	sth,_ := time.ParseDuration("-1h")
	fmt.Println("1小时前的时间：",time.Now().Add(sth))

	// 获取2天前的时间
	oldTime := time.Now().AddDate(0, 0, -2)
	fmt.Println("两天前的时间为：", oldTime)

	//获取两个月前的时间
	oldTime = time.Now().AddDate(0, -2, 0)
	fmt.Println("两个月前的时间为：", oldTime)
}

//获取指定时间后的时间
func GetTimeAfter() {
	// 获取50秒后的时间，方式1
	st,_ := time.ParseDuration("50s")
	fmt.Println("50秒之后的时间：",time.Now().Add(st))

	// 获取1分钟后的时间，n秒前则是time.Second * n，方式2
	t := time.Now().Add(time.Minute * 1)
	fmt.Println("一分钟后的时间：",t)

	//获取1小时后的时间
	sth,_ := time.ParseDuration("1h")
	fmt.Println("1小时之后的时间：",time.Now().Add(sth))

	// 获取当前时间2天后的时间
	newTime := time.Now().AddDate(0, 0, 2)
	//newTime 的结果为时间time类型
	fmt.Println("当前时间2天后的时间", newTime)

	//获取当前时间2月后的时间
	newTime = time.Now().AddDate(0, 2, 0)
	fmt.Println("当前时间2月后的时间", newTime)
}

func GetTimeDelta() {
	t1 := time.Now()
	//设置期间经历了50秒时间
	t2 := time.Now().Add(time.Second * 50)
	fmt.Println("t2与t1相差：",t2.Sub(t1))		//t2与t1相差： 50s
}