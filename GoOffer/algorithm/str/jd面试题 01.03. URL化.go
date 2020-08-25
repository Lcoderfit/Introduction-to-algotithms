package str

import "strings"

//方法一：一行
func ReplaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}

//方法二：byte数组, 先统计空格，然后扩容，最后替换
func ReplaceSpaces1(S string, length int) string {
	//截取最左边长度为length的部分
	sBytes := []byte(S)[:length]
	//先计算空格有多少，根据空格扩容
	cnt := 0
	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] == ' ' {
			cnt++
		}
	}
	//lenght原本含有一个空格，所以只需要+2*cnt
	ret := make([]byte, length+2*cnt)
	j := 0
	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] == ' ' {
			ret[j], ret[j+1], ret[j+2] = '%', '2', '0'
			j += 3
		} else {
			ret[j] = sBytes[i]
			j++
		}
	}
	return string(ret)
}