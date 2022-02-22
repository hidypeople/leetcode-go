package task0169

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//
// Constraints:
//   n == nums.length
//   1 <= n <= 5 * 10^4
//   -2^31 <= nums[i] <= 2^31 - 1
//
// Results:
//   Runtime: 12 ms, faster than 98.97% of Go online submissions for Majority Element.
//   Memory Usage: 6 MB, less than 95.89% of Go online submissions for Majority Element.
func majorityElement(nums []int) int {
	majorityCount, majorityNum := 0, 0
	for _, num := range nums {
		if majorityCount == 0 {
			majorityCount, majorityNum = 1, num
		} else if num == majorityNum {
			majorityCount++
		} else {
			majorityCount--
		}
	}
	return majorityNum
}
