package task0438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAnagrams(t *testing.T) {
	assert.Equal(t, []int{0, 6}, findAnagrams("cbaebabacd", "abc"))
}

func Test_findAnagrams2(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, findAnagrams("abab", "ab"))
}
