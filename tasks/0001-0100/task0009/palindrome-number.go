package task0009

// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// For example, 121 is palindrome while 123 is not.
//
// Constraints:
//   -2^31 <= x <= 2^31 - 1
//
// Results (unstable from 4ms to 20ms):
//   Runtime: 4 ms, faster than 98.30% of Go online submissions for Palindrome Number.
//   Memory Usage: 5.1 MB, less than 81.56% of Go online submissions for Palindrome Number.
func isPalindrome(x int) bool {
	if x <= 0 {
		return x == 0
	}
	var dig int
	curr := x
	rev := 0
	for curr > 0 {
		dig = curr % 10
		rev = rev*10 + dig
		curr = curr / 10
	}
	return x == rev
}
