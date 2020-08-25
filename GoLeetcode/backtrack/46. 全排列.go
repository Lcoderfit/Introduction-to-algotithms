package backtrack


func permute(nums []int) [][]int {
	var ret [][]int
	var trace []int
	visited := make([]bool, len(nums))

	var backtrace func(trace []int)
	backtrace = func(trace []int) {
		if len(nums) == len(trace) {
			tmp := make([]int, len(trace))
			copy(tmp, trace)
			ret = append(ret, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			trace = append(trace, nums[i])
			visited[i] = true
			backtrace(trace)
			visited[i] = false
			trace = trace[:len(trace) - 1]
		}
	}
	backtrace(trace)
	return ret
}