package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenVerifyCode(n int) string {
	code := ""
	// 每次调用该函数都创建一个"活种"
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		// (*rand).Intn(n int) 返回[0, 10)范围间的一个随机数
		x := rnd.Intn(10)
		code += strconv.Itoa(x)
	}
	return code
}