package binaryFind
/*
392. 判断子序列.py
时间复杂度：O(m+n)
空间复杂度：O(1)
*/

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	for (i < len(s)) && (j < len(t)) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	if i == len(s) {
		return true
	}
	return false
}

