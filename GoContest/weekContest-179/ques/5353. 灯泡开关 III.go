package ques


func NumTimesAllBlue(light []int) int {
	n := len(light)
	cnt := 0
	status := make([]int, n)
	for i := 0; i < n; i++ {
		if check(light, i, status) {
			cnt++
		}
		status[i] = 1
	}
	return cnt
}

func check(light []int, i int, status []int) bool {
	for j := 0; j < i; j++ {
		if light[j] != 1 {
			return false
		}
	}
	cnt := 0
	for j := i+1; j < len(light); j++ {
		if status[j] == 1 {
			cnt++
		}
	}
	if cnt == 0 || cnt == (len(light)-i-1) {
		return true
	}
	return false
}
