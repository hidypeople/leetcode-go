package task0004

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
//
// Constraints:
//   nums1.length == m
//   nums2.length == n
//   0 <= m <= 1000
//   0 <= n <= 1000
//   1 <= m + n <= 2000
//   -106 <= nums1[i], nums2[i] <= 106
//
// Results:
//   Runtime: 8 ms, faster than 97.82% of Go online submissions for Median of Two Sorted Arrays.
//   Memory Usage: 5.3 MB, less than 78.69% of Go online submissions for Median of Two Sorted Arrays.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	i1, i2 := 0, 0
	middle := (n1 + n2) / 2 // [3,5 -> 4], [3,6] -> 4
	for i1+i2 < middle {
		has_i1 := i1 < n1
		has_i2 := i2 < n2
		if has_i1 && has_i2 {
			x1 := nums1[i1]
			x2 := nums2[i2]
			if x1 <= x2 {
				i1++
			} else {
				i2++
			}
		} else if has_i1 {
			i1++
		} else if has_i2 {
			i2++
		}
	}

	isOdd := (n1+n2)%2 == 1
	median := float64(0)
	medianPrev := float64(0)

	getPreviousValue := func(isNums1 bool, mainIdx, secIdx int) float64 {
		var _mainNums, _secNums = nums1, nums2
		if !isNums1 {
			_mainNums, _secNums = nums2, nums1
		}

		_secIdx := secIdx - 1
		if len(_secNums) == 0 {
			_secIdx = -1
		}

		_mainIdxPrev := mainIdx - 1
		if _mainIdxPrev >= 0 && _secIdx >= 0 {
			return float64(mathMax(_mainNums[_mainIdxPrev], _secNums[_secIdx]))
		} else if _mainIdxPrev >= 0 {
			return float64(_mainNums[_mainIdxPrev])
		} else if _secIdx >= 0 {
			return float64(_secNums[_secIdx])
		} else {
			return float64(_mainNums[mainIdx])
		}
	}

	if i1 < n1 && i2 < n2 {
		x1, x2 := nums1[i1], nums2[i2]
		if x1 < x2 {
			median = float64(x1)
			medianPrev = getPreviousValue(true, i1, i2)
		} else {
			median = float64(x2)
			medianPrev = getPreviousValue(false, i2, i1)
		}
	} else if i1 >= n1 {
		median = float64(nums2[i2])
		medianPrev = getPreviousValue(false, i2, i1)
	} else if i2 >= n2 {
		median = float64(nums1[i1])
		medianPrev = getPreviousValue(true, i1, i2)
	}

	if isOdd {
		return median
	} else {
		return (median + medianPrev) / 2
	}
}

func mathMax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
