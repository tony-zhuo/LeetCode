package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		}
	}

	return stack[0]
}
