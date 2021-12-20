package task0140

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	assert.Equal(t, []string{
		"cats and dog", "cat sand dog",
	}, wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

func Test_wordBreak2(t *testing.T) {
	assert.Equal(t, []string{
		"pine apple pen apple", "pineapple pen apple", "pine applepen apple",
	}, wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
}

func Test_wordBreak3(t *testing.T) {
	assert.Equal(t, []string{}, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func Test_wordBreak4(t *testing.T) {
	assert.Equal(t,
		[]string{
			"a a a a", "aa a a", "a aa a", "aaa a", "a a aa", "aa aa", "a aaa",
		},
		wordBreak("aaaa", []string{"a", "aa", "aaa"}))
}
