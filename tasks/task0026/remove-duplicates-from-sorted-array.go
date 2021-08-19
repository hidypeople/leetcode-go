package task0026

// Given an integer array nums sorted in non-decreasing order, remove the duplicates
// in-place such that each unique element appears only once. The relative order of
// the elements should be kept the same.
// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result. It does
// not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying
// the input array in-place with O(1) extra memory.
//
// Constraints:
// 0 <= nums.length <= 3 * 10^4
// -100 <= nums[i] <= 100
// nums is sorted in non-decreasing order.
//
// Results:
//   Runtime: 4 ms, faster than 99.44% of Go online submissions for Remove Duplicates from Sorted Array.
//   Memory Usage: 4.4 MB, less than 34.64% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniqueIdx := 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if nums[uniqueIdx] != num {
			uniqueIdx++
			nums[uniqueIdx] = num
		}
	}
	return uniqueIdx + 1
}
