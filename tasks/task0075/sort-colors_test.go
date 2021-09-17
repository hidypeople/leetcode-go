package task0075

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortColors(t *testing.T) {
	vals := []int{2, 0, 2, 1, 1, 0}
	sortColors(vals)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, vals)
}

func Test_sortColors2(t *testing.T) {
	vals := []int{2, 0, 1}
	sortColors(vals)
	assert.Equal(t, []int{0, 1, 2}, vals)
}

func Test_sortColors3(t *testing.T) {
	vals := []int{0}
	sortColors(vals)
	assert.Equal(t, []int{0}, vals)
}
