package task0042

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trap(t *testing.T) {
	assert.Equal(t, 0, trap([]int{0, 1, 1}))
	assert.Equal(t, 0, trap([]int{1, 1}))
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
