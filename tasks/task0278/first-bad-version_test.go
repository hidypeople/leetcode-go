package task0278

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstBadVersion(t *testing.T) {
	assert.Equal(t, 1, firstBadVersion(5))
}

func Test_firstBadVersion2(t *testing.T) {
	assert.Equal(t, 4, firstBadVersion_ForTest(5, func(x int) bool {
		return x >= 4
	}))
}

func Test_firstBadVersion3(t *testing.T) {
	assert.Equal(t, 1, firstBadVersion_ForTest(1, func(x int) bool {
		return x >= 1
	}))
}
