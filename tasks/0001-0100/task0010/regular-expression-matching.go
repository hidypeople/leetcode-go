package task0010

// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).
//
// Constraints:
//   1 <= s.length <= 20
//   1 <= p.length <= 30
//   s contains only lowercase English letters.
//   p contains only lowercase English letters, '.', and '*'.
//   It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Regular Expression Matching.
//   Memory Usage: 2.4 MB, less than 64.37% of Go online submissions for Regular Expression Matching.
func isMatch(s string, p string) bool {
	var dfs func(i, j int) bool
	cache := map[[2]int]bool{}
	dfs = func(i, j int) bool {
		if v, ok := cache[[2]int{i, j}]; ok {
			return v
		}
		if i >= len(s) && j >= len(p) {
			return true
		}
		if j >= len(p) {
			return false
		}
		match := false
		if i < len(s) && (s[i] == p[j] || p[j] == '.') {
			match = true
		}
		if (j+1) < len(p) && p[j+1] == '*' {
			cache[[2]int{i, j}] = dfs(i, j+2) || (match && dfs(i+1, j))
			return cache[[2]int{i, j}]
		}
		if match {
			cache[[2]int{i, j}] = dfs(i+1, j+1)
			return cache[[2]int{i, j}]
		}
		cache[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0)
}
