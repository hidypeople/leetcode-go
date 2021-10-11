package task0089

// An n-bit gray code sequence is a sequence of 2n integers where:
// Every integer is in the inclusive range [0, 2^n - 1],
// The first integer is 0,
// An integer appears no more than once in the sequence,
// The binary representation of every pair of adjacent integers differs by exactly one bit, and
// The binary representation of the first and last integers differs by exactly one bit.
// Given an integer n, return any valid n-bit gray code sequence.
//
// Constraints:
//   1 <= n <= 16
//
// Results:
//   Runtime: 8 ms, faster than 98.18% of Go online submissions for Gray Code.
//   Memory Usage: 6.6 MB, less than 100.00% of Go online submissions for Gray Code.
func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	nLen := 2 << (n - 1)
	result := make([]int, nLen)
	copy(result, grayCode(n-1))
	for i := nLen>>1 - 1; i >= 0; i-- {
		result[nLen-i-1] = (1 << (n - 1)) | result[i]
	}
	return result
}
