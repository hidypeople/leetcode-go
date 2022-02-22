package task0191

// Write a function that takes an unsigned integer and returns the
// number of '1' bits it has (also known as the Hamming weight).
// Note:
//   Note that in some languages, such as Java, there is no unsigned integer type.
//   In this case, the input will be given as a signed integer type.
//   It should not affect your implementation, as the integer's internal binary representation is
//   the same, whether it is signed or unsigned.
//   In Java, the compiler represents the signed integers using 2's complement notation.
//   Therefore, in Example 3, the input represents the signed integer. -3.
//
// Constraints:
//   The input must be a binary string of length 32.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
//   Memory Usage: 1.9 MB, less than 75.62% of Go online submissions for Number of 1 Bits.
func hammingWeight(num uint32) int {
	var result int
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			result++
		}
		num >>= 1
	}
	return result
}
