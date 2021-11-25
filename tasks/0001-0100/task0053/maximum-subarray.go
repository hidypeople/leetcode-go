package task0053

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.
//
// Constraints:
//   1 <= nums.length <= 3 * 10^4
//  -10^5 <= nums[i] <= 10^5
//
// Result:
//   Runtime: 92 ms, faster than 97.60% of Go online submissions for Maximum Subarray.
//   Memory Usage: 9.5 MB, less than 27.12% of Go online submissions for Maximum Subarray.
func maxSubArray(nums []int) int {
	maxSub := nums[0]
	currSum := 0

	for _, num := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum += num
		if currSum > maxSub {
			maxSub = currSum
		}
	}
	return maxSub
}
