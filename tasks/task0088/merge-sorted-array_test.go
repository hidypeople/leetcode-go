package task0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	nums1, nums2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}

func Test_merge2(t *testing.T) {
	nums1, nums2 := []int{1}, []int{}
	merge(nums1, 1, nums2, 0)
	assert.Equal(t, []int{1}, nums1)
}

func Test_merge3(t *testing.T) {
	nums1, nums2 := []int{0}, []int{1}
	merge(nums1, 0, nums2, 1)
	assert.Equal(t, []int{1}, nums1)
}

func Test_merge4(t *testing.T) {
	nums1, nums2 := []int{1, 2, 10, 0, 0, 0}, []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 2, 5, 6, 10}, nums1)
}

func Test_merge5(t *testing.T) {
	nums1, nums2 := []int{4, 5, 6, 0, 0, 0}, []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums1)
}

func Test_merge6(t *testing.T) {
	nums1, nums2 := []int{4, 0, 0, 0, 0, 0}, []int{1, 2, 3, 5, 6}
	merge(nums1, 1, nums2, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums1)
}

func Test_merge7(t *testing.T) {
	nums1, nums2 := []int{1, 2, 4, 5, 6, 0}, []int{3}
	merge(nums1, 5, nums2, 1)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums1)
}
