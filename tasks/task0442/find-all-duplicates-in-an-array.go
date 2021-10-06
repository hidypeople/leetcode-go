package task0442

// Given an integer array nums of length n where all the integers of nums are in the range [1, n]
// and each integer appears once or twice, return an array of all the integers that appears twice.
// You must write an algorithm that runs in O(n) time and uses only constant extra space.
//
// Constraints:
//   n == nums.length
//   1 <= n <= 10^5
//   1 <= nums[i] <= n
//   Each element in nums appears once or twice.
//
// Results:
//   Runtime: 48 ms, faster than 86.81% of Go online submissions for Find All Duplicates in an Array. *Unstable*
//   Memory Usage: 7.8 MB, less than 50.55% of Go online submissions for Find All Duplicates in an Array.
func findDuplicates(nums []int) []int {
	// CPU: O(n), Memory: constant size
	// The idea is simple: we'll keep the 1 bit on the higher half of the interger
	// to indicate the duplicate, so "nums[i] > (1 << 16)" means that we have already met "i"
	duplicates := make([]int, 0)
	shift := 1 << 16
	var val int
	for _, v := range nums {
		val = v%shift - 1
		if nums[val]>>16 > 0 {
			duplicates = append(duplicates, val+1)
			continue
		}
		nums[val] += shift
	}
	return duplicates
}
