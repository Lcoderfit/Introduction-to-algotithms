package utils

import (
	"unicode"
)

//判断字符串中是否含有中文字符
func HasChineseChar(str string) bool {
	for _, s := range str {
		if unicode.Is(unicode.Han, s) {
			return true
		}
	}
	return false
}