package task0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	array := []int{1, 1, 2}
	result := removeDuplicates(array)
	assert.Equal(t, 2, result)
	assert.Equal(t, []int{1, 2}, array[:2])
}

func Test_removeDuplicates2(t *testing.T) {
	array := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(array)
	assert.Equal(t, 5, result)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, array[:5])
}

func Test_removeDuplicates3(t *testing.T) {
	array := []int{}
	result := removeDuplicates(array)
	assert.Equal(t, 0, result)
	assert.Equal(t, []int{}, array)
}
