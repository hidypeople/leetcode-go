package task0038

import "strings"

// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
// countAndSay(1) = "1"
// countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1),
// which is then converted into a different digit string.
// To determine how you "say" a digit string, split it into the minimal number
// of groups so that each group is a contiguous section all of the same character.
// Then for each group, say the number of characters, then say the character.
// To convert the saying into a digit string, replace the counts with a number
// and concatenate every saying.
//
// Constraints:
//   1 <= n <= 30
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Count and Say.
//   Memory Usage: 2.5 MB, less than 86.55% of Go online submissions for Count and Say.
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	// Run recursively
	sayResult := countAndSay(n - 1)

	// Calc next sayResult
	builder := new(strings.Builder)
	var prevCount byte = 1
	prevChar := sayResult[0]
	for i := 1; i < len(sayResult); i++ {
		currChar := sayResult[i]
		if currChar != prevChar {
			// sequence interrupted:
			builder.WriteByte('0' + prevCount)
			builder.WriteByte(prevChar)
			prevCount = 1
			prevChar = currChar
		} else {
			prevCount++
		}
	}
	builder.WriteByte('0' + prevCount)
	builder.WriteByte(prevChar)
	return builder.String()
}
