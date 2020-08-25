package str

import (
	"strconv"
	"strings"
)

//自解法
func compareVersion(version1 string, version2 string) int {
	sArr1, sArr2 := strings.Split(version1, "."), strings.Split(version2, ".")
	delta := len(sArr1) - len(sArr2)
	if delta < 0 {
		delta = -delta
	}
	if len(sArr1) < len(sArr2) {
		for i := 0; i < delta; i++ {
			sArr1 = append(sArr1, "0")
		}
	} else {
		for i := 0; i < delta; i++ {
			sArr2 = append(sArr2, "0")
		}
	}
	maxLen := max(len(sArr1), len(sArr2))
	for i := 0; i < maxLen; i++ {
		m, _ := strconv.Atoi(sArr1[i])
		n, _ := strconv.Atoi(sArr2[i])
		if m < n {
			return -1
		} else if m > n {
			return 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//适合大字符串的解法（字符转整数出现溢出）
func CompareVersion1(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	//将两个字符slice修改成一样长的
	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v1) > len(v2) {
		v2 = append(v2, "0")
	}

	length := len(v1)
	for i := 0; i < length; i++ {
		//去除左边的前导0
		vs1 := strings.TrimLeft(v1[i], "0")
		vs2 := strings.TrimLeft(v2[i], "0")
		//将短的字符前面补0, 这样就可以从高位到低位比较
		for len(vs1) < len(vs2) {
			vs1 = "0" + vs1
		}
		for len(vs1) > len(vs2) {
			vs2 = "0" + vs2
		}
		//从高位到低位比较字符
		for j := 0; j < len(vs1); j++ {
			if vs1[j] > vs2[j] {
				return 1
			} else if vs1[j] < vs2[j] {
				return -1
			}
		}
	}
	return 0
}