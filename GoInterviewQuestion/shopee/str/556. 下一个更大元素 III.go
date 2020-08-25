package str

import (
	"math"
	"strconv"
)

//最优解
func NextGreaterElement(n int) int {
	s := strconv.Itoa(n)
	sByte := []byte(s)
	length := len(sByte)
	//找到第一个小于右边数字的位置，这个位置即为需要更换的位置
	i := length - 2
	for ; i >= 0 && sByte[i] >= sByte[i+1]; i-- {}
	if i < 0 {
		return -1
	}

	//从最低位开始找，找到第一个大于i位置处数字的数，然后替换这两个数字
	j := length - 1
	for ; j > i && sByte[j] <= sByte[i]; j-- {}
	//因为i的右边肯定存在比位置i大的数，所以此处不需要考虑找不到sByte[j] > sByte[i]的情况
	sByte[i], sByte[j] = sByte[j], sByte[i]

	//此时i位置右边数字为从大到小排列，需要逆序为从小到大排序
	for m, n := i+1, length - 1; m < n; m, n = m+1, n-1 {
		sByte[m], sByte[n] = sByte[n], sByte[m]
	}
	//n经过变换后，可能会超出int32范围
	ret, _ := strconv.Atoi(string(sByte))
	if ret > math.MaxInt32 {
		return -1
	}
	return ret
}

//自己想的版本
func NextGreaterElement1(n int) int {
	if n < 10 {
		return -1
	}
	s := strconv.Itoa(n)
	sBytes := []byte(s)
	i := 0
	for i = len(sBytes) - 2; i >= 0; i-- {
		if sBytes[i] < sBytes[i+1] {
			break
		}
	}
	if i < 0 {
		return -1
	}
	j := 0
	for j = len(sBytes) - 1; j > i; j-- {
		if sBytes[j] > sBytes[i] {
			sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
			break
		}
	}
	if j == i {
		return -1
	}

	//将末尾部分逆序
	for m, n := i+1, len(sBytes) - 1; m < n; m, n = m+1, n-1 {
		sBytes[m], sBytes[n] = sBytes[n], sBytes[m]
	}
	ret, _ := strconv.Atoi(string(sBytes))
	if ret > math.MaxInt32 {
		return -1
	}
	return ret
}
