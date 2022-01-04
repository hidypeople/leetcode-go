package task1009

// The complement of an integer is the integer you get when you flip all the 0's to 1's
// and all the 1's to 0's in its binary representation.
// For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
// Given an integer n, return its complement.
//
// Constraints:
//   0 <= n < 10^9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Complement of Base 10 Integer.
//   Memory Usage: 2.1 MB, less than 16.00% of Go online submissions for Complement of Base 10 Integer.
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	pow, complement := 0, 0
	for n > 0 {
		complement |= (^n & 1) << pow
		n = n >> 1
		pow++
	}
	return complement
}
