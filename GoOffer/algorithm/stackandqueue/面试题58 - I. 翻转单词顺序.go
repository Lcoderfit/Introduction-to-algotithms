package stackandqueue

import "strings"

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.Trim(s, " ")
	str := strings.Split(s, " ")
	strList := make([]string, 0)
	for _, st := range str {
		if st != "" {
			strList = append(strList, st)
		}
	}
	for i, j := 0, len(strList) - 1; i < j; i, j = i + 1, j-1 {
		strList[i], strList[j] = strList[j], strList[i]
	}
	return strings.Join(strList, " ")
}
