package task0207

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canFinish(t *testing.T) {
	assert.True(t, canFinish(2, [][]int{{1, 0}}))
}

func Test_canFinish2(t *testing.T) {
	assert.False(t, canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func Test_canFinish3(t *testing.T) {
	assert.False(t, canFinish(3, [][]int{{0, 2}, {1, 2}, {2, 0}}))
}

func Test_canFinish4(t *testing.T) {
	assert.False(t, canFinish(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}))
}

func Test_canFinish5(t *testing.T) {
	assert.True(t, canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}
