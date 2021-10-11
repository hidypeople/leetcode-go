package task0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePaths(t *testing.T) {
	assert.Equal(t, 28, uniquePaths(3, 7))
}

func Test_uniquePaths2(t *testing.T) {
	assert.Equal(t, 3, uniquePaths(3, 2))
}

func Test_uniquePaths3(t *testing.T) {
	assert.Equal(t, 6, uniquePaths(3, 3))
}

func Test_uniquePaths4(t *testing.T) {
	assert.Equal(t, 1916797311, uniquePaths(51, 9))
}
