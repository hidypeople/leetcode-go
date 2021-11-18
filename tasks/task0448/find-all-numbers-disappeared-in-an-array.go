package task0448

// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
//
// Constraints:
//   n == nums.length
//   1 <= n <= 10^5
//   1 <= nums[i] <= n
//
// Results:
//   Runtime: 44 ms, faster than 100.00% of Go online submissions for Find All Numbers Disappeared in an Array.
//   Memory Usage: 7.5 MB, less than 92.45% of Go online submissions for Find All Numbers Disappeared in an Array.
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	count := make([]int, n)
	for _, num := range nums {
		count[num-1]++
	}
	// I don't append result by one item on each iteration,
	// instead I put values one-by-one into large array and cut it off
	// at the end. It consumes more memory, but faster.
	result := make([]int, n)
	resultIdx := 0
	for i := 1; i <= n; i++ {
		if count[i-1] == 0 {
			result[resultIdx] = i
			resultIdx++
		}
	}
	return result[:resultIdx]
}
