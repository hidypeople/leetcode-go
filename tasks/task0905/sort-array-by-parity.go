package task0905

// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
// Return any array that satisfies this condition.
//
// Constraints:
//   1 <= nums.length <= 5000
//   0 <= nums[i] <= 5000
//
// Results:
//   Runtime: 3 ms, faster than 97.07% of Go online submissions for Sort Array By Parity.
//   Memory Usage: 5.1 MB, less than 14.23% of Go online submissions for Sort Array By Parity.
func sortArrayByParity(nums []int) []int {
	result := make([]int, len(nums))
	evenCnt, oddCnt := 0, len(nums)-1
	for _, num := range nums {
		if num%2 == 0 {
			result[evenCnt] = num
			evenCnt++
		} else {
			result[oddCnt] = num
			oddCnt--
		}
	}
	return result
}
