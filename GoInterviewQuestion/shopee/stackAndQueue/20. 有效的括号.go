package stackAndQueue


//自写版
func IsValid(s string) bool {
	//用辅助栈，如果遇到左括号，一律入栈
	//遇到右括号，取栈顶元素，如果匹配则出栈，不匹配直接返回false
	if s == "" {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				t := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				if (t == '(' && v != ')') || (t == '[' && v != ']') || (t == '{' && v != '}') {
					return false
				}
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

//最优版
func IValid(s string) bool {
	//用辅助栈，如果遇到左括号，一律入栈
	//遇到右括号，取栈顶元素，如果匹配则出栈，不匹配直接返回false
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' {
			stack = append(stack, ')')
		} else if v == '[' {
			stack = append(stack, ']')
		} else if v == '{' {
			stack = append(stack, '}')
		} else {
			//栈顶元素是最内层括号配对的右括号
			if len(stack) == 0 || stack[len(stack) - 1] != v {
				return false
			} else {
				//匹配，出栈
				stack = stack[:len(stack) - 1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}