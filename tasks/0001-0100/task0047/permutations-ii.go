package task0047

// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
//
// Constraints:
//   1 <= nums.length <= 8
//   -10 <= nums[i] <= 10
//
// Results:
//   Runtime: 4 ms, faster than 89.41% of Go online submissions for Permutations II. *Unstable*
//   Memory Usage: 6.4 MB, less than 19.41% of Go online submissions for Permutations II.
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	result := [][]int{}
	for i := 0; i < n; i++ {
		curr := nums[i]
		if i > 0 {
			j := i - 1
			for j >= 0 {
				if nums[j] == curr {
					break
				}
				j--
			}
			if j != -1 {
				continue
			}
		}
		// get the array that doesn't contain curr element:
		others := make([]int, n-1)
		for j, k := 0, 0; j < n; j++ {
			if j == i {
				continue
			}
			others[k] = nums[j]
			k++
		}
		// do recursive permutation for others and merge into result
		for _, permutation := range permuteUnique(others) {
			result = append(result, append([]int{curr}, permutation...))
		}
	}
	return result
}
