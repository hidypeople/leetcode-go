package task0057

// You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi]
// represent the start and the end of the ith interval and intervals is sorted in ascending order by
// starti. You are also given an interval newInterval = [start, end] that represents the start and
// end of another interval.
// Insert newInterval into intervals such that intervals is still sorted in ascending order by starti
// and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
// Return intervals after the insertion.
//
// Constraints:
//   0 <= intervals.length <= 10^4
//   intervals[i].length == 2
//   0 <= starti <= endi <= 10^5
//   intervals is sorted by starti in ascending order.
//   newInterval.length == 2
//   0 <= start <= end <= 10^5
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Insert Interval.
//   Memory Usage: 4.8 MB, less than 86.81% of Go online submissions for Insert Interval.
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	result := [][]int{}
	i := 0
	for i < len(intervals) {
		interval := intervals[i]
		if right < interval[0] {
			// if we're on the right hand of the new interval: exit
			break
		}

		if left > interval[1] {
			// if we're on the left hand of the new interval: just add current
			result = append(result, interval)
			i++
			continue
		}

		// expand interval
		if interval[0] < left {
			left = interval[0]
		}
		if interval[1] > right {
			right = interval[1]
		}
		i++
	}
	// append new interval
	result = append(result, []int{left, right})
	// append res
	result = append(result, intervals[i:]...)
	return result
}
