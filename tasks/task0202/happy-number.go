package task0202

// Write an algorithm to determine if a number n is happy.
// A happy number is a number defined by the following process:
//   Starting with any positive integer, replace the number by the sum of the squares of its digits.
//   Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
//   Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
//
// Constraints:
//   1 <= n <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
//   Memory Usage: 2.1 MB, less than 72.64% of Go online submissions for Happy Number.
func isHappy(n int) bool {
	previous := map[int]bool{}
	for n > 2 {
		x, next := n, 0
		for x > 0 {
			next += (x % 10) * (x % 10)
			x = x / 10
		}
		if previous[next] || n == 1 {
			// if we came into the cycle => return
			return n == 1
		}
		previous[next] = true
		n = next
	}
	return n == 1
}
