package task0005

// Given a string s, return the longest palindromic substring in s.
//
// Constraints:
//   1 <= s.length <= 1000
//   s consist of only digits and English letters.
//
// Results:
//   Runtime: 4 ms, faster than 96.51% of Go online submissions for Longest Palindromic Substring.
//   Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.
func longestPalindrome(s string) string {
	N := len(s)
	if N < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < N; i++ {
		len1 := growLeftRight(s, i, i)
		len2 := growLeftRight(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func growLeftRight(s string, left int, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
