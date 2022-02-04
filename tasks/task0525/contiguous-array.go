package task0525

// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.
//
// Constraints:
//   1 <= nums.length <= 10^5
//   nums[i] is either 0 or 1.
//
// Results:
//   Runtime: 80 ms, faster than 100.00% of Go online submissions for Contiguous Array.
//   Memory Usage: 7.4 MB, less than 65.22% of Go online submissions for Contiguous Array.
func findMaxLength(nums []int) int {
	// O(n)
	maxLength, sum := 0, 0
	// Let:
	//   0 do sum--
	//   1 do sum++
	// So array:
	// [1, 0, 1, 1, 1, 0, 0, 0,  0,  0]
	// will have sums
	// [1, 0, 1, 2, 3, 2, 1, 0, -1, -2]
	// And we need to find the maximum length between two equal numbers

	// Zero indexes contains the first index where we met sum number
	zeroIndexes := map[int]int{0: -1}
	for i, num := range nums {
		sum += 2*num - 1 // fn(x) = [0 => -1, 1 => 1]
		prevZeroIndex, ok := zeroIndexes[sum]
		if !ok {
			// If we didn't met this sum yet - remember the index
			zeroIndexes[sum] = i
		} else if i-prevZeroIndex > maxLength {
			// If we met this sum - compare with the first occurence and check for maximum
			maxLength = i - prevZeroIndex
		}
	}
	return maxLength
}
