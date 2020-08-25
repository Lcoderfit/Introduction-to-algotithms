package str

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	var ret int64
	//处理正负号
	flag := true
	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		flag = false
		str = str[1:]
	}
	//去除前导0
	var k int
	for k = 0; k < len(str); k++ {
		if str[k] != '0' {
			break
		}
	}
	str = str[k:]
	//字符转数字
	//int32类型最多只有11位（20多亿）
	//这里多取一位，因为原字符串可能有12位字符数字，但光取前11位可能比int32要小
	for i := 0; i < 12  && i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		//i to s没有错误，s to i可能存在错误，所以Atoi函数返回result, err两个值
		ret = ret*10 + int64(str[i] - '0')
	}
	if !flag {
		ret = -ret
	}
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	if ret < math.MinInt32 {
		return math.MinInt32
	}
	return int(ret)
}