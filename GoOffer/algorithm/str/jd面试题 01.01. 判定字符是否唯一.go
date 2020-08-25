package str


//bitMap法找重复
func IsUnique(astr string) bool {
	bitMap := 0
	for i := 0; i < len(astr); i++ {
		if (bitMap & (1 << (astr[i] - 'a'))) != 0 {
			return false
		}
		//必须在astr[i] - 'a'左右两边加上括号，否则会先计算1 << astr[i]
		bitMap |= 1 << (astr[i] - 'a')
	}
	return true
}