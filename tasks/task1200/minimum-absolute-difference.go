package task1200

import "sort"

// Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.
// Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
//   a, b are from arr
//   a < b
//   b - a equals to the minimum absolute difference of any two elements in arr
//
// Constraints:
//   2 <= arr.length <= 10^5
//   -10^6 <= arr[i] <= 10^6
//
// Results:
//   Runtime: 68 ms, faster than 97.96% of Go online submissions for Minimum Absolute Difference.
//   Memory Usage: 8.8 MB, less than 22.45% of Go online submissions for Minimum Absolute Difference.
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDistance := arr[1] - arr[0]
	result := [][]int{{arr[0], arr[1]}}
	for i := 1; i < len(arr)-1; i++ {
		dist := arr[i+1] - arr[i]
		if dist == minDistance {
			result = append(result, []int{arr[i], arr[i+1]})
		} else if dist < minDistance {
			result = [][]int{{arr[i], arr[i+1]}}
			minDistance = dist
		}
	}
	return result
}
