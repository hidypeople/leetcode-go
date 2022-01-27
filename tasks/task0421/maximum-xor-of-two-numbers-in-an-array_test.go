package task0421

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaximumXOR(t *testing.T) {
	assert.Equal(t, 28, findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}

func Test_findMaximumXOR2(t *testing.T) {
	assert.Equal(t, 127, findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}))
}
