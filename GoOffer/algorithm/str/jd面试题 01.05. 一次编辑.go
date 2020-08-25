package str

//双指针，题解最优版
func OneEditAway(first string, second string) bool {
	l1, l2 := len(first), len(second)
	if l1 - l2 > 1 || l1 - l2 < -1 {
		return false
	}

	i := 0
	for i < l1 && i < l2 {
		//如果first和second满足要求，则当出现一个字符不等时，剩下的字符一定相等
		if first[i] != second[i] {
			if l1 > l2 {
				return first[i+1:] == second[i:]
			} else if l1 < l2 {
				return first[i:] == second[i+1:]
			} else {
				return first[i+1:] == second[i+1:]
			}
		}
		i++
	}
	return true
}