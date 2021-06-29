/*
方法1：贪心法（上下限）
时间复杂度：O(n)
空间复杂度：O(1)

方法2：潮汐法
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:
*/
package medium

func CheckValidString(s string) bool {
	l, h := 0, 0
	for _, c := range s {
		if c == '(' {
			l++
			h++
		} else if c == '*' {
			if l > 0 {
				l--
			}
			h++
		} else {
			if l > 0 {
				l--
			}
			h--
		}
		if h < 0 {
			return false
		}
	}
	return l == 0
}

func CheckValidString1(s string) bool {
	// rightBracketCount表示可能作为右括号的星号的数量
	h, startCount, rightBracketCount := 0, 0, 0
	for _, c := range s {
		if c == '(' {
			h++
		} else if c == ')' {
			h--
			if h < 0 {
				if h+startCount < 0 {
					return false
				} else {
					startCount += h
					h = 0
				}
			}
		} else {
			startCount++
			rightBracketCount++
		}
		// 这一步的目的，其实是为了括号匹配有序性，因为如果'(**(', 此时有两个星号，h=1， 能作为右括号的星号只能为1个，多出来的一个星号是不能与最右边的'('匹配的
		// 所以需要先取h与rightBracketCount的最小值，将多出来的作为右括号的星号去除
		rightBracketCount = Min(rightBracketCount, startCount)
		rightBracketCount = Min(rightBracketCount, h)
	}
	return rightBracketCount >= h
}
