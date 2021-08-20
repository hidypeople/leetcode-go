package task0028

import "strings"

// Implement strStr().
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
// Clarification:
// What should we return when needle is an empty string? This is a great question to ask during an interview.
// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
//
// Constraints:
//   0 <= haystack.length, needle.length <= 5 * 104
//   haystack and needle consist of only lower-case English characters.
//
// Result:
//   Runtime: 4 ms, faster than 66.29% of Go online submissions for Implement strStr().
//   Memory Usage: 2.5 MB, less than 49.09% of Go online submissions for Implement strStr().
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// Btw:
// "strings.Index(haystack, needle)" gets the 0ms result...
func strStr_Fastest(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
