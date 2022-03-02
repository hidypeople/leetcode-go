package task0392

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	assert.True(t, isSubsequence("abc", "ahbgdc"))
}

func Test_isSubsequence2(t *testing.T) {
	assert.False(t, isSubsequence("axc", "ahbgdc"))
}
