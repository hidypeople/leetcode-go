package task0996

import (
	"math"
	"sort"
)

// An array is squareful if the sum of every pair of adjacent elements is a perfect square.
// Given an integer array nums, return the number of permutations of nums that are squareful.
// Two permutations perm1 and perm2 are different if there is some index i such that perm1[i] != perm2[i].
//
// Constraints:
//   1 <= nums.length <= 12
//   0 <= nums[i] <= 10^9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Squareful Arrays.
//   Memory Usage: 2.1 MB, less than 71.43% of Go online submissions for Number of Squareful Arrays.
func numSquarefulPerms(nums []int) int {
	sort.Ints(nums)
	counter := 0
	var result *int = &counter
	permute(nums, 0, 0, result)
	return *result
}

func permute(nums []int, index int, prev int, result *int) {
	// if we recursively came to the end: this is good permutation
	if index >= len(nums) {
		*result = *result + 1
		return
	}
	done := map[int]bool{}
	for i := index; i < len(nums); i++ {
		// do not repeat permutations with the same valued element as we've done before
		if done[nums[i]] {
			continue
		}
		done[nums[i]] = true
		if index == 0 || (index > 0 && isPerfectSquare(nums[i]+prev)) {
			nums[i], nums[index] = nums[index], nums[i]
			permute(nums, index+1, nums[index], result)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
}

func isPerfectSquare(a int) bool {
	square := int(math.Sqrt(float64(a)))
	return a == square*square
}
