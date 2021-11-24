package task0986

// You are given two lists of closed intervals, firstList and secondList,
// where firstList[i] = [starti, endi] and secondList[j] = [startj, endj].
// Each list of intervals is pairwise disjoint and in sorted order.
// Return the intersection of these two interval lists.
// A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.
// The intersection of two closed intervals is a set of real numbers that
// are either empty or represented as a closed interval. For example, the
// intersection of [1, 3] and [2, 4] is [2, 3].
//
// Constraints:
//   0 <= firstList.length, secondList.length <= 1000
//   firstList.length + secondList.length >= 1
//   0 <= starti < endi <= 10^9
//   endi < starti+1
//   0 <= startj < endj <= 10^9
//   endj < startj+1
//
// Results:
//   Runtime: 16 ms, faster than 97.62% of Go online submissions for Interval List Intersections.
//   Memory Usage: 7 MB, less than 91.67% of Go online submissions for Interval List Intersections.
func intervalIntersection(arr1 [][]int, arr2 [][]int) [][]int {
	n1, n2 := len(arr1), len(arr2)
	if n1 == 0 || n2 == 0 {
		return [][]int{}
	}

	result := [][]int{}
	i1, i2 := 0, 0
	for i1 < n1 && i2 < n2 {
		l1, r2 := arr1[i1][0], arr2[i2][1]
		if r2 < l1 {
			// arr2: [....l2....r2..............]
			// arr1: [..............l1....r1....]
			i2++
			continue
		}
		r1, l2 := arr1[i1][1], arr2[i2][0]
		if r1 < l2 {
			// arr2: [..............l2....r2....]
			// arr1: [....l1....r1..............]
			i1++
			continue
		}

		// The arrays intersects
		l := l1
		if l2 > l1 {
			l = l2
		}
		r := r1
		if r2 < r1 {
			r = r2
		}
		result = append(result, []int{l, r})
		if r1 == r {
			i1++
		}
		if r2 == r {
			i2++
		}
	}
	return result
}
