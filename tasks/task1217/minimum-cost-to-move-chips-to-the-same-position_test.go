package task1217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCostToMoveChips(t *testing.T) {
	assert.Equal(t, 1, minCostToMoveChips([]int{1, 2, 3}))
}

func Test_minCostToMoveChips2(t *testing.T) {
	assert.Equal(t, 2, minCostToMoveChips([]int{2, 2, 2, 3, 3}))
}

func Test_minCostToMoveChips3(t *testing.T) {
	assert.Equal(t, 1, minCostToMoveChips([]int{1, 1000000000}))
}
