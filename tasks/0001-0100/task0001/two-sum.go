package task0001

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Constraints:
//   2 <= nums.length <= 104
//   -109 <= nums[i] <= 109
//   -109 <= target <= 109
//   Only one valid answer exists.
//
// Results:
//   Runtime: 4 ms, faster than 97.43% of Go online submissions for Two Sum.
//   Memory Usage: 4.2 MB, less than 61.84% of Go online submissions for Two Sum.
func twoSum(nums []int, target int) []int {
	mapPrevious := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := mapPrevious[num]; ok {
			return []int{idx, i}
		} else {
			mapPrevious[target-num] = i
		}
	}
	return []int{}
}
