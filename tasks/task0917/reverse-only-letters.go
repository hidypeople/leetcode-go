package task0917

// Given a string s, reverse the string according to the following rules:
// All the characters that are not English letters remain in the same position.
// All the English letters (lowercase or uppercase) should be reversed.
// Return s after reversing it.
//
// Constraints:
//   1 <= s.length <= 100
//   s consists of characters with ASCII values in the range [33, 122].
//   s does not contain '\"' or '\\'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Only Letters.
//   Memory Usage: 2.2 MB, less than 11.43% of Go online submissions for Reverse Only Letters.
func reverseOnlyLetters(s string) string {
	isChar := func(chr byte) bool {
		return ('a' <= chr && chr <= 'z') || ('A' <= chr && chr <= 'Z')
	}

	sBytes := []byte(s)
	left, right := 0, len(s)-1

	for left < right {

		// Skip non char symbols
		skip := false
		sLeft := sBytes[left]
		if !isChar(sLeft) {
			left++
			skip = true
		}

		sRight := sBytes[right]
		if !isChar(sRight) {
			right--
			skip = true
		}
		if skip {
			continue
		}

		// Swap chars
		sBytes[left], sBytes[right] = sBytes[right], sBytes[left]
		left++
		right--
	}
	return string(sBytes)
}
