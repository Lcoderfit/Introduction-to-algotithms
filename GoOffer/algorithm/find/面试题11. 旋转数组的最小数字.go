package find


func MinArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	i, j := 0, len(numbers) - 1
	for i < j {
		mid := (i+j)/2
		if numbers[mid] > numbers[j] {
			i = mid + 1
		} else if numbers[mid] < numbers[0] {
			j = mid
		} else {
			j--
		}
	}
	return numbers[i]
}