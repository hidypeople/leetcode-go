package task0032

// Given a string containing just the characters '(' and ')',
// find the length of the longest valid (well-formed) parentheses substring.
//
// Constraints:
//   0 <= s.length <= 3 * 10^4
//   s[i] is '(', or ')'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Valid Parentheses.
//   Memory Usage: 3.7 MB, less than 22.75% of Go online submissions for Longest Valid Parentheses.
func longestValidParentheses(s string) int {
	stack := []int{} // Stack contains the index of open '('
	start := -1
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// ')'
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					// if stack is not empty -> we came to the end of some valid sequence:
					// use index stored in stack to calculate the length of this sequence
					result = max(result, i-stack[len(stack)-1])
				} else {
					// if stack is empty -> we came to zero state
					result = max(result, i-start)
				}
			} else {
				// Out of stack (too many ')')
				start = i
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
