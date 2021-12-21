package task0231

// Given an integer n, return true if it is a power of two. Otherwise, return false.
// An integer n is a power of two, if there exists an integer x such that n == 2^x.
//
// Constraints:
//   -2^31 <= n <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Power of Two.
//   Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Power of Two.
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	x := n
	for x > 0 {
		if x&1 == 1 && x>>1 > 0 {
			return false
		}
		x = x >> 1
	}
	return true
}
