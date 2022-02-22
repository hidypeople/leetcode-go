package task0164

// Given an integer array nums, return the maximum difference between two successive elements
// in its sorted form. If the array contains less than two elements, return 0.
// You must write an algorithm that runs in linear time and uses linear extra space.
//
// Constraints:
//   1 <= nums.length <= 10^5
//   0 <= nums[i] <= 10^9
//
// Results:
//   Runtime: 124 ms, faster than 92.86% of Go online submissions for Maximum Gap.
//   Memory Usage: 10.4 MB, less than 61.90% of Go online submissions for Maximum Gap.
func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	const MAX, MIN = 1_000_000_001, -1_000_000_001

	// Find min and max
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}
	if maxVal == minVal {
		return 0
	}

	// Prepare n-1 buckets
	bucketsMin := make([]int, n-1)
	bucketsMax := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		bucketsMin[i] = MAX
		bucketsMax[i] = MIN
	}

	// int gap = (int)Math.ceil((double)(max - min)/(num.length - 1));
	gap := int(float64(maxVal-minVal)/float64(n-1)) + 1

	var bucketIdx int
	for _, num := range nums {
		if num == minVal || num == maxVal {
			continue
		}
		// Put number into min and max buckets
		bucketIdx = (num - minVal) / gap
		bucketsMin[bucketIdx] = min(num, bucketsMin[bucketIdx])
		bucketsMax[bucketIdx] = max(num, bucketsMax[bucketIdx])
	}

	// Now non-empty buckets contains min-max range for it's numbers
	// We know that elements with maximum gap cannot be in he same bucket,
	// so we need to compare max value in previous bucket with min value
	// from next non-empty bucket.
	maxGap, previous := MIN, minVal
	for i := 0; i < n-1; i++ {
		if bucketsMin[i] == MAX {
			// empty bucket
			continue
		}
		// min value minus the previous value is the current gap
		maxGap = max(maxGap, bucketsMin[i]-previous)
		// update previous bucket value
		previous = bucketsMax[i]
	}
	maxGap = max(maxGap, maxVal-previous)
	return maxGap
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
