package task0016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	assert.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 1))
}
