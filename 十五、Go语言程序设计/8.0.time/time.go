package main

import (
	"fmt"
	"time"
)

func main() {
}

func PrintTime() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
}

func PrintTimestamp() {
	// 时间戳是自1970年1月1日到当前时间的总毫秒数
	now := time.Now()
	t1 := now.Unix()     // 时间戳
	t2 := now.UnixNano() // 纳秒时间戳
	fmt.Println(t1)
	fmt.Println(t2)
}

func TimestampToTime(timestamp int64) {
	// 第二个参数相当于一个标志位，一般传0
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	fmt.Println(timeObj.Year())
	fmt.Println(timeObj.Month())
}

// 代表两个时间点之间经过的间隔
// Duration类型是一个int64类型的别名
// const (
//     Nanosecond  Duration = 1
//     Microsecond          = 1000 * Nanosecond
//     Millisecond          = 1000 * Microsecond
//     Second               = 1000 * Millisecond
//     Minute               = 60 * Second
//     Hour                 = 60 * Minute
// )

func Duration() {
	now := time.Now()
	// 加时间间隔, 一个小时以后可以用-time.Hour
	fmt.Println(now.Add(time.Hour))

	hourAge := now.Add(-time.Hour)

	// Sub：求两个时间之间的差值（now - hourAge）
	// 如果超过了int64能表示的最大值/最小值，则返回最大值/最小值
	subTime := now.Sub(hourAge)
	fmt.Println(subTime)

	// 判断两个时间是否相等, 返回bool值
	// 会自动考虑时区的影响，即使不同时区也可以正确比较时间
	now.Equal(hourAge)

	// 判断now是否在hourAge之前
	now.Before(hourAge)

	// 判断now是否在hourAge之后
	now.After(hourAge)

	// 定时器, 本质上是一个通道(channel)
	timer := time.Tick(10 * time.Hour)
	for t := range timer {
		// 打印任务执行时的当前时间
		fmt.Println(t) // 每个小时都会执行该任务
	}
}

// 时间格式化，用的是 2006 1 2 3 4 5 PM Mon Jan 来格式化(2006年1月2号15点04分5秒 PM 星期一 一月)
// 如果格式化为12小时方法，需指定PM(时间为上午，格式化字符串也只需要使用PM)
func TimeFormatter() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-01 15:04:05 Mon Jan"))
	// 格式化为12小时制表示
	fmt.Println(now.Format("2006-01-01 15:04:05 PM Mon Jan"))

	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02"))
}

// 将字符串转换为Timestamp
func StringToTimeStamp(s string) int64 {
	timeObj, err := time.Parse("2006-01-02", s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Println(timeObj)
	// 转换为时间戳
	return timeObj.Unix()
}

// 睡眠
func SleepFunc() {
	time.Sleep(time.Second)

	n := 1000
	// 如果定义一个上面的n，则n会识别为int类型，而Duration是一个int64类型，所以会报错，需要显示转换
	// 如果直接传入 time.Sleep(100 * time.Second)则没有问题，因为程序会自动转换为int64类型
	// time.Sleep(100) 也没有问题，会自动转换类型，表示100纳秒
	time.Sleep(time.Duration(n) * time.Hour)
}

func UTCTime() {
	// 获取本地时间，即当前时区的时间
	now := time.Now()

	// 返回的是一个UTC时间，而不是当前时区的时间
	t, err := time.Parse("2006-01-02 15:04:05", "2021-07-17 18:37:30")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
	// 按照东八区的时区和格式解析一个字符串
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("fuck sheet")
		return
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-07-17 18:37:30", loc)
	if err != nil {
		fmt.Println("parse time failed,")
	}
	fmt.Println(timeObj)

}
