package task0055

// You are given an integer array nums. You are initially positioned at the array's
// first index, and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.
//
// Constraints:
//   1 <= nums.length <= 10^4
//   0 <= nums[i] <= 10^5
//
// Results:
//   Runtime: 60 ms, faster than 69.36% of Go online submissions for Jump Game.
//   Memory Usage: 7 MB, less than 85.71% of Go online submissions for Jump Game.
func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}

	maxFar := 0
	nextStep := nums[0]
	for i := 1; i < n; i++ {
		if i <= nextStep {
			// Calc the maximum next jump distance after current step
			if i+nums[i] > maxFar {
				maxFar = i + nums[i]
				if maxFar >= n-1 {
					return true
				}
			}
			if i == nextStep {
				// if nextStep doesn't moves forward - no way
				if nextStep == maxFar {
					return false
				}
				nextStep = maxFar
				maxFar = 0
			}
		}
	}
	return false
}
