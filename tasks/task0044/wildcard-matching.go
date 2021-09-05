package task0040

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
//   '?' Matches any single character.
//   '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).
//
// Constraints:
//   0 <= s.length, p.length <= 2000
//   s contains only lowercase English letters.
//   p contains only lowercase English letters, '?' or '*'.
//
// Results:
//   Runtime: 4 ms, faster than 95.86% of Go online submissions for Wildcard Matching.
//   Memory Usage: 3 MB, less than 91.72% of Go online submissions for Wildcard Matching.
func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	is, ip, asterisk, match := 0, 0, -1, 0
	for is < lenS {
		if ip < lenP && p[ip] == '*' {
			// If current pattern character is asterisk: we need to remember the position for both strings
			// and increment pattern index
			match = is
			asterisk = ip
			ip++
		} else if ip < lenP && (s[is] == p[ip] || p[ip] == '?') {
			// Non-asterisk chars match: just increment both indexes
			is++
			ip++
		} else if asterisk >= 0 {
			// If we are in the asterisk zone (but chars are not matching):
			// 1. We need to move forward on the "s" string
			// 2. We need to reset P to the next of asterisk character
			match++
			is = match
			ip = asterisk + 1
		} else {
			return false
		}
	}
	// Skip tail asterisks...
	for ip < lenP && p[ip] == '*' {
		ip++
	}
	return ip == lenP
}
