package task0150

import "strconv"

// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
// Note that division between two integers should truncate toward zero.
// It is guaranteed that the given RPN expression is always valid. That means the expression
// would always evaluate to a result, and there will not be any division by zero operation.
//
// Constraints:
//   1 <= tokens.length <= 10^4
//   tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Evaluate Reverse Polish Notation.
//   Memory Usage: 4.2 MB, less than 91.10% of Go online submissions for Evaluate Reverse Polish Notation.
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			x, y := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var res int
			switch token {
			case "+":
				res = x + y
			case "-":
				res = x - y
			case "*":
				res = x * y
			case "/":
				res = x / y
			}
			stack[len(stack)-1] = res
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
