package task0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 3}, []int{2, 7}))
	assert.Equal(t, 1.5, findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{1, 2, 3}, []int{3, 4, 5, 6}))
	assert.Equal(t, 4.5, findMedianSortedArrays([]int{}, []int{3, 4, 5, 6}))
	assert.Equal(t, 4.5, findMedianSortedArrays([]int{3, 4, 5, 6}, []int{}))
	assert.Equal(t, 6.0, findMedianSortedArrays([]int{100, 200, 300}, []int{3, 4, 5, 6}))
	assert.Equal(t, 100.0, findMedianSortedArrays([]int{3, 4, 5, 6, 1000, 10001}, []int{100, 200, 300}))
	assert.Equal(t, 4.5, findMedianSortedArrays([]int{}, []int{3, 4, 5, 6}))
	assert.Equal(t, 0.0, findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	assert.Equal(t, 1.0, findMedianSortedArrays([]int{}, []int{1}))
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{}, []int{2}))
}
