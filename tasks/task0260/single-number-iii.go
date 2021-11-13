package task0260

// Given an integer array nums, in which exactly two elements appear only once and all the
// other elements appear exactly twice. Find the two elements that appear only once.
// You can return the answer in any order.
// You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.
//
// Constraints:
//   2 <= nums.length <= 3 * 10^4
//   -2^31 <= nums[i] <= 2^31 - 1
//   Each integer in nums will appear twice, only two integers will appear once.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Single Number III.
//   Memory Usage: 4 MB, less than 100.00% of Go online submissions for Single Number III.
func singleNumber(nums []int) []int {
	sum := 0
	for _, x := range nums {
		sum ^= x
	}
	// The sum contains the result of expression (a ^ b) where a and b are the result numbers
	// So the '1' means that this bit is different for a and b, if we'll make a number that
	// contains only this bit equals to '1' (others are '0') => it will give different result
	// for expressions "a & 0x00...1...0" and "b & 0x00...1...0", so we'll be able to split
	// the nums into 2 groups, each group will contain only one number
	// e.g.
	// 0x10110110 <- sum = a ^ b
	// 0x00000010 <- sum2 = sum & -sum (extract the "last bit number")
	// The only one expression "0x00000010 & a", "0x00000010 & b" will give "0"
	sum2 := sum & -sum
	a := 0
	for _, x := range nums {
		if x&sum2 != 0 {
			a ^= x
		}
	}
	return []int{a, a ^ sum}
}
