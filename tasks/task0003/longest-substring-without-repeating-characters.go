package task0003

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Given a string s, find the length of the longest substring without repeating characters.
//
// Constraints:
//   0 <= s.length <= 5 * 104
//   s consists of English letters, digits, symbols and spaces.
//
// Results:
//   Runtime: 4 ms, faster than 90.69% of Go online submissions for Longest Substring Without Repeating Characters.
//   Memory Usage: 3.2 MB, less than 44.02% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	max := 0
	left := 0
	for right, c := range s {
		seenVal, isSeen := seen[c]
		if isSeen {
			left = Max(seenVal, left)
		}
		max = Max(max, right-left+1)
		seen[c] = right + 1
	}
	return max
}
