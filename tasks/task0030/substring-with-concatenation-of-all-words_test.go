package task0030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSubstring(t *testing.T) {
	result := findSubstring("aaaaAA_BB_CC_AA_CC_BB_aaaaaaCC_AA_BB_CC_", []string{"AA_", "BB_", "CC_", "CC_"})
	assert.Equal(t, []int{7, 10, 28}, result)
}

func Test_findSubstring2(t *testing.T) {
	result := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	assert.Equal(t, []int{0, 9}, result)
}

func Test_findSubstring3(t *testing.T) {
	result := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	assert.Equal(t, []int{}, result)
}

func Test_findSubstring4(t *testing.T) {
	result := findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
	assert.Equal(t, []int{6, 9, 12}, result)
}

func Test_findSubstring5(t *testing.T) {
	result := findSubstring("aaaaaaaaaaaaaa", []string{"aa", "aa"})
	assert.Equal(t, []int{0, 2, 4, 6, 8, 10, 1, 3, 5, 7, 9}, result)
}
