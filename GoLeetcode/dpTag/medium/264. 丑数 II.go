/*
方法1：动态规划
时间复杂度：O(C)
空间复杂度：O(C)

case1:
10
r:
12
*/
package medium

// Ugly 丑数类
type Ugly struct {
	nums []int
}

// UglyConstructor 丑数类的构造函数
func UglyConstructor() Ugly {
	ugly := Ugly{nums: make([]int, 1)}
	ugly.GenerateNumberList()
	return ugly
}

// UglyNumberList 初始化丑数列表，长度为1690
func (that *Ugly) GenerateNumberList() {
	that.nums[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i <= 1690; i++ {
		next := Min(that.nums[i2]*2, that.nums[i3]*3, that.nums[i5]*5)
		that.nums = append(that.nums, next)

		if that.nums[i2]*2 == next {
			i2++
		} else if that.nums[i3]*3 == next {
			i3++
		} else {
			i5++
		}
	}
}

// NthUglyNumber 计算第n个丑数
func (that *Ugly) NthUglyNumber(n int) int {
	if n > 1690 {
		return -1
	}
	return that.nums[n]
}

func NthUglyNumber2(n int) int {
	nums := []int{1}
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x2, x3, x5 := nums[i2]*2, nums[i3]*3, nums[i5]*5
		next := Min(x2, x3, x5)
		nums = append(nums, next)

		// 为了避免出现重复元素，需要用if，而不是else if
		if x2 == next {
			i2++
		}
		if x3 == next {
			i3++
		}
		if x5 == next {
			i5++
		}
	}
	return nums[n-1]
}
