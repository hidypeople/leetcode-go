package task1094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_carPooling(t *testing.T) {
	assert.False(t, carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
}

func Test_carPooling2(t *testing.T) {
	assert.True(t, carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
}
