package problems

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, alpha := range s {
		switch alpha {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if pop != alpha {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
