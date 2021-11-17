package task0128

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
//
// Constraints:
//   0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
//
// Results:
//   Runtime: 36 ms, faster than 95.79% of Go online submissions for Longest Consecutive Sequence.
//   Memory Usage: 8.9 MB, less than 49.51% of Go online submissions for Longest Consecutive Sequence.
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Map contains the length of sequence for current number
	numsTable := map[int]int{}
	longest := 1
	for _, num := range nums {
		if _, ok := numsTable[num]; !ok {
			// left and right element sequences length
			left := numsTable[num-1]
			right := numsTable[num+1]

			// current sequence length is the sum of left, right sequences + self
			sum := left + right + 1
			if longest < sum {
				longest = sum
			}

			// Set new extended sequence length for left and right boundary:
			numsTable[num] = sum
			numsTable[num-left] = sum
			numsTable[num+right] = sum
		}
	}

	return longest
}
