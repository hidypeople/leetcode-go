package task0020

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
// An input string is valid if:
//   Open brackets must be closed by the same type of brackets.
//   Open brackets must be closed in the correct order.
//
// Constraints:
//   1 <= s.length <= 104
//   s consists of parentheses only '()[]{}'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
//   Memory Usage: 2 MB, less than 77.59% of Go online submissions for Valid Parentheses.
func isValid(s string) bool {
	if len(s)%2 > 0 {
		return false
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		chr := s[i]
		if chr == '{' || chr == '[' || chr == '(' {
			// stack.push
			stack = append(stack, chr)
		} else {
			stackLen := len(stack)
			if stackLen == 0 {
				return false
			}
			// stack.pop
			chr2 := stack[stackLen-1]
			stack = stack[:stackLen-1]
			if (chr2 == '{' && chr != '}') || (chr2 == '[' && chr != ']') || (chr2 == '(' && chr != ')') {
				return false
			}
		}
	}
	return len(stack) == 0
}
