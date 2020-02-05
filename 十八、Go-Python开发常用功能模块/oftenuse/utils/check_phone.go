package utils

import "regexp"

func CheckPhone(phone string) bool {
	if len(phone) <= 0 {
		return false
	}
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone); !m {
		return false
	}
	return true
}
