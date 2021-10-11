package task0070

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Constraints:
//   1 <= n <= 45
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
//   Memory Usage: 1.9 MB, less than 68.04% of Go online submissions for Climbing Stairs.
func climbStairs(n int) int {
	// We need to calculate Fibonacci(n)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n0 := 1
	n1 := 2
	var x int
	for i := 3; i <= n; i++ {
		x = n0 + n1
		n0, n1 = n1, x
	}
	return n1
}
