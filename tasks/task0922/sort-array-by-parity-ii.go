package task0922

// Given an array of integers nums, half of the integers in nums are odd, and the other half are even.
// Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.
// Return any answer array that satisfies this condition.
//
// Constraints:
//   2 <= nums.length <= 2 * 10^4
//   nums.length is even.
//   Half of the integers in nums are even.
//   0 <= nums[i] <= 1000
//
// Results:
//   Runtime: 16 ms, faster than 100.00% of Go online submissions for Sort Array By Parity II.
//   Memory Usage: 6.4 MB, less than 100.00% of Go online submissions for Sort Array By Parity II.
func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	// i - index of even positions, j - odd
	i, j := 0, 1
	for i < n && j < n {
		if nums[i]%2 == 0 {
			// if even number on the even position: ok
			i += 2
		} else if nums[j]%2 == 1 {
			// if odd number on the odd position: ok
			j += 2
		} else {
			// Otherwise: swap them
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
