package str

import "fmt"

//hashtable, 滑动窗口
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) || len(s) == 0 || len(p) == 0 {
		return []int{}
	}
	ret := make([]int, 0)
	sArr, pArr := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		//将p的字符进行hashTable处理
		pArr[p[i]-'a']++
		//将s的前len(p)个字符进行hashtable处理,初始化滑动窗口值
		sArr[s[i]-'a']++
	}

	fmt.Println(pArr)
	fmt.Println(sArr)
	//固定滑动窗口长度
	i, j := 0, len(p) - 1
	for j < len(s) {
		if match(sArr, pArr) {
			ret = append(ret, i)
		}

		if j == len(s) - 1 {
			break
		}
		//滑动窗口向右移动，减去左指针位置对应字符的次数，添加右指针对应字符的次数
		sArr[s[i]-'a']--
		i++
		j++
		sArr[s[j]-'a']++
	}
	return ret
}

func match(sArr, pArr [26]int) bool {
	for i := 0; i < len(sArr); i++ {
		//窗口中的字母出现次数与p相等
		if sArr[i] != pArr[i] {
			return false
		}
	}
	return true
}