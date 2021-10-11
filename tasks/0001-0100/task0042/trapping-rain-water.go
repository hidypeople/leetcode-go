package task0042

// Given n non-negative integers representing an elevation map where the width of
// each bar is 1, compute how much water it can trap after raining.
//
// Constraints:
//   n == height.length
//   1 <= n <= 2 * 10^4
//   0 <= height[i] <= 10^5
//
// Results:
//   Runtime: 4 ms, faster than 84.15% of Go online submissions for Trapping Rain Water.
//   Memory Usage: 3.6 MB, less than 14.52% of Go online submissions for Trapping Rain Water.
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	sum := 0
	maxL, maxR := 0, 0
	l, r := 0, n-1
	for l <= r {
		hl, hr := height[l], height[r]
		if hl <= hr {
			if hl >= maxL {
				maxL = hl
			} else {
				sum += maxL - hl
			}
			l++
		} else {
			if hr >= maxR {
				maxR = hr
			} else {
				sum += maxR - hr
			}
			r--
		}
	}
	return sum
}
