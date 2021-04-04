/*
方法1：位运算
时间复杂度：O(n)
空间复杂度：O(1)

方法2：动态规划+位运算(最高有效位)
时间复杂度：O(n)
空间复杂度：O(1)

方法3：动态规划+位运算(最低有效位)
时间复杂度：O(n)
空间复杂度：O(1)

方法4：动态规划+位运算(最低设置位)
时间复杂度：O(n)
空间复杂度：O(1)

case1:

r:

*/
package medium

// CountBits 通过布莱恩-肯尼跟算法得到每个数中的二进制表示中1的个数
func CountBits(num int) []int {
	ans := []int{0}
	for i := 1; i <= num; i++ {
		count := 0
		t := i
		for t > 0 {
			count++
			t = (t - 1) & t
		}
		ans = append(ans, count)
	}
	return ans
}

// CountBits2 最高有效位
func CountBits2(num int) []int {
	bits := make([]int, num+1)
	highBit := 0
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

// CountBits3 最低有效位
func CountBits3(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + (i & 1)
	}
	return bits
}

// CountBits4 最低设置位
func CountBits4(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
