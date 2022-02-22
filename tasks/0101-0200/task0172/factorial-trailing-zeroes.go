package task0172

// Given an integer n, return the number of trailing zeroes in n!.
// Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.
//
// Constraints:
//   0 <= n <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Factorial Trailing Zeroes.
//   Memory Usage: 2 MB, less than 49.23% of Go online submissions for Factorial Trailing Zeroes.
func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	// All zeroes comes from numbers like X = 5*y
	// We need to calculate how many 5 factors in the sequence [1..n]
	return n/5 + trailingZeroes(n/5)
}
