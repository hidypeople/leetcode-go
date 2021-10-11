package task0055

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump0(t *testing.T) {
	assert.True(t, canJump([]int{1}))
}

func Test_canJump(t *testing.T) {
	assert.True(t, canJump([]int{2, 3, 1, 1, 4}))
}

func Test_canJump2(t *testing.T) {
	assert.False(t, canJump([]int{3, 2, 1, 0, 4}))
}

func Test_canJump3(t *testing.T) {
	assert.False(t, canJump([]int{5, 4, 3, 0, 0, 0, 4}))
}

func Test_canJump4(t *testing.T) {
	assert.True(t, canJump([]int{1, 10, 0, 2, 0, 0, 0, 4}))
}

func Test_canJump5(t *testing.T) {
	assert.False(t, canJump([]int{1, 1, 0, 12, 0, 0, 0, 4}))
}

func Test_canJump6(t *testing.T) {
	assert.True(t, canJump([]int{1, 1, 1, 1, 1, 1, 1}))
}

func Test_canJump7(t *testing.T) {
	assert.True(t, canJump([]int{1, 1, 2, 2, 0, 1, 1}))
}

// []
