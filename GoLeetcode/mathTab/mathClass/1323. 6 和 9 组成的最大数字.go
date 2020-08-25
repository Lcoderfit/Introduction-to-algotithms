package mathClass

import (
	"math"
	"strconv"
)

/*
时间复杂度：O(logn)，其中n是num的数值大小
空间复杂度：O(logn)。
 */
//数字转字符串
func Maximum69Number (num int) int {
	sTmp := strconv.Itoa(num)
	sNum := []byte(sTmp)
	for i := 0; i < len(sNum); i++ {
		if sNum[i] == '6' {
			sNum[i] = '9'
			break
		}
	}
	ret, _ := strconv.Atoi(string(sNum))
	return ret
}


/*
时间复杂度：O(logn)，其中n是num的数值大小
空间复杂度：O(1)。
*/
//找出最高位的数字为6的位置
func Maximum69Number1 (num int) int {
	ret := num
	flag, count := -1, -1
	for num != 0 {
		count++
		if num%10 == 6 {
			flag = count
		}
		num = num/10
	}
	if flag == -1 {
		return ret
	}
	return ret + 3*int(math.Pow(10.0, float64(flag)))
}