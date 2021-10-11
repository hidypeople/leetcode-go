package task0080

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such
// that each unique element appears at most twice. The relative order of the elements should be kept the same.
// Since it is impossible to change the length of the array in some languages, you must instead have the
// result be placed in the first part of the array nums. More formally, if there are k elements after removing
// the duplicates, then the first k elements of nums should hold the final result. It does not matter what
// you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
//
// Constraints:
//   1 <= nums.length <= 3 * 10^4
//   -10^4 <= nums[i] <= 10^4
//   nums is sorted in non-decreasing order.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted Array II.
//   Memory Usage: 3.2 MB, less than 25.62% of Go online submissions for Remove Duplicates from Sorted Array II.
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	var countOfDuplicates int
	var prevVal int
	uniqIdx := 0
	for i, leftVal := range nums {
		if leftVal != prevVal || i == 0 {
			// If current element is new element: reset duplicate counter and update prevVal
			countOfDuplicates = 1
			prevVal = leftVal
			nums[uniqIdx] = leftVal
			uniqIdx++
		} else {
			if countOfDuplicates < 2 {
				// If we've met less than 3 duplicate elements => keep moving
				nums[uniqIdx] = leftVal
				uniqIdx++
				countOfDuplicates++
			}
		}
	}
	return uniqIdx
}
