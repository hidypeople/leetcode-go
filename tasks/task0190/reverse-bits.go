package task0190

// Reverse bits of a given 32 bits unsigned integer.
// Note:
//   Note that in some languages, such as Java, there is no unsigned integer type.
//   In this case, both input and output will be given as a signed integer type.
//   They should not affect your implementation, as the integer's internal binary
//   representation is the same, whether it is signed or unsigned.
//   In Java, the compiler represents the signed integers using 2's complement notation.
//   Therefore, in Example 2 above, the input represents the signed integer -3 and
//   the output represents the signed integer -1073741825.
//
// Constraints:
//   The input must be a binary string of length 32
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Bits.
//   Memory Usage: 2.5 MB, less than 85.96% of Go online submissions for Reverse Bits.
func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num>>i)&1
	}
	return result
}
