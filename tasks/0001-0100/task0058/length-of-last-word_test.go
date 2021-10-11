package task0058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
}

func Test_lengthOfLastWord2(t *testing.T) {
	assert.Equal(t, 4, lengthOfLastWord("   fly me   to   the moon  "))
}

func Test_lengthOfLastWord3(t *testing.T) {
	assert.Equal(t, 6, lengthOfLastWord("luffy is still joyboy"))
}
