/*
时间复杂度：O(1)
空间复杂度：O(1)
*/
package mathClass

func DistributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)
	i := 0
	for candies > 0 {
		if candies-i-1 >= 0 {
			ret[i%num_people] += i + 1
		} else {
			ret[i%num_people] += candies
		}
		//candies = candies - (i + 1)
		candies -= i + 1
		i++
	}
	return ret
}