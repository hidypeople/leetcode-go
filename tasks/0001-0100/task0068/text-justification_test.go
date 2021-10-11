package task0068

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fullJustify(t *testing.T) {
	assert.Equal(t,
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
		fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

func Test_fullJustify2(t *testing.T) {
	assert.Equal(t,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
		fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
}

func Test_fullJustify3(t *testing.T) {
	assert.Equal(t,
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
		fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}

func Test_fullJustify4(t *testing.T) {
	assert.Equal(t,
		[]string{
			"aaaaaaaaaa",
			"aaaaaaaaa ",
			"aaa  aa aa",
			"aaaaaaa   ",
		},
		fullJustify([]string{
			"aaaaaaaaaa",
			"aaaaaaaaa",
			"aaa", "aa", "aa",
			"aaaaaaa"}, 10))
}
