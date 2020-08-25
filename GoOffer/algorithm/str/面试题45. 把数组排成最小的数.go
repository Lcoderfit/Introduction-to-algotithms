package str

import (
	"sort"
	"strconv"
	"strings"
)


//方法一：库函数
func MinNumber(nums []int) string {
	sArr := make([]string, len(nums))
	for i, v := range nums {
		sArr[i] = strconv.Itoa(v)
	}
	//匿名函数返回false，则交换sArr[i]和sArr[j], 返回true则不交换
	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] + sArr[j] < sArr[j] + sArr[i]
	})

	//字符串缓冲区
	bd := strings.Builder{}
	for _, s := range sArr {
		bd.WriteString(s)
	}
	return bd.String()

	//或者直接用下面的
	//return strings.Join(sArr, "")
}

//方法二：常规方法
func MinNumber2(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	length := len(nums)
	sArr := make([]string, length)
	for i := range nums {
		sArr[i] = strconv.Itoa(nums[i])
	}
	for i := 0; i < length; i++ {
		for j := i+1; j < length; j++ {
			f, b := sArr[i], sArr[j]
			if f + b > b + f {
				sArr[i], sArr[j] = sArr[j], sArr[i]
			}
		}
	}
	ret := ""
	for _, s := range sArr {
		ret += s
	}
	return ret
}