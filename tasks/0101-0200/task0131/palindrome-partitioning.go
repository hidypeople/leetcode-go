package task0131

type StackStruct struct {
	idx    int
	length int
	path   []string
}

// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.
// A palindrome string is a string that reads the same backward as forward.
//
// Constraints:
//   1 <= s.length <= 16
//   s contains only lowercase English letters.
//
// Results:
//   Runtime: 284 ms, faster than 93.02% of Go online submissions for Palindrome Partitioning.
//   Memory Usage: 29 MB, less than 11.63% of Go online submissions for Palindrome Partitioning.
func partition(s string) [][]string {
	n := len(s)
	if n == 1 {
		return [][]string{{s}}
	}

	// DP: rows means "starting symbol", columns: substring length
	// e.g. s = "abbab", dp[1, 2] means substring s[1:1+2] if true - it is palindrom
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true // strings with length of 1 are palindroms
		if i > 0 && s[i-1] == s[i] {
			dp[i-1][1] = true // strings with length of 2 and s[i]==s[i+1] are palindroms
		}
	}
	for length := 3; length <= n; length++ {
		// left char index
		for l := 0; l <= n-length; l++ {
			// right char index
			r := l + length - 1
			dp[l][length-1] = true
			if s[l] == s[r] {
				// if left and right chars are equal -> we need to compare l+1 <-> r-1 result
				dp[l][length-1] = dp[l+1][length-3]
			} else {
				dp[l][length-1] = false
			}
		}
	}

	// Use BFS to build the result from DP structure
	result := make([][]string, 0)
	stack := []StackStruct{} // [2]int <=> [index, length]
	for i := 0; i < n; i++ {
		if dp[0][i] {
			stack = append(stack, StackStruct{
				idx: 0, length: i + 1, path: []string{s[:i+1]},
			})
		}
	}
	for len(stack) > 0 {
		stackItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nextIdx := stackItem.idx + stackItem.length
		if nextIdx >= n {
			result = append(result, stackItem.path)
			continue
		}
		for i := 0; i+nextIdx < n; i++ {
			if dp[nextIdx][i] {
				newPath := make([]string, len(stackItem.path)+1)
				copy(newPath, stackItem.path)
				newPath[len(newPath)-1] = s[nextIdx : nextIdx+i+1]
				stack = append(stack, StackStruct{
					idx: nextIdx, length: i + 1, path: newPath,
				})
			}
		}
	}
	return result
}
