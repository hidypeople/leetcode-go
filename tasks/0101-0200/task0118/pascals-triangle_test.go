package task0118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	assert.Equal(t, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, generate(5))
}

func Test_generate2(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, generate(1))
}
