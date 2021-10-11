package task0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{}))
	assert.Equal(t, "", longestCommonPrefix([]string{""}))
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))

}
