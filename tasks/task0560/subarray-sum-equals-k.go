package task0560

// Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.
//
// Constraints:
//   1 <= nums.length <= 2 * 10^4
//   -1000 <= nums[i] <= 1000
//   -10^7 <= k <= 10^7
//
// Results:
//   Runtime: 36 ms, faster than 99.08% of Go online submissions for Subarray Sum Equals K.
//   Memory Usage: 7.7 MB, less than 68.13% of Go online submissions for Subarray Sum Equals K.
func subarraySum(nums []int, k int) int {
	counts := map[int]int{k: 1}
	sum, result := 0, 0
	for _, num := range nums {
		sum += num
		if count, ok := counts[sum]; ok {
			result += count
		}
		counts[k+sum]++
	}
	return result
}
