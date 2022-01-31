package task1672

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumWealth(t *testing.T) {
	assert.Equal(t, 6, maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}

func Test_maximumWealth2(t *testing.T) {
	assert.Equal(t, 10, maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
}

func Test_maximumWealth3(t *testing.T) {
	assert.Equal(t, 17, maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
