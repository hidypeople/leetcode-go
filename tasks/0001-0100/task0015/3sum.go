package task0015

import "sort"

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
//
// Constraints:
//   0 <= nums.length <= 3000
//   -10^5 <= nums[i] <= 10^5
//
// Results:
//   Runtime: 28 ms, faster than 92.04% of Go online submissions for 3Sum.
//   Memory Usage: 7.7 MB, less than 52.96% of Go online submissions for 3Sum.
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		x := nums[i]
		if x > 0 {
			break
		}
		resultLen := len(result)
		if resultLen > 0 && result[resultLen-1][0] == x {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := x + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res := []int{x, nums[l], nums[r]}
				// println(x, nums[l], nums[r])
				result = append(result, res)
				for l < r && nums[l] == res[1] {
					l++
				}
				for l < r && nums[r] == res[2] {
					r--
				}
			}
		}
	}
	return result
}
