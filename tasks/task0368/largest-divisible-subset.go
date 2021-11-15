package task0368

import (
	"sort"
)

// Given a set of distinct positive integers nums, return the largest subset answer
// such that every pair (answer[i], answer[j]) of elements in this subset satisfies:
// answer[i] % answer[j] == 0, or
// answer[j] % answer[i] == 0
// If there are multiple solutions, return any of them.
//
// Constraints:
//   1 <= nums.length <= 1000
//   1 <= nums[i] <= 2 * 10^9
//   All the integers in nums are unique.
//
// Results:
//   Runtime: 36 ms, faster than 100.00% of Go online submissions for Largest Divisible Subset.
//   Memory Usage: 3 MB, less than 100.00% of Go online submissions for Largest Divisible Subset.
func largestDivisibleSubset(nums []int) []int {
	type DpItem struct {
		next int
		val  int
	}
	// Sort items to ensure that next value cannot be divisor of previous
	sort.Ints(nums)

	// dp[i] will contain 2 values:
	// - best previous value index
	// - best previous number of divisors
	dp := make([]DpItem, len(nums))
	dpStartIdx, dpStartVal := -1, 0
	for i, num := range nums {

		// Look behind to the previous values (for all of those values the best result is known
		// from the previous iterations)
		maxIdx, maxVal := -1, 0
		for j := i - 1; j >= 0; j-- {
			if num%nums[j] == 0 {
				// if previous value is divident of current num and it has more dividents
				// than current best result: remember it
				if maxVal < dp[j].val {
					maxIdx = j
					maxVal = dp[j].val
				}
			}
		}
		if maxVal+1 > dpStartVal {
			// the best of the best result
			dpStartVal = maxVal + 1
			dpStartIdx = i
		}
		// store previous index with the most dividents and number of dividents for the current item
		dp[i] = DpItem{maxIdx, maxVal + 1}
	}

	// e.g.
	// nums = [1,2,3,4,8]
	// dpStartIdx = 4
	// dp = [{idx:-1,1}, {idx:0, 2}, {idx:0, 2}, {idx:1, 3}, {idx:3, 4}]
	// That means the best result is: nums[4] -> nums[3] -> nums[1] -> nums[0]
	result := []int{}
	i := dpStartIdx
	for i >= 0 {
		result = append(result, nums[i])
		i = dp[i].next
	}
	return result
}
