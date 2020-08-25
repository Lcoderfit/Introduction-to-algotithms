package str

import (
	"math"
	"strings"
)

func StrToInt(str string) int {
	//1.清除左右两边的空格
	str = strings.TrimSpace(str)
	//2.判断清除空格后是否为空字符
	if len(str) == 0 {
		return 0
	}
	//3.判断第一个字符是否为非法字符
	if !isDigit(str[0]) && str[0] != '+' && str[0] != '-' {
		return 0
	}
	//4.flag记录数字的正负
	flag := true
	if str[0] == '-' {
		flag = false
	}
	//5.如果第一个字符不是数字，说明第一个字符是'+'或者'-', 将第一个字符截掉
	if !isDigit(str[0]) {
		str = str[1:]
	}
	//6.将所有前缀'0'清除
	for i := 0; i < len(str); i++ {
		if str[i] != '0' {
			str = str[i:]
			break
		}
	}
	//7.32位最大数的位数为10位，如果str长度超过10位，只要截取前11位置即可
	length := min(len(str), 11)
	//8.将字符转成整数
	var ret int64
	for i := 0; i < length; i++ {
		if !isDigit(str[i]) {
			break
		}
		ret = ret*10 + int64(str[i]-'0')
	}
	//9.数字正负性
	if !flag {
		ret = -ret
	}
	//10.超过32位表示范围
	if ret > math.MaxInt32 {
		ret = math.MaxInt32
	}
	if ret < math.MinInt32 {
		ret = math.MinInt32
	}
	return int(ret)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isDigit(r byte) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}