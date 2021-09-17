package task0075

// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects
// of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.
//
// Constraints:
//   n == nums.length
//   1 <= n <= 300
//   nums[i] is 0, 1, or 2.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
//   Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Sort Colors.
func sortColors(nums []int) {
	// Complexity: O(2N)
	colorsCount := [3]int{0, 0, 0}
	for _, num := range nums {
		colorsCount[num]++
	}
	i := 0
	for color, count := range colorsCount {
		for j := 0; j < count; j++ {
			nums[i] = color
			i++
		}
	}
}
