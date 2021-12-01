package task0137

// Given an integer array nums where every element appears three times except for one,
// which appears exactly once. Find the single element and return it.
// You must implement a solution with a linear runtime complexity and use only constant extra space.
//
// Constraints:
//   1 <= nums.length <= 3 * 10^4
//   -2^31 <= nums[i] <= 2^31 - 1
//   Each element in nums appears exactly three times except for one element which appears once.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Single Number II.
//   Memory Usage: 3.4 MB, less than 79.63% of Go online submissions for Single Number II.
func singleNumber(nums []int) int {
	var result int32 = 0
	for i := 0; i < 32; i++ {
		// Calculate the number of '1' for each bit
		sum_i := 0
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				sum_i++
			}
		}
		// If the number of '1' is 3*n => this bit is '0' in the result number
		// If the number of '1' is 3*n+1 => this bit is '1' in the result number
		if sum_i%3 == 1 {
			// Put the '1' bit to the proper position
			result |= 1 << i
		}
	}
	return int(result)
}

func singleNumber2(nums []int) int {
	var bit1, bit2 int
	for _, n := range nums {
		bit1 = (^bit2) & (bit1 ^ n)
		bit2 = (^bit1) & (bit2 ^ n)
	}
	return bit1
}
