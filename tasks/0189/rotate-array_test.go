package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(array, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, array)
}

func TestRotate2(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Rotate(array, 4)
	assert.Equal(t, []int{7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6}, array)
}

func TestRotate3(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	Rotate(array, 3)
	assert.Equal(t, []int{6, 7, 8, 0, 1, 2, 3, 4, 5}, array)
}

func TestRotate4(t *testing.T) {
	array := []int{}
	Rotate(array, 3)
	assert.Equal(t, []int{}, array)
}

func TestRotate5(t *testing.T) {
	array := []int{1}
	Rotate(array, 3)
	assert.Equal(t, []int{1}, array)
}

func TestRotate6(t *testing.T) {
	array := []int{1,2,3}
	Rotate(array, 3)
	assert.Equal(t, []int{1,2,3}, array)
}
