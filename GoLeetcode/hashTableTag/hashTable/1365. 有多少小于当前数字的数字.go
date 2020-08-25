/*
2 <= nums.length <= 500
0 <= nums[i] <= 100
时间复杂度：O(s+n)
空间复杂度：O(s+n)
*/
package hashTable

//频次数组+前缀和
func SmallNumberThanCurrent(nums []int) []int {
	cnt := make([]int, 101)
	//频次数组
	for _, v := range nums {
		cnt[v]++
	}
	//设置前缀和
	for i := 1; i < 101; i++ {
		cnt[i] += cnt[i-1]
	}

	//整合结果
	n := len(nums)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		//nums[i] == 0时，由于没有比0更小的数，所以就是cnt[0]
		if nums[i] != 0 {
			ret[i] = cnt[nums[i] - 1]
		}
	}

	return ret
}
