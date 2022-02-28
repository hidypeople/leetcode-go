package task0228

import (
	"sort"
	"strconv"
)

// You are given a sorted unique integer array nums.
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges, and there is no
// integer x such that x is in one of the ranges but not in nums.
// Each range [a,b] in the list should be output as:
// "a->b" if a != b
// "a" if a == b
//
// Constraints:
//   0 <= nums.length <= 20
//   -2^31 <= nums[i] <= 2^31 - 1
//   All the values of nums are unique.
//   nums is sorted in ascending order.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Summary Ranges.
//   Memory Usage: 1.9 MB, less than 73.96% of Go online submissions for Summary Ranges.
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	sort.Ints(nums)
	result := []string{}
	prev, left := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == prev+1 {
			prev = num
			// continue range
			continue
		}
		if left == prev {
			result = append(result, strconv.Itoa(left))
		} else {
			result = append(result, strconv.Itoa(left)+"->"+strconv.Itoa(prev))
		}
		left = num
		prev = num
	}
	if left == prev {
		result = append(result, strconv.Itoa(left))
	} else {
		result = append(result, strconv.Itoa(left)+"->"+strconv.Itoa(prev))
	}
	return result
}
