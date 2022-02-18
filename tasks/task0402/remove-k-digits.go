package task0402

import "strings"

// Given string num representing a non-negative integer num, and an integer k,
// return the smallest possible integer after removing k digits from num.
//
// Constraints:
//   1 <= k <= num.length <= 10^5
//   num consists of only digits.
//   num does not have any leading zeros except for the zero itself.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove K Digits.
//   Memory Usage: 3 MB, less than 73.08% of Go online submissions for Remove K Digits.
func removeKdigits(num string, k int) string {
	n := len(num)
	if k >= n {
		return "0"
	}

	// This stack will collect all accepted digits
	stack := []byte{}
	i := 0
	for i < n {
		// Remove the top items from stack until current number is lesser that stack top
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		// Put current number to stack
		stack = append(stack, num[i])
		i++
	}

	if k > 0 {
		// If we've completed the cycle, but there is still some digits to remove:
		// remove the rest of them
		stack = stack[:len(stack)-k]
	}

	// Build the result string (+skip leading zeroes)
	countAdded := 0
	builder := new(strings.Builder)
	for _, chr := range stack {
		if countAdded == 0 && chr == '0' {
			continue
		}
		countAdded++
		builder.WriteByte(chr)
	}
	if countAdded == 0 {
		return "0"
	}
	return builder.String()
}
