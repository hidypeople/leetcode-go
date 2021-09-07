package task0198

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob(t *testing.T) {
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
}

func Test_rob2(t *testing.T) {
	assert.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
}
