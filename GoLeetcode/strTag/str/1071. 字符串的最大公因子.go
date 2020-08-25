package str

//方法一：数学规律
/*
时间复杂度：O(n)，字符串拼接比较是否相等需要O(n)的时间复杂度，求两个字符串长度的最大公约数需要 O(log n) 的时间复杂度，
所以总时间复杂度为 O(n+logn)=O(n) 。

空间复杂度：O(n)，程序运行时建立了中间变量用来存储 str1 与 str2 的相加结果
 */
func GcdOfStrings(str1 string, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	gcd := Gcd(len(str1), len(str2))
	if str1 + str2 == str2 + str1 {
		return str1[:gcd]
	}
	return ""
}

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

//方法二
/*
时间复杂度：O(n)，其中n是两个字符串的长度范围，即len1+len2
​判断最大公约数长度的前缀串是否符合条件需要 O(n)的时间复杂度，求两个字符串长度的最大公约数需要O(logn)的时间复杂度，
所以总时间复杂度为 O(n+log n)=O(n) 。
空间复杂度：O(n)，比较的过程中需要创建一个长度为 O(n)的临时字符串变量，所以需要额外O(n)的空间。
 */
func GcdOfStrings1(str1 string, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	gcd := Gcd(len(str1), len(str2))
	subStr := str1[:gcd]

	s1, s2 := "", ""
	for i := 0; i < len(str1)/gcd; i++ {
		s1 += subStr
	}
	for i := 0; i < len(str2)/gcd; i++ {
		s2 += subStr
	}
	if s1 == str1 && s2 == str2 {
		return subStr
	}
	return ""
}