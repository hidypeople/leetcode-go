package task0242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	assert.True(t, isAnagram("nagaram", "anagram"))
}

func Test_isAnagram2(t *testing.T) {
	assert.False(t, isAnagram("zagaram", "anagram"))
}

func Test_isAnagram3(t *testing.T) {
	assert.True(t, isAnagram("a", "a"))
}
