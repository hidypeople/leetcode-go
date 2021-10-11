package task0096

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numTrees(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
}

func Test_numTrees2(t *testing.T) {
	assert.Equal(t, 1, numTrees(1))
}
