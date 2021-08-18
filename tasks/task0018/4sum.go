package task0018

import "sort"

// Given an array nums of n integers, return an array of all the
// unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
// 0 <= a, b, c, d < n
// a, b, c, and d are distinct.
// nums[a] + nums[b] + nums[c] + nums[d] == target
// You may return the answer in any order.
//
// Constraints:
//   1 <= nums.length <= 200
//   -10^9 <= nums[i] <= 10^9
//   -10^9 <= target <= 10^9
//
// Results:
//   Runtime: 4 ms, faster than 96.75% of Go online submissions for 4Sum.
//   Memory Usage: 2.8 MB, less than 68.18% of Go online submissions for 4Sum.
func fourSum(nums []int, target int) [][]int {
	results := [][]int{}
	n := len(nums)
	if n < 4 {
		return results
	}
	sort.Ints(nums)

	// Go through most left number
	left := 0
	for left < n-3 {
		// Go through most right number
		right := n - 1
		for right > left+2 {
			// Last two numbers are coming to each other (and try to reach out target value)
			left2, right2 := left+1, right-1
			for left2 < right2 {
				sum := nums[left] + nums[left2] + nums[right2] + nums[right]
				if sum < target {
					left2++
				} else if sum > target {
					right2--
				} else {
					results = append(results, []int{nums[left], nums[left2], nums[right2], nums[right]})
					left2++
					for left2 < n-2 && nums[left2] == nums[left2-1] {
						left2++
					}
					right2--
					for right2 > 1 && nums[right2] == nums[right2+1] {
						right2--
					}
				}
			}
			right--
			for right > 2 && nums[right] == nums[right+1] {
				right--
			}
		}
		left++
		for left < n-3 && nums[left] == nums[left-1] {
			left++
		}
	}
	return results
}
