package task0088

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored
// inside the array nums1. To accommodate this, nums1 has a length of m + n, where the
// first m elements denote the elements that should be merged, and the last n elements
// are set to 0 and should be ignored. nums2 has a length of n.
//
// Constraints:
//   nums1.length == m + n
//   nums2.length == n
//   0 <= m, n <= 200
//   1 <= m + n <= 200
//   -10^9 <= nums1[i], nums2[j] <= 10^9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
//   Memory Usage: 2.4 MB, less than 76.77% of Go online submissions for Merge Sorted Array.
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, i1, i2 := m+n-1, m-1, n-1
	for i >= 0 {
		if i1 < 0 {
			nums1[i] = nums2[i2]
			i2--
		} else if i2 < 0 {
			nums1[i] = nums1[i1]
			i1--
		} else if nums1[i1] > nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
		i--
	}
}
