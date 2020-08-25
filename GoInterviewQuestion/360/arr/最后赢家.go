package arr


func LastWinner(array []int, m int) int {
	if len(array) == 1 {
		return 0
	}
	maxVal := 0
	pos := 0
	for i := 0; i < len(array); i++ {
		if maxVal < array[i] {
			maxVal = array[i]
			pos = i
		}
	}
	if m >= len(array) - 1 {
		if pos == 0 {
			return m
		} else {
			return pos+m-1
		}
	}
	count := 0
	cur := 0
	for {
		if array[0] < array[1] {
			cur = 1
			array = append(array[1:], array[0])
		} else {
			cur++
			tmp := array[1]
			array = append(array[0:1], array[2:]...)
			array = append(array, tmp)
		}
		count++
		if cur == m {
			break
		}
	}
	return count
}