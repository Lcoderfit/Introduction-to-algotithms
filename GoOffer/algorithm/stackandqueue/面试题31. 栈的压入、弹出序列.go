package stackandqueue

import "strings"

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}

	j := 0
	stack := make([]int, 0)
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for j < len(popped) && len(stack) != 0 && stack[len(stack) - 1] == popped[j] {
			stack = stack[:len(stack) - 1]
			j++
		}
	}
	return len(stack) == 0
}

func test() {
	strings.Join()

}