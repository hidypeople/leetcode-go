package task0017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	assert.Equal(t, []string{}, letterCombinations(""))
	assert.Equal(t, []string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
	}, letterCombinations("23"))
	assert.Equal(t, []string{
		"a", "b", "c",
	}, letterCombinations("2"))
}
