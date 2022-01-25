package task0941

// Given an array of integers arr, return true if and only if it is a valid mountain array.
// Recall that arr is a mountain array if and only if:
//   arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
//   arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//   arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//
// Constraints:
//   1 <= arr.length <= 10^4
//   0 <= arr[i] <= 10^4
//
// Results:
//   Runtime: 20 ms, faster than 98.44% of Go online submissions for Valid Mountain Array.
//   Memory Usage: 6.7 MB, less than 99.61% of Go online submissions for Valid Mountain Array.
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	isGoingUp, peakIndex := true, -1
	prev := arr[0]
	var curr int
	for i := 1; i < n; i++ {
		curr = arr[i]
		if prev == curr {
			return false
		} else if !isGoingUp && prev < curr {
			return false
		} else if prev > curr {
			isGoingUp = false
			if peakIndex == -1 {
				peakIndex = i - 1
			}
		}
		prev = curr
	}
	return peakIndex > 0 && peakIndex < n-1
}
