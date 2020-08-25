package doublePtr


/*

时间复杂度：O(n)，n是nums1的长度
空间复杂度：O(1)。
*/
func TwoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return nil
	}
	i, j := 0, len(numbers) - 1
	for i < j {
		if numbers[i] + numbers[j] < target {
			i++
		} else if numbers[i] + numbers[j] > target {
			j--
		} else {
			break
		}
	}
	return []int{i+1, j+1}
}