/*
时间复杂度：O(n)
空间复杂度：O(n)
*/
package hashTable


func DistributeCandies(candies []int) int {
	//bool类型默认为false
	hash := make(map[int]bool, 0)
	count := 0
	for i := 0; i < len(candies); i++ {
		if !hash[candies[i]]  {
			count++
			hash[candies[i]] = true
		}
	}

	return min(count, len(candies)/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
