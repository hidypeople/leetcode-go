package task0084

// Given an array of integers heights representing the histogram's bar height where the width of
// each bar is 1, return the area of the largest rectangle in the histogram.
//
// Constraints:
//   1 <= heights.length <= 10^5
//   0 <= heights[i] <= 10^4
//
// Results:
//   Runtime: 80 ms, faster than 97.92% of Go online submissions for Largest Rectangle in Histogram.
//   Memory Usage: 9.2 MB, less than 54.86% of Go online submissions for Largest Rectangle in Histogram.
func largestRectangleArea(heights []int) int {
	// Complexity: O(N)
	n := len(heights)
	leftBoundaryArray, rightBoundaryArray := make([]int, n), make([]int, n)
	leftBoundaryArray[0] = -1
	rightBoundaryArray[n-1] = n

	// leftBoundaryArray contains the indexes of heights that at the left of i-th position and less than heights[i]
	//
	// e.g. [1, 3, 4, 4, 5, 4, 1], heights[5] is 4, so the closest value that less than 4 is 3,
	// so leftBoundaryArray[5] will contain the index of "3" =>  leftBoundaryArray[5] = 1
	for i := 1; i < n; i++ {
		p := i - 1
		// we can reuse previous values: if we met index where value is larger than current,
		// we can do the jump the left boundary of this value and do not check the values
		// in interval [leftBoundaryArray[p]+1, p-1]
		//
		// e.g. [1, 3, 12, 11, 10, 4, 1], heights[5] == 4.
		//   1. The previous sibling value is height[4] == 10 and it is larger than 4
		//   2. We take leftBoundaryArray[4] == 1 and we can jump directly to the index 1 because
		//      we know that interval [2, 4] contains the numbers greater than 4
		for p >= 0 && heights[p] >= heights[i] {
			p = leftBoundaryArray[p]
		}
		leftBoundaryArray[i] = p
	}

	// rightBoundaryArray is similar to leftBoundaryArray
	for i := n - 2; i >= 0; i-- {
		p := i + 1
		for p < n && heights[p] >= heights[i] {
			p = rightBoundaryArray[p]
		}
		rightBoundaryArray[i] = p
	}

	maxHeight := 0
	for i := 0; i < n; i++ {
		maxHeight = max(maxHeight, heights[i]*(rightBoundaryArray[i]-leftBoundaryArray[i]-1))
	}

	return maxHeight
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
