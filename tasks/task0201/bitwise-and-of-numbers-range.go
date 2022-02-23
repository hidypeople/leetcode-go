package task0201

// Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.
//
// Constraints:
//   0 <= left <= right <= 231 - 1
//
// Results:
//   Runtime: 4 ms, faster than 98.90% of Go online submissions for Bitwise AND of Numbers Range.
//   Memory Usage: 5.1 MB, less than 100.00% of Go online submissions for Bitwise AND of Numbers Range.
func rangeBitwiseAnd(left int, right int) int {
	powOf2, result, diff := 1, right&left, right-left
	for diff >= powOf2 {
		result &= ^powOf2
		powOf2 <<= 1
	}
	return result
}
