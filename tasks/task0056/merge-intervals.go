package task0056

import "sort"

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
// Constraints:
//   1 <= intervals.length <= 10^4
//   intervals[i].length == 2
//   0 <= starti <= endi <= 10^4
//
// Results:
//   Runtime: 8 ms, faster than 92.38% of Go online submissions for Merge Intervals. *Unstable*
//   Memory Usage: 4.6 MB, less than 96.63% of Go online submissions for Merge Intervals.
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	// sort intervals by left part
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	currentResult := intervals[0]
	result := [][]int{currentResult}
	i := 1
	for i < len(intervals) {
		interval := intervals[i]
		i++
		if currentResult[1] < interval[0] {
			result = append(result, interval)
			currentResult = interval
			continue
		}
		if interval[1] > currentResult[1] {
			currentResult[1] = interval[1]
		}
	}
	return result
}
