package task0152

// Given an integer array nums, find a contiguous non-empty subarray within the array that has the
// largest product, and return the product.
// The test cases are generated so that the answer will fit in a 32-bit integer.
// A subarray is a coontiguous subsequence of the array.
//  The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// Constraints:
//   1 <= nums.length <= 2 * 10^4
//   -10 <= nums[i] <= 10
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Product Subarray.
//   Memory Usage: 3.4 MB, less than 19.79% of Go online submissions for Maximum Product Subarray.
func maxProduct(nums []int) int {
	// currMax/currMin stores the max/min product of subarray that ends with the current number nums[i]
	num, n, result, currMax, currMin := 0, len(nums), nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		num = nums[i]
		if num < 0 {
			// negative value => swap min and max
			currMax, currMin = currMin, currMax
		}

		// Update current max
		if num > currMax*num {
			currMax = num
		} else {
			currMax = num * currMax
		}

		// Update current min
		if num < currMin*num {
			currMin = num
		} else {
			currMin = num * currMin
		}

		// the newly computed max value is a candidate for our global result
		if result < currMax {
			result = currMax
		}
	}
	return result
}
