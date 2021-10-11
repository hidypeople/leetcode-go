package task0046

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
//
// Constraints:
//   1 <= nums.length <= 6
//   -10 <= nums[i] <= 10
//   All the integers of nums are unique.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
//   Memory Usage: 3.5 MB, less than 6.58% of Go online submissions for Permutations.
func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	result := [][]int{}
	for i := 0; i < n; i++ {
		// get the array that doesn't contain nums[i] element:
		others := make([]int, n-1)
		for j, k := 0, 0; j < n; j++ {
			if j == i {
				continue
			}
			others[k] = nums[j]
			k++
		}
		// do recursive permutation for others and merge into result
		for _, permutation := range permute(others) {
			result = append(result, append([]int{nums[i]}, permutation...))
		}
	}
	return result
}
