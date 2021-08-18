package task0016

import "sort"

// Given an integer array nums of length n and an integer target, find three
// integers in nums such that the sum is closest to target.
// Return the sum of the three integers.
// You may assume that each input would have exactly one solution.
//
// Constraints:
//   3 <= nums.length <= 1000
//   -1000 <= nums[i] <= 1000
//   -10^4 <= target <= 10^4
//
// Results:
//   Runtime: 4 ms, faster than 96.37% of Go online submissions for 3Sum Closest.
//   Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for 3Sum Closest.
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	diff := 100000
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return target - diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
