package str

import "math"

//自写版
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ret := make([]byte, 0)
	minLength := math.MaxInt64
	for _, v := range strs {
		if minLength > len(v) {
			minLength = len(v)
		}
	}
	for i := 0; i < minLength; i++ {
		t := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if t != strs[j][i] {
				return string(ret)
			}
		}
		ret = append(ret, t)
	}
	return string(ret)
}

//空间优化
func LongestCommonPrefix1(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			//如果strs[0]是公共前缀，且他是最短的，则直接返回strs[0]
			//如果strs[0]不是最短的，那么位置为i时，s[0...i-1]个字符就是最长公共前缀
			if i == len(strs[j]) || c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}