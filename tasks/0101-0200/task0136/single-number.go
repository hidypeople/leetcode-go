package task0134

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.
//
// Constraints:
//   1 <= nums.length <= 3 * 10^4
//   -3 * 10^4 <= nums[i] <= 3 * 10^4
//   Each element in the array appears twice except for one element which appears only once.
//
// Results:
//   Runtime: 12 ms, faster than 99.05% of Go online submissions for Single Number.
//   Memory Usage: 6.2 MB, less than 58.44% of Go online submissions for Single Number.
func singleNumber(nums []int) int {
	result := 0
	for _, x := range nums {
		result ^= x
	}
	return result
}
