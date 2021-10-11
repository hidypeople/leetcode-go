package task0030

import "sort"

// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
// If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
// The replacement must be in place and use only constant extra memory.
//
// Constraints:
//   1 <= nums.length <= 100
//   0 <= nums[i] <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Next Permutation.
//   Memory Usage: 2.5 MB, less than 12.32% of Go online submissions for Next Permutation.
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	n := len(nums)

	// Find the tail subsequence that decreases:
	// e.g. nums=[1,6,2,3,5,4,0] => tail sequence will be [5,4,0], tailIndex=4
	tailIndex := n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			tailIndex = i + 1
			break
		}
		if i == 0 {
			tailIndex = 0
			break
		}
	}

	// If whole nums has desc order -> sort in increasing order
	if tailIndex == 0 {
		sort.Ints(nums)
		return
	}

	// Find the value
	// e.g. nums=[1,6,2,3,5,4,0], tail=[5,4,0], pre_tail_value=3, so we need to find
	// the closest to the end value that is greater than pre_tail_value and swap them
	// [1,6,2, !3! ,5, !4! ,0] => [1,6,2, !4! ,5, !3! ,0]
	changingValue := nums[tailIndex-1]
	tilePointIndex := n - 1
	for i := n - 1; i >= tailIndex; i-- {
		if nums[i] > changingValue {
			tilePointIndex = i
			break
		}
	}
	// Swap
	nums[tailIndex-1], nums[tilePointIndex] = nums[tilePointIndex], nums[tailIndex-1]

	// Sort tile (it is sorted in reverse order): we need to reverse it
	left, right := tailIndex, n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
