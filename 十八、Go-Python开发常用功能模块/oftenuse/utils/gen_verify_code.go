package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 生成n位验证码
func GenVerifyCode(n int) string {
	code := ""
	// 每次调用该函数都创建一个"活种"
	// int31n(n int32) int32：返回一个int32位的[0-n)的随机数
	// code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		// (*rand).Intn(n int) 返回[0, 10)范围间的一个随机数
		x := rnd.Intn(10)
		code += strconv.Itoa(x)
	}
	return code
}

// 生成6位验证码
func GenVerifyCode1() string {
	//rand.Seed(seed int64)方法，如果seed的值一样，那么每次调用产生的随机序列是一样的，
	//只有每次调用时seed的值不一样，才会产生不一样的随机序列
	return fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}