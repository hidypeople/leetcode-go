package task0127

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ladderLength(t *testing.T) {
	assert.Equal(t, 5, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func Test_ladderLength2(t *testing.T) {
	assert.Equal(t, 0, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
