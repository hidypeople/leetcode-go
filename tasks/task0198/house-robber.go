package task0198

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
// the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and
// it will automatically contact the police if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can
// rob tonight without alerting the police.
//
// Constraints:
//   1 <= nums.length <= 100
//   0 <= nums[i] <= 400
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
//   Memory Usage: 2.1 MB, less than 76.32% of Go online submissions for House Robber.
func rob(nums []int) int {
	results := make([]int, len(nums)+2)
	results[0] = 0
	results[1] = 0
	for i, num := range nums {
		// In the results[i+2] we keep the maximum amount of money for step i
		if num+results[i] > results[i+1] {
			// if current step and pre-previous maximum is better than previous maximum - take it
			results[i+2] = num + results[i]
		} else {
			// otherwise - take previous maximum
			results[i+2] = results[i+1]
		}
	}
	return results[len(nums)+1]
}
