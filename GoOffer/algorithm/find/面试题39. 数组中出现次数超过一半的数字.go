/*
摩尔投票法，有则+1，没有则-1
 */
package find

func MajorityElement(nums []int) int {
	ret, votes := 0, 0
	for _, v := range nums {
		if votes == 0 {
			ret = v
		}
		if ret == v {
			votes++
		} else {
			votes--
		}
	}
}
