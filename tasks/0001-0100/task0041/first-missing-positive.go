package task0041

// Given an unsorted integer array nums, return the smallest missing positive integer.
// You must implement an algorithm that runs in O(n) time and uses constant extra space.
//
// Constraints:
//   1 <= nums.length <= 5 * 10^5
//   -2^31 <= nums[i] <= 2^31 - 1
//
// Results:
//   Runtime: 116 ms, faster than 84.70% of Go online submissions for First Missing Positive.
//   Memory Usage: 25.2 MB, less than 50.18% of Go online submissions for First Missing Positive.
func firstMissingPositive(nums []int) int {
	// Move all positive numbers to the beginning of array
	posCount := 0
	hasOne := false
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > 0 {
			if num == 1 {
				hasOne = true
			}
			nums[posCount] = num
			posCount++
		}
	}
	if !hasOne {
		return 1
	}

	done := 0
	for i := 0; i < posCount; i++ {
		idx := abs(nums[i]) - 1
		if idx < posCount && nums[idx] > 0 {
			nums[idx] = -nums[idx]
			done++
		}
	}
	if done == posCount {
		return posCount + 1
	}
	for i := 0; i < posCount; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return 0
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
