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
