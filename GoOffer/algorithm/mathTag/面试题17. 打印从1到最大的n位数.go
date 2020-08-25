package mathTag


func PrintNumbers(n int) []int {
	end := 9
	for i := 1; i < n; i++ {
		end = end*10 + 9
	}
	// end := int(math.Pow10(n)) - 1
	ret := make([]int, 0)
	for i := 1; i <= end; i++ {
		ret = append(ret, i)
	}
	return ret
}