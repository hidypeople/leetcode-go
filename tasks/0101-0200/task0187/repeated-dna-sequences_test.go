package task0187

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	assert.Equal(t, []string{"AAAAACCCCC", "CCCCCAAAAA"},
		findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

func Test_findRepeatedDnaSequences2(t *testing.T) {
	assert.Equal(t, []string{"AAAAAAAAAA"},
		findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
