package task0149

// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane,
// return the maximum number of points that lie on the same straight line.
//
// Constraints:
//   1 <= points.length <= 300
//   points[i].length == 2
//   -10^4 <= xi, yi <= 10^4
//   All the points are unique.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Max Points on a Line.
//   Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Max Points on a Line.
func maxPoints(p [][]int) int {
	n := len(p)
	if n < 3 {
		return n
	}
	max, count, x0, y0, x1, y1, x2, y2 := 2, 0, 0, 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		// first point
		x0, y0 = p[i][0], p[i][1]
		for j := i + 1; j < n; j++ {
			// second point
			x1, y1 = p[j][0], p[j][1]
			count = 2
			for k := j + 1; k < n; k++ {
				// try to find another points on the first-second point line
				x2, y2 = p[k][0], p[k][1]
				// (x1-x0)/(y1-y0) = (x2-x0)/(y2-y0) => (x1-x0)*(y2-y0) = (x2-x0)*(y1-y0)
				if (x1-x0)*(y2-y0) == (x2-x0)*(y1-y0) {
					count++
				}
			}
			if max < count {
				max = count
			}
		}
	}
	return max
}
