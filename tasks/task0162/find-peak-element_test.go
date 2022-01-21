package task0162

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
	assert.Equal(t, 2, findPeakElement([]int{1, 2, 3, 1}))
}

func Test_findPeakElement2(t *testing.T) {
	assert.Equal(t, 5, findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

func Test_findPeakElement3(t *testing.T) {
	assert.Equal(t, 2, findPeakElement([]int{1, 2, 5, 4, 3, 2, 1}))
}

func Test_findPeakElement4(t *testing.T) {
	assert.Equal(t, 4, findPeakElement([]int{1, 2, 3, 4, 5, 2, 1}))
}
