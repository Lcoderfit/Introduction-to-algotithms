package utils

import (
	"fmt"
	"regexp"
	"unicode"
)

/*
	regSpecial := "[^`~!@#$^&*()=|{}':;',\\[\\].<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？+-]"
	ok, _ := regexp.MatchString(regSpecial, "lcoder12`4541")
	fmt.Println(ok)
	//对模式进行编译
	res := regexp.MustCompile(regSpecial)
	//n为-1表示返回全部结果，n == 2表示返回匹配到的两个结果
	fmt.Println(res.FindAllString("lcoder12`4541", 2))
 */
func CheckPassword(password string) (bool, error) {
	if len(password) < 8 || len(password) > 16 {
		return false, fmt.Errorf("密码长度必须为8~16位")
	}

	//判断是否含有特殊字符
	regSpecial := "[`~!@#$^&*()=|{}':;',\\[\\].<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？+-]"
	if ok, _ := regexp.MatchString(regSpecial, password); ok {
		return false, fmt.Errorf("密码中不能含有特殊字符")
	}

	//判断密码中是否含有中文
	for _, s := range password {
		if unicode.Is(unicode.Han, s) {
			return false, fmt.Errorf("密码中不能包含中文")
		}
	}

	//判断是否含有大写、小写字母、数字中的至少两种
	regUpper := `[A-Z]`
	regLower := `[a-z]`
	regNumber := `[0-9]`
	//在[abc]前面加上^则表示取反：[^abc]
	matchCount := 0
	if ok, _ := regexp.MatchString(regUpper, password); ok {
		matchCount++
	}
	if ok, _ := regexp.MatchString(regLower, password); ok {
		matchCount++
	}
	if ok, _ := regexp.MatchString(regNumber, password); ok {
		matchCount++
	}

	if matchCount < 2 {
		return false, fmt.Errorf("必须包含数字、字母大、小写中的至少两种")
	}

	return true, nil
}
