package task01413

// Given an array of integers nums, you start with an initial positive value startValue.
// In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
// Return the minimum positive value of startValue such that the step by step sum is never less than 1.
//
// Constraints:
//   1 <= nums.length <= 100
//   -100 <= nums[i] <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Value to Get Positive Step by Step Sum.
//   Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Minimum Value to Get Positive Step by Step Sum.
func minStartValue(nums []int) int {
	x := 1
	sum := 1
	for _, n := range nums {
		sum = sum + n
		if sum < 1 {
			x += 1 - sum
			sum += 1 - sum
		}
	}
	return x
}
