package task0011

// Given n non-negative integers a1, a2, ..., an , where each represents a point
// at coordinate (i, ai). n vertical lines are drawn such that the two endpoints
// of the line i is at (i, ai) and (i, 0). Find two lines, which, together with
// the x-axis forms a container, such that the container contains the most water.
// Notice that you may not slant the container.
//
// Constraints:
//   n == height.length
//   2 <= n <= 105
//   0 <= height[i] <= 104
//
// Results:
//   Runtime: 80 ms, faster than 83.62% of Go online submissions for Container With Most Water.
//   Memory Usage: 8.8 MB, less than 14.86% of Go online submissions for Container With Most Water.
func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}

	res := 0
	l, r := 0, n-1
	level := 0
	for l < r {
		hl, hr := height[l], height[r]
		if hl <= hr {
			if hl > level {
				level = hl
				sum := (r - l) * level
				if sum > res {
					res = sum
				}
			}
			l++
		} else {
			if hr > level {
				level = hr
				sum := (r - l) * level
				if sum > res {
					res = sum
				}
			}
			r--
		}
	}
	return res
}
