package task0532

import "sort"

// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
// 0 <= i < j < nums.length
// |nums[i] - nums[j]| == k
// Notice that |val| denotes the absolute value of val.
//
// Constraints:
//   1 <= nums.length <= 10^4
//   -10^7 <= nums[i] <= 10^7
//   0 <= k <= 10^7
//
// Results:
//   Runtime: 4 ms, faster than 97.78% of Go online submissions for K-diff Pairs in an Array.
//   Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for K-diff Pairs in an Array.
func findPairs(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)

	count, l, r := 0, 0, 1
	var diff int
	var numL int
	for l < n {
		numL = nums[l]
		r = l + 1
		for r < n {
			diff = nums[r] - numL
			if diff >= k {
				if diff == k {
					count++
				}
				break
			} else {
				r++
			}
		}
		for l+1 < n && numL == nums[l+1] {
			l++
		}
		l++
	}
	return count
}
