package task0189

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Could you do it in-place with O(1) extra space?
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
