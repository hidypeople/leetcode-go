package task0978

// Given an integer array arr, return the length of a maximum size turbulent subarray of arr.
// A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.
// More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent if and only if:
// For i <= k < j:
//   arr[k] > arr[k + 1] when k is odd, and
//   arr[k] < arr[k + 1] when k is even.
// Or, for i <= k < j:
//   arr[k] > arr[k + 1] when k is even, and
//   arr[k] < arr[k + 1] when k is odd.
//
// Constraints:
//   1 <= arr.length <= 4 * 10^4
//   0 <= arr[i] <= 10^9
//
// Result:
//   Runtime: 60 ms, faster than 100.00% of Go online submissions for Longest Turbulent Subarray.
//   Memory Usage: 7.2 MB, less than 100.00% of Go online submissions for Longest Turbulent Subarray.
func maxTurbulenceSize(arr []int) int {
	// complexity is O(len(arr))
	result := 1
	left, right := 0, 0
	prev := arr[0]
	prevDirection := 0 // Direction sign: ( 1 => prev < curr | 0 => prev == curr | -1: prev > curr )
	for i := 1; i < len(arr); i++ {
		curr := arr[i]

		if curr == prev {
			// equality => break current turbulence sequence
			result = max(right-left+1, result)
			prevDirection = 0
			left, right = i, i
			prev = curr
			continue
		}

		// Area: curr != prev

		if prevDirection == 0 {
			// Previous comparison was equality => we just need to set up direction sign and continue:
			if prev < curr {
				prevDirection = 1
			} else {
				prevDirection = -1
			}
			right = i
		} else if prevDirection == 1 {
			// previous step was ASC => we expect that current step will be DESC
			if prev < curr {
				// current step is ASC too => break current turbulence sequence
				result = max(right-left+1, result)
				left, right = i-1, i
				prevDirection = 1
			} else {
				// current step is DESC => continue sequence
				right = i
				prevDirection = -1
			}
		} else {
			// previous step was DESC => we expect that current step will be ASC
			if prev > curr {
				// current step is DESC too => break current turbulence sequence
				result = max(right-left+1, result)
				left, right = i-1, i
				prevDirection = -1
			} else {
				// current step is ASC => continue sequence
				right = i
				prevDirection = 1
			}
		}
		prev = curr
	}
	// Apply current incompleted result and return
	return max(right-left+1, result)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
