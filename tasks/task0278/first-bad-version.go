package task0278

// You are a product manager and currently leading a team to develop a new product.
// Unfortunately, the latest version of your product fails the quality check.
// Since each version is developed based on the previous version, all the versions
// after a bad version are also bad.
// Suppose you have n versions [1, 2, ..., n] and you want to find out the first
// bad one, which causes all the following ones to be bad.
// You are given an API bool isBadVersion(version) which returns whether version
// is bad. Implement a function to find the first bad version. You should minimize
// the number of calls to the API.
//
// Constraints:
//   1 <= bad <= n <= 231 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.
//   Memory Usage: 1.9 MB, less than 30.50% of Go online submissions for First Bad Version.
func firstBadVersion(n int) int {
	// binary search
	left, right := 1, n
	for left <= right {
		middle := left + (right-left)/2
		if !isBadVersion(middle) {
			// middle is not bad => go to the right half
			left = middle + 1
		} else {
			// middle is bad => go to the left half including middle
			if right == middle {
				// if middle == right => the only one element left. Return it
				return right
			}
			right = middle
		}
	}
	return 0
}

// Mock "isBadVersion" for main solution
func isBadVersion(version int) bool {
	return version >= 1
}

// This function has the same as firstBadVersion, but gets isBadVersion as parameter.
// It is used only for running unit tests
func firstBadVersion_ForTest(n int, isBadVersion func(version int) bool) int {
	// binary search
	left, right := 1, n
	for left <= right {
		middle := left + (right-left)/2
		if !isBadVersion(middle) {
			// middle is not bad => go to the right half
			left = middle + 1
		} else {
			// middle is bad => go to the left half including middle
			if right == middle {
				// if middle == right => the only one element left. Return it
				return right
			}
			right = middle
		}
	}
	return 0
}
