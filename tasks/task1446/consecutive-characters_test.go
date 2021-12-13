package task1446

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPower(t *testing.T) {
	assert.Equal(t, 2, maxPower("leetcode"))
	assert.Equal(t, 5, maxPower("abbcccddddeeeeedcba"))
	assert.Equal(t, 5, maxPower("triplepillooooow"))
	assert.Equal(t, 11, maxPower("hooraaaaaaaaaaay"))
	assert.Equal(t, 1, maxPower("tourist"))
}
