package task0189

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Could you do it in-place with O(1) extra space?
//
// Constraints:
//   1 <= nums.length <= 10^5
//   -2^31 <= nums[i] <= 2^31 - 1
//   0 <= k <= 10^5
//
// Results:
//   Runtime: 28 ms, faster than 96.30% of Go online submissions for Rotate Array.
//   Memory Usage: 8.1 MB, less than 62.59% of Go online submissions for Rotate Array.
func rotate(nums []int, k int) {
	n := len(nums)
	if k == 0 || k == n {
		return
	}

	left := n
	// starting point
	start := 0
	for start < k && left > 0 {
		i := start
		var iNext int = (i + k) % n
		prev := nums[i]
		for iNext > start {
			temp := nums[iNext]
			nums[iNext] = prev
			left--
			if left < 0 {
				return
			}
			prev = temp
			i = iNext
			iNext = (i + k) % n
		}
		nums[start] = prev
		left--
		start++
	}
}
