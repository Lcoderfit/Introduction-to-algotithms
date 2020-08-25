package doublePtr

import "unicode"


/*
双指针解法-前后
时间复杂度：O(n)，n是s的长度
空间复杂度：O(1)。
*/
func IsPalindrome(s string) bool {
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}

	i, j := 0, len(s) - 1
	for i < j {
		vi, vj := rune(s[i]), rune(s[j])
		if !isValid(vi) && !isValid(vj) {
			i++
			j--
		} else if !isValid(vi) {
			i++
		} else if !isValid(vj) {
			j--
		} else if unicode.ToUpper(vi) == unicode.ToUpper(vj) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}