package task0179

import (
	"sort"
	"strconv"
	"strings"
)

// Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
// Since the result may be very large, so you need to return a string instead of an integer.
//
// Constraints:
//   1 <= nums.length <= 100
//   0 <= nums[i] <= 10^9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Largest Number.
//   Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Largest Number.
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	allZeroes := true
	for i, num := range nums {
		if num > 0 {
			allZeroes = false
		}
		strs[i] = strconv.Itoa(num)
	}
	if allZeroes {
		return "0"
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[j]+strs[i] < strs[i]+strs[j]
	})
	builder := new(strings.Builder)
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
