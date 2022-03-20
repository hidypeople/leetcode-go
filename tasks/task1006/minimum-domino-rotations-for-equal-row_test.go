package task1006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDominoRotations(t *testing.T) {
	assert.Equal(t, 2, minDominoRotations(
		[]int{2, 1, 2, 4, 2, 2},
		[]int{5, 2, 6, 2, 3, 2},
	))
}

func Test_minDominoRotations2(t *testing.T) {
	assert.Equal(t, -1, minDominoRotations(
		[]int{3, 5, 1, 2, 3},
		[]int{3, 6, 3, 3, 4},
	))
}

func Test_minDominoRotations3(t *testing.T) {
	assert.Equal(t, 0, minDominoRotations(
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
	))
}

func Test_minDominoRotations4(t *testing.T) {
	assert.Equal(t, 2, minDominoRotations(
		[]int{6, 1, 6, 4, 6, 6},
		[]int{5, 6, 2, 6, 3, 6},
	))
}
