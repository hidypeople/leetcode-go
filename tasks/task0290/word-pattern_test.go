package task0290

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordPattern(t *testing.T) {
	assert.Equal(t, true, wordPattern("abba", "dog cat cat dog"))
}

func Test_wordPattern2(t *testing.T) {
	assert.Equal(t, false, wordPattern("abba", "dog cat cat fish"))
}

func Test_wordPattern3(t *testing.T) {
	assert.Equal(t, false, wordPattern("aaaa", "dog cat cat dog"))
}

func Test_wordPattern4(t *testing.T) {
	assert.Equal(t, false, wordPattern("abba", "dog dog dog dog"))
}

func Test_wordPattern5(t *testing.T) {
	assert.Equal(t, false, wordPattern("abc", "dog cat dog"))
}

func Test_wordPattern6(t *testing.T) {
	assert.Equal(t, false, wordPattern("abc", "abc"))
}

func Test_wordPattern7(t *testing.T) {
	assert.Equal(t, true, wordPattern("abc", "b c a"))
}
