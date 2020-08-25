package trace


var ret [][]int

//一：回溯
func Permute(nums []int) [][]int {
	ret = make([][]int, 0)
	trace := make([]int, 0)
	backtrace(nums, trace)
	return ret
}

func backtrace(nums, trace []int) {
	if len(trace) == len(nums) {
		//注意，这里一定要再另开辟一个slice复制trace的内容
		//因为
		tmp := make([]int, len(nums))
		copy(tmp, trace)
		ret = append(ret, tmp)
		return
	}

	//遍历选择列表
	var j int
	for i := 0; i < len(nums); i++ {
		for j = 0; j < len(trace); j++ {
			if trace[j] == nums[i] {
				break
			}
		}
		//j != len(trace)说明trace中存在与nums[i]相等的元素
		if j != len(trace) {
			continue
		}
		//做选择
		trace = append(trace, nums[i])
		backtrace(nums, trace)
		//取消做选择
		trace = trace[:len(trace) - 1]
	}
}


//递归, 需要消耗更多的内存
func Permute1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		buf := make([]int, len(nums) - 1)
		copy(buf[:i], nums[:i])
		copy(buf[i:], nums[i+1:])
		tail := Permute1(buf)
		for _, v := range tail {
			//嵌套append
			ret = append(ret, append(v, nums[i]))
		}
	}
	return ret
}

//三：回溯简化
func permute3(nums []int) [][]int {
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