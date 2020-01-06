package utils

import (
	"math/rand"
	"strconv"
	"time"
)

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