package task0026

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The relative order of the elements may be changed.
// Since it is impossible to change the length of the array in some languages, you must
// instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input
// array in-place with O(1) extra memory.
//
// Constraints:
//   0 <= nums.length <= 100
//   0 <= nums[i] <= 50
//   0 <= val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
//   Memory Usage: 2.2 MB, less than 99.76% of Go online submissions for Remove Element.
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	resultIndex := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != val {
			nums[resultIndex] = num
			resultIndex++
		}
	}
	return resultIndex
}
