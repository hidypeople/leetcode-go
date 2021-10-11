package task0022

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// Constraints:
//   1 <= n <= 8
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Generate Parentheses.
//   Memory Usage: 3.7 MB, less than 16.51% of Go online submissions for Generate Parentheses.
func generateParenthesis(n int) []string {
	return generate("", n, 0, 0)
}

func generate(currString string, n, numOpen, numClose int) []string {
	result := []string{}
	// If we have available open parentheses
	if numOpen < n {
		result = append(
			result,
			generate(currString+"(", n, numOpen+1, numClose)...)
	}
	// If we have available close parentheses and that number is less than available open parentheses
	if numClose < numOpen && numClose < n {
		result = append(
			result,
			generate(currString+")", n, numOpen, numClose+1)...)
	}
	// We came to the end: append whole string
	if numOpen == n && numClose == n {
		result = append(result, currString)
	}
	return result
}
