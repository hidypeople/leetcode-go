package task0167

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number. Let these two numbers
// be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer
// array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
//
// Constraints:
//   2 <= numbers.length <= 3 * 10^4
//   -1000 <= numbers[i] <= 1000
//   numbers is sorted in non-decreasing order.
//   -1000 <= target <= 1000
//   The tests are generated such that there is exactly one solution.
//
// Results:
//   Runtime: 4 ms, faster than 92.55% of Go online submissions for Two Sum II - Input Array Is Sorted.
//   Memory Usage: 3.1 MB, less than 57.59% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSum(numbers []int, target int) []int {
	done := map[int]int{}
	for i, number := range numbers {
		diff := target - number
		if j, ok := done[diff]; ok {
			return []int{j + 1, i + 1}
		}
		done[number] = i
	}
	return []int{}
}
