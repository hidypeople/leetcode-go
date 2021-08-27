package task0996

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSquarefulPerms(t *testing.T) {
	assert.Equal(t, 2, numSquarefulPerms([]int{1, 17, 8}))
}

func Test_numSquarefulPerms2(t *testing.T) {
	assert.Equal(t, 1, numSquarefulPerms([]int{2, 2, 2}))
}

func Test_numSquarefulPerms3(t *testing.T) {
	assert.Equal(t, 0, numSquarefulPerms([]int{99, 62, 10, 47, 53, 9, 83, 33, 15, 24}))
}

func Test_numSquarefulPerms4(t *testing.T) {
	assert.Equal(t, 0, numSquarefulPerms([]int{75, 91, 39, 33, 39, 39, 69, 20, 43, 38, 48, 29}))
}

func Test_numSquarefulPerms5(t *testing.T) {
	assert.Equal(t, 1, numSquarefulPerms([]int{1, 1, 8, 1, 8}))
}

func Test_numSquarefulPerms6(t *testing.T) {
	assert.Equal(t, 6, numSquarefulPerms([]int{1, 8, 8, 1, 8}))
}

func Test_numSquarefulPerms7(t *testing.T) {
	assert.Equal(t, 3, numSquarefulPerms([]int{80, 1, 80, 1, 3, 6, 3}))
}

func Test_numSquarefulPerms8(t *testing.T) {
	assert.Equal(t, 3, numSquarefulPerms([]int{1, 8, 8}))
}
