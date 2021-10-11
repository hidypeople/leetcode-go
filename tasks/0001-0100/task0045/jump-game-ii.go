package task0045

// Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.
// You can assume that you can always reach the last index.
//
// Constraints:
//   1 <= nums.length <= 104
//   0 <= nums[i] <= 1000
//
// Results:
//   Runtime: 8 ms, faster than 99.42% of Go online submissions for Jump Game II.
//   Memory Usage: 6.1 MB, less than 49.85% of Go online submissions for Jump Game II.
func jump(nums []int) int {
	numsLength := len(nums)
	if numsLength < 2 {
		return 0
	}
	if nums[0] >= numsLength-1 {
		return 1
	}
	numberOfSteps := 1
	farIndex := 0
	// The current maximum step size:
	maxStepIndex := nums[0]
	for i := 1; i < numsLength; i++ {
		if i <= maxStepIndex {
			if i+nums[i] > farIndex {
				farIndex = i + nums[i]
				if farIndex >= numsLength-1 {
					return numberOfSteps + 1
				}
			}
		}
		// If we reach the maximum step size: update next step to the most far index
		if i == maxStepIndex {
			maxStepIndex = farIndex
			farIndex = 0
			numberOfSteps++
		}
	}
	return numberOfSteps
}
