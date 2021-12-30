package task1015

// Given a positive integer k, you need to find the length of the smallest positive integer
// n such that n is divisible by k, and n only contains the digit 1.
// Return the length of n. If there is no such n, return -1.
// Note: n may not fit in a 64-bit signed integer.
//
// Constraints:
//   1 <= k <= 10^5
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Smallest Integer Divisible by K.
//   Memory Usage: 2.1 MB, less than 50.00% of Go online submissions for Smallest Integer Divisible by K.
func smallestRepunitDivByK(k int) int {
	if k%2 == 2 || k%5 == 0 {
		return -1
	}

	// (a + b) % k = (a % k + b % k) % k
	//
	// 11%k = (10 % k + 1 % k) % k = ((1 % k) * 10 + 1) % k
	// 111%k = (100 % k + 10 % k + 1 % k) % k = ((((1 % k) * 10 + 1) % k) * 10 + 1) % k
	r := 0
	for n := 1; n <= k; n++ {
		r = (r*10 + 1) % k
		if r == 0 {
			return n
		}
	}
	return -1
}
