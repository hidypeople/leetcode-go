package task0151

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	assert.Equal(t, "the sky is blue", reverseWords("blue is sky the"))
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"))
}

func Test_reverseWords2(t *testing.T) {
	assert.Equal(t, "world hello", reverseWords("  hello world  "))
}

func Test_reverseWords3(t *testing.T) {
	assert.Equal(t, "example good a", reverseWords("  a good   example"))
}

func Test_reverseWords4(t *testing.T) {
	assert.Equal(t, "Alice Loves Bob", reverseWords("  Bob    Loves  Alice   "))
}

func Test_reverseWords5(t *testing.T) {
	assert.Equal(t, "bob like even not does Alice", reverseWords("Alice does not even like bob"))
}
