package mathTag


import "strings"

func IsNumber(s string) bool {
	s = strings.TrimSpace(s)
	foundE, foundPoint := false, false
	foundDigit := false

	for i, c := range s {
		if isDigit(c) {
			foundDigit = true
			continue
		}
		switch c {
		case '+', '-':
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		case 'E', 'e':
			if foundE || !foundDigit {
				return false
			}
			foundDigit = false
			foundE = true
		case '.':
			if foundPoint || foundE {
				return false
			}
			foundPoint = true
		default:
			return false
		}
	}

	return foundDigit
}

func isDigit(c rune) bool {
	v := int(c-'0')
	return v >= 0 && v <= 9
}